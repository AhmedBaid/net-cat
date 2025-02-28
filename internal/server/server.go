package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"chat-app/internal/helpers"
	"chat-app/internal/validators"
	"chat-app/utils"
)

func Server(conn net.Conn) {
	conn.Write([]byte(utils.WelcomeMessage))
	var username string
	var err error

	for {
		username, err = bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading username:", err)
			CloseConnection(conn, "")
			return
		}
		username = strings.TrimSpace(username)

		if !validators.Valid(username) {
			conn.Write([]byte("Only letters are allowed in the username.\n"))
			conn.Write([]byte("Enter your name: "))
			continue
		}

		if !validators.ValidateLength(username) {
			conn.Write([]byte("Username must be between 3 and 15 characters.\n"))
			conn.Write([]byte("Enter your name: "))
			continue
		}

		break
	}

	utils.MU.Lock()
	utils.Clients[conn] = username
	utils.MU.Unlock()

	helpers.Broadcasting(fmt.Sprintf("\n%s has joined the chat... 👋\n", username), conn)

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected:", err)
			CloseConnection(conn, username)
			return
		}

		message = strings.TrimSpace(message)
		helpers.Broadcasting(fmt.Sprintf("\n[%s] [%s]: %s\n", utils.Time, username, message), conn)
	}
}

func CloseConnection(conn net.Conn, username string) {
	utils.MU.Lock()
	if username != "" {
		helpers.Broadcasting(fmt.Sprintf("\n%s has left the chat... 🚪\n", username), conn)
		delete(utils.Clients, conn)
	}
	utils.Counter--
	fmt.Println("final:", utils.Counter)
	utils.MU.Unlock()

	conn.Close()
}
