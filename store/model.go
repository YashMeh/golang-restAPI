package store

//Book struct
type Book struct {
	ID    int     `json:"id"`
	Isbn  string  `json:"isbn"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

// //Author struct
// type Author struct {
// 	FirstName string `json:"firstname"`
// 	LastName  string `json:"lastname"`
// }

//Books is ...
// var Books []Book
