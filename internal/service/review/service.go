package review

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	datareview "github.com/HermanSetiawan77777/JakMall/datasource/review"
	response "github.com/HermanSetiawan77777/JakMall/internal/httpserver/response"
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

func CalculatedSummary() response.TotalSummary {
	data := GetData()
	var payload response.TotalSummary
	payload.TotalReviews = len(data)
	payload.OneStar = 0
	payload.TwoStar = 0
	payload.ThreeStar = 0
	payload.FourStar = 0
	payload.FiveStar = 0

	for _, v := range data {
		if v.Rating == 1 {
			payload.OneStar = payload.OneStar + 1
		}
		if v.Rating == 2 {
			payload.TwoStar = payload.OneStar + 1
		}
		if v.Rating == 3 {
			payload.ThreeStar = payload.OneStar + 1
		}
		if v.Rating == 4 {
			payload.FourStar = payload.OneStar + 1
		}
		if v.Rating == 5 {
			payload.FiveStar = payload.OneStar + 1
		}
	}
	payload.AverageRatings = float64(payload.OneStar+(2*payload.TwoStar)+(3*payload.ThreeStar)+
		(4*payload.FourStar)+(5*payload.FiveStar)) / float64(payload.TotalReviews)
	return payload
}
