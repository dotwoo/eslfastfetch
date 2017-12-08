package main

import (
	"fmt"
	"os"
	. "strings"
	"time"

	"github.com/leesper/holmes"
)

func sleep(n int) {
	time.Sleep(time.Second * time.Duration(n))
}

func trace(msg string) func() {
	start := time.Now()
	//log.Printf("enter %s", msg)
	return func() {
		holmes.Debugln("exit ", msg, time.Since(start))
	}
}

func mkdirs(s string) {
	err := os.MkdirAll(s, 0777)
	if err != nil {
		holmes.Fatalln(err.Error())
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

func IsHtml(url string) bool {
	exts := []string{"htm", "html"}
	for _, ext := range exts {
		if Contains(ToLower(url), "."+ext) {
			return true
		}
	}
	return false
}
