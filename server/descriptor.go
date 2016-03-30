package server

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"net"
	"time"

	"github.com/danielkrainas/shrugmud/logging"
)

type Descriptor struct {
	uuid         string
	conn         *net.TCPConn
	outputBuffer chan string
	CtrlState    CtrlState
	State        DescState
}

type DescState uint16

const (
	StateUnknown = DescState(0)
	StateWelcome = DescState(1)
	StateLogin   = DescState(2)
	StatePlaying = DescState(3)
	StateClose   = DescState(4)
)

func canonicalToState(canonical string) DescState {
	switch canonical {
	case "welcome":
		return StateWelcome
	case "login":
		return StateLogin
	case "playing":
		return StatePlaying
	case "close":
		return StateClose
	case "unknown":
		return StateUnknown
	}

	return StateUnknown
}

func stateToCanonical(state DescState) string {
	switch state {
	case StateWelcome:
		return "welcome"
	case StateLogin:
		return "login"
	case StatePlaying:
		return "playing"
	case StateClose:
		return "close"
	case StateUnknown:
		return "unknown"
	}

	return "unknown"
}

func createUuid(conn *net.TCPConn) string {
	data := fmt.Sprintf("%s@%s", conn.RemoteAddr().String(), time.Now().String())
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

func newDescriptor(conn *net.TCPConn) *Descriptor {
	return &Descriptor{
		uuid:         createUuid(conn),
		conn:         conn,
		outputBuffer: make(chan string),
	}
}

func (d *Descriptor) Write(v ...interface{}) error {
	_, err := d.conn.Write([]byte(fmt.Sprint(v...)))
	if err != nil {
		return err
	}

	return nil
}

func (d *Descriptor) Writef(format string, v ...interface{}) error {
	return d.Write(fmt.Sprintf(format, v...))
}

func (d *Descriptor) HandleOut() {
	for d.State != StateClose {
		msg := <-d.outputBuffer
		d.Write(msg)
	}

	d.dispose()
}

func (d *Descriptor) HandleInput(router *CtrlRouter) {
	reader := bufio.NewReader(d.conn)
	for d.State != StateClose {
		line, _, err := reader.ReadLine()
		if err != nil {
			logging.Error.Fatal(err)
			continue
		}

		err = router.Dispatch(string(line), d)
		if err != nil {
			logging.Error.Fatal(err)
		}
	}
}

func (d *Descriptor) Handle(router *CtrlRouter) {
	go d.HandleOut()
	go d.HandleInput(router)
}

func (d *Descriptor) Close() {
	d.State = StateClose
}

func (d *Descriptor) dispose() {
	d.conn.Close()
}
