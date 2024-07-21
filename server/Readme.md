# LRU Cache with Golang Backend and React Frontend

## Backend Setup

1. **Clone the repository:**

    ```sh
    git clone <repository-url>
    cd <repository-directory>
    ```

2. **Install dependencies:**

    ```sh
    go mod tidy
    ```

3. **Run the server:**

    ```sh
    go run
    ```

## API Endpoints

### GET /cache/{key}
Retrieve a value based on the key.

### POST /cache
Add a key/value pair with an expiration time.
```json
{
    "key": "exampleKey",
    "value": "exampleValue",
    "expiration": 5 // in seconds
}
