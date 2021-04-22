package structtojson

import (
	"encoding/json"
	"fmt"
	"time"
)

// DNSRecordSpec represents the specification (desired state) of a particular DNSRecord
type DNSRecordSpec struct {
	URL        string    `json:"url"`
	RecordType string    `json:"record_type"`
	Domain     string    `json:"domain"`
	Content    string    `json:"content"`
	Proxied    bool      `json:"proxied,omitempty"`
	ZoneID     string    `json:"cloudflare_zone_id"`
	Removed    bool      `json:"removed"`
	CreatedOn  time.Time `json:"created_on,omitempty"`
	UpdatedOn  time.Time `json:"updated_on,omitempty"`
}

func TestDNSRecordSpecStructToJson() {
	spec := &DNSRecordSpec{
		RecordType: "A",
		Domain:     "abc.xyz.com",
		Content:    "1.2.3.4",
		Removed:    false,
	}
	b, err := json.Marshal(spec)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
