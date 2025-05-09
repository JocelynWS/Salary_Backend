package routes

import (
	"net/http"
	"strconv"

	"salary_api_ver1/internal/control"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func InitServer() {
	r := gin.Default()
	//r.StaticFile("/frontend", "../static/index.html")

	// API GET
	// Load giao diện HTML từ file
	r.LoadHTMLFiles("/home/jocelyn/salary_api_ver1/static/index.html")
	r.Static("/static", "./static") // Nếu sau này có thêm file CSS/JS

	// Giao diện người dùng
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// POST /calculate: Tính lương từ form tay
	r.POST("/calculate", func(c *gin.Context) {
		var req struct {
			GrossSalary float64 `json:"gross"`
			Dependents  int     `json:"dependents"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "detail": err.Error()})
			return
		}

		netSalary := control.CalculateNetSalary(req.GrossSalary, req.Dependents)
		c.JSON(http.StatusOK, gin.H{
			"gross":      req.GrossSalary,
			"dependents": req.Dependents,
			"net_salary": int(netSalary),
		})
	})

	// API POST
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		f, _ := file.Open()
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

	r.Run(":8081")
}
