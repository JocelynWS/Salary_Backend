package routes

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"salary_api_ver1/internal/control"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize/v2"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Load giao diện HTML từ file
	r.LoadHTMLFiles("frontend/index.html")
	r.Static("/static", "./frontend") // Nếu sau này có thêm file CSS/JS

	// Giao diện người dùng
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		f, _ := file.Open()
		defer f.Close()

		xls, err := excelize.OpenReader(f)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read Excel"})
			return
		}

		rows, _ := xls.GetRows("Sheet1")
		results := []gin.H{}
		for i, row := range rows {
			if i == 0 || len(row) < 3 {
				continue
			}
			gross, err := strconv.ParseFloat(row[1], 64)
			if err != nil {
				continue
			}
			dependents, err := strconv.Atoi(row[2])
			if err != nil {
				dependents = 0
			}
			net := control.CalculateNetSalary(gross, dependents)
			results = append(results, gin.H{
				"name":       row[0],
				"gross":      gross,
				"dependents": dependents,
				"net":        int(net),
			})
		}

		c.JSON(http.StatusOK, gin.H{"salaries": results})
	})

	return r
}

func TestRootNoParams(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Welcome")
}

func TestRootWithParams(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/?gross=10000000&dependents=1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "gross_salary")
	assert.Contains(t, w.Body.String(), "net_salary")
}

func TestUploadEndpoint(t *testing.T) {
	router := setupRouter()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	filePath := "./testdata.xlsx"
	fileWriter, _ := writer.CreateFormFile("file", filepath.Base(filePath))
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("failed to open test file: %v", err)
	}
	defer file.Close()
	io.Copy(fileWriter, file)
	writer.Close()

	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "salaries")
}
