package dto

type HostInfo struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type BruteList struct {
	User string
	Pass string
}
