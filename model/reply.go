package model

type Reply struct {
	Model
	ThreadId int
	UserId   int
	Body     string
}
