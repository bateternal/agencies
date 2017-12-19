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
	head                 *Service
	sizeFather           int
}

type Agency struct {
	AgencyName string
	services   []*Service
	order      *MaxHeap
	next       *Agency
}

type LinkListAgency struct {
	head *Agency
}

type LinkListService struct {
	head *Service
	size int
}

func (ll *LinkListService) RemoveAt(i int) {
	if i < 0 || i > ll.size {
		fmt.Errorf("INDEX out of bounds")
	}
	node := ll.head
	if i == 0 {
		ll.head = node.next
	} else {
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

// IndexOf returns the position of the Order t
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
	if s.head == nil {
		s.head = service
	} else {
		last := s.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = service
	}
}

func (a *Agency) AddOffer(service *Service) {
	if service.sizeFather == 0 {
		service.sizeFather = 1
	} else {
		service.sizeFather++
	}
	a.services = append(a.services, service)
}

func (a *Agency) Delete(service *Service, ll *LinkListService) {
	a.services = remove(a.services, service)
	service.sizeFather--
	if service.sizeFather == 0 {
		var index int
		index = ll.IndexOf(service.ServiceName)
		ll.RemoveAt(index)
		service.head = nil
		service = nil
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
	fmt.Println()
}

func (ll *LinkListService) ListServices() {
	head := ll.head
	if head == nil {
		fmt.Println("This list is empty!")
	} else {
		for {
			fmt.Println(head.ServiceName + "  ")
			if head.head != nil {
				fmt.Println("[")
				head.ListSubServices()
				fmt.Println("]")
			}
			if head.next == nil {
				break
			}
			head = head.next
		}
	}
}

func (s *Service) ListSubServices() {
	head := s.head
	if head == nil {
		return
	} else {
		for {
			fmt.Println(head.ServiceName + "  ")
			if head.head != nil {
				fmt.Println("[")
				head.ListSubServices()
				fmt.Println("]")
			}
			if head.next == nil {
				break
			}
			head = head.next
		}
	}
}

func (a *Agency) Order(s *Service, Lvl int,c string) {
	if a.order == nil{
		order := MaxHeap{}
		a.order = &order
	}
	mHeap := a.order
	i :=  Lvl*1000 - (len(mHeap.Orders) + 1)
	order := Order{s,i,c}
	mHeap.Insert(order)
}

func (a *Agency) ListOrder() {
	mHeap := a.order
	if mHeap.Size() == 0 {
		fmt.Println("ListOrder is empty")
		return
	}
	j := 1
	for{
		s := mHeap.ExctractMax()
		fmt.Print(j)
		fmt.Println(": " + s.action.ServiceName)
		j++
		if mHeap.Size() == 0{
			break
		}
	}
}


func (ll *LinkListService) Search(serviceName string) *Service {
	head := ll.head
	if head == nil {
		return nil
	}
	for {
		if head.ServiceName == serviceName {
			return head
		}
		if head.head != nil{
			ss := head.Search(serviceName)
			if ss != nil{
				return ss
			}
		}
		if head.next == nil {
			break
		}
		head = head.next
	}
	return nil
}

func (s *Service) Search(serviceName string) *Service {
	head := s.head
	for {
		if head.ServiceName == serviceName {
			return head
		}
		if head.head != nil{
			ss := head.Search(serviceName)
			if ss != nil{
				return ss
			}
		}
		if head.next == nil {
			break
		}
		head = head.next
	}
	return nil
}

func (ll *LinkListAgency) Search(agencyName string) *Agency {
	head := ll.head
	if head == nil {
		return nil
	}
	for {
		if head.AgencyName == agencyName {
			return head
		}
		if head.next == nil {
			break
		}
		head = head.next
	}
	return nil
}
