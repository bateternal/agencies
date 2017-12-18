package main

import (
	"./DataStructure"

	"bufio"
	"os"
	"fmt"
	"strings"
)


var linkListService DataStructure.LinkListService
var linkListAgency DataStructure.LinkListAgency
var maxHeap DataStructure.MaxHeap
func main() {
	linkListService = DataStructure.LinkListService{}
	linkListAgency = DataStructure.LinkListAgency{}
	for {
		command := start()
		if len(command) == 3 && command[0] == "add" && command[1] == "service" {
			AddService(command[2])
		} else if len(command) == 5 && command[0] == "add" && command[1] == "subservice" && command[3] == "to" {
			AddSubService(command[4], command[2])
		}else if len(command) == 5 && command[0] == "add" && command[1] == "offer" && command[3] == "to" {
			AddOffer(command[2],command[4])
		}else{
			error()
		}
	}
}

func start() []string {
	maxHeap = *DataStructure.New()
	reader := bufio.NewReader(os.Stdin)
	//fmt.Println(reflect.TypeOf(reader))
	fmt.Print("Enter Command: ")
	text, _ := reader.ReadString('\n')
	text = string(text)
	com := strings.Split(text, " ")
	return com
}

func AddService (serviceName string){
	s := DataStructure.Service{ServiceName:serviceName}
	service := linkListService.Search(serviceName)
	if service == nil {
		linkListService.AddService(&s)
	}else{
		fmt.Println("This ServiceName already used!")
	}
}

func AddSubService (serviceName string,subServiceName string){
	service := linkListService.Search(serviceName)
	if service != nil {
		Service := DataStructure.Service{ServiceName: subServiceName}
		service.AddSubService(&Service)
	}else{
		fmt.Println(serviceName + " not found!")
	}

}

func AddOffer(serviceName string,agencyName string){
	service := linkListService.Search(serviceName)
	agency := linkListAgency.Search(agencyName)
	if service == nil && agency == nil{
		fmt.Println(agencyName + " and " + serviceName + " not found!")
	}else if service == nil{
		fmt.Println(serviceName + " not found!")
	}else if agency == nil{
		fmt.Println(agencyName + "not found!")
	}else{
		agency.AddOffer(service)
	}
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

