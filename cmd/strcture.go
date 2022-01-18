package cmd

type Response struct {
	AccessToken string `json:"accessToken"`
}
type Configuration struct {
	Server      string `json:"server"`
	AccessToken string `json:"accessToken"`
	Active      bool
}
type ConfigList struct {
	ConfigList []Configuration `json:"configList"`
}

func (c Configuration) Init() Configuration {
	c.Active = true
	return c
}
