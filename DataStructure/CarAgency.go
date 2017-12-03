package DataStructure

import "sync"

type list interface{
	IndexOf()
	IsEmpty()
	Size()
	String()
	Head()
}


type service struct {
	CarName              string
	CarModel             string
	Description          string
	TechnicalDescription string
	Expense              string
	next                 *service
	downFirst            *service
	sizeDown             int
	sizeFather           int
}


type agency struct{
	services []*service
	next *agency
}

type linklist struct{
	head interface{}
	size int
	lock sync.RWMutex
}