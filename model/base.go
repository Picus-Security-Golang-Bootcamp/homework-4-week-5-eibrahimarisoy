package model

type Args struct {
	Sort   string
	Order  string
	Offset string
	Limit  string
	Search string
}

type Data struct {
	TotalData    int64
	FilteredData int64
	Data         []Book
}
