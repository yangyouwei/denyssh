package hostsdeny

import "os"

var CONTENT string
var FILE_DIR string = "/var/log/secure"
var sshd string = "sshd:"

func Tracefile(str_content string)  {
	fd,_:=os.OpenFile(FILE_DIR,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	buf:=[]byte(sshd+str_content+"\n")
	fd.Write(buf)
	fd.Close()
}
