# Babalaas Web Server
The web server for the Babalaas mobile application.  
- Written in Go using modules
- REST API 
- CRUD functionality for the Babalaas entities
- Connects to PostgreSQL Database

## Running Locally
1. Be sure to have Go 1.19+ installed 
2. Clone this repository
3. Create a dotenv file named ".env" in the root directory
    1. Copy and paste the text from the .env template below
    2. Replace the template values with actual environment variable values
4. Within the web-server directory execute the following command:
    ```BASH
    $ docker-compose up
    ```
5. You can access the API through the following address: http://localhost:8080/api/
6. Verify API functionality by performing the following GET command and recieving a JSON response: http://localhost:8080/api/posts

## .env File Template
```text
CONNECTION_STRING="postgres://{user}:{password}@{hostname}:{port}/{database-name}"
PORT=8080
```