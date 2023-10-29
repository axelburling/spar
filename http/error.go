package http

import (
	"fmt"
	"net"
)

func Error(conn net.Conn, status int, errorMessage string) (int, error) {
	n, err := Resp(conn, status, "text/plain", fmt.Sprintf("%d %s", status, errorMessage))

	if err != nil {
		Error(conn, 500, "Internal server error")
	}

	return n, err
}
