package readpasswd

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"strings"
)

var USER_NAME []string

func GETUSER() {
	//逐行读取文本
	fi, err := os.Open("D:\\passwd.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		//判断后缀是否为可登陆用户
		if strings.HasSuffix(string(a), "/bin/bash"){
			user := strings.Split(string(a),":")  //取用户名
			USER_NAME = append(USER_NAME,user[0])
		}
	}
}