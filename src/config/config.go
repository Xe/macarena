/*
Package config manages the configuration for macarena.

The following is an example configuration file:

        {
                "networks": [
                        {
                                "name": "ShadowNET",
                                "host": "127.0.0.1",
                                "port": 5335,
                                "ssl":  false,
                                "pass": "foobang"
                        },
                        {
                                "name": "ShadowNET-2",
                                "host": "127.0.0.1",
                                "port": 5336,
                                "ssl":  false,
                                "pass": "foobang"
                        }
                ],
               "myinfo": {
                        "nick": "Foobang",
                        "user": "bar",
                        "real": "fake info"
                },
                "channels": ["#test", "#spam"]
        }
*/
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Validator is an interface to see if a configuration element is valid.
type Validator interface {
	Validate() (correct bool)
}

// Config manages the configuration for macarena. It is a container for the
// various data users may want to configure.
type Config struct {
	Networks []Network `json:"networks"` // Networks to connect to
	MyInfo   Info      `json:"myinfo"`   // Information about the bot
	Channels []string  `json:"channels"` // Channels to relay
}

// Network is a container representing an IRC network.
type Network struct {
	Name         string `json:"name"` // Name of network for logging
	Host         string `json:"host"` // Hostname of server to connect to
	Port         int    `json:"port"` // Port of server to connect to
	UseSSL       bool   `json:"ssl"`  // Use SSL?
	ServicesPass string `json:"pass"` // Services password (/ns IDENTIFY)
}

// Info is info about the bot.
type Info struct {
	Nick  string `json:"nick"` // Nickname of bot
	User  string `json:"user"` // Username of bot
	Gecos string `json:"real"` // Realname of bot
}

// LoadFile wraps Load with opening and closing a file.
func LoadFile(path string) (cfg Config, err error) {
	fin, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer fin.Close()

	return Load(fin)
}

// Load returns a Config based on the given io.Reader or an error.
func Load(fin io.Reader) (cfg Config, err error) {
	data, err := ioutil.ReadAll(fin)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err
	}

	if !cfg.MyInfo.Validate() {
		return Config{}, errors.New("Invalid information about the bot")
	}

	for id, network := range cfg.Networks {
		if !network.Validate() {
			return Config{}, fmt.Errorf("Network id %d (%s) failed validation", id, network.Name)
		}
	}

	if len(cfg.Channels) == 0 {
		return Config{}, errors.New("No channels to join")
	}

	return
}

// Validate validates information about an Info.
func (i Info) Validate() bool {
	return i.Gecos != "" && i.Nick != "" && i.User != ""
}

// Validate validates information about a Network.
func (n Network) Validate() bool {
	return n.Name != "" && n.Host != "" && n.Port != 0
}
