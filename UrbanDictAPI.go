// package UrbanDict

// import (
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"strings"

// 	"golang.org/x/net/html"
// 	"gopkg.in/xmlpath.v2"
// )

// //FetchDef returns the results in an array
// func FetchDef(word string) []byte {
// 	baseURL := "https://www.urbandictionary.com/define.php?"
// 	resp, err := http.Get(baseURL + word)
// 	if err != nil {
// 		log.Print("err")
// 		log.Print(err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	return body
// }

package UrbanDict

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/xmlpath.v2"
)

//FetchDef returns the results in an array
func FetchDef(word string) []string {
	baseURL := "https://www.urbandictionary.com/define.php?"
	resp, err := http.Get(baseURL + word)
	if err != nil {
		log.Print("err")
		log.Print(err)
	}
	defer resp.Body.Close()

	//read, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//reader = strings.NewReader(read)
	xmlroot, xmlerror := xmlpath.ParseHTML(resp.Body)
	if xmlerror != nil {
		log.Fatal(err)
	}

	xpath1 := `//a[@class="word"]`
	path1 := xmlpath.MustCompile(xpath1)

	iter := path1.Iter(xmlroot)
	var body []string
	for iter.Next() {
		fmt.Println(iter.Node().String())
		body = append(body, iter.Node().String())

	}

	return body
}
