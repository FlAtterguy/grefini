package main

import (
	"net"
	"encoding/json"
	"bufio"
	"fmt"
)

type User struct {
	Nickname string
	Username string
	Host	 string
}

type Message struct {
	Source	 User
	Command	 string
	Target	 string
	Text	 string
}

const MESSAGE = "PRIVMSG"
const NOTICE  = "NOTICE"

var sock net.Conn

func initmods() {
	initquote()
	initmacro()
	initball()
	initmeta()
	initdoge()
	initurl()
	initsed()
	initmamma()
	initrandst()
	fmt.Println("flotrshi-port - All modules loaded!")
}

func handle(msg Message) {
	quote(msg)
	macro(msg)
	ball(msg)
	meta(msg)
	doge(msg)
	url(msg)
	logSearch(msg)
	sed(msg)
	mamma(msg)
	randst(msg)
}

func main () {
	// Init flotrshi-port modules
	initmods()

	// Connect to the proxy
	var err error
	sock, err = net.Dial("tcp","127.0.0.1:4314")
	if err != nil { panic(err) }
	defer sock.Close()

	in := bufio.NewReader(sock)
	for {
		bytes, _, err := in.ReadLine()
		if err != nil { break }

		var msg Message
		err = json.Unmarshal(bytes,&msg)
		if err != nil { fmt.Printf("ERROR reading JSON: %s\r\n",err.Error()) }
		// Dispatch to flotrshi-port modules
		handle(msg)
	}
}

func send(msg Message) {
	bytes, _ := json.Marshal(msg)
	fmt.Fprintf(sock,string(bytes)+"\r\n")
}
