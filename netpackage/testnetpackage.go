package netpackage

import (
	"fmt"
	"net"
)

func TestNetPackage() {
	host := "gautamtrivedi.wpengine.com"
	fmt.Println(net.LookupHost(host))
	fmt.Println(net.LookupAddr(host))
	fmt.Println(net.LookupCNAME(host))
	fmt.Println(net.LookupIP(host))
	fmt.Println(net.LookupNS(host))
	fmt.Println(net.LookupMX(host))
	/*
	[52.216.143.75] <nil>
	[] lookup gtrivedi.com: invalid address
	gtrivedi.com. <nil>
	[52.216.143.75] <nil>

	*/

	if whoIsResults, err := WhoIs("104.154.54.137"); err != nil {
		fmt.Errorf("Error: ", err)
	} else {
		fmt.Println(whoIsResults)
		netNameKey := "NetName:"
		netName := GetValueByKey(netNameKey, whoIsResults)
		fmt.Println(netName)
	}
}