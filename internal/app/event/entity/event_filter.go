package entity

// QueryFilter adalah struct untuk filter dan sorting
type QueryFilter struct {
	Sort   SortQuery
	Filter FilterQuery
	Search string `query:"search"`
}

// SortQuery adalah struct untuk sorting
type SortQuery struct {
	By     string 	`query:"sort_by"`
	Order  string 	`query:"order"`
}

// FilterQuery adalah struct untuk filter
type FilterQuery struct {
	Price         float64 	`query:"price"`
	Location      string	`query:"location"`
	StartDate     string	`query:"start_date"`
	EndDate       string	`query:"end_date"`
	Status        string	`query:"status"`
}