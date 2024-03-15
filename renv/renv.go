package renv

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

var envMode = flag.String("envMode", "", "env mode")
var ParseAtLocationParam = flag.String("parse_at_location", "", "parse at location")

func ParseCmd(v interface{}) {
	Parse(*envMode, *ParseAtLocationParam, v)
}

func Parse(env, parseAtLocation string, v interface{}) {
	var fileName string
	if env == "" {
		fileName = "./.env.local.yaml"
	} else {
		fileName = fmt.Sprintf("./.env.%s.yaml", env)
	}
	if parseAtLocation != "" {
		fileName = strings.TrimPrefix(fileName, "./")
		fileName = filepath.Join(parseAtLocation, fileName)

	}
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Println(fileName)
		panic(fmt.Sprintf("missing env file %v", fileName))
	}

	ParseAtLocation(fileName, v)

}
func ParseAtLocation(fileName string, v interface{}) {
	raw, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	// parse to v
	err = yaml.Unmarshal(raw, v)
	if err != nil {
		panic(err)
	}
}
