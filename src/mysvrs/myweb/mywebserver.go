package myweb

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/nybuxtsui/log"
	"mysvrs/util"
)

var startTime = time.Now()
var visitcount int64

// 主处理：根据FORM表单信息关键字method判断业务功能
func Start(httpaddr string, httpsaddr string) {
	http.HandleFunc("/mysvrs", mysvrshandler)
	http.HandleFunc("/ver", verhandler)
	http.HandleFunc("/now", nowhandler)

	err := http.ListenAndServe(httpaddr, nil)
	if err != nil {
		log.Error("Start err: %s", err)
	}
}

func mysvrshandler(w http.ResponseWriter, r *http.Request) {
}

func verhandler(w http.ResponseWriter, r *http.Request) {
	var wel string = fmt.Sprintf("Welcom to mysvrs, you are the %dth visitor.", visitcount)
	fmt.Fprintln(w, wel)
	fmt.Fprintln(w, util.GetVersion())
	fmt.Fprintln(w, startTime)
}

func nowhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, time.Now().Format("2006-01-02 15:04:05"))
}

func refreshVisitCount() {
	for {
		time.Sleep(10 * time.Second)
	}
}

func Init() {
	visitfile, err := os.Open("visitcount.conf")
	if err != nil {
		log.Error("Init error: %s", err)
	}
	var buf []byte = make([]byte, 10)
	length, _ := visitfile.Read(buf)
	visitfile.Close()
	log.Debug("visitcount: %v, %s, %d", buf, string(buf[:length-1]), length)
	go refreshVisitCount()
}
