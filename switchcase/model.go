package switchcase

import "time"

// DomainStatusReport represents the full report with install information and contains all domains and their statuses
type DomainStatusReport struct {
	Complete    bool           `json:"complete"`
	Id          string         `json:"id"`
	InstallName string         `json:"install_name"`
	InstallIP   string         `json:"install_ip"`
	Admin       bool           `json:"admin"`
	Timestamp   time.Time      `json:"-"`
	Domains     []DomainStatus `json:"domains"`
}

// DomainStatus represents an individual domain
type DomainStatus struct {
	Name              string `json:"name"`
	CName             string `json:"cname"`
	ARecord           string `json:"a_record"`
	Complete          bool   `json:"complete"`
	Result            string `json:"result"`
	DnsProvider       string `json:"dns_provider"`
	InstallName       string `json:"install_name"`
	ClusterID         string `json:"cluster_id"`
	HTTPForwarded     string `json:"http_forwarded"`
	HTTPForwardedList string `json:"http_forwarded_list"`
}

//Copy performs a deep copy of the DomainStatusReport instance
func (report *DomainStatusReport) Copy() *DomainStatusReport {

	c := DomainStatusReport{
		Complete:    report.Complete,
		Id:          report.Id,
		InstallName: report.InstallName,
		InstallIP:   report.InstallIP,
		Admin:       report.Admin,
		Timestamp:   report.Timestamp,
	}

	var cda []DomainStatus
	for _, domain := range report.Domains {
		cd := domain.Copy(report.Admin)
		cda = append(cda, *cd)
	}

	c.Domains = cda
	return &c
}

//CheckReportCompletion returns true if the all the domains in the report are in a complete state
func (report *DomainStatusReport) CheckReportCompletion() bool {
	for _, domain := range report.Domains {
		if !domain.Complete {
			return false
		}
	}
	return true
}

// Copy preforms a deep copy of a DomainStatus instance
func (ds DomainStatus) Copy(admin bool) *DomainStatus {
	c := DomainStatus{
		Name:        ds.Name,
		Complete:    ds.Complete,
		Result:      ds.Result,
		DnsProvider: ds.DnsProvider,
		CName:       ds.CName,
		ARecord:     ds.ARecord,
	}

	if admin {
		c.InstallName = ds.InstallName
		c.ClusterID = ds.ClusterID
		c.HTTPForwarded = ds.HTTPForwarded
		c.HTTPForwardedList = ds.HTTPForwardedList
	}

	return &c
}
