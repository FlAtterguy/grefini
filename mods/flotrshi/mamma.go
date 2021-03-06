package main

import (
	"math/rand"
	"fmt"
	"strings"
)

var mreplies []string = []string { "X non da sbocchi", "Con te non faremo gli stessi errori che abbiamo fatto con X", "Vai a studiare X!" , "Finirai a spalare X!", "X ti ha rincoglionito!", "Non ti meriti X", "Tuo padre lavora tutto il giorno e tu, invece, fai X", "DOV'E' Y?", "EH MA NON E' ORA CHE TI TROVI Y?!" }

func initmamma() {
	fmt.Println("Mamma di Jhy pronta! (!mamma)")
}

func mamma(msg Message) {
	if msg.Command == MESSAGE {
		if len(msg.Text) < 6 { return }
		if msg.Text[0:6] == "!mamma" {
			n := rand.Intn(len(mreplies))
			final := strings.Replace(mreplies[n],"X",msg.Text[7:],-1)
			if strings.Index(mreplies[n],"Y") >= 0 {
				final = strings.Replace(mreplies[n],"Y",strings.ToUpper(msg.Text[7:]),1)
			}
			send(Message{
				Command: MESSAGE,
				Target : msg.Target,
				Text   : final,
			})
			return
		}
	}
}
