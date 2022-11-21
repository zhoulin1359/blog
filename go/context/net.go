package context

import (
	"context"
	"fmt"
	"net"
	"time"
)

func Client(ctx context.Context) error {

	conn, err := net.Dial("tcp", "127.0.0.1:9001")
	if err != nil {
		// 连接的时候出现错误
		return err
	}
	// 当函数返回的时候关闭连接
	defer func() {
		_ = conn.Close()
	}()
	i := 0
	for {
		i++
		_, err = conn.Write([]byte("ping:" + time.Now().String())) // 发送数据
		if err != nil {
			return err
		}
		buf := [512]byte{}
		// 读取服务端发送的数据
		n, err := conn.Read(buf[:])
		if err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("第%d次,客户端接收服务端发送的数据:%s", i, string(buf[:n])))
		time.Sleep(time.Second)
	}
}
