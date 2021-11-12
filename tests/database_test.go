package main

import "testing"

func TestConnectToDB(t *testing.T) {
	if configPathPtr == nil || *configPathPtr == "" {
		t.Log("config flag has to be set! usage: --config=/path/to/config")
		t.Fail()
		return
	}
}
