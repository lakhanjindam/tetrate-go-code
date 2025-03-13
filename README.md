# Tetrate go code samples

Created a Go GIN REST API server which runs on `localhost:8080`

## Run application

```bash
go run main.go
```

### Available Endpoints

- **GET /api/v1/ping**

  - Description: Check if the server is running
  - Example:

  ```bash
  curl http://localhost:8080/api/v1/ping
  ```

- **POST /api/v1/country**
  - Description: get country details
  - Example:
  ```bash
  curl http://localhost:8080/api/v1/country?name={country-name}
  ```
