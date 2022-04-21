package tcp

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	listener,err := net.Listen("tcp","127.0.0.1:8000")
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for  {
		conn,err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		//处理用户请求 新建协程
		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn) {
	defer conn.Close()
	//获取客户端网络信息
	addr := conn.RemoteAddr().String()
	fmt.Println(addr," connect successfully")
	fmt.Println(addr)

	buf := make([]byte,2048)
	for  {
		//读取数据
		n,err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf[:n]))

		if "exit"== string(buf[:n-1]){
			return
		}
		//数据转换为大写再发送给用户
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}
