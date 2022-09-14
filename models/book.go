package models

import "github.com/mitchellh/mapstructure"

type Book struct {
	ID          int    `json:"id"`
	AuthorID    int    `json:"author_id"`
	Title       string `json:"title"`
	CoverImage  string `json:"cover_image"`
	Pages       int    `json:"pages"`
	ReleaseDate string `json:"releaseDate"`
	Isbn        string `json:"isbn"`
}

type Books []Book

var books Books

func InitBookData() *Books {
	books = make(Books, 8)
	bookData := []interface{}{
		map[string]interface{}{
			"id":          1,
			"author_id":   1,
			"title":       "Oliver Twist",
			"cover_image": "",
			"pages":       234,
			"releaseDate": "1839",
			"isbn":        "wt345",
		},
		map[string]interface{}{
			"id":          2,
			"author_id":   1,
			"title":       "Hard Times",
			"cover_image": "",
			"pages":       300,
			"releaseDate": "1854",
			"isbn":        "jk654",
		},
		map[string]interface{}{
			"id":          3,
			"author_id":   3,
			"title":       "Hamlet",
			"cover_image": "",
			"pages":       160,
			"releaseDate": "1603",
			"isbn":        "po789",
		},
		map[string]interface{}{
			"id":          4,
			"author_id":   2,
			"title":       "IT",
			"cover_image": "",
			"pages":       500,
			"releaseDate": "2017",
			"isbn":        "yu098",
		},
		map[string]interface{}{
			"id":          5,
			"author_id":   4,
			"title":       "Norwegian Wood",
			"cover_image": "http://t1.gstatic.com/images?q=tbn:ANd9GcQvJJDi2mzwg9v_PlmKYL31gXIz0kvAObJak7DVFPcD_mJTIyec",
			"pages":       296,
			"releaseDate": "1987",
			"isbn":        "hj846",
		},
		map[string]interface{}{
			"id":          6,
			"author_id":   4,
			"title":       "Kafka on the Shore",
			"cover_image": "http://t0.gstatic.com/images?q=tbn:ANd9GcRHFU_j93PPsbQGqoaZJnHH6-Emk_sIxG823SkoRTL0nvdEP41f",
			"pages":       505,
			"releaseDate": "2002",
			"isbn":        "op012",
		},
		map[string]interface{}{
			"id":          7,
			"author_id":   4,
			"title":       "After Dark",
			"cover_image": "http://t3.gstatic.com/images?q=tbn:ANd9GcQBMNA8A19vQpNY4bkgadsLhiRUFqBKwKAA6ANrw8VEtOiPrYQJ",
			"pages":       208,
			"releaseDate": "2004",
			"isbn":        "cv456",
		},
		map[string]interface{}{
			"id":          8,
			"author_id":   4,
			"title":       "1Q84",
			"cover_image": "http://t0.gstatic.com/images?q=tbn:ANd9GcTBQZSg-b2LFkLlt9fWndS1w8SONabDZBHf0dtdb3-bqcuKxduL",
			"pages":       928,
			"releaseDate": "2009",
			"isbn":        "al207",
		},
	}

	mapstructure.Decode(bookData, &books)

	return &books
}

func GetSingleBook(id int) Book {
	var book Book
	for _, v := range books {
		if v.ID != id {
			continue
		}
		book = v
	}
	return book
}

func AddBook(authorId, pages int, title, coverImage, releaseDate, isbn string) Book {
	book := Book{
		ID:          len(books) + 1,
		AuthorID:    authorId,
		Title:       title,
		CoverImage:  coverImage,
		Pages:       pages,
		ReleaseDate: releaseDate,
		Isbn:        isbn,
	}
	return book
}

func UpdateBook(id, authorId, pages int, title, coverImage, releaseDate, isbn string) Book {
	book := Book{
		ID:          len(books) + 1,
		AuthorID:    authorId,
		Title:       title,
		CoverImage:  coverImage,
		Pages:       pages,
		ReleaseDate: releaseDate,
		Isbn:        isbn,
	}
	return book
}
