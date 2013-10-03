package ipfilter

import (
	"testing"
)

func TestIPFilter(t *testing.T) {
	var f IPFilter
	if err := f.AddIPNetString("192.168.1.0/24"); err != nil {
		t.Errorf("Add: 192.168.1.0/24, %s", err.Error())
	}
	if err := f.AddIPString("192.168.100.3"); err != nil {
		t.Errorf("Add: 192.168.100.3, %s", err.Error())
	}
	if !f.FilterIPString("192.168.1.3") {
		t.Error("Filter Fail: 192.168.1.3")
	}
	if !f.FilterIPString("192.168.100.3") {
		t.Error("Filter Fail: 192.168.100.3")
	}
}

func TestLoad(t *testing.T) {
	var f IPFilter
	var data = []byte("192.168.1.0/24 192.168.1.3")
	if err := f.Load(data); err != nil {
		t.Error(err)
	}
	if !f.FilterIPString("192.168.1.3") {
		t.Error("Filter Fail: 192.168.1.3")
	}
}
