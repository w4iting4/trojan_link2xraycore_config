package main

type Config struct {
	Log       Log                    `json:"log"`
	API       API                    `json:"api"`
	Inbounds  []Inbound              `json:"inbounds"`
	Outbounds []Outbound             `json:"outbounds"`
	Policy    Policy                 `json:"policy"`
	Routing   Routing                `json:"routing"`
	Stats     map[string]interface{} `json:"stats"`
}

type Log struct {
	Access   string `json:"access"`
	DNSLog   bool   `json:"dnsLog"`
	Error    string `json:"error"`
	Loglevel string `json:"loglevel"`
}

type API struct {
	Tag      string   `json:"tag"`
	Services []string `json:"services"`
}

type Inbound struct {
	Tag            string                 `json:"tag"`
	Listen         *string                `json:"listen"`
	Port           int                    `json:"port"`
	Protocol       string                 `json:"protocol"`
	Settings       map[string]interface{} `json:"settings"`
	StreamSettings *StreamSettings        `json:"streamSettings"`
	Sniffing       *Sniffing              `json:"sniffing"`
}

type Outbound struct {
	Tag            string           `json:"tag"`
	Protocol       string           `json:"protocol"`
	Settings       OutboundSettings `json:"settings"`
	StreamSettings *StreamSettings  `json:"streamSettings"`
}

type OutboundSettings struct {
	Servers []Server `json:"servers"`
}

type Server struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Password string `json:"password"`
}

type StreamSettings struct {
	Network     string       `json:"network"`
	Security    string       `json:"security"`
	TLSSettings *TLSSettings `json:"tlsSettings"`
	TCPSettings *TCPSettings `json:"tcpSettings"`
}

type TLSSettings struct {
	ServerName    string   `json:"serverName"`
	ALPN          []string `json:"alpn"`
	Fingerprint   string   `json:"fingerprint"`
	AllowInsecure bool     `json:"allowInsecure"`
}

type TCPSettings struct {
	Header Header `json:"header"`
}

type Header struct {
	Type string `json:"type"`
}

type Sniffing struct {
	// Define fields if needed
}

type Policy struct {
	Levels map[string]Level `json:"levels"`
	System System           `json:"system"`
}

type Level struct {
	StatsUserDownlink bool `json:"statsUserDownlink"`
	StatsUserUplink   bool `json:"statsUserUplink"`
}

type System struct {
	StatsInboundDownlink  bool `json:"statsInboundDownlink"`
	StatsInboundUplink    bool `json:"statsInboundUplink"`
	StatsOutboundDownlink bool `json:"statsOutboundDownlink"`
	StatsOutboundUplink   bool `json:"statsOutboundUplink"`
}

type Routing struct {
	DomainStrategy string `json:"domainStrategy"`
	Rules          []Rule `json:"rules"`
}

type Rule struct {
	Type        string   `json:"type"`
	InboundTag  []string `json:"inboundTag"`
	OutboundTag string   `json:"outboundTag"`
	IP          []string `json:"ip"`
	Protocol    []string `json:"protocol"`
}
