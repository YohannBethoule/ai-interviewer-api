package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"interviewer-api/domain"
	"interviewer-api/service"
	"net/http"
	"os"
)

type startRequest struct {
	JobTitle            string `form:"jobTitle"`
	YearsOfExperience   uint   `form:"yearsOfExperience"`
	InterviewerPosition string `form:"interviewerPosition"`
	JobDescription      string `form:"jobDescription"`
}

func testPrompt() gin.HandlerFunc {
	return func(c *gin.Context) {
		message := c.Query("message")
		openaiClient := service.NewOpenAiClient(os.Getenv("OPENAI_API_KEY"))
		answer, err := openaiClient.SendRequest(message)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.String(http.StatusOK, answer)
	}
}

func startConversation() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request startRequest
		if errBindRequest := c.ShouldBindQuery(&request); errBindRequest != nil {
			fmt.Sprintln("ERROR reading request : %s", errBindRequest.Error())
			c.String(http.StatusBadRequest, errBindRequest.Error())
		}
		jobSummaryPrompt := domain.GenerateJobSummaryPrompt(request.JobDescription)
		openaiClient := service.NewOpenAiClient(os.Getenv("OPENAI_API_KEY"))
		answer, err := openaiClient.SendRequest(jobSummaryPrompt)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.String(http.StatusOK, answer)
	}
}
