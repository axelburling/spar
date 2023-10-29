package http

import (
	"fmt"
	"net"
)

func Resp(conn net.Conn, status int, contentType string, body string) (int, error) {
	str := fmt.Sprintf("HTTP/1.1 %d OK\r\nContent-Type: %s\r\nContent-Length: %d\r\n\r\n%s", status, contentType, len(body), body)

	return conn.Write([]byte(str))
}
