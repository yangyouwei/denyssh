package main

import (
	"github.com/yangyouwei/denyssh/readsecure"
	"github.com/yangyouwei/denyssh/hostsdeny"
)

var ip_list map[string]int
var LOGS chan string = make(chan string)

func main()  {
	readsecure.LOGS_DIR = "D:\\passwd.txt"
	record_ip()
	timer_number()
}

func record_ip() {
	go readsecure.READ_LOG(LOGS)
	for   {
		s := <- LOGS
		ip_addr := readsecure.GET_IP(s)
		if ip_addr == "" {
			continue
		}
		for a = range ip_list {
			if _,ok := ip_list; ok{
				//modify ip num+1
			}else {
				//append in to map num=1
			}
		}
	}
}

//for a long time and count number
func timer_number()  {
	if number = 5 {
		hostsdeny.Tracefile() //write to hosts.deny
	}
	return
}