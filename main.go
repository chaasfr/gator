package main

import (
	"fmt"

	"github.com/chaasfr/gator/internal/config"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	conf.SetUser("laBaguette")

	conf, err = config.Read()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(conf)
}