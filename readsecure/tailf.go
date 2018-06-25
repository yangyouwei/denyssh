package readsecure

import (
	"fmt"
	"github.com/hpcloud/tail"
)


var LOGS_DIR string

func READ_LOG(logs chan  string) {
	t, _ := tail.TailFile(LOGS_DIR, tail.Config{Follow: true})
	for line := range t.Lines {
		logs <- fmt.Sprint(line.Text)
	}
}

