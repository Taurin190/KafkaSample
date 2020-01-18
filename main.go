package main

import (
	"fmt"
	"os"
	"os/signal"
)

func usage() {
	fmt.Println("usage: go run main.go <command>")
	fmt.Println("")
	fmt.Println("followings are commands:")
	fmt.Println("\tpserver: create producer server")
	fmt.Println("\tcprinter: create consumer print process")
	os.Exit(1)
}

func main() {

	usage()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	fmt.Println("go-kafka-example start.")

	<-signals

	fmt.Println("go-kafka-example stop.")
}
