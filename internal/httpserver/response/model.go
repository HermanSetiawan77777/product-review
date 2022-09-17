package response

type TotalSummary struct {
	TotalReviews   int     `json:"total_reviews"`
	AverageRatings float64 `json:"average_ratings"`
	FiveStar       int     `json:"5_star"`
	FourStar       int     `json:"4_star"`
	ThreeStar      int     `json:"3_star"`
	TwoStar        int     `json:"2_star"`
	OneStar        int     `json:"1_star"`
}
