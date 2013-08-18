package ipfilter

import (
	"testing"
)

func TestIPFilter(t *testing.T) {
	var f IPFilter
	if err := f.AddIPNetString("192.168.1.0/24"); err != nil {
		t.Fatalf("Add: 192.168.1.0/24, %s", err.Error())
	}
	if err := f.AddIPString("192.168.100.3"); err != nil {
		t.Fatalf("Add: 192.168.100.3, %s", err.Error())
	}
	if !f.FilterIPString("192.168.1.3") {
		t.Fatal("Filter Failed: 192.168.1.3")
	}
	if !f.FilterIPString("192.168.100.3") {
		t.Fatal("Filter Failed: 192.168.100.3")
	}
}
