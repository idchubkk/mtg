package proxy

import (
	"io"
	"log"
	"net"
	"my-tg-proxy/config"
)

type Proxy struct {
	cfg *config.Config
}

func NewProxy(cfg *config.Config) *Proxy {
	return &Proxy{cfg: cfg}
}

func (p *Proxy) Start() error {
	listener, err := net.Listen("tcp", p.cfg.BindTo)
	if err != nil {
		return err
	}
	log.Println("MTProto Proxy listening on", p.cfg.BindTo)

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}

		go p.handleClient(clientConn)
	}
}

func (p *Proxy) handleClient(clientConn net.Conn) {
	defer clientConn.Close()

	// 这里我们直接转发到 Telegram 官方节点
	tgServer := "149.154.167.50:443"
	serverConn, err := net.Dial("tcp", tgServer)
	if err != nil {
		log.Println("Failed to connect to Telegram:", err)
		return
	}
	defer serverConn.Close()

	// 启动双向转发
	go io.Copy(serverConn, clientConn)
	io.Copy(clientConn, serverConn)
}
