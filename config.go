package main

import (
	"encoding/json"
	"io/ioutil"
)

type BotConfig struct {
	Name           string   `json:"name"`
	Server         string   `json:"server,omitempty"`
	SSL            bool     `json:"ssl,omitempty"`
	Nick           string   `json:"nick,omitempty"`
	Gecos          string   `json:"gecos,omitempty"`
	Sasl           bool     `json:"sasl,omitempty"`
	User           string   `json:"user,omitempty"`
	SaslUsername   string   `json:"sasl_username,omitempty"`
	SaslPassword   string   `json:"sasl_password,omitempty"`
	Autojoin       string   `json:"autojoin,omitempty"`
	Debug          bool     `json:"debug,omitempty"`
	Prefix         string   `json:"prefix,omitempty"`
	ServerPassword string   `json:"server_password,omitempty"`
	Modes          string   `json:"modes,omitempty"`
	Admins         []string `json:"admins,omitempty"`
}

func (b *IRCBot) LoadConfig() error {
	var conf = BotConfig{}
	cfile, err := ioutil.ReadFile("config.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(cfile, &conf)
	if err != nil {
		return err
	}
	b.config = conf
	return nil
}

func (b *BotConfig) IsAdmin(checkAdmin string) bool {
	for _, admin := range b.Admins {
		if admin[0:2] == "$a:" {
			if admin[3:] == checkAdmin {
				return true
			}
		}
	}
	return false
}
