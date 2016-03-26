package server

import (
	"crypto/md5"
	"fmt"
	"net"
	"time"
)

type Descriptor struct {
	uuid  string
	conn  *net.TCPConn
	State DescState
}

type DescState uint16

const (
	StateClose          = 0
	StateWelcome        = 1
	StatePromptName     = 2
	StateConfirmName    = 3
	StatePromptPassword = 4
	StateMotd           = 5
	StatePlaying        = 6
)

func createUuid(conn *net.TCPConn) string {
	data := fmt.Sprintf("%s@%s", conn.RemoteAddr().String(), time.Now().String())
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

func newDescriptor(conn *net.TCPConn) *Descriptor {
	return &Descriptor{
		uuid:  createUuid(conn),
		conn:  conn,
		State: StateWelcome,
	}
}

func (d *Descriptor) Write()

func (d *Descriptor) HandleOut() {
	for d.State != StateClose {
		msg := <- d.
	}

	d.dispose()
}

func (d *Descriptor) HandleCommands() {

}

func (d *Descriptor) Handle(toServer) {

}

func (d *Descriptor) Close() {
	d.State = StateClose
}

func (d *Descriptor) dispose() {
	d.conn.Close()
}
