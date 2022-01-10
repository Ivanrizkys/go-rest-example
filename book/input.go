package book

type Input struct {
	Title string
	Price int
	// * menangkap json dengan underscore
	SubTittle string `json:"sub_title"`
}

type Validator struct {
	Title string      `json:"title" binding:"required"`
	Price interface{} `json:"price" binding:"required,number"`
	Email string      `json:"email" binding:"required"`
}
