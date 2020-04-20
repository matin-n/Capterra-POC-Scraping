package main

import (
	"flag"
	"fmt"
	"github.com/headzoo/surf"
	"strings"
	"time"
)

var link = flag.String("link", "https://www.capterra.com/p/132513/Breeze-ChMS/", "what is the link to the URL?")

func main() {

	flag.Parse()

	pageURL, logoImageURL, businessName, website := scrapeLink(*link)

	fmt.Println("Page URL:", pageURL)
	fmt.Println("Business Name:", businessName)
	fmt.Println("Logo Image URL:", logoImageURL)
	fmt.Println("Website:", website)

	fmt.Scanf("h")

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

	/*
		fmt.Println("Page URL:", pageURL)
		fmt.Println("Business Name:", businessName)
		fmt.Println("Logo Image URL:", logoImageURL)
		fmt.Println("Website:", website)
	*/

	return pageURL, logoImageURL, businessName, website
}
