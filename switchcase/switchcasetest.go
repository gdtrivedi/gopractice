package switchcase

import (
	"fmt"
	"strings"
)

type SwitchCaseIn struct {
	Hostnames []string
	Install   *string
}

type httpResults struct {
	forwardedList    string
	forwarded        string
	installName      string
	clusterID        string
	edge             string
	httpError        bool
	httpErrorMessage string
}

type dnsRequest struct {
	reportId     string
	installName  string
	installIP    string
	domainStatus DomainStatus
}

// DNSResult the output of the lookup
type DNSResult struct {
	cname       string
	domainIP    string
	dnsProvider string
	error       bool
}

const (
	okWPEngine        = "ok_wpengine"
	okCNAME           = "ok_cname"
	okNotCNAME        = "ok_not_cname"
	okProxied         = "ok_proxied"
	okGlobalIngress   = "ok_global_ingress"
	errorWrongInstall = "error_wrong_install"
	errorNotWPE       = "error_not_wpe"
	errorDoesNotExist = "error_does_not_exist"
	errorHTTP         = "error_http"
	errorProxiedByWPE = "error_proxied_by_wpe"
)

func SwitchCaseTest() {
	dnsResult := DNSResult{
		cname:       "",
		domainIP:    "",
		dnsProvider: "",
		error:       true,
	}
	httpResult := httpResults{
		forwardedList:    "",
		forwarded:        "",
		installName:      "",
		clusterID:        "",
		edge:             "",
		httpError:        false,
		httpErrorMessage: "",
	}
	dnsReq := dnsRequest{
		reportId:    "",
		installName: "",
		installIP:   "",
		domainStatus: DomainStatus{
			Name:              "",
			CName:             "",
			ARecord:           "",
			Complete:          false,
			Result:            "",
			DnsProvider:       "",
			InstallName:       "",
			ClusterID:         "",
			HTTPForwarded:     "",
			HTTPForwardedList: "",
		},
	}
	fmt.Println("IfElse vs SwitchCase: ", computeStatusWithIfElse(dnsResult, httpResult, dnsReq) == computeStatusWithSwitchCase(dnsResult, httpResult, dnsReq))

	dnsResult.error = false
	httpResult.httpError = true
	fmt.Println("IfElse vs SwitchCase: ", computeStatusWithIfElse(dnsResult, httpResult, dnsReq) == computeStatusWithSwitchCase(dnsResult, httpResult, dnsReq))

	httpResult.httpError = false
	httpResult.installName = ""
	fmt.Println("IfElse vs SwitchCase: ", computeStatusWithIfElse(dnsResult, httpResult, dnsReq) == computeStatusWithSwitchCase(dnsResult, httpResult, dnsReq))

	httpResult.installName = "i1"
	dnsReq.installName = "i2"
	fmt.Println("IfElse vs SwitchCase: ", computeStatusWithIfElse(dnsResult, httpResult, dnsReq) == computeStatusWithSwitchCase(dnsResult, httpResult, dnsReq))

	dnsReq.installName = "i1"
	httpResult.forwarded = "forwarded"
	fmt.Println("IfElse vs SwitchCase: ", computeStatusWithIfElse(dnsResult, httpResult, dnsReq) == computeStatusWithSwitchCase(dnsResult, httpResult, dnsReq))

	httpResult.forwarded = ""
	httpResult.edge = "GI"
	fmt.Println("IfElse vs SwitchCase: ", computeStatusWithIfElse(dnsResult, httpResult, dnsReq) == computeStatusWithSwitchCase(dnsResult, httpResult, dnsReq))

	httpResult.edge = ""
	dnsResult.domainIP = "1.2.3.4"
	dnsReq.installIP = "1.2.3.1"
	fmt.Println("IfElse vs SwitchCase: ", computeStatusWithIfElse(dnsResult, httpResult, dnsReq) == computeStatusWithSwitchCase(dnsResult, httpResult, dnsReq))

	dnsReq.installIP = "1.2.3.4"
	dnsReq.installName = "myinstall"
	dnsResult.cname = dnsReq.installName + ".wpengine.com."
	fmt.Println("IfElse vs SwitchCase: ", computeStatusWithIfElse(dnsResult, httpResult, dnsReq) == computeStatusWithSwitchCase(dnsResult, httpResult, dnsReq))

	dnsReq.installName = "myinstall1"
	dnsReq.domainStatus.Name = dnsReq.installName + ".wpengine.com"
	fmt.Println("IfElse vs SwitchCase: ", computeStatusWithIfElse(dnsResult, httpResult, dnsReq) == computeStatusWithSwitchCase(dnsResult, httpResult, dnsReq))

	httpResult.edge = "GI"
	dnsResult.domainIP = "1.2.3.4"
	dnsReq.installIP = "1.2.1.2"
	httpResult.installName = "i1"
	dnsReq.installName = "i1"
	dnsReq.domainStatus.Name = ""
	fmt.Println("IfElse vs SwitchCase: ", computeStatusWithIfElse(dnsResult, httpResult, dnsReq) == computeStatusWithSwitchCase(dnsResult, httpResult, dnsReq))
}

// computeStatus parses the dnsLookup and httpLookup results and returns the status
func computeStatusWithIfElse(dnsResult DNSResult, httpResult httpResults, dnsRequest dnsRequest) string {

	status := okNotCNAME

	if dnsResult.error {
		status = errorDoesNotExist

	} else if httpResult.httpError {
		status = errorHTTP

	} else if httpResult.installName == "" {
		status = errorNotWPE

	} else if httpResult.installName != dnsRequest.installName {
		status = errorWrongInstall

	} else if httpResult.forwarded != "" {
		status = errorProxiedByWPE

	} else if httpResult.edge == "GI" {
		status = okGlobalIngress

	} else if dnsResult.domainIP != dnsRequest.installIP {
		status = okProxied

	} else if dnsResult.cname == dnsRequest.installName+".wpengine.com." {
		status = okCNAME
	}

	if strings.HasSuffix(dnsRequest.domainStatus.Name, "wpengine.com") {
		status = okWPEngine
	}

	return status
}

// computeStatus parses the dnsLookup and httpLookup results and returns the status
func computeStatusWithSwitchCase(dnsResult DNSResult, httpResult httpResults, dnsRequest dnsRequest) string {

	status := okNotCNAME

	switch {
	case dnsResult.error:
		status = errorDoesNotExist
		break
	case httpResult.httpError:
		status = errorHTTP
		break
	case httpResult.installName == "":
		status = errorNotWPE
		break
	case httpResult.installName != dnsRequest.installName:
		status = errorWrongInstall
		break
	case httpResult.forwarded != "":
		status = errorProxiedByWPE
		break
	case httpResult.edge == "GI":
		status = okGlobalIngress
		break
	case dnsResult.domainIP != dnsRequest.installIP:
		status = okProxied
		break
	case dnsResult.cname == dnsRequest.installName+".wpengine.com.":
		status = okCNAME
		break
	}
	if strings.HasSuffix(dnsRequest.domainStatus.Name, "wpengine.com") {
		status = okWPEngine
	}

	return status
}
