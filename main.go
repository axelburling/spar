package main

//server

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	f "github.com/axelburling/spar/formater"
	g "github.com/axelburling/spar/gen"
	h "github.com/axelburling/spar/http"
	p "github.com/axelburling/spar/parser"
	personnummer "github.com/personnummer/go/v3"
)

func Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello\n"))
	})
}

func main() {
	log.SetFlags(log.Lshortfile)

	cer, err := tls.LoadX509KeyPair("./tls/server.crt", "./tls/server.key")
	if err != nil {
		log.Println(err)
		return
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}, ClientAuth: tls.RequestClientCert, InsecureSkipVerify: true}
	ln, err := tls.Listen("tcp", "localhost:443", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()

	// http.ServeTLS(ln, Handler(), "./tls/server.crt", "./tls/server.key")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	err := conn.(*tls.Conn).Handshake()

	if err != nil {
		log.Println(err)
		return
	}

	state := conn.(*tls.Conn).ConnectionState()
	for i, cert := range state.PeerCertificates {
		subject := cert.Subject
		issuer := cert.Issuer
		log.Printf(" %d s:/C=%v/ST=%v/L=%v/O=%v/OU=%v/CN=%s", i, subject.Country, subject.Province, subject.Locality, subject.Organization, subject.OrganizationalUnit, subject.CommonName)
		log.Printf("   i:/C=%v/ST=%v/L=%v/O=%v/OU=%v/CN=%s", issuer.Country, issuer.Province, issuer.Locality, issuer.Organization, issuer.OrganizationalUnit, issuer.CommonName)
	}

	buffer := make([]byte, 2048)

	reqLine, err := conn.Read(buffer)

	if reqLine == 0 {
		log.Println("Connection closed")
		return
	}

	if err != nil {
		log.Println(err)
		return
	}

	n := bytes.Index(buffer, []byte{0})
	message := string(buffer[:n+1])

	em := 0
	method := ""

	for i, line := range strings.Split(message, "\n") {

		if i == 0 {
			method = strings.Split(line, " ")[0]
		}

		l := strings.TrimSpace(line)

		if l == "" {
			em = i
			break
		}
	}

	if method != "POST" {

		_, err := h.Resp(conn, 405, "text/plain; charset=utf-8", "Method not allowed")
		if err != nil {
			h.Error(conn, 500, "Internal server error")
			return
		}
		return
	}

	xml := ""

	// loop over message and print each line
	for i, line := range strings.Split(message, "\n") {
		if i > em {
			xml += line
		}
	}

	req, err := p.Parse([]byte(xml))

	if err != nil {
		h.Error(conn, 400, "Bad request")
		return
	}

	pin, err := personnummer.New(req.PersonId, &personnummer.Options{})

	if err != nil {
		h.Error(conn, 400, "Bad request")
		return
	}

	pe, err := g.Gen(pin)

	if err != nil {
		h.Error(conn, 400, "Bad request")
		return
	}

	res := &f.SEnvelope{
		SBody: f.SBody{
			Ns17SPARPersonsokningSvar: f.Ns17SPARPersonsokningSvar{
				Ns17PersonsokningSvarsPost: *pe,
			},
		},
	}

	bResp, err := f.FormatResponse(*res)

	if err != nil {
		h.Error(conn, 400, "Bad request")
		return
	}

	fmt.Println(string(bResp))

	_, err = h.Resp(conn, 200, "text/xml", string(bResp))

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("Sent %d bytes\n", n)

}
