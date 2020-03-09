package view

import "gin-gorm-rails-like-sample-api/model/entity"

func convertToViewShops(before []*entity.Shop) []shop {
	after := make([]shop, len(before))
	for i, p := range before {
		after[i] = convertToViewShop(p)
	}
	return after
}

func convertToViewShop(before *entity.Shop) shop {
	return shop{
		ID:              before.ID,
		ShopName:        before.ShopName,
		ShopDescription: before.ShopDescription,
	}
}

func convertToViewBooks(before []*entity.Book) []book {
	after := make([]book, len(before))
	for i, p := range before {
		after[i] = convertToViewBook(p)
	}
	return after
}

func convertToViewBook(before *entity.Book) book {
	return book{
		ID:              before.ID,
		BookName:        before.BookName,
		BookDescription: before.BookDescription,
		Popularity:      before.Popularity,
	}
}
