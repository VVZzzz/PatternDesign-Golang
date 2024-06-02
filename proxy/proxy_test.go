package proxy

import "testing"

func TestProxy(t *testing.T) {
	var realSubj Subject
	realSubj = &Proxy{}
	ret := realSubj.Do()

	if ret != "proxy start real proxy end" {
		t.Fatal("TestProxy fail")
	}

}
