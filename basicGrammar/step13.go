package main

import (
	"net"
	"fmt"
	"errors"
)

func main() {
	isIp, err := privateIp("12.23.10.25")
	if err != nil {
		panic(err)
	}
	fmt.Println(isIp)
	ipLookUP()
}

func ipLookUP(){
	ip := net.ParseIP("127.0.0.1")
	ips, err := net.LookupIP("localhost")
	if err != nil {
		panic(err)
	}

	for _, v := range ips {
		if v.Equal(ip) {
			fmt.Println(v)
		}
	}
}

func privateIp(ip string)(bool, error){
	var err error
	private := false
	IP := net.ParseIP(ip)
	if IP == nil {
		err = errors.New("Invalid IP")
	}else {
		_, private24BitBlock, _ := net.ParseCIDR("10.0.0.0/8")
		_, private20BitBlock, _ := net.ParseCIDR("172.16.0.0/12")
		_, private16BitBlock, _ := net.ParseCIDR("192.168.0.0/16")
		private = private24BitBlock.Contains(IP) || private20BitBlock.Contains(IP) || private16BitBlock.Contains(IP)
	}

	return private, err
}
