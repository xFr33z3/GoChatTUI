package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

type BMessage struct {
	sender      net.Conn
	text        string
	messageType string
}

var (
	clients   = make(map[net.Conn]bool)
	nicknames = make(map[net.Conn]string)
	broadcast = make(chan BMessage)

	ServerIP string = "0.0.0.0:8000"
)

var mu sync.Mutex

func sendMessage(conn net.Conn, msg string) {
	_, err := conn.Write([]byte(msg + "\n"))
	if err != nil {
		log.Println("Error sending to client:", err)
		conn.Close()
		mu.Lock()
		delete(clients, conn)
		delete(nicknames, conn)
		mu.Unlock()
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		delete(clients, conn)
	}()

	clients[conn] = true
	nicknames[conn] = conn.RemoteAddr().String()
	reader := bufio.NewReader(conn)

	file, err := os.Open("motd.txt")
	if err != nil {
	} else {
		sendMessage(conn, "")
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			sendMessage(conn, line)
		}
	}

	for {
		data, err := reader.ReadString('\n')
		msg := strings.TrimSpace(data)

		if err != nil {
			log.Println("Client '"+nicknames[conn]+"' disconnected:", err)
			broadcast <- BMessage{conn, "[red]Client '" + nicknames[conn] + "' disconnected:[white]", "user"}
			return
		}
		if strings.HasPrefix(msg, "/") {
			args := strings.Split(msg, " ")
			command := args[0]
			if command == "/nick" {
				if len(args) <= 1 {
					sendMessage(conn, "[yellow]SERVER >>> Not enough arguments"+"\n")
				} else {
					sendMessage(conn, "[yellow]SERVER >>> Username changed to "+args[1]+"\n")
					mu.Lock()
					nicknames[conn] = strings.TrimSpace(args[1])
					mu.Unlock()
				}
			}
		} else {
			// Process and broadcast message to the group chat
			broadcast <- BMessage{conn, msg + "[white]", "user"}
		}
	}
}

func broadcaster() {
	for {
		message := <-broadcast
		mu.Lock()
		senderName := nicknames[message.sender]
		clientsCopy := make([]net.Conn, 0, len(clients))
		for c := range clients {
			clientsCopy = append(clientsCopy, c)
		}
		mu.Unlock()
		for _, conn := range clientsCopy {
			_, err := conn.Write([]byte("[white]" + senderName + ": " + message.text + "\n"))
			if err != nil {
				log.Println("Error sending to client:", err)
				conn.Close()
				delete(clients, conn)
			}
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ServerIP)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	fmt.Println("Server listening on port 8000")

	go broadcaster()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
