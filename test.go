package main

import (
	"fmt"
	"github.com/godaner/awdb-golang/awdb-golang"
	"log"
	"net"
)

func main() {
	db, err := awdb.Open("C:/Users/admin/Desktop/allawdb/IP_city_single_WGS84_awdb.awdb")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	ip := net.ParseIP("166.111.4.100")

	var record interface{}
	err = db.Lookup(ip, &record)
	if err != nil {
		log.Fatal(err)
	}
	var resMap = record.(map[string]interface{})
	fmt.Printf("accuracy:%s", resMap["accuracy"])
	fmt.Println()
	fmt.Printf("%s", record)
}
