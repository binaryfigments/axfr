package main

import (
	"encoding/json"
	"fmt"

	dnsaxfr "github.com/binaryfigments/axfr"
)

func main() {
	// hostname := "zonetransfer.me"
	hostname := "databyte.nl"
	nameserver := "8.8.8.8"

	axfrdata := dnsaxfr.Get(hostname, nameserver)

	json, err := json.MarshalIndent(axfrdata, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", json)
}
