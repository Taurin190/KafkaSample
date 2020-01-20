package main

import (
	"flag"
	"fmt"
	"os"

	"./consumer"
	"./producer/infrastructure"
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
	flag.Parse()
	kafkaServers := []string{"kafkaServers", "localhost:32770", "kafka address"}
	args := flag.Args()
	if len(args) == 0 {
		usage()
	}
	fmt.Printf("args : %s\n", args[0])
	if args[0] == "pserver" {
		infrastructure.Router.Run()
	} else if args[0] == "cprinter" {
		consumer.ConsumerPrint(kafkaServers)
	} else {
		usage()
	}
}
