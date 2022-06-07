package main

import (
	"log"
	"os"
	"time"
	"fmt"
)

const MYFILE = "webtail.log"

func init() {
//        file := "./" +"webtail"+ ".log"
        logFile, err := os.OpenFile(MYFILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
        if err != nil {
                panic(err)
        }
        log.SetOutput(logFile) // 将文件设置为log输出的文件

//        log.SetPrefix("[webtail]")
//        log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
        return
}

func main() {

	for {
		log.Printf("starting main....")
		time.Sleep(time.Second)
		readFile(MYFILE)

                
	}
	//  log.Panicf("Oh, system error when hacker login")
	//  log.Fatalf("Danger! hacker has login")
}

func readFile(fname string) {
    logFile, err := os.Open(fname)
    if err != nil {
        panic(err)
    }
    defer logFile.Close()

    buf := make([]byte, 39)
    stat, err := os.Stat(MYFILE)
    start := stat.Size() - 39 
    _, err = logFile.ReadAt(buf, start)
    if err == nil {
        fmt.Printf("%s\n", buf)
    }


}
