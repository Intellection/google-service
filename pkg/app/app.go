package app

import (
	"encoding/json"
	"google-service/pkg/google"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	r.GET("/healthz", health)
	r.GET("/spreadsheet", spreadsheet())
	r.Run()
}

func health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func spreadsheet() func(*gin.Context) {
	sheetsService := google.SheetsService()

	return func(c *gin.Context) {

		spreadsheetID := c.Request.URL.Query().Get("spreadsheetId")
		if spreadsheetID == "" {
			c.JSON(400, gin.H{"error": "Must supply a spreadsheetId URL param"})
			return
		}

		spreadsheetRange := c.Request.URL.Query().Get("range")
		if spreadsheetRange == "" {
			c.JSON(400, gin.H{"error": "Must supply a range URL param. See here: https://support.google.com/docs/answer/6208276?hl=en"})
			return
		}

		resp, err := sheetsService.Spreadsheets.Values.Get(spreadsheetID, spreadsheetRange).Do()
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Display pulled data
		if len(resp.Values) == 0 {
			c.JSON(400, gin.H{"error": "No data was found for spreadsheet " + spreadsheetID})
			return
		}

		b, err := json.Marshal(resp.Values)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.Data(http.StatusOK, "application/json", b)
	}
}
