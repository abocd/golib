package glog

import (
	"testing"
	//"fmt"
	//"strings"
	"os"
	"sync"
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
		ShowLevel:  "",
		SaveLevel:  LevelWarn,
		Flag:       0,
		MaxLogSize: 1000,
		NeedFlush:  true,
		GzLog:      true,
	})
	var wg sync.WaitGroup
	for i:=1;i<=50300;i++ {
		wg.Add(1)
		go func(i int) {
			gg.Info(i, "this is info")
			//gg.Error(i, "show error", strings.Repeat(" this is a test from golang", 1))
			gg.Error(i, "show error")
			gg.Warn(i, "this is warn")
			wg.Done()
		}(i)
		//time.Sleep(2*time.Second)
	}
	//time.Sleep(time.Second*2)
	wg.Wait()
	gg.Flush()
}