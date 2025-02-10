package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bugReport struct {
	Id            string `json:"id"`
	Subject       string `json:"subject"`
	TypeId        int    `json:"type_id"`
	DateSubmitted string `json:"date_submitted"`
	Description   string `json:"description"`
}

// TODO: Remove and place in unit test
var bugReports = []bugReport{
	{Id: "1", Subject: "Large files do not send", TypeId: 1, DateSubmitted: "2021-01-01", Description: "When I try to send large files, the app crashes."},
	{Id: "2", Subject: "Chatroom closes with errors", TypeId: 2, DateSubmitted: "2021-01-02", Description: "When I try to open the chatroom, it closes with errors."},
	{Id: "3", Subject: "Text gets cut off in chat room", TypeId: 3, DateSubmitted: "2021-01-03", Description: "When I type a long message, it gets cut off."},
}

func getBugReports(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, bugReports)
}

func saveBugReport(c *gin.Context) {
	var newBugReport bugReport
	if err := c.BindJSON(&newBugReport); err != nil {
		log.Println("New bug report added errored: ", err)
		return
	}
	bugReports = append(bugReports, newBugReport)
	log.Println("New bug report added: ", newBugReport)
	c.IndentedJSON(http.StatusCreated, newBugReport)
}

func main() {
	log.Println("Showing bug reports: ", bugReports)

	router := gin.Default()
	router.GET("/bug-reports", getBugReports)
	router.POST("/bug-reports", saveBugReport)
	router.Run("localhost:8080")
}
