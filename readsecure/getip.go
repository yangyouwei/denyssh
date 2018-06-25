package readsecure

import (
	"strings"
	"fmt"
)

func GET_IP(logs string)(s string) {
	a := strings.Split(logs," ")
	b := fmt.Sprint(a[4])
	//判断是否为sshd日志
	if strings.HasPrefix(b,"sshd") {
		//fmt.Println("sshd logs")
		sshd_logs := strings.Split(logs,"]: ")
		logs_content := fmt.Sprint(sshd_logs[1])
		//判断那种类型的失败，用户名或密码
		if strings.HasPrefix(logs_content,"Invalid user ") {
			//fmt.Println("invalid user")
			s = INVALID_USER_IP(logs_content)
			return s
		}else if strings.HasPrefix(logs_content,"Failed password for root from") {
			//fmt.Println("failed password")
			s = FAILED_PASSWORD_IP(logs_content)
			return s
		}
		return
	}
	return
}

//获取invalid user ip
func INVALID_USER_IP(logs string)(invalid_user_ip string)  {
	invalid_user_ip = fmt.Sprint(strings.Split(logs," ")[4])
	return invalid_user_ip
}
//获取failed password ip
func FAILED_PASSWORD_IP(logs string) (failed_password_ip string){
	failed_password_ip = fmt.Sprint(strings.Split(logs," ")[5])
	return failed_password_ip
}
