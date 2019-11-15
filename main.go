package main

import (
	"flag"

	"github.com/mplulu/rano"
)

var token = flag.String("token", "", "token")
var groupId = flag.String("groupId", "", "groupId")
var message = flag.String("message", "", "message")

func main() {
	flag.Parse()

	r := rano.NewRano(*token, []string{*groupId})
	r.Send(*message)
}
