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
	router      *CtrlRouter
}

func New(serverConfig *config.ServerConfig, router *CtrlRouter) *Server {
	server := &Server{
		Port:        serverConfig.Port,
		Host:        serverConfig.Host,
		Descriptors: list.New(),
		router:      router,
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

	logging.Info.Printf("%s started on port %d.", realm.Name(), server.Port)
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			logging.Error.Fatal(err)
			return
		}

		d := newDescriptor(conn)
		d.Handle(server.router)
		server.Descriptors.PushFront(d)
		logging.Info.Printf("connection accepted: %s@%s", conn.RemoteAddr().Network(), conn.RemoteAddr().String())
	}
}
