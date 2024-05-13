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

func main() {
	router := gin.Default()
	router.GET("/maddadjokes", getJoke)

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		return
	}
}