package urbandict

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/xmlpath.v2"
)

//FetchDef returns the results in an array
func FetchDef(word string) []string {
	baseURL := "https://www.urbandictionary.com/define.php?term="

	//http.Get returns a Response object and error if any
	resp, err := http.Get(baseURL + word)
	if err != nil {
		log.Print(err)
	}
	//closing resp.Body using defer
	defer resp.Body.Close()

	//the response body is parsed using the xmlpath package
	xmlroot, xmlerror := xmlpath.ParseHTML(resp.Body)
	if xmlerror != nil {
		log.Fatal(err)
	}

	//xpaths defined and compiled
	xpathMeaning := `//div[@class="meaning"]`
	xpathUpCount := `//div[@class="count"]/a[@class="up"]`
	xpathDownCount := `//div[@class="count"]/a[@class="down"]`

	pathMeaning := xmlpath.MustCompile(xpathMeaning)
	pathUpcount := xmlpath.MustCompile(xpathUpCount)
	pathDowncount := xmlpath.MustCompile(xpathDownCount)

	//defining JSON structure to be returned
	type Details struct {
		meaning   string
		upcount   string
		downcount string
	}
	var body []string

	//iterated through all the nodes to get all meanings
	iterMeaning := pathMeaning.Iter(xmlroot)
	iterUpCount := pathUpcount.Iter(xmlroot)
	iterDownCount := pathDowncount.Iter(xmlroot)

	var meaning, upCount, downCount []string
	for iterMeaning.Next() {
		meaning = append(meaning, iterMeaning.Node().String()) 
	}
	for	iterUpCount.Next(){
		upCount = append(upCount, iterUpCount.Node().String()) 
	}
	
	for	iterDownCount.Next(){
		downCount = append(downCount, iterDownCount.Node().String()) 
	}
		
	
		data := Details{meaning, upCount, downCount}
		b, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(b)
	}
	return body

}
