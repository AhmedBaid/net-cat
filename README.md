# Net-Cat

A simple TCP-based chat application built using Golang. This project allows multiple clients to connect to a server and communicate in real time.

## Features
- Multi-client support
- Usernames for each client
- Welcome message broadcasted to all users when a new user joins
- Message timestamps
- Proper handling of client disconnections
- Server logging system

## Installation & Setup

### Prerequisites
- Go (latest version recommended)

### Clone the Repository
```sh
git clone https://github.com/yourusername/net-cat.git
cd net-cat
```

### Run the Server
```sh
go run TCPchat/main.go
```

## Project Structure
```
NET-CAT/
├── internal/
│   ├── helpers/
│   │   ├── Broadcating.go
│   ├── logger/
│   │   ├── logger.go
│   ├── server/
│   │   ├── server.go
│   ├── validators/
│   │   ├── validators.go
├── TCPchat/
│   ├── main.go
├── utils/
│   ├── utils.go
├── .gitignore
├── go.mod
├── network.md
├── README.md
```

## Usage
1. Start the server.
2. Start multiple clients.
3. Enter a username when prompted.
4. Start chatting!

## Example Output
```
Welcome to the chat server!
Enter your username: Alice
Welcome 👋 Alice
[2025-02-26 10:30:00] [Alice]: Hello everyone!
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
This project is open-source and available under the MIT License.

