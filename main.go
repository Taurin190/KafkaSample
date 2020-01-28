package main

import (
	"flag"
	"fmt"
	"os"

	"./config"
	"./consumer"
	"./producer/infrastructure"
)

func usage() {
	fmt.Println("usage: go run main.go <command>")
	fmt.Println("")
	fmt.Println("followings are commands:")
	fmt.Println("\tpserver: create producer server")
	fmt.Println("\tcprinter: create consumer print process")
	fmt.Println("\tcpreserver: create consumer preserving process (please run mongodb on local:27017)")
	os.Exit(1)
}

func main() {
	conf := config.GetConfig()
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		usage()
	}
	fmt.Printf("args : %s\n", args[0])
	if args[0] == "pserver" {
		infrastructure.Router.Run()
	} else if args[0] == "cprinter" {
		c := consumer.GetConsumerPrinter(*conf)
		c.Run()
	} else if args[0] == "cpreserver" {
		c := consumer.GetConsumerPreserver(*conf)
		c.Run()
	} else {
		usage()
	}
}
