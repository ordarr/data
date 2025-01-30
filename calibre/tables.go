package calibre

type Identifier struct {
	ID   uint `gorm:"primary_key"`
	Book uint
	Type string
	Val  string
}

func (Identifier) TableName() string {
	return "identifiers"
}

type BookAuthorLink struct {
	Book   uint
	Author uint
}

func (BookAuthorLink) TableName() string {
	return "books_authors_link"
}

type Series struct {
	ID   uint `gorm:"primary_key"`
	Name string
}

func (Series) TableName() string {
	return "series"
}

type BookSeriesLink struct {
	Book   uint
	Series uint
}

func (BookSeriesLink) TableName() string {
	return "books_series_link"
}

type Book struct {
	ID          uint `gorm:"primary_key"`
	Title       string
	SeriesIndex float32
	Authors     []Author     `gorm:"many2many:books_authors_link;foreignKey:ID;joinForeignKey:Book;References:ID;joinReferences:Author"`
	Series      []Series     `gorm:"many2many:books_series_link;foreignKey:ID;joinForeignKey:Book;References:ID;joinReferences:Series"`
	Identifiers []Identifier `gorm:"foreignKey:Book"`
}

func (Book) TableName() string {
	return "books"
}

type Author struct {
	ID    uint `gorm:"primary_key"`
	Name  string
	Books []Book `gorm:"many2many:books_authors_link;foreignKey:ID;joinForeignKey:Author;References:ID;joinReferences:Book"`
}

func (Author) TableName() string {
	return "authors"
}
