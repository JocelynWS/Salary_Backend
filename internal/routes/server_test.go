package routes

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Giả lập server nhưng không gọi r.Run()
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.POST("/calculate", func(c *gin.Context) {
		var req struct {
			GrossSalary float64 `json:"gross"`
			Dependents  int     `json:"dependents"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "detail": err.Error()})
			return
		}

		netSalary := req.GrossSalary - float64(req.Dependents)*4400000 // giả định đơn giản
		c.JSON(http.StatusOK, gin.H{
			"gross":      req.GrossSalary,
			"dependents": req.Dependents,
			"net_salary": int(netSalary),
		})
	})

	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		if file == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File missing"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "file received"})
	})

	return r
}

func TestCalculateAPI(t *testing.T) {
	router := setupRouter()

	body := map[string]interface{}{
		"gross":      20000000,
		"dependents": 2,
	}
	jsonValue, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/calculate", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), `"net_salary":`)
}

func TestUploadAPI(t *testing.T) {
	router := setupRouter()

	file, err := os.CreateTemp("", "test.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, err := writer.CreateFormFile("file", file.Name())
	assert.NoError(t, err)
	part.Write([]byte("dummy content")) // vì không cần xử lý thực

	writer.Close()

	req, _ := http.NewRequest("POST", "/upload", &buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), `"file received"`)
}
