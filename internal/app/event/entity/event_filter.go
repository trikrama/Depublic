package entity

type FilterTicketsRequest struct {
	Search      string `json:"search" param:"search"`
	Location    string `json:"location" param:"location"`
	Category    string `json:"category" param:"category"`
	StartTime   string `json:"startTime" param:"startTime"`
	EndTime     string `json:"endTime" param:"endTime"`
	MinPrice    string `json:"minPrice" param:"minPrice"`
	MaxPrice    string `json:"maxPrice" param:"maxPrice"`
	SortBy      string `json:"sortBy" param:"sortBy"`
}


func NewFilter(search, location, category, startTime, endTime, minPrice, maxPrice, sortBy string) *FilterTicketsRequest {
	return &FilterTicketsRequest{
		Search:      search,
		Location:    location,
		Category:    category,
		StartTime:   startTime,
		EndTime:     endTime,
		MinPrice:    minPrice,
		MaxPrice:    maxPrice,
		SortBy:      sortBy,
	}
}