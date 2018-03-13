package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

type Config struct{
	Workspace string `json:"worksapce"`
	Name string `json:"name"`
	Age int64 `json:"age"`
}

// get toml data
func (config *Config) LoadConfig(path string)(*Config, error){
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, fmt.Errorf("LoadConfig: %v", err)
	}
	return config, nil
}
func flags(){
	var cfg *Config
	// -conf=config/real.toml or -conf=config/dev.toml
	flag.StringVar(&configPath, "conf", "", "config path(dev or real).")
	flag.Parse()
	fmt.Println(configPath)

	// get tomml data
	c, err := cfg.LoadConfig(configPath)
	if err != nil {
		fmt.Println(err)
	}
    //fmt.Println(c.Workspace)
    fmt.Println(c.Name)
    fmt.Println(c.Age)

}

func main() {
	flags()
}
