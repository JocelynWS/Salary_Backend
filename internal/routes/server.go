package routes

import (
	"net/http"
	"strconv"

	"salary-api/internal/control"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func InitServer() {
	r := gin.Default()
	r.StaticFile("/frontend", "../static/index.html")

	// API GET
	r.GET("/", func(c *gin.Context) {
		grossStr := c.Query("gross")
		dependentStr := c.Query("dependents")

		if grossStr == "" {
			c.JSON(http.StatusOK, gin.H{
				"message": "Welcome! Please provide ?gross=amount&dependents=n to calculate net salary.",
			})
			return
		}

		gross, err := strconv.ParseFloat(grossStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gross amount"})
			return
		}

		dependents, err := strconv.Atoi(dependentStr)
		if err != nil || dependents < 0 {
			dependents = 0
		}

		net := control.CalculateNetSalary(gross, dependents)
		c.JSON(http.StatusOK, gin.H{
			"gross_salary": gross,
			"dependents":   dependents,
			"net_salary":   int(net),
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
