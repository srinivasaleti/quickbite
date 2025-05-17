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

When you work in local you might want to have a db running inside a docker. To start db

```bash
make compose-quickbitedb-up
```

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

## Data Seeding for Categories and Products

When the server starts, it runs the seeder which automatically adds or updates data in the database.

### Where the Data Lives

- Categories data is in: `server/internal/domain/product/seeder/data/categories.yaml`
  - Here, **category name** is used as the key to match data between the database and YAML file.
- Products data is in: `server/internal/domain/product/seeder/data/products.yaml`
  - Here, **externalID** is used as the key to match data between the database and YAML file.

Both files have the details of products and categories in simple YAML format.

### Adding New Data

Feel free to add new categories or products to these YAML files. When the server restarts, the seeder will update the database with your new data automatically.
