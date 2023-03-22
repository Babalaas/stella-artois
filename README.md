# Babalaas Application Server
The server for the Babalaas social media mobile application.
- Written in Go using modules
- Connects to PostgreSQL container with test data
- Designed using principles from the Go Clean Architecture

## Running Locally
1. Be sure to have Go 1.19, Docker, and Docker Compose installed 
2. Clone this repository
3. Within the stella-artois directory execute the following command:
    ```BASH
    $ docker-compose up --build
    ```
4. You can access the server through the following address: http://localhost:8080/
5. Verify server functionality by performing the following GET command and recieving a "pong" response: http://localhost:8080/ping

## .env File Template
```text
CONNECTION_STRING="postgres://{user}:{password}@{hostname}:{port}/{database-name}"
PORT=8080
GIN_MODE="debug"
```