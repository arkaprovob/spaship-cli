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

type Environment struct {
	Name              string `json:"name"`
	UpdateRestriction bool   `json:"updateRestriction"`
	Exclude           bool   `json:"exclude"`
}

type SpashipMapping struct {
	WebsiteVersion string        `json:"websiteVersion"`
	WebsiteName    string        `json:"websiteName"`
	Environments   []Environment `json:"environments"`
	SpaName        string        `json:"name"`
	Route          string        `json:"mapping"`
}

func (c Configuration) Init() Configuration {
	c.Active = false
	return c
}

func (c *ConfigList) AddConfig(config Configuration) []Configuration {
	c.ConfigList = append(c.ConfigList, config)
	return c.ConfigList
}

func (e Environment) Init() Environment {
	e.UpdateRestriction = false
	e.Exclude = false
	return e
}

func (s SpashipMapping) Init() SpashipMapping {
	s.WebsiteVersion = "v1"
	return s
}

func (s *SpashipMapping) AddEnvironment(e Environment) []Environment {
	s.Environments = append(s.Environments, e)
	return s.Environments
}
