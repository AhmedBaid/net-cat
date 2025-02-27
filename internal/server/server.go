package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

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
		fmt.Println("eded", username)
		if err != nil {
			fmt.Println("Error reading username:", err)
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
	Time := time.Now().Format("2006-01-02 15:04:05")

	utils.MU.Lock()
	utils.Clients[conn] = username
	utils.MU.Unlock()

	helpers.Broadcasting(fmt.Sprintf("\n%s has joined the chat... 👋\n", username), conn)
	helpers.Broadcasting(fmt.Sprintf("[%s] [%s]: ", Time, username), conn)

	for {
		conn.Write([]byte(fmt.Sprintf("[%s] [%s]: ", Time, username)))

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected:", err)
			break
		}

		message = strings.TrimSpace(message)

		helpers.Broadcasting(fmt.Sprintf("\n[%s] [%s]: %s\n", Time, username, message), conn)
		helpers.Broadcasting(fmt.Sprintf("[%s] [%s]: ", Time, username), conn)
	}

	utils.MU.Lock()
	delete(utils.Clients, conn)
	utils.MU.Unlock()
	helpers.Broadcasting(fmt.Sprintf("\n%s has left the chat... 🚪", username), conn)
	conn.Close()
}
