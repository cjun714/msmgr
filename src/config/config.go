package config

import (
	"os"
)

// Port web server port
const Port int = 8080

// YamlPath .yaml file path
const YamlPath string = "./yamls/"

// AppName application name
const AppName string = "go.demo"

// EurekaURL Eureka url
var EurekaURL = "http://127.0.0.1:8761"

func init() {
	EurekaURL = os.Args[1]
}
