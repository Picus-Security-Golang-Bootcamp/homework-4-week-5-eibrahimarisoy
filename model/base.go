package model

// for getting query parameters
type Args struct {
	Sort   string
	Order  string
	Offset string
	Limit  string
	Search string
}

// for pagination response
type Data struct {
	FilteredData int64
	Limit        int64
	Offset       int64
	Data         []Book
}

// for buy controller
type Quantity struct {
	Amount uint `json:"amount"`
}
