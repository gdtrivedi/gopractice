package main

type PayloadReader struct {
	Messages []Message `json:"messages"`
	Qos int `json:"qos"`
	TopicFilter string `json:"topicFilter"`
}

type Message struct {
	Format string `json:"format"`
	Payload Payload `json:"payload"`
	Qos int `json:"qos"`
	Timestamp int64 `json:"timestamp"`
	Topic string `json:"topic"`
}

type Payload struct {
	Payload map[string]interface{} `json:"payload"`
	Timestamp string `json:"timestamp"`
}

type DHCPPayload struct {
	Timestamp string `json:"timestamp,omitempty"`
	Tid string `json:"tid,omitempty"`
	Client    string   `json:"client,omitempty"`
	Connect   string	`json:"connect,omitempty"`
	Device string `json:"device,omitempty"`
	Expires string `json:"expires,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Ip string `json:"ip,omitempty"`
	Mac string `json:"mac,omitempty"`
	Ssid string `json:"ssid,omitempty"`
}