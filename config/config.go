package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

/*
This config file structure works well if you want to roll your own config fies.
Use three separate files:
  1. env.yaml (global settings)
  2. ini.yaml (private passwords)
  3. and settings.yaml (application settings)
And put them in separate server dirs:
  1. development
  2. staging
  3. production

Put symlinks in:
  /etc/<project_name>/
    /development
      env.yaml # <-- these are symlinks
      ini.yaml
      settings.yaml
    /staging
      env.yaml
      ini.yaml
      settings.yaml
    /production
      env.yaml
      ini.yaml
      settings.yaml
Symlinks point to actual files located here:
  /<project dir>/config/ignore # <-- /ignore in .gitignore to keep files out of repo
    /development
      env.yaml # <-- actual files
      ini.yaml
      settings.yaml
    /staging
      env.yaml
      ini.yaml
      settings.yaml
    /production
      env.yaml
      ini.yaml
      settings.yaml

  Switch between server environments use:
  SERVER=[development|staging|production] go run main.go
*/

type ServerEnv int

const (
	Development ServerEnv = iota
	Staging
	Production
)

func (s ServerEnv) String() string {
	switch s {
	case Development:
		return "development"
	case Staging:
		return "staging"
	case Production:
		return "production"
	}
	return "unknown"
}

var (
	Conf *Config
)

const (
	settingsDir = "/etc/db_controller_example"
)

func init() {
	// load all settings files, func Load() at bottom of page.
	Conf = Load()
}

// struct can be passed to middleware and handlers that need it.
type Config struct {
	// #######################################
	// Env
	// #######################################
	Check           string `yaml:"check"`
	Server          string `yaml:"server"`
	DevelopmentBase string `yaml:"developmentBase"`
	StagingBase     string `yaml:"stagingBase"`
	ProductionBase  string `yaml:"productionBase"`
	ApiPort         string `yaml:"apiPort"`
	// #######################################
	// INI
	// #######################################
	RedisCredentials struct {
		Netw string `yaml:"netw"`
		Addr string `yaml:"addr"`
		Pass string `yaml:"pass"`
		Rurl string `yaml:"urul"`
		User string `yaml:"user"`
	} `yaml:"redisDb"`
	PostgresCredentials struct {
		Dnam string `yaml:"dnam"`
		Host string `yaml:"host"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Port string `yaml:"port"`
	} `yaml:"postgresDb"`
	DoSpaces struct {
		UseSpaces    string `yaml:"useSpaces"`
		AccessKey    string `yaml:"accessKey"`
		Secret       string `yaml:"secret"`
		BucketName   string `yaml:"bucketName"`
		CustomDomain string `yaml:"customDomain"`
		RegionName   string `yaml:"regionName"`
		EndpointUrl  string `yaml:"endpointUrl"`
		VanityUrl    string `yaml:"vanityUrl"`
	} `yaml:"doSpaces"`
	// #######################################
	// Settings
	// #######################################
	// api
	ThresholdTime int64 `yaml:"thresholdTime"`
	ThresholdHits int64 `yaml:"thresholdHits"`
	HoldCounter   int64 `yaml:"holdCounter"`
	HoldTime      int64 `yaml:"holdTime"`
	// media dirs
	TempFiles    string `yaml:"tempFiles"`
	MediaRoot    string `yaml:"mediaRoot"`
	UploadPrefix string `yaml:"uploadPrefix"`
}

// env.yaml file should be placed in same dir as this go file
func (c *Config) GetEnv() {
	// build the dirs
	server := os.Getenv("SERVER")
	if server == "" {
		server = "development"
	}
	fp := filepath.Join(settingsDir, strings.ToLower(server), "env.yaml")
	// check for the ini.yaml file
	if _, err := os.Stat(fp); err != nil {
		fmt.Printf("Error %v\n", err)
		panic("Cannot find env.yaml file.")
	}
	// read it
	yf, err := ioutil.ReadFile(fp)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		panic("Cannot read the ini.yaml file.")
	}
	// unmarshal it
	err = yaml.Unmarshal(yf, c)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		panic("Cannot unmarshal the ini.yaml file.")
	}
	// set all key/value pairs
	idx := map[string]string{
		"CHECK":           c.Check,
		"SERVER":          c.Server,
		"DEVLOPMENT_BASE": c.DevelopmentBase,
		"STAGING_BASE":    c.StagingBase,
		"PRODUCTION_BASE": c.ProductionBase,
		"APIPORT":         c.ApiPort,
	}
	for k, v := range idx {
		err = os.Setenv(k, v)
		if err != nil {
			fmt.Println("Error loading env vars.", err)
			log.Fatal("Error loading env vars.")
		}
	}
}

func (c *Config) GetIni() {
	// build the dirs
	server := os.Getenv("SERVER")
	fp := filepath.Join(settingsDir, strings.ToLower(server), "ini.yaml")
	// check for the ini.yaml file
	if _, err := os.Stat(fp); err != nil {
		fmt.Printf("Error %v\n", err)
		panic("Cannot find ini.yaml file.")
	}
	// read it
	yf, err := ioutil.ReadFile(fp)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		panic("Cannot read the ini.yaml file.")
	}
	// unmarshal it
	err = yaml.Unmarshal(yf, c)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		panic("Cannot unmarshal the ini.yaml file.")
	}
}

func (c *Config) GetSettings() {
	// build the dirs
	server := os.Getenv("SERVER")
	fp := filepath.Join(settingsDir, strings.ToLower(server), "settings.yaml")
	// check for the settings.yaml file
	if _, err := os.Stat(fp); err != nil {
		fmt.Printf("Error %v\n", err)
		panic("Cannot find settings.yaml file.")
	}
	// read it
	yf, err := ioutil.ReadFile(fp)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		panic("Cannot read the settings.yaml file.")
	}
	// unmarshal it
	err = yaml.Unmarshal(yf, c)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		panic("Cannot unmarshal the settings.yaml file.")
	}
}
func Load() *Config {
	c := new(Config)
	c.GetEnv()
	c.GetIni()
	c.GetSettings()
	return c
}
