package main

import (
	"flag"
	"log"

	"github.com/sonicLee/learngo/testFsNotify/watch"
)

func main() {
	fileName := flag.String("conf", ".", "配置文件路径")
	flag.Parse()

	log.Println("定时任务载入开始")

	watch.WatchFile(*fileName, func() {
		println("hahaha")
	})

	select {
	default:
	}
	log.Println("定时任务载入完毕")
}
