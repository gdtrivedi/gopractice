package netpackage

import (
	"io/ioutil"
	"net"
	"strings"
	"time"
)

const (
	//todo: move these to config
	ipWhoisServer = "whois.iana.org"
	whoisPort     = "43"
)

// WhoIs kicks off the whoIs query for a given IPV4 address and returns the results as a string
func WhoIs(domain string) (result string, err error) {

	result, err = query(domain)
	if err != nil {
		return
	}

	key := "whois:"
	server := GetValueByKey(key, result)
	if server == "" {
		return
	}

	tmpResult, err := query(domain, server)
	if err != nil {
		return
	}

	result += tmpResult

	return
}

func query(domain string, servers ...string) (result string, err error) {
	var server string

	if len(servers) == 0 || servers[0] == "" {
		server = ipWhoisServer
	} else {
		server = strings.ToLower(servers[0])
		if server == "whois.arin.net" {
			domain = "n + " + domain
		}
	}

	conn, e := net.DialTimeout("tcp", net.JoinHostPort(server, whoisPort), time.Second*2)
	if e != nil {
		err = e
		return
	}

	defer conn.Close()
	conn.Write([]byte(domain + "\r\n"))
	conn.SetReadDeadline(time.Now().Add(time.Second * 2))

	buffer, e := ioutil.ReadAll(conn)
	if e != nil {
		err = e
		return
	}

	result = string(buffer)

	return
}

// GetValueByKey returns the value that corresponds with the given key from the whoIs results string
func GetValueByKey(key string, whoIsResults string) string {

	start := strings.Index(whoIsResults, key)
	if start == -1 {
		return ""
	}

	start += len(key)
	end := strings.Index(whoIsResults[start:], "\n")
	return strings.TrimSpace(whoIsResults[start : start+end])
}
