package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetEndpoint(endpoint string) (string, error) {
	response, err := http.Get(endpoint)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	stringResponse := string(responseData)
	fmt.Println(stringResponse)
	return stringResponse, nil
}