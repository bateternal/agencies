package main

import (
	"./DataStructure"
)

func main() {
	ll := DataStructure.LinkListAgency{}
	a := DataStructure.Agency{AgencyName:"Car"}
	b := DataStructure.Agency{AgencyName:"def"}
	ll.AddAgency(&a)
	ll.AddAgency(&b)
	l := DataStructure.LinkListService{}
	c := DataStructure.Service{ServiceName:"c"}
	d := DataStructure.Service{ServiceName:"d"}
	e := DataStructure.Service{ServiceName:"e"}
	h := DataStructure.Service{ServiceName:"h"}
	i := DataStructure.Service{ServiceName:"i"}
	l.AddService(&c)
	//d.AddSubService(&e)
	l.AddService(&d)
	l.AddService(&h)
	d.AddSubService(&e)
	d.AddSubService(&i)
	l.ListServices()
}

