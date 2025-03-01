package main

import (
	"fmt"
	"net"
	"os"

	"chat-app/internal/logger"
	"chat-app/internal/server"
	"chat-app/utils"
)

func main() {
	logFile, err := logger.Logger()
	if err != nil {
		fmt.Println("Error setting up log file:", err)
		return
	}
	defer logFile.Close()
	port := ":8989"
	if len(os.Args) > 2 {
		logger.ErrorLogger.Println("[USAGE]: ./TCPChat $port")
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	if len(os.Args) == 2 {
		port = ":" + os.Args[1]
	}

	listner, err := net.Listen("tcp", port)
	if err != nil {
		logger.ErrorLogger.Println("error in listening  : ", err)
		fmt.Println("error in listening  : ", err)
		return
	}
	defer listner.Close()
	logger.InfoLogger.Printf("Server runing at port %s\n", port)
	fmt.Printf("Server runing at port %s\n", port)
	for {
		conn, err := listner.Accept()
		if err != nil {
			logger.ErrorLogger.Println("error in accepting  : ", err)
			fmt.Println("error in accepting  : ", err)
			return
		}
		utils.MU.Lock()
		if utils.Counter > 10 {
			logger.InfoLogger.Println("Chat is full. Try later...")
			conn.Write([]byte("Chat is full. Try later...\n"))
			conn.Close()
			utils.MU.Unlock()
			continue
		}
		utils.Counter++
		logger.InfoLogger.Println("new connection accepted", conn.LocalAddr())

		utils.MU.Unlock()

		go server.Server(conn)
	}
}
