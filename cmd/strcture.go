/*
Copyright © 2022 Arkaprovo Bhattacharjee <apb@live.in>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
