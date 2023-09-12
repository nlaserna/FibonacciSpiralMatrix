package structs

type MatrixResponse struct {
	Timestamp int64   `json:"ts"`
	Rows      [][]int `json:"rows"`
}
