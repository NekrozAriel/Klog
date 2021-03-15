package klog

import (
	"fmt"
	"log"
	"os"
	"time"
)

// logger 输出至目录下的.log文件，简单测试用。
var logger *log.Logger

// NewLog 写入init()内以使用。log文件会输出至当前目录下。
func NewKlog(fileName string) *log.Logger {
	fileTime := fmtTime()
	file, err := os.OpenFile(fmt.Sprint("./", fileName, fileTime, ".log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil
	}
	logger = log.New(file, "", log.LstdFlags|log.Llongfile)
	return logger
}

func Prtln(s string) {
	logger.Println(s)
}

func Prtf(format string, v ...interface{}) {
	logger.Printf(format, v...)
}

func Ftlln(s string) {
	logger.Fatalln(s)
}

func Ftlf(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}

func Pncln(s string) {
	logger.Panicln(s)
}

func Pncf(format string, v ...interface{}) {
	logger.Panicf(format, v...)
}

// fmtTime 格式化当前时间，获取标题需要的时间格式
func fmtTime() string {
	return fmt.Sprintf("%d%02d%02d%02d%02d%02d", time.Now().Year(),
		time.Now().Month(), time.Now().Day(), time.Now().Hour(),
		time.Now().Minute(), time.Now().Second())
}
