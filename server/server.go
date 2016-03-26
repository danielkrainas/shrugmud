package server

import (
	"container/list"
	"net"
	"strconv"

	"github.com/danielkrainas/shrugmud/config"
	"github.com/danielkrainas/shrugmud/logging"
	"github.com/danielkrainas/shrugmud/world"
)

type Server struct {
	Port        int
	Host        string
	Descriptors *list.List
}

func New(serverConfig *config.ServerConfig) *Server {
	server := &Server{
		Port:        serverConfig.Port,
		Host:        serverConfig.Host,
		Name:        serverConfig.Name,
		Descriptors: list.New(),
	}

	server.Descriptors.Init()
	return server
}

func (server *Server) Start(realm world.Realm) {
	logging.Trace.Println("Server.Start: enter")
	defer logging.Trace.Println("Server.Start: exit")

	addr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(server.Host, strconv.Itoa(server.Port)))
	if err != nil {
		logging.Error.Fatal(err)
		return
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		logging.Error.Fatal(err)
		return
	}

	logging.Info.Printf("%s started on port %d.", server.Host, server.Port)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			logging.Error.Fatal(err)
			return
		}

		d := newDescriptor(conn)
		d.Handle()
		server.Descriptors.PushFront(d)
		logging.Info.Printf("connection accepted: %s@%s", conn.RemoteAddr().Network(), conn.RemoteAddr().String())
	}
}
