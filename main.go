package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
)

type Data struct {
	Jokes []Joke `json:"jokes"`
}

type Joke struct {
	Q string `json:"q"`
	A string `json:"a"`
}

type Healthcheck struct {
	Status string `json:"status"`
}

func loadData() Data {
	jsonFile, err := os.Open("data/jokes.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var payload Data
	json.Unmarshal(byteValue, &payload)
	return payload
}

func getRandomJoke(data Data) Joke {
	randomIndex := rand.Intn(len(data.Jokes))
	pick := data.Jokes[randomIndex]
	return pick
}

func getJoke(c *gin.Context) {
	data := loadData()
	joke := getRandomJoke(data)
	c.IndentedJSON(http.StatusOK, joke)
}

func getSlackJoke(c *gin.Context) {
	data := loadData()
	joke := getRandomJoke(data)
	val := fmt.Sprintf("*%s*\n\t%s", joke.Q, joke.A)

	c.IndentedJSON(http.StatusOK, map[string]any{
		"response_type": "in_channel", "type": "mrkdwn", "text": val})

}

func healthCheck(c *gin.Context) {
	ok := Healthcheck{
		Status: "ok",
	}
	c.IndentedJSON(http.StatusOK, ok)
}

func main() {
	router := gin.Default()
	router.GET("/maddadjokes", getJoke)
	router.POST("/maddadjokes/slack", getSlackJoke)

	router.GET("/maddadjokes/health", healthCheck)

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		return
	}
}
