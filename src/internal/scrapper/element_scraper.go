package scrapper

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Element struct {
	Name     string `json:"name"`
	ImageUrl string `json:"imageUrl,omitempty"`
	Link     string `json:"link,omitempty"`
	Category string `json:"category,omitempty"`
}

func ScrapElements() ([]Element, error) {
	var allElements []Element
	baseURL := "https://little-alchemy.fandom.com/wiki/Category:Little_Alchemy_2"

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	for char := 'A'; char <= 'Z'; char++ {
		url := fmt.Sprintf("%s?from=%c", baseURL, char)

		if len(allElements) >= 200 {
			break
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Printf("Error creating request: %v", err)
			continue
		}

		req.Header.Set("User-Agent", "FireBoyWaterGirl-AlchemyProject/1.0")

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Error making request: %v", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			log.Printf("Status code error: %d %s", resp.StatusCode, resp.Status)
			continue
		}


		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Printf("Error parsing HTML: %v", err)
			continue
		}

		doc.Find("#mw-content-text > div.category-page__members > div > ul > li > a").Each(func(i int, s *goquery.Selection) {
			var el Element

			nameCell := s.Text()
			if nameCell != "" {
				el.Name = strings.TrimSpace(nameCell)
			}

			link, exists := s.Attr("href")
			if exists && link != "" {
				el.Link = "https://little-alchemy.fandom.com" + link
			}


			imgEl := s.Find("a")
			if imgEl.Length() > 0 {
				imgSrc, exists := imgEl.Attr("href")
				if exists && imgSrc != "" {
					el.ImageUrl = imgSrc
				}
			}

			if el.Name != "" && !isDuplicate(allElements, el.Name) {
				allElements = append(allElements, el)
			}
		})
	}

	for i := range allElements {
		err := ScrapElementDetails(&allElements[i])
		if err != nil {
			log.Printf("Error scraping element details for %s: %v", allElements[i].Name, err)
		}
	}


	if len(allElements) < 10 {
		log.Printf("Warning: Only scraped %d elements, might be incomplete", len(allElements))
	}

	return allElements, nil
}

func isDuplicate(elements []Element, name string) bool {
	for _, el := range elements {
		if el.Name == name {
			return true
		}
	}
	return false
}

func ScrapElementDetails(element *Element) error {
	resp, err := http.Get(element.Link)
	if err != nil {
		log.Printf("Error fetching element details: %v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("Status code error: %d %s", resp.StatusCode, resp.Status)
		return fmt.Errorf("failed to fetch details for %s", element.Name)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	categoryEl := doc.Find("#mw-content-text > div.mw-content-ltr.mw-parser-output > aside > div > div > ul > li > a")
	if categoryEl.Length() > 0 {
		element.Category = categoryEl.Text()
	}

	imgEl := doc.Find("#mw-content-text > div.mw-content-ltr.mw-parser-output > aside > figure > a > img")
	if imgEl.Length() > 0 {
		imgSrc, exists := imgEl.Attr("src")
		if exists && imgSrc != "" {
			if strings.HasPrefix(imgSrc, "/") {
				element.ImageUrl = "https://static.wikia.nocookie.net" + imgSrc
			} else {
				element.ImageUrl = imgSrc
			}
		}
	}

	return nil
}