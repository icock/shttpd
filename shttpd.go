// This source code is dedicated to the Public Domain.

// Package shttpd is a static http server serving files under current directory,
// i.e. `python2 -m SimpleHTTPServer` in Go.
package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"log"
)

// #include "sysexits.h"
import "C"

func main() {
	var port uint16
	if len(os.Args) == 1 {
		port = 8000
	} else if len(os.Args) == 2 {
		var err error
		port, err = parseArgument(os.Args[1])
		if err != nil {
			usage()
		}
	} else {
		usage()
	}

	workingDirectory, _ := os.Getwd()

	fmt.Printf("Serving HTTP on port %d (http://127.0.0.1:8000/) ...", port)
	err := http.ListenAndServe(
		":"+strconv.Itoa(int(port)),
		http.FileServer(http.Dir(workingDirectory)))
	log.Fatal(err)
}

func parseArgument(port string) (uint16, error) {
	ret, err := strconv.ParseInt(port, 10, 16)
	return uint16(ret), err
}
func usage() {
	fmt.Println("Usage: shttpd [PORT]")
	fmt.Println("PORT must be an integer between 0 and 65535")
	os.Exit(C.EX_USAGE)
}
