package main

import (
	"clients/telegram"
	"log"
)

func main() {
	tg, err := telegram.New()
	if err != nil {
		log.Fatal(err)
	}
	err = tg.Start()
	if err != nil {
		log.Fatal(err)
	}
}
