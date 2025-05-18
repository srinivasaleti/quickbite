# Developer Guide

This project uses:

- React (with TypeScript) for UI
- Go (version 1.23 or above) for backend server
- Node.js version 20 or above
- Yarn for frontend package manager
- Docker
- Makefile for commands

---

## Requirements

Make sure you have these installed before starting:

- Node.js >= 20
- Go >= 1.23
- Yarn (you can install using `npm install -g yarn`)
- Make (comes pre-installed on Linux/macOS)
- Docker

---

# Start App in Local Production Mode

To run the app like production on your local system, follow the steps below:

---

### Prerequisites

Before running the command, make sure:

- You have **Go 1.23** installed
- You have **Node 20** installed
- Docker

---

### Run the App

```bash
make app-up
```

### What This Command Does

- Uses Docker to run everything
- Builds the UI and puts the files in: `server/web/assets`
- Builds the Go server
- Runs the app inside Docker on port 8080
- Exposes the app to your system on port 18080

Access app on: http://localhost:18080

**_NOTE: If you are doing this first time it might take sometime 5-10mins. As it will install all go,ui packages and pull docker images._**

# Start App in Development Mode

When working in local development, you might want to run the database inside Docker.

---

### Start the Database

To start the database using Docker, run:

```bash
make compose-quickbitedb-up
```

### Setup Environment

-The backend needs a .env file for configuration. I have already provided a `.env.sample` file.

- It has a key called: QUICKBITE_DB_CONNECTION_STRING
- If you are using the command above to run the DB, you don’t need to change anything.
- If you are using a different DB, you can update the value as needed.

### Add backend url to UI

- I have created a custom hook to handle all API calls: `ui/src/common/hooks/useApi.ts`
- By default, it uses the base URL: `/api`
- In development mode, **UI and Server run on different ports**
- Because of this, the **UI cannot talk to the backend directly**
- Please **change the `baseUrl`** to point to your backend port (for example: `http://localhost:8080/api`)

```
const api = axios.create({
  baseURL: " http://localhost:<server_port>",
  headers: {
    "Content-Type": "application/json",
  },
});
```

### Start Server and UI Together

Run this command to start both server and UI:

```
make app-dev SERVER_PORT=<server_port>
```

Replace <server_port> with your desired port(or leave it empty to use default 8080)

- **What This Command Does**
- Starts the Go server on given port (default: 8080)

- Starts the React UI on port 5173

### Open in you browser

- UI: http://localhost:5173
- Server: http://localhost:8080
- or http://localhost:<server_port>

# Stop Server

To stop the running server:

```bash
make app-down
```

This will:

- Stop the Go server on given server port (default: 8080)
- Stop the React UI on port 5173

# Run server tests

I have added tests to only backend. I have added both unit tests & integration tests using `testcontainers`.

---

### Prerequisites

Before running the command, make sure:

- You have **Go 1.23** installed
- Docker

---

### Run the tests

```bash
make go-coverage
```

# Data Seeding for Categories and Products

When the server starts, it runs the seeder which automatically adds or updates data in the database.

### Where the Data Lives

- Categories data is in: `server/internal/domain/product/seeder/data/categories.yaml`
  - Here, **category name** is used as the key to match data between the database and YAML file.
- Products data is in: `server/internal/domain/product/seeder/data/products.yaml`
  - Here, **externalID** is used as the key to match data between the database and YAML file.

Both files have the details of products and categories in simple YAML format.

### Adding New Data

Feel free to add new categories or products to these YAML files. When the server restarts, the seeder will update the database with your new data automatically.

# Coupons Folder

This folder **must contain** the following files:

- `couponbase1.gz`
- `couponbase2.gz`
- `couponbase3.gz`

- Please add aboves file in this folder.

---

### Required for Coupon Validation

These files are **required** for the coupon filtering system to work correctly.  
Make sure they exist **before running the service**.

If **any file is missing**, the app will **not** be able to validate coupons.

---

### Important Note for Testing Coupons

The coupon files are **very large** and take time to load after the server starts.  
So when testing with coupons, **please wait 10–20 seconds** after starting the server.  
This ensures all coupons are properly loaded.

---

### Discounts on Coupon code

The coupon codes `HAPPYHOURS` and `BUYGETONE` are **not present** in the original coupon files.  
However, they are listed in the challenge description here:  
[oolio-group/kart-challenge](https://github.com/oolio-group/kart-challenge)

So I have **explicitly added support** for them:

- **`HAPPYHOURS`** – Applies an **18% discount** on the total order.
- **`BUYGETONE`** – Gives the **lowest priced item for free**.

### Note on Coupon Codes

In the challenge README here:  
[oolio-group/kart-challenge](https://github.com/oolio-group/kart-challenge/blob/advanced-challenge/backend-challenge/README.md),  
some coupon codes are mentioned under **Valid Promo Codes**, such as:

- `HAPPYHRS`
- `FIFTYOFF`

However, the README **does not provide clear details** about what these coupons should actually do.  
So, I made the following decisions:

- If the coupon code is **valid** (i.e., found in the coupon file), I apply a **10% discount** on the total order.
- If the coupon code is **invalid** (e.g., `SUPER100`), I throw an **error** saying the coupon is not valid.

This approach ensures all valid coupons have **some discount behavior**, while clearly handling invalid inputs.

If more detailed information is added in the future, the logic can be updated accordingly.
