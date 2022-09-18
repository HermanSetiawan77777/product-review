package review

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	datareview "github.com/HermanSetiawan77777/JakMall/datasource/review"
	response "github.com/HermanSetiawan77777/JakMall/internal/httpserver/response"
	serviceprod "github.com/HermanSetiawan77777/JakMall/internal/service/product"
)

func GetData() []datareview.Review {
	file, err := ioutil.ReadFile("../datasource/review/reviews.json")
	if err != nil {
		fmt.Println("Error when opening file: ", err)
	}
	var payload []datareview.Review
	err = json.Unmarshal(file, &payload)
	if err != nil {
		fmt.Println("Error during Unmarshal file: ", err)
	}
	return payload
}

func CalculatedSummary(name string) response.TotalSummary {
	data := GetData()
	var payload response.TotalSummary
	tempmap := map[int]int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
	}
	payload.TotalReviews = 0
	for _, v := range data {
		if name != "" && v.ProductId == serviceprod.GetDetail(name).Id {
			tempmap[v.Rating] = tempmap[v.Rating] + 1
			payload.TotalReviews = payload.TotalReviews + 1
		}
		if name == "" {
			tempmap[v.Rating] = tempmap[v.Rating] + 1
			payload.TotalReviews = payload.TotalReviews + 1
		}

	}
	payload.OneStar = tempmap[1]
	payload.TwoStar = tempmap[2]
	payload.ThreeStar = tempmap[3]
	payload.FourStar = tempmap[4]
	payload.FiveStar = tempmap[5]
	payload.AverageRatings = float64(payload.OneStar+(2*payload.TwoStar)+(3*payload.ThreeStar)+
		(4*payload.FourStar)+(5*payload.FiveStar)) / float64(payload.TotalReviews)
	return payload
}
