package main

import (
	"flag"
	"fmt"
	"mplulu/srano/renv"
	"os"
	"strings"

	"github.com/mplulu/rano"
)

var messageArg = flag.String("message", "", "message")

type ENV struct {
	Token   string `yaml:"token"`
	GroupId string `yaml:"group_id"`
}

func main() {
	flag.Parse()
	var env *ENV
	homePath, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	renv.Parse("", fmt.Sprintf("%v/.srano", homePath), &env)

	message := *messageArg
	if message == "" {
		message = strings.Join(os.Args[1:], " ")
	}

	r := rano.NewRano(env.Token, []string{env.GroupId})
	err = r.Send(message)
	if err != nil {
		panic(err)
	}
}
