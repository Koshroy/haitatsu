package main

import (
	"fmt"
	"github.com/Koshroy/imap"
)

func main() {
	fmt.Println("hello world!")
	imap.HandleFunc("ehlo", func(s string, reply *imap.ImapReply) { reply.Status.Write([]byte("received " + s))})
	imap.ListenAndServe("localhost:13337")
}
