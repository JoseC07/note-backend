package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type summary struct {
	ID              string `json:"id"`
	Summary         string `json:"summary"`
	AudioPlayBack   string `json:"audioPlayBack"`
	AudioLength     string `json:"audioLength"`
	AudioTranscript string `json:"audioTranscript"`
	AudioUrl        string `json:"audioUrl"`
	AudioTitle      string `json:"audioTitle"`
	AudioDate       string `json:"audioDate"`
}

var summaries = []summary{
	{ID: "1", Summary: "Summary 1", AudioPlayBack: "AudioPlayBack 1", AudioLength: "AudioLength 1", AudioTranscript: "AudioTranscript 1", AudioUrl: "AudioUrl 1", AudioTitle: "AudioTitle 1", AudioDate: "AudioDate 1"},
	{ID: "2", Summary: "Summary 2", AudioPlayBack: "AudioPlayBack 2", AudioLength: "AudioLength 2", AudioTranscript: "AudioTranscript 2", AudioUrl: "AudioUrl 2", AudioTitle: "AudioTitle 2", AudioDate: "AudioDate 2"},
	{ID: "3", Summary: "Summary 3", AudioPlayBack: "AudioPlayBack 3", AudioLength: "AudioLength 3", AudioTranscript: "AudioTranscript 3", AudioUrl: "AudioUrl 3", AudioTitle: "AudioTitle 3", AudioDate: "AudioDate 3"},
}

func getSummaries(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, summaries)
}

func postSummaries(c *gin.Context) {
	var newSummary summary

	//Call Bind JSON to bing the recieved JSON to new Summary

	if err := c.BindJSON(&newSummary); err != nil {
		return
	}

	summaries = append(summaries, newSummary)
	c.IndentedJSON(http.StatusCreated, newSummary)

}

func getSummaryByID(c *gin.Context) {
	idToCheck := c.Param("id")

	//loop over the list of summaries looking for
	// a summary whose ID

	for _, value := range summaries {
		if value.ID == idToCheck {
			c.IndentedJSON(http.StatusOK, value)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"messge": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/summaries", getSummaries)
	router.GET("/summaries/:id", getSummaryByID)
	router.POST("/summaries", postSummaries)

	router.Run("localhost:8080")

}
