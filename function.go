package main

import (
	"fmt"
	"log"
	"os"
	. "strings"
	"time"
)

func fuck(e error) {
	if e != nil {
		log.Println(fmt.Sprintf("fucking: %v", e))
		sleep(3)
	}
}

func fatal(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func sleep(n int) {
	time.Sleep(time.Second * time.Duration(n))
}

func trace(msg string) func() {
	start := time.Now()
	//log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func mkdirs(s string) {
	err := os.MkdirAll(s, 0777)
	if err != nil {
		fatal(err)
	} else {
		fmt.Printf("Create Directory OK!")
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func IsMp3(url string) bool {
	exts := []string{"mp3"}
	for _, ext := range exts {
		if Contains(ToLower(url), "."+ext) {
			return true
		}
	}
	return false
}

func print(s interface{}) {
	fmt.Printf("%v\n", s)
}
