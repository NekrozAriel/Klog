package klog

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// logger 输出至目录下的.log文件，简单测试用。
var logger *log.Logger

// NewLog 写入init()内以使用。log文件会输出至控制台和当前目录下。
func NewKlog(fName string, t time.Time) {
	fTime := fmtTime(t)
	file, err := os.OpenFile(fmt.Sprint("./", fName, fTime, ".log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic("创建.log文件失败。")
	}
	logger = log.New(file, "", log.LstdFlags|log.Llongfile)
	mw := io.MultiWriter(os.Stdout, file)
	logger.SetOutput(mw)
}

func Prtln(v ...interface{}) {
	logger.Println(v...)
}

func Prtf(format string, v ...interface{}) {
	logger.Printf(format, v...)
}

func Ftlln(v ...interface{}) {
	logger.Fatalln(v...)
}

func Ftlf(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}

func Pncln(v ...interface{}) {
	logger.Panicln(v...)
}

func Pncf(format string, v ...interface{}) {
	logger.Panicf(format, v...)
}

// fmtTime 格式化当前时间，获取标题需要的时间格式
func fmtTime(t time.Time) string {
	return fmt.Sprintf("%d%02d%02d%02d%02d%02d", t.Year(),
		t.Month(), t.Day(), t.Hour(),
		t.Minute(), t.Second())
}
