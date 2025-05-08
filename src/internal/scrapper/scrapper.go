package scrapper

import (
	"Tubes2_BE_FireBoyWaterGirl/src/types"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func ScrapRecipe() *types.RecipeGraph{
	
	result := types.NewRecipeGraph()

	url := "https://little-alchemy.fandom.com/wiki/Elements_(Little_Alchemy_2)"
    res, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
		return result
    }
    defer res.Body.Close()

    if res.StatusCode != 200 {
        log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
    }

    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        log.Fatal(err)
    }

	doc.Find("table tbody tr").EachWithBreak(func(i int, s *goquery.Selection) bool {

		if (i < 2){
			return true
		}

		tds := s.Find("td")
		if tds.Length() < 2 {
			return true
		}

		element := tds.Eq(0).Find("a").Text()

		if(element == "Time" || element == "Archeologist" || element == "Ruins "){
			return true
		}
		
		tds.Eq(1).Find("li").Each(func(i int, li *goquery.Selection) {
			ingredients := []string{}
	
			li.Find("a").Each(func(j int, a *goquery.Selection) {
				text := a.Text()
				if (text == ""){
					return
				}
				
				ingredients = append(ingredients, a.Text())
			})
			if (ingredients[0] == "Time" || ingredients[0] == "Ruins" || ingredients[1] == "Time" ||  ingredients[1] == "Ruins"){

				return
			}
			result.AddRecipe(element, ingredients[0], ingredients[1])
		
		})
	
		return true
	})

	return result
}