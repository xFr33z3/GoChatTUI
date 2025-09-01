# GoChatTUI

A real-time terminal-based chat application written in Go, featuring a clean TUI (Terminal User Interface) and server-client architecture.

## 🚀 Features

- **Multi-client support**: Multiple users can connect and chat simultaneously
- **Real-time messaging**: Instant message broadcasting to all connected clients
- **Terminal User Interface**: Clean, scrollable chat interface built with [tview](https://github.com/rivo/tview)
- **Nickname system**: Users can set custom nicknames with `/nick <name>` command
- **Message of the Day (MOTD)**: Welcome message displayed when clients connect
- **Keyboard navigation**: Use arrow keys to scroll through chat history
- **TCP-based networking**: Reliable connection handling with automatic cleanup

## 📁 Project Structure

```
GoChatTUI/
├── client/           # TUI chat client
│   ├── client.go     # Main client application
│   ├── go.mod        # Client dependencies
│   └── go.sum        # Dependency checksums
├── server/           # Chat server
│   ├── server.go     # TCP server implementation
│   ├── go.mod        # Server dependencies
│   └── motd.txt      # Message of the Day file
└── README.md         # This file
```

## 🛠️ Prerequisites

- Go 1.24.3 or later
- Terminal with support for ANSI colors (recommended)

## 📦 Installation

1. Clone the repository:
```bash
git clone https://github.com/xFr33z3/GoChatTUI.git
cd GoChatTUI
```

2. Install server dependencies:
```bash
cd server
go mod tidy
```

3. Install client dependencies:
```bash
cd ../client
go mod tidy
```

## 🚀 Usage

### Starting the Server

1. Navigate to the server directory:
```bash
cd server
```

2. Run the server:
```bash
go run server.go
```

The server will start listening on `localhost:8000` by default.

### Connecting with the Client

1. Open a new terminal and navigate to the client directory:
```bash
cd client
```

2. Run the client:
```bash
go run client.go
```

3. The client will automatically connect to the server and display the MOTD.

### Chat Commands

- `/nick <username>`: Set your nickname
- **Arrow Keys (↑/↓)**: Scroll through chat history
- **Enter**: Send message
- **Ctrl+C**: Disconnect and exit

## 🎮 Example Usage

1. Start the server in one terminal
2. Start multiple client instances in separate terminals
3. Set nicknames: `/nick Alice`, `/nick Bob`
4. Start chatting! Messages will be broadcast to all connected clients

## 🏗️ Architecture

### Server (`server.go`)
- Handles multiple concurrent TCP connections
- Manages client nicknames and connection state
- Broadcasts messages to all connected clients
- Implements graceful connection cleanup
- Serves MOTD from `motd.txt` file

### Client (`client.go`)
- TUI interface built with `tview` library
- Separate goroutines for sending and receiving messages
- Real-time message display with scroll support
- Input handling for commands and messages

## 🔧 Dependencies

### Server
- Go standard library only

### Client
- [tview](https://github.com/rivo/tview) - Terminal UI library
- [tcell](https://github.com/gdamore/tcell) - Low-level terminal interface

## 🐛 Troubleshooting

### Common Issues

**Connection refused**: Make sure the server is running before starting the client.

**Port already in use**: If port 8000 is busy, modify the port in both server and client code.

**Display issues**: Ensure your terminal supports ANSI colors and has sufficient size.

### Debug Mode

To enable debug logging, modify the log level in the source code or check terminal output for error messages.
