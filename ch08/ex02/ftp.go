package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", "localhost:21")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		server := &ftpServer{ctrlConn: conn, cd: "."}
		go server.start()
	}
}

type ftpServer struct {
	ctrlConn   net.Conn
	clientAddr string
	cd         string
}

type command struct {
	name  string
	param string
}

func (c command) String() string {
	return c.name + " " + c.param
}

func (s *ftpServer) start() {
	defer s.ctrlConn.Close()
	r := bufio.NewReader(s.ctrlConn)
	s.response(ServiceReady)
	for {
		cmd, err := readCommand(r)
		if err != nil {
			if err == io.EOF {
				log.Println("End.")
				return
			}
			log.Println(err)
			return
		}
		log.Printf("cmd: %v\n", cmd)
		switch cmd.name {
		case USER:
			s.response(LoggedIn)
		case QUIT:
			s.response(ServiceClosing)
			return
		case TYPE:
			if cmd.param == "A" {
				s.response(CommandOkay)
			} else {
				s.response(NotImplForParam)
			}
		case MODE:
			if cmd.param == "S" {
				s.response(CommandOkay)
			} else {
				s.response(NotImplForParam)
			}
		case PORT:
			s.port(cmd.param)
		case RETR:
			s.retrieve(cmd.param)
		case STOR:
			s.store(cmd.param)
		case NOOP:
			s.response(CommandOkay)
		default:
			s.response(NotImplemented)

		}
	}
}

func (s *ftpServer) response(code int) {
	fmt.Fprintf(s.ctrlConn, "%d\n", code)
}

func readCommand(r *bufio.Reader) (command, error) {
	line, _, err := r.ReadLine()
	if err != nil {
		return command{}, err
	}
	cmd := string(line)
	sp := strings.Index(cmd, " ")
	if sp < 0 {
		return command{cmd, ""}, nil
	}
	return command{strings.ToUpper(cmd[:sp]), cmd[sp+1:]}, nil
}

func (s *ftpServer) port(param string) {
	params := strings.Split(param, ",")
	if len(params) != 6 {
		s.response(SyntaxErrorInParam)
		return
	}
	ip := strings.Join(params[:4], ".")
	portUpper, err1 := strconv.Atoi(params[4])
	portLower, err2 := strconv.Atoi(params[5])
	if err1 != nil || err2 != nil {
		s.response(SyntaxErrorInParam)
	}
	port := portUpper*256 + portLower
	s.clientAddr = fmt.Sprintf("%s:%d", ip, port)
	s.response(CommandOkay)
}

func (s *ftpServer) retrieve(param string) {
	path := path.Join(s.cd, param)
	file, err := os.Open(path)
	if err != nil {
		s.response(FileNotFound)
		return
	}
	defer file.Close()
	s.response(OpenDataConn)
	conn, err := net.Dial("tcp", s.clientAddr)
	if err != nil {
		s.response(CantOpenDataConn)
		return
	}
	defer conn.Close()
	_, err = io.Copy(conn, file)
	if err != nil {
		s.response(TransferAborted)
		return
	}

	s.response(ActionCompleted)
}
func (s *ftpServer) store(param string) {
	path := path.Join(s.cd, param)

	s.response(OpenDataConn)
	conn, err := net.Dial("tcp", s.clientAddr)
	if err != nil {
		s.response(CantOpenDataConn)
		return
	}
	defer conn.Close()
	file, err := os.Create(path)
	if err != nil {
		s.response(FileUnavailable)
		return
	}
	defer file.Close()
	_, err = io.Copy(file, conn)
	if err != nil {
		s.response(TransferAborted)
		return
	}

	s.response(ActionCompleted)
}

const (
	// must be implemented for minimum implementation
	USER = "USER"
	QUIT = "QUIT"
	PORT = "PORT"
	TYPE = "TYPE"
	MODE = "MODE"
	RETR = "RETR"
	STOR = "STOR"
	NOOP = "NOOP"

	CWD  = "CWD"
	CDUP = "CDUP"
	LIST = "LIST"
)

const (
	OpenDataConn       = 150
	CommandOkay        = 200
	ServiceReady       = 220
	ServiceClosing     = 221
	LoggedIn           = 230
	ActionCompleted    = 250
	CantOpenDataConn   = 425
	FileUnavailable    = 450
	TransferAborted    = 426
	SyntaxErrorInParam = 501
	NotImplemented     = 502
	NotImplForParam    = 504
	FileNotFound       = 550
)
