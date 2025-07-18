# Go HTTP Server

A simple HTTP/1.1 server implemented in Go, capable of serving multiple clients concurrently. This project is a learning exercise to understand the fundamentals of HTTP, TCP servers, and request handling in Go.

## Features

- Handles multiple client connections concurrently
- Parses and responds to basic HTTP/1.1 requests
- Serves static files or sample responses
- Minimal dependencies, easy to understand and extend

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.20 or higher installed on your system

### Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/go-http-server.git
   cd go-http-server
   ```

2. **Build and run the server:**
   ```sh
   ./start.sh #If you have MacOS/Linux
   ./win_start.sh #If you have Windows
   ```

## Usage

By default, the server listens on a specified port (see `main.go` for configuration). You can test it by opening your browser and navigating to `http://localhost:4122` (or the port you configured).

You can also use `curl`:
```sh
curl http://localhost:4122/
```

## Project Structure

```
go-http-server/
  ├── app/
  │   ├── main.go             # Main server implementation
  ├── go.mod                  # Go module file
  ├── README.md               # Project documentation
  ├── start.sh                # Script to start the server in MacOS/Linux
  └── win_start.sh            # Script to start the server in Windows
```

## Learning Resources

- [HTTP Protocol](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol)
- [HTTP/1.1 Request Syntax](https://www.w3.org/Protocols/rfc2616/rfc2616-sec5.html)
- [Go net package](https://pkg.go.dev/net)

## Contributing

Contributions, issues, and feature requests are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License.

---

*Happy coding!*
