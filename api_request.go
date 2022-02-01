package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func return_link(chat_id int64) string {
	url := goDotEnvVariable("SCALEO_URL")
	method := "POST"
	client := &http.Client{}

	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var users Users
	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].ChatID != chat_id {
			continue
		}
		var rawText = fmt.Sprintf(`{
	 "email": "%s",
	    "password": "%s"
		}`, users.Users[i].Email, users.Users[i].Password)
		payload := strings.NewReader(rawText)
		req, err := http.NewRequest(method, url, payload)
		if err != nil {
			fmt.Println(err)
			return "something went wrong"
		}
		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return "something went wrong"
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return "something went wrong"
		}

		var info Info
		json.Unmarshal(body, &info)
		link := info.Info.Link

		return link

	}
	return "no user found"
}
