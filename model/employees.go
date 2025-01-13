package model

type Employee struct {
	ID         string  `json:"id" bson:"_id,omitempty"`
	Name       string  `json:"name" bson:"name"`
	Department string  `json:"department" bson:"department"`
	Status     string  `json:"status" bson:"status"`
	Salary     float64 `json:"salary" bson:"salary"`
}
