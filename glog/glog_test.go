package glog

import "testing"

func TestAsset(t *testing.T) {
	Asset("断言","%s","很好")
	Asset("断言2","这里")
}

func TestError(t *testing.T) {
	Error("错了","%s","对的")
}