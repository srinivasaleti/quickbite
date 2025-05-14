# Developer Guide

This project uses:

- React (with TypeScript) for UI
- Go (version 1.23 or above) for backend server
- Node.js version 20 or above
- Yarn for frontend package manager
- Makefile for commands

---

## Requirements

Make sure you have these installed before starting:

- Node.js >= 20
- Go >= 1.23
- Yarn (you can install using `npm install -g yarn`)
- Make (comes pre-installed on Linux/macOS)

---

## Start App in Development Mode

To start both the server and the UI:

```bash
make app-dev SERVER_PORT=<server_port>
```

This will:

- Start the Go server on given server port (default: 8080)
- Start the React UI on port 5173

Now you can access:

- UI : http://localhost:5173
- Server: http://localhost:8080 or http://localhost:<server_port>

## Stop Server

To stop the running server:

```bash
make app-down
```

This will:

- Stop the Go server on given server port (default: 8080)
- Stop the React UI on port 5173
