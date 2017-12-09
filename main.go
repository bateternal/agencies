package main

import (
	"./DataStructure"

	"bufio"
	"os"
	"fmt"
	"strings"
	_ "go/types"
	"reflect"
)


var linkListService DataStructure.LinkListService
var linkListAgency DataStructure.LinkListAgency
func main() {
	command := start()
	select {
	case len(command)==3 && command[0]=="add" && command[1] == "service":
		AddService(command[2])
	}
}

func start() []string {
	linkListService = DataStructure.LinkListService{}
	linkListAgency = DataStructure.LinkListAgency{}
	reader := bufio.NewReader(os.Stdin)
	//fmt.Println(reflect.TypeOf(reader))
	fmt.Print("Enter Command: ")
	text, _ := reader.ReadString('\n')
	text = string(text)
	com := strings.Split(text, "")
	return com
}

func AddService (serviceName string){
	s := DataStructure.Service{ServiceName:serviceName}
	linkListService.AddService(&s)
}

func AddSubservice (serviceName string,subServiceName string){
	service := linkListService.Search(serviceName)
	Service := DataStructure.Service{ServiceName:subServiceName}
	service.AddSubService(&Service)
}

func error(){
	fmt.Println("invalid command!")
}


//func main() {
//	ll := DataStructure.LinkListAgency{}
//	a := DataStructure.Agency{AgencyName:"Car"}
//	b := DataStructure.Agency{AgencyName:"def"}
//	ll.AddAgency(&a)
//	ll.AddAgency(&b)
//	l := DataStructure.LinkListService{}
//	c := DataStructure.Service{ServiceName:"c"}
//	d := DataStructure.Service{ServiceName:"d"}
//	e := DataStructure.Service{ServiceName:"e"}
//	h := DataStructure.Service{ServiceName:"h"}
//	i := DataStructure.Service{ServiceName:"i"}
//
//
//	l.AddService(&c)
//	l.AddService(&d)
//	l.AddService(&h)
//	d.AddSubService(&e)
//	d.AddSubService(&i)
//	a.AddOffer(&c)
//	a.Delete(&c,&l)
//	l.ListServices()
//
//	a.Order(&e,0)
//	a.Order(&h,1)
//	a.Order(&d,0)
//}

