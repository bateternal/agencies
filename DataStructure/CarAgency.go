package DataStructure

import (
	"fmt"
)

type Service struct {
	ServiceName          string
	CarModel             string
	Description          string
	TechnicalDescription string
	Expense              string
	next                 *Service
	dNext                *Service
	sizeFather           int
}

type Agency struct {
	AgencyName string
	services   []*Service
	order0     []*Service
	order1     []*Service
	order2     []*Service
	next       *Agency
}

type LinkListAgency struct {
	head *Agency
}

type LinkListService struct {
	head *Service
	size int
}

func (ll *LinkListService) RemoveAt(i int){
	if i < 0 || i > ll.size {
		fmt.Errorf("Index out of bounds")
	}
	node := ll.head
	if i == 0{
		ll.head = node.next
	}else {
		j := 0
		for j < i-1 {
			j++
			node = node.next
		}
		remove := node.next
		node.next = remove.next
	}
	ll.size--
	fmt.Println("deleted")
}

// IndexOf returns the position of the Item t
func (ll *LinkListService) IndexOf(ServiceName string) int {
	node := ll.head

	j := 0
	for {
		if node.ServiceName == ServiceName {
			return j
		}
		if node.next == nil {
			return -1
		}
		node = node.next
		j++
	}
}


func remove(s []*Service, i *Service) []*Service {
	var j int
	f := len(s)
	for j = 0; j < len(s); j++ {
		if s[j] == i {
			f = j
		}
	}
	if f == len(s) {
		fmt.Println("service not found!")
		return s
	}
	s[f] = s[len(s)-1]
	return s[:len(s)-1]
}

func (ll *LinkListAgency) AddAgency(agency *Agency) {
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
}

func (ll *LinkListService) AddService(service *Service) {
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
}

func (s *Service) AddSubService(service *Service) {
	if s.dNext == nil {
		s.dNext = service
	} else {
		last := s.dNext
		for {
			if last.dNext == nil {
				break
			}
			last = last.dNext
		}
		last.dNext = service
	}
}

func (a *Agency) AddOffer(service *Service) {
	if service.sizeFather == 0{
		service.sizeFather = 1
	}else {
		service.sizeFather++
	}
	a.services = append(a.services, service)
}

func (a *Agency) Delete(service *Service,ll *LinkListService) {
	a.services = remove(a.services, service)
	service.sizeFather--
	if service.sizeFather == 0{
		var index int
		index = ll.IndexOf(service.ServiceName)
		ll.RemoveAt(index)
		head := service
		service = nil
		for{
			if head.dNext == nil{
				break
			}
			headI := head.dNext
			head.dNext = nil
			head = headI
		}
	}
}

func (ll *LinkListAgency) ListAgencies() {
	head := ll.head
	if head == nil {
		fmt.Println("This list is empty!")
	} else {
		for {
			fmt.Print(head.AgencyName + " , ")
			if head.next == nil {
				break
			}
			head = head.next
		}
	}
}

func (ll *LinkListService) ListServices() {
	head := ll.head
	if head == nil {
		fmt.Println("This list is empty!")
	} else {
		for {
			fmt.Print(head.ServiceName + "[ ")
			if head.dNext != nil {
				headP := head.dNext
				for {
					fmt.Print(headP.ServiceName + " ,")
					if headP.dNext == nil {
						break
					}
					headP = headP.dNext
				}
			}
			fmt.Print("]\n")
			if head.next == nil{
				break
			}
			head = head.next
		}
	}
}

func (s *Service) ListSubServices() {
	head := s.dNext
	if head == nil {
		fmt.Println("not exist!")
	} else {
		for {
			fmt.Print(head.ServiceName + " ,")
			if head.dNext == nil {
				break
			}
			head = head.dNext
		}
	}
	fmt.Println()
}

func (a *Agency) Order(s *Service,Lvl int) {
	switch Lvl {
	case 0:
		a.order0 = append(a.order0, s)
	case 1:
		a.order1 = append(a.order1, s)
	case 2:
		a.order2 = append(a.order2, s)
	}
}

func (a *Agency) ListOrder(){
	i:=0
	for ; i<len(a.order0);i++{
		fmt.Println(a.order0[i].ServiceName)
	}
	i=0
	for ; i<len(a.order1);i++{
		fmt.Println(a.order1[i].ServiceName)
	}
	i=0
	for ; i<len(a.order2);i++{
		fmt.Println(a.order2[i].ServiceName)
	}
}

func (ll *LinkListService) Search(serviceName string) *Service{
	head := ll.head
	for {
		if head.ServiceName == serviceName{
			return head
		}
		if head.next == nil{
			break
		}
		head = head.next
	}
	 fmt.Println("This service does not exist!")
	 return nil
}

func (ll *LinkListAgency) Search(agencyName string) *Agency{
	head := ll.head
	if head == nil{
		fmt.Println("no service already exist!")
		return nil
	}
	for {
		if head.AgencyName == agencyName{
			return head
		}
		if head.next == nil{
			break
		}
		head = head.next
	}
	fmt.Println("This service does not exist!")
	return nil
}