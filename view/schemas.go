package view

type shop struct {
	ID              uint64 `json:"id"`
	ShopName        string `json:"shop_name"`
	ShopDescription string `json:"shop_description"`
}

type bookWithRank struct {
	ID              uint64  `json:"id"`
	BookName        string  `json:"book_name"`
	BookDescription string  `json:"book_description"`
	Sales           uint    `json:"sales"`
	Rank            *string `json:"rank"`
}
