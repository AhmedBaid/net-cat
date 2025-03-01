package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

	"chat-app/internal/helpers"
	"chat-app/internal/logger"
	"chat-app/internal/validators"
	"chat-app/utils"
)

func Server(conn net.Conn) {
	conn.Write([]byte(utils.Cyan + utils.WelcomeMessage + utils.Reset))
	var username string
	var err error

	for {
		username, err = bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			logger.ErrorLogger.Printf("❌ Error reading username || connection closed: %v", err)
			fmt.Println("❌ Error reading username || connection closed:", err)
			CloseConnection(conn, "")
			return
		}
		username = strings.TrimSpace(username)

		if !validators.ValidName(username) {
			conn.Write([]byte("🚫 Invalid username... Only printable characters allowed\n"))
			time.Sleep(1 * time.Second)
			conn.Write([]byte("Enter your name again: "))
			continue
		}
		if !validators.ValidateLength(username) {
			conn.Write([]byte("🚫 Invalid username... The name should be between 3 and 15 letters\n"))
			time.Sleep(1 * time.Second)
			conn.Write([]byte("Enter your name again: "))
			continue
		}
		if !validators.SameName(username) {
			conn.Write([]byte("🚫 Invalid username... The name already exists\n"))
			time.Sleep(1 * time.Second)
			conn.Write([]byte("Enter your name again: "))
			continue
		}

		break
	}

	utils.MU.Lock()
	utils.Clients[conn] = username
	utils.MU.Unlock()
	logger.InfoLogger.Printf("✅ %s has joined the chat...", username)
	helpers.Broadcasting(utils.Green+fmt.Sprintf("\n%s has join the chat...\n"+utils.Reset, username), conn)

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			logger.ErrorLogger.Printf("❌ Client disconnected || connection closed: %v", err)
			fmt.Println("❌ Client disconnected || connection closed:", err)
			CloseConnection(conn, username)
			return
		}
		message = strings.TrimSpace(message)
		if !validators.ValidMessage(message) || !validators.ValidateLengthMessage(message) {
			time.Sleep(1 * time.Second)
			conn.Write([]byte(fmt.Sprintf("[%s] [%s]:", time.Now().Format("15:04:05"), username)))
			continue
		}
		helpers.Broadcasting(fmt.Sprintf("\n[%s] [%s]: %s\n", time.Now().Format("15:04:05"), username, message), conn)
	}
}

func CloseConnection(conn net.Conn, username string) {
	utils.MU.Lock()
	if username != "" {
		logger.InfoLogger.Printf("🚪 %s has left the chat...", username)
		helpers.Broadcasting(utils.Red+fmt.Sprintf("\n%s has left the chat...\n"+utils.Reset, username), conn)
		delete(utils.Clients, conn)
	}
	utils.Counter--
	utils.MU.Unlock()
	conn.Close()
}
