package cmd

type Response struct {
	AccessToken string `json:"accessToken"`
	Identifier  string `json:"identifier"`
}
type Configuration struct {
	Server      string `json:"server"`
	AccessToken string `json:"accessToken"`
	Active      bool   `json:"active"`
	Alias       string `json:"alias"`
}
type ConfigList struct {
	ConfigList []Configuration `json:"configList"`
}

func (c Configuration) Init() Configuration {
	c.Active = false
	return c
}
func (c *ConfigList) AddConfig(config Configuration) []Configuration {
	c.ConfigList = append(c.ConfigList, config)
	return c.ConfigList
}
