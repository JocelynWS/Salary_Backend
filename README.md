# SALARY-API: Gross to Net Salary Conversion Application

This application provides a simple API to convert Gross salary to Net salary based on the number of dependents. It includes two main endpoints:

* **GET /:** Calculates Net salary from Gross salary and the number of dependents passed via query parameters.
* **POST /upload:** Calculates Net salary for multiple individuals from an uploaded Excel file.

## Requirements

* **Go:** Version 1.18 or higher (for `go.mod` support).
* **Web Browser:** To interact with the user interface (if you have added one).

## Running the Application

1.  **Clone the repository (if you have one):**
    ```bash
    git clone [https://github.com/JocelynWS/NetEase](https://github.com/JocelynWS/NetEase)
    cd salary-api
    ```

2.  **Download dependencies:**
    ```bash
    go mod tidy
    go mod download
    ```

3.  **Run the backend server:**
    ```bash
    go run cmd/main.go
    ```
    The server will start and listen on port `http://localhost:8081`.

4.  **Open the user interface (frontend):**
    Open your web browser and navigate to `http://localhost:8081/frontend`.

## API Usage

### 1. Calculate Net Salary (GET API)

* **Endpoint:** `/`
* **Method:** `GET`
* **Parameters (query parameters):**
    * `gross`: The Gross salary amount (required).
    * `dependents`: The number of dependents (required, can be 0).
* **Example:**
    `http://localhost:8081/?gross=70000000&dependents=0`
* **Response (JSON):**
    ```json
    {
        "gross_salary": 70000000,
        "dependents": 0,
        "net_salary": 52987500
    }
    ```

### 2. Calculate Net Salary from File (API POST)

* **Endpoint:** `/upload`
* **Method:** `POST`
* **Headers:** `Content-Type: multipart/form-data`
* **Body (form-data):**
    * `file`: Select an Excel (`.xlsx`) file containing salary data. The file is expected to have columns in the following order: `Name`, `Gross Salary`, `Number of Dependents` (the header row may be present or absent; the current code ignores the header).

    


