package request

type Employee struct {
	Count int `json:"count" binding:"min=1"`
}
