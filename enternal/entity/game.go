package entity

type Book struct {
	Id           int
	Name         string
	FirstElement int
}

type Element struct {
	Id   int
	Desc string
	Type int
	Next map[int]string
}
