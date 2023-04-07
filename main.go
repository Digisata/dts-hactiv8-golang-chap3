package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	for {
		water, wind := rand.Intn(100), rand.Intn(100)
		postData(water, wind)
		time.Sleep(15 * time.Second)
	}
}

func postData(water, wind int) {
	data := map[string]interface{}{
		"water": water,
		"wind":  wind,
	}

	reqJson, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(reqJson))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))
	resData := map[string]int{}
	err = json.Unmarshal(body, &resData)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resData)
	waterStatus := getStatus(resData["water"], 5, 6, 8)
	windStatus := getStatus(resData["wind"], 6, 7, 15)

	fmt.Println("status water: ", waterStatus)
	fmt.Println("status wind: ", windStatus)
}

func getStatus(value, low, medium, high int) string {
	if value < low {
		return "aman"
	} else if value >= medium && value <= high {
		return "siaga"
	} else {
		return "bahaya"
	}
}
