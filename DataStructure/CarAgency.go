package DataStructure

import (
	"sync"
	"fmt"
)

type Service struct {
	ServiceName          string
	CarModel             string
	Description          string
	TechnicalDescription string
	Expense              string
	next                 *Service
	dNext            *Service
	sizeDown             int
	sizeFather           int
	lock sync.RWMutex
}


type Agency struct {
	AgencyName string
	services   []*Service
	next       *Agency
	lock sync.RWMutex
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


func remove(s []*Service, i *Service) []*Service {
	var j int
	f := len(s)
	for j=0 ;j<len(s);j++{
		if s[j] == i {
			f = j
		}
	}
	if f == len(s){
		fmt.Println("service not found!")
		return s
	}
	s[f] = s[len(s)-1]
	return s[:len(s)-1]
}


func (ll *LinkListAgency) AddAgency(agency *Agency){
	ll.lock.Lock()

	if ll.head == nil {
		ll.head = agency
	} else {
		last := ll.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = agency
	}
	ll.size++
	ll.lock.Unlock()
}

func (ll *LinkListService) AddService(service *Service){
	ll.lock.Lock()

	if ll.head == nil {
		ll.head = service
	} else {
		last := ll.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = service
	}
	ll.size++
	ll.lock.Unlock()
}

func (s *Service) AddSubService(service *Service){
	s.lock.Lock()
	if s.dNext == nil{
		s.dNext = service
	}else{
		last := s.dNext
		for{
			if last.dNext == nil{
				break
			}
			last = last.dNext
		}
		last.dNext = service
	}
	s.sizeDown++
	s.lock.Lock()
}

func (a *Agency) AddOffer(service *Service) *Agency{
	a.services = append(a.services,service)
	return a
}

func (a *Agency) Delete(service *Service) *Agency{
	a.services = remove( a.services , service)
	return a
}



