package entity

// QueryFilter adalah struct untuk filter dan sorting
type QueryFilter struct {
	Sort   SortQuery
	Filter FilterQuery
	Search string
}

// SortQuery adalah struct untuk sorting
type SortQuery struct {
	By     string
	Order  string
}

// FilterQuery adalah struct untuk filter
type FilterQuery struct {
	Price         float64
	Location      string
	StartDate     string
	EndDate       string
	Status        string
}