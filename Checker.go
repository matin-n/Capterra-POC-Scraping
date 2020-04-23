package main

import (
	"fmt"
	"github.com/headzoo/surf"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {

	var startID, endID int // listing ID's

	startID = 132510
	endID = 132520

	for i := startID; i <= endID; i++ {

		if validPage("https://www.capterra.com/p/"+strconv.Itoa(i)+"/CMS/") == true {

			pageURL, logoImageURL, businessName, website := scrapeLink("https://www.capterra.com/p/" + strconv.Itoa(i) + "/CMS/")

			fmt.Println("Page URL:", pageURL)
			fmt.Println("Business Name:", businessName)
			fmt.Println("Logo Image URL:", logoImageURL)
			fmt.Println("Website:", website)
			fmt.Println()
		}

	}

}

func scrapeLink(link string) (string, string, string, string) {
	bow := surf.NewBrowser()

	bow.SetTimeout(10 * time.Second)

	err := bow.Open(link)
	if err != nil {
		panic(err)
	}

	//1. Page URL (https://www.capterra.com/p/9014/Trakstar-Performance-Appraisal-Software/)
	//2. Logo image URL (https://cdn0.capterra-static.com/logos/150/2002547-1487627051.png)
	//3. Business Name: Trakstar 360 Degree Feedback
	//4. Website: http://www.trakstar.com";

	var pageURL, logoImageURL, businessName, website string

	pageURL = bow.Url().String()

	businessName = bow.Dom().Find(".DesktopProductHeader__ProductHeading-sc-1w230hs-2").Text() // find business name using CSS selector

	// iterate thru all the images and find the logo
	for _, image := range bow.Images() {
		if strings.Contains(image.URL.String(), "https://cdn0.capterra-static.com/logos/150/") {
			logoImageURL = image.URL.String()
			break
		}
	}

	bow.Click(".kSyqVF") // visit website button

	website = bow.Url().String() // company website URL

	return pageURL, logoImageURL, businessName, website
}

func validPage(url string) bool {

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "	Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3782.0 Safari/537.36 Edg/76.0.152.0")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if strings.Contains(string(body), "Browse All Business Software Directories at Capterra") {
		return false
	} else if strings.Contains(string(body), "<title>403 Forbidden</title></head>") {
		return false
	} else {
		return true
	}

}
