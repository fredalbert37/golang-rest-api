package main

import (
	"fmt"
	"github.com/fredalbert37/golang-rest-api/config"
)

func main() {
	conf := config.GetConfig()
	fmt.Println(conf)
}
