package UrbanDict

import (
	"io/ioutil"
	"log"
	"net/http"
)

//FetchDef returns the results in an array
func FetchDef(word string) []byte {
	baseURL := "https://www.urbandictionary.com/define.php?"
	resp, err := http.Get(baseURL + word)
	if err != nil {
		log.Print("err")
		log.Print(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("err")
		log.Print(err)
	}
	return body
}
