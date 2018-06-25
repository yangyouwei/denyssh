package main

import (
	"denyssh/readsecure"
	"fmt"
)
var LOGS chan string = make(chan string)

func main()  {
	readsecure.LOGS_DIR = "D:\\passwd.txt"
	go readsecure.READ_LOG(LOGS)
	for   {
		s := <- LOGS
		ip_addr := readsecure.GET_IP(s)
		if ip_addr == "" {
			continue
		}
		fmt.Println(ip_addr)
	}

}
