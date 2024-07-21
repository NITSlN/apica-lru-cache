# LRU Cache Application

This project implements an LRU (Least Recently Used) cache with a Go backend and a React frontend. The backend provides a RESTful API and a WebSocket for real-time updates, while the frontend displays the cache items and allows users to add new items.

## Project Structure

- **server**: Contains the Go backend implementation.
- **client**: Contains the React frontend implementation.

## Features

### Backend

- **LRU Cache Implementation**: A thread-safe LRU cache implemented in Go using a hash map and doubly linked list.
- **API Endpoints**:
  - `GET /cache/{key}`: Retrieve a value by key.
  - `POST /cache`: Add a new key-value pair with an expiration time.
  - `DELETE /cache/{key}`: Delete a key-value pair.
- **WebSocket**: Provides real-time updates of the cache state.

### Frontend

- **React Application**: Displays cache items and allows adding new items.
- **Real-time Updates**: Uses WebSocket to receive real-time updates from the backend.
- **Tailwind CSS**: For styling the components.

## Prerequisites

- **Go**: Install from [golang.org](https://golang.org/doc/install).
- **Node.js**: Install from [nodejs.org](https://nodejs.org/).

## Setup

### Backend

1. **Navigate to the backend directory**:

    ```sh
    cd server
    ```

2. **Install dependencies**:

    ```sh
    go mod tidy
    ```

3. **Run the server**:

    ```sh
    go run ./cmd/server
    ```

4. **Server will be running at**:

    ```
    http://localhost:8000
    ```

### Frontend

1. **Navigate to the frontend directory**:

    ```sh
    cd client
    ```

2. **Install dependencies**:

    ```sh
    npm install
    ```

3. **Start the development server**:

    ```sh
    npm start
    ```

4. **Application will be running at**:

    ```
    http://localhost:3000
    ```

## Backend Code Structure

- **main.go**: Entry point of the application, sets up the server and routes.
- **cache.go**: Implements the LRU cache with set, get, and delete operations.
- **handlers.go**: Contains the HTTP handlers for the API endpoints and WebSocket.
- **router.go**: Sets up the routes for the API endpoints.
- **structs.go**: Defines the data structures used in the application.

## Frontend Code Structure

- **src/App.js**: Main application component.
- **src/components/CacheList.js**: Component to display cache items.
- **src/components/CacheItem.js**: Component to render individual cache item.
- **src/index.js**: Entry point of the React application.
- **src/styles/tailwind.css**: Tailwind CSS configuration.

## API Endpoints

- **GET /cache/{key}**: Retrieve a value by key.

    ```sh
    curl http://localhost:8000/cache/{key}
    ```

- **POST /cache**: Add a new key-value pair with an expiration time.

    ```sh
    curl -X POST http://localhost:8000/cache \
         -H "Content-Type: application/json" \
         -d '{"key":"mykey", "value":"myvalue", "expiration":3600}'
    ```

- **DELETE /cache/{key}**: Delete a key-value pair.

    ```sh
    curl -X DELETE http://localhost:8000/cache/{key}
    ```

## WebSocket

- **ws://localhost/ws**: Connect to the WebSocket to receive real-time updates of the cache state.

## Usage

### Adding a New Cache Item

1. Open the frontend application in your browser at `http://localhost:3000`.
2. Fill in the key, value, and expiration time fields.
3. Click "Add in cache" to add the item to the cache.

### Viewing Cache Items

1. The frontend application will display the cache items in a table.
2. The expiration time will update in real-time as received from the WebSocket.

### Deleting a Cache Item

1. Use the DELETE API endpoint to remove a cache item by key.
