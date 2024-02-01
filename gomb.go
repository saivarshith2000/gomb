package gomb

import (
	"net"
)

type Gomb struct {
	adminServer *AdminServer
	brokerPort  int
}

func NewGomb(adminPort, brokerPort int, store *Store) *Gomb {
	adminServer := NewAdminServer(adminPort, store)
	return &Gomb{
		adminServer: adminServer,
		brokerPort:  brokerPort,
	}
}

func (g *Gomb) Start(errChan chan error) {
	// TODO: Start both admin server and message server
	g.adminServer.Start()
	// addr := fmt.Sprintf(":%d", b.brokerPort)
	// listener, err := net.Listen("tcp", addr)
	// if err != nil {
	// 	errChan <- err
	// 	return
	// }
	// for {
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		errChan <- err
	// 		return
	// 	}
	// 	b.handleClient(conn)
	// }
}

func (g *Gomb) handleClient(conn net.Conn) {

}
