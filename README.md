# SALARY-API: Gross to Net Salary Conversion Application

This application provides a simple yet powerful API to convert **Gross salary** to **Net salary**, factoring in the number of dependents. It's built with **Go (Golang)** and supports both individual and batch salary calculations via Excel uploads.

Ideal for:
- HR Systems
- Payroll Applications
- Admin Tools

---

## ✨ Features

- `GET /` — Displays optional user interface (if implemented)
- `POST /calculate` — Computes net salary from gross salary & dependents
- `POST /upload` — Upload Excel file with multiple entries for batch processing

---
### Project structure
.github/  
&nbsp;&nbsp;&nbsp;&nbsp;└── workflows/  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;└── ci.yaml  
cmd/  
&nbsp;&nbsp;&nbsp;&nbsp;└── main.go  
internal/  
&nbsp;&nbsp;&nbsp;&nbsp;├── control/  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;└── calculateSalary.go  
&nbsp;&nbsp;&nbsp;&nbsp;├── model/  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;└── salary.go  
&nbsp;&nbsp;&nbsp;&nbsp;└── routes/  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;├── server_test.go  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;├── server.go  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;└── testdata.xlsx  
static/  
&nbsp;&nbsp;&nbsp;&nbsp;└── index.html  
test/  
&nbsp;&nbsp;&nbsp;&nbsp;└── testdata.xlsx  
Dockerfile  
go.mod  
go.sum  
README.md



---

## 🚀 Requirements

| Component        | Version/Requirement     |
|------------------|--------------------------|
| Go               | 1.18 or higher            |
| Docker           | 20.10+ (optional)         |
| Web Browser      | To view UI (if applicable) |

---

## 🏃 Running the Application

### 1. Clone the Repository

```bash
git clone https://github.com/JocelynWS/Salary_Backend
cd Salary_Backend
```

### 2. Download Dependencies

```bash
go mod tidy
go mod download
```
### 3. Run the Backend Server

```bash
go run cmd/main.go
```
The server will be available at: http://localhost:8081

## 🔌 API Usage (use Postman to test)

**1. Calculate Net Salary (Single Entry)**  
**Endpoint:** `/calculate`  

**Method:** POST  

**Body (form-data):**  
- `gross`: Gross salary (required)  
- `dependents`: Number of dependents (required)  

**2. Calculate Net Salary from Excel File**  
**Endpoint:** `/upload`  

**Method:** POST  

**Headers:**  
`Content-Type: multipart/form-data`  

**Body (form-data):**  
- `file`: Excel `.xlsx` file with columns:  
  - Name  
  - Gross Salary  
  - Number of Dependents  

## ⚙️ CI/CD via GitHub Actions

**GitHub Actions config file:** `.github/workflows/ci.yaml`

**Triggers:**  
- On push to `master`  
- On pull request to `master`  

**Workflow Steps:**  
1. Checkout code  
2. Set up Go (e.g., v1.23)  
3. Run tests: `go test -v ./...`  
4. Build binary: `salary_api`  
5. (Optional) Build & push Docker image to Docker Hub  

### 🐳 Docker Deployment

**Build & Run Locally**  
```bash
docker build -t salary-api:latest .
docker run -d -p 8081:8081 --name salary_backend salary-api:latest
```
### Push to Docker Hub

```bash
docker tag salary-api jocelyn33/salary-api:v2
docker push jocelyn33/salary-api:v2
```
### Pull & Run from Docker Hub

```bash
docker pull jocelyn33/salary-api:latest
docker run -d -p 8081:8081 --name salary_backend jocelyn33/salary-api:latest
```
### ☁️ Deploy to Render

**Option 1: Native Go Deployment**

1. Log in to Render  
2. Click **New + → Web Service**  
3. Connect to your GitHub repository  
4. Set:  
   - **Name:** `salary-api-v2`  
   - **Branch:** `master`  
5. Click **Create Web Service**  

Render will build and deploy your app.

**Option 2: Docker Deployment**

- Ensure your Docker image is pushed to Docker Hub

On Render:

1. Click **New + → Web Service**  
2. Choose **Deploy an existing Dockerfile**  
3. Use Docker image: `jocelyn33/salary-api:v2`  

Render builds and deploys your container.

Your app will be available at a public URL like:  
`https://salary-api-v2.onrender.com/`




