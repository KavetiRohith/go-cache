package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/KavetiRohith/go-cache/cache"
)

type ServerOpts struct {
	ListenAddr string
}

type Server struct {
	ServerOpts
	cache *cache.Cache
}

func NewServer(opts ServerOpts, c *cache.Cache) *Server {
	return &Server{
		ServerOpts: opts,
		cache:      c,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return fmt.Errorf("listener error: %s", err)
	}

	log.Printf("server starting on [%s]\n", s.ListenAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error %s\n", err)
			continue
		}

		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer func() {
		conn.Close()
		log.Printf("connection closed: %s\n", conn.RemoteAddr())
	}()

	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		s.parseCommand(conn, scanner.Bytes())
	}
}

func (s *Server) parseCommand(conn net.Conn, rawCmd []byte) {
	var (
		parts   = strings.Fields(string(rawCmd))
		len_cmd = len(parts)
	)

	if len_cmd < 2 {
		conn.Write([]byte("message must atleast have command and key\n"))
		return
	}

	var (
		cmd = parts[0]
		key = parts[1]
	)

	switch cmd {
	case "SET":
		switch len_cmd {
		case 3:
			val := parts[2]
			s.handleSet(conn, key, val)
		case 4:
			val := parts[2]
			ttl := parts[3]
			s.handleSetWithTTL(conn, key, val, ttl)
		default:
			conn.Write([]byte("SET message must atleast have key and value\n"))
		}
	case "GET":
		s.handleGet(conn, key)
	case "DEL":
		s.handleDel(conn, key)
	case "HAS":
		s.handleHas(conn, key)
	}
}

func (s *Server) handleSet(conn net.Conn, key string, val string) {
	err := s.cache.Set(key, val)
	if err != nil {
		conn.Write([]byte(err.Error()))
	}

	conn.Write([]byte("Success\n"))
	log.Printf("SET %s %s\n", key, val)
}

func (s *Server) handleSetWithTTL(conn net.Conn, key string, val string, ttl string) {
	parsedTTL, err := strconv.Atoi(ttl)
	if err != nil {
		conn.Write([]byte("Invalid TTl\n"))
	}
	err = s.cache.SetWithTTL(key, val, time.Duration(parsedTTL)*time.Second)
	if err != nil {
		conn.Write([]byte(err.Error()))
	}

	conn.Write([]byte("Success\n"))
	log.Printf("SET %s %s exp: %v seconds\n", key, val, parsedTTL)
}

func (s *Server) handleGet(conn net.Conn, key string) {
	val, err := s.cache.Get(key)
	if err != nil {
		conn.Write([]byte(""))
		return
	}

	conn.Write([]byte(fmt.Sprintf("%s\n", val)))
	log.Printf("GET %s %s\n", key, val)
}

func (s *Server) handleDel(conn net.Conn, key string) {
	err := s.cache.Delete(key)
	if err != nil {
		conn.Write([]byte(err.Error()))
	}

	conn.Write([]byte("Success\n"))
	log.Printf("DEL %s\n", key)
}

func (s *Server) handleHas(conn net.Conn, key string) {
	isPresent := s.cache.Has(key)
	if isPresent {
		conn.Write([]byte("Yes\n"))
	} else {
		conn.Write([]byte("No\n"))
	}
	log.Printf("HAS %s %v\n", key, isPresent)
}
