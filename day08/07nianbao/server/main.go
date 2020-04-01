package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	proto "github.com/Altair-cyx/day08/07nianbao/protocol"
)

// socket_stick/server/main.go

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	// var buf [1024]byte
	for {
		// n, err := reader.Read(buf[:])
		// if err == io.EOF {
		// 	break
		// }
		// if err != nil {
		// 	fmt.Println("read from client failed, err:", err)
		// 	break
		// }
		// recvStr := string(buf[:n])
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode failed,err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", msg)
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
