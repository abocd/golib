package glog

import (
	"testing"
	//"fmt"
	//"strings"
	"time"
	"os"
)

func TestAsset(t *testing.T) {
	Asset("%s","it is good")
	Asset("this",LongFile ,ShortFile)
	Info("bool",0&(0|1))
}

func TestError(t *testing.T) {
	Error("error","error")
}

func TestAll(t *testing.T){
	//fmt.Println("test all")
	os.Remove("../temp/1.log")
	gg := NewGLogFile("../temp/1.log",&Glog{
		ShowLevel: LevelError,
		SaveLevel: LevelWarn,
		Flag:0,
		MaxLogSize:10000000,
	})
	for i:=1;i<=1000;i++ {
		go func(i int) {
			gg.Info(i, "this is info")
			//gg.Error(i, "show error", strings.Repeat(" this is a test from golang", 1))
			gg.Error(i, "show error")
			gg.Warn(i, "this is warn")
		}(i)
	}
	time.Sleep(time.Second*2)
	gg.Flush()
}