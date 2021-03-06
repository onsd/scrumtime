package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

// NewAppFromYaml returns the app data from provided yaml file
func NewAppFromYaml(path string) (*App, error) {
	app := new(App)

	// Get data from provided yaml file
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		errstr := fmt.Sprintf("Error while reading config file: %v", err)
		return nil, errors.New(errstr)
	}
	err = yaml.Unmarshal(raw, app)

	if err != nil {
		return nil, err
	}

	err = app.Validate()
	if err != nil {
		return nil, err
	}

	return app, nil
}

// App represents the app's configuration
type App struct {
	Messengers map[string]*Messenger `yaml:"messengers"`
	Schedules  map[string]*Schedule  `yaml:"schedules"`
}

// Validate validates an app config
func (a *App) Validate() error {
	if len(a.Schedules) == 0 {
		return fmt.Errorf("config file doesn't contain schedules")
	}

	if len(a.Messengers) == 0 {
		return fmt.Errorf("config file doesn't contain messengers")
	}

	for _, s := range a.Schedules {
		for _, msgr := range s.Messengers {
			if _, ok := a.Messengers[msgr]; !ok {
				return fmt.Errorf("messenger %s not defined", msgr)
			}
		}
	}

	return nil
}

// addIndent adds an indentation after each new line in the provided string
func addIndent(s string) string {
	return strings.Replace(s, "\n", "\n\t", -1)
}

// Schedule represents the configuration of a single schedule
type Schedule struct {
	Messengers []string `yaml:"messengers"`
	Schedule   string   `yaml:"schedule"`
	Message    string   `yaml:"message"`
}

// Messenger represents a messenger config
type Messenger struct {
	Platform string `yaml:"platform"`
	ChatID   string `yaml:"chat_id"`
	APIKey   string `yaml:"api_key"`
}

func truncateString(str string, num int) string {
	bnoden := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		bnoden = str[0:num] + "..."
	}
	return bnoden
}
