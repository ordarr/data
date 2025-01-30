package calibre

func AllAuthors() *[]Author {
	var author []Author
	Calibre.Model(&Author{}).Find(&author)
	return &author
}

func FindByID(id uint) *Author {
	var author Author
	Calibre.Model(&Author{ID: id}).First(&author)
	return &author
}

func FindByName(name string) *Author {
	var author Author
	Calibre.Model(&Author{}).Where("name = ?", name).First(&author)
	return &author
}

//func GetBooksForAuthor(id core.Ids) *[]Book {
//	var author Author
//	var stats []koreader.Book
//	calibre.Calibre.Preload("Books").Preload("Books.Series").Preload("Books.Identifiers").Model(&calibre.Author{ID: id.Calibre}).First(&author)
//	var sb strings.Builder
//	sb.WriteString("authors like '%")
//	sb.WriteString(author.Name)
//	sb.WriteString("%'")
//	koreader.Calibre.Preload("StatsPages").Model(&koreader.Book{}).Where(sb.String()).Find(&stats)
//	var statsMap = make(map[string]*koreader.Book)
//	for k := range stats {
//		statsMap[stats[k].Title] = &stats[k]
//	}
//	i := len(author.Books)
//	var books = make([]Book, i)
//	for j := range author.Books {
//		calibreBook := author.Books[j]
//		koreaderBook := statsMap[calibreBook.Title]
//		idents := calibreBook.Identifiers
//		ids := make(map[string]string)
//		for x := range idents {
//			ids[idents[x].Type] = idents[x].Val
//		}
//		book := Book{
//			ID: BookId{
//				Calibre: calibreBook.ID,
//			},
//			Title:       calibreBook.Title,
//			Identifiers: ids,
//			Authors: []Author{{
//				ID:   author.ID,
//				Name: author.Name,
//			}},
//		}
//		if len(calibreBook.Series) > 0 {
//			book.Series.Name = calibreBook.Series[0].Name
//			book.Series.Number = calibreBook.SeriesIndex
//		}
//		if koreaderBook != nil {
//			book.Pages = koreaderBook.Pages
//			book.ID.Koreader = koreaderBook.ID
//			book.Stats = Stats{
//				LastOpen:       koreaderBook.LastOpen,
//				TotalReadTime:  koreaderBook.TotalReadTime,
//				TotalReadPages: koreaderBook.TotalReadPages,
//				Percentage:     (100 * koreaderBook.TotalReadPages) / koreaderBook.Pages,
//			}
//		}
//		books[j] = book
//	}
//	return &books
//}
