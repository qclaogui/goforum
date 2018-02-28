package model

//Reply model
type Reply struct {
	Model
	ThreadID int
	UserID   int
	Body     string
}
