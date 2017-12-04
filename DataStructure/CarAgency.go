package DataStructure

import (
	"sync"
)

type list interface{
	createList()
	Append()
	Insert()
	IndexOf()
	IsEmpty()
	String()
	Head()
	Size()
}


type Service struct {
	ServiceName          string
	CarModel             string
	Description          string
	TechnicalDescription string
	Expense              string
	next                 *Service
	downFirst            *Service
	sizeDown             int
	sizeFather           int
}


type Agency struct {
	AgencyName string
	services   []*Service
	next       *Agency
}

type LinkListAgency struct{
	head *Agency
	size int
	lock sync.RWMutex
}

type LinkListService struct{
	head *Service
	size int
	lock sync.RWMutex
}


func (ll *LinkListAgency) Append(Name string){
	ll.lock.Lock()
	agency := Agency{Name,nil,nil}
	if ll.head == nil {
		ll.head = &agency
	} else {
		last := ll.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = &agency
	}
	ll.size++
	ll.lock.Unlock()
}

func (ll LinkListService) Append(serviceName string,carModel string,description string,technicalDescription string,expense string){
	ll.lock.Lock()
	service := Service{serviceName,carModel,description,technicalDescription,expense,nil,nil,nil,nil}
	if ll.head == nil {
		ll.head = &service
	} else {
		last := ll.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = &service
	}
	ll.size++
	ll.lock.Unlock()
}

