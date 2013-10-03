package ipfilter

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
)

type IPFilter struct {
	ipnets []net.IPNet
	ips    []net.IP
}

func (f *IPFilter) FilterIP(ip net.IP) bool {
	for _, item := range f.ipnets {
		if item.Contains(ip) {
			return true
		}
	}
	for _, item := range f.ips {
		if item.Equal(ip) {
			return true
		}
	}
	return false
}

func (f *IPFilter) FilterIPString(s string) bool {
	ip := net.ParseIP(s)
	if ip == nil {
		return false
	}
	return f.FilterIP(ip)
}

func (f *IPFilter) AddIPNet(item net.IPNet) {
	f.ipnets = append(f.ipnets, item)
}

func (f *IPFilter) AddIPNetString(s string) error {
	_, ipnet, err := net.ParseCIDR(s)
	if err != nil {
		return err
	}
	f.AddIPNet(*ipnet)
	return nil
}

func (f *IPFilter) AddIP(ip net.IP) {
	f.ips = append(f.ips, ip)
}

func (f *IPFilter) AddIPString(s string) error {
	ip := net.ParseIP(s)
	if ip == nil {
		return errors.New("Parse IP Error: " + s)
	}
	f.AddIP(ip)
	return nil
}

func (f *IPFilter) Load(data []byte) error {
	for _, item := range bytes.Fields(data) {
		if bytes.IndexByte(item, '/') < 0 {
			if err := f.AddIPString(string(item)); err != nil {
				return err
			}
		} else {
			if err := f.AddIPNetString(string(item)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (f *IPFilter) LoadFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}
	return f.Load(data)
}

func (f *IPFilter) LoadHttp(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return f.Load(data)
}
