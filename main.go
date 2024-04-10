package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

// wikipedia.org is blocked sni
var host = "wikipedia.org"

//var host = "test.org"

type MC struct {
	net.Conn
}

func (c *MC) Write(b []byte) (n int, err error) {
	bn := len(b)
	fmt.Println(bn)
	var wn, wp int
	var sn = 77
	for i := 0; i < bn; i += sn {
		wp = i + sn
		if wp > bn {
			wp = bn
		}
		fmt.Println(string(b[i:wp]))
		wn, err = c.Conn.Write(b[i:wp])
		time.Sleep(time.Millisecond * 500)
		n += wn
		if err != nil {
			return
		}
	}
	return
}

func main() {
	c, err := net.DialTimeout("tcp", "185.15.59.224:443", time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.(*net.TCPConn).SetNoDelay(true)
	tc := tls.Client(&MC{c}, &tls.Config{InsecureSkipVerify: true, ServerName: host})
	err = tc.Handshake()
	if err != nil {
		fmt.Println("Handshake:", err)
		return
	}
	fmt.Println("Handshake ok")
}
