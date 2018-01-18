package urbandict

import (
	"log"
	"net/http"

	"gopkg.in/xmlpath.v2"
)

//defining JSON structure to be returned
type Details struct {
	Meaning   string
	Upcount   string
	Downcount string
}

//FetchDef returns the results in an array
func FetchDef(word string) []Details {
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
	xpathUpCount := `//div[@class="thumbs"]/a[@class="up"]`
	xpathDownCount := `//div[@class="thumbs"]/a[@class="down"]`

	pathMeaning := xmlpath.MustCompile(xpathMeaning)
	pathUpcount := xmlpath.MustCompile(xpathUpCount)
	pathDowncount := xmlpath.MustCompile(xpathDownCount)

	//iterated through all the nodes to get all meanings
	iterMeaning := pathMeaning.Iter(xmlroot)
	iterUpCount := pathUpcount.Iter(xmlroot)
	iterDownCount := pathDowncount.Iter(xmlroot)

	//Making an array of elements for Meanings, upcount and downcount
	var meaning, upCount, downCount []string
	for iterMeaning.Next() {
		meaning = append(meaning, iterMeaning.Node().String())
	}
	for iterUpCount.Next() {
		upCount = append(upCount, iterUpCount.Node().String())
	}

	for iterDownCount.Next() {
		downCount = append(downCount, iterDownCount.Node().String())
	}

	//iterating through the elements and making an array of it
	var d []Details
	for i := 0; i < len(meaning); i++ {
		d = append(d, Details{meaning[i], upCount[i], downCount[i]})
	}
	return d

}
