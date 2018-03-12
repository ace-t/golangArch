package main

import (
	"time"
	"github.com/BurntSushi/toml"
	"fmt"
)

type Config struct {
	Age int
	Cats []string
	Pi float64
	Perfection []int
	DOB time.Time // requires `import time`
}

/*
toml 연동해보기 https://github.com/BurntSushi/toml
 */
func main() {
	var conf Config
	if _, err := toml.DecodeFile("config/configTest.toml", &conf); err != nil {
		// handle error
		fmt.Println(err)
	}
	fmt.Println(conf.Age)
	fmt.Println(conf.Cats)
	fmt.Println(conf.Pi)
	fmt.Println(conf.Perfection)
	fmt.Println(conf.DOB)
}
