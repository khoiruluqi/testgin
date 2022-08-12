package domains

type Book struct {
	Id    			int `gorm:"primaryKey"`
	Title 			string
	Price 			int
	AuthorID 		int
  	//Author   		Author
}