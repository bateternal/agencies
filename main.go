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
		command,text := start()
		if len(command) == 3 && command[0] == "add" && command[1] == "service" {
			AddService(command[2][:len(command[2])-1])
		} else if len(command) == 5 && command[0] == "add" && command[1] == "subservice" && command[3] == "to" {
			AddSubService(command[4][:len(command[4])-1], command[2])
		}else if len(command) == 5 && command[0] == "add" && command[1] == "offer" && command[3] == "to" {
			AddOffer(command[2],command[4][:len(command[4])-1])
		}else if len(command) == 4 && command[0] == "delete" && command[2] == "from" {
			DeleteOffer(command[1],command[3][:len(command[3])-1])
		}else if len(command) == 3 && command[0] == "add" && command[1] == "agency" {
			AddAgency(command[2][:len(command[2])-1])
		}else if len(command) == 2 && command[0] == "list" && command[1][:len(command[1])-1] == "agencies" {
			ListAgencies()
		}else if len(command) == 2 && command[0] == "list" && command[1][:len(command[1])-1] == "services"{
			ListServices()
		}else if len(command) == 4  && command[0] == "list" && command[1] == "services" && command[2] == "from"{
			ListSubServices(command[3][:len(command[3])-1])
		}else if len(command) == 8 && command[0] == "order" && command[2] == "to" && command[4] == "by" && command[6] == "with"{
			AddOrder(command[1],command[3],command[5],command[7][:len(command[7])-1])
		}else if len(command) == 3 && command[0] == "list" && command[1] == "orders" {
			Execute(command[2][:len(command[2])-1])
		}else{
			error(text)
		}
	}
}

func start() ([]string,string) {
	maxHeap = *DataStructure.New()
	reader := bufio.NewReader(os.Stdin)
	//fmt.Println(reflect.TypeOf(reader))
	fmt.Print("root@ubuntu:~# ")
	text, _ := reader.ReadString('\n')
	text = string(text)
	com := strings.Split(text, " ")
	return com,text
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
	sub := linkListService.Search(subServiceName)
	if service != nil  && sub == nil{
		Service := DataStructure.Service{ServiceName: subServiceName}
		service.AddSubService(&Service)
	}else if sub != nil{
		fmt.Println(subServiceName + " already used!")
	}else{
		fmt.Println(serviceName , " not found!")
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

func DeleteOffer(serviceName string, agencyName string){
	service := linkListService.Search(serviceName)
	agency := linkListAgency.Search(agencyName)
	if service == nil && agency == nil{
		fmt.Println(agencyName + " and " + serviceName + " not found!")
	}else if service == nil{
		fmt.Println(serviceName + " not found!")
	}else if agency == nil{
		fmt.Println(agencyName + "not found!")
	}else{
		agency.Delete(service,&linkListService)
	}
}

func AddAgency(agencyName string){
	agency := DataStructure.Agency{AgencyName:agencyName}
	linkListAgency.AddAgency(&agency)
}

func ListAgencies(){
	linkListAgency.ListAgencies()
}

func ListServices(){
	linkListService.ListServices()
}

func ListSubServices(serviceName string){
	service := linkListService.Search(serviceName)
	if service == nil{
		fmt.Println(serviceName + " not found!")
	}else{
		service.ListSubServices()
	}
}

func AddOrder(serviceName string, agencyName string,customer string,lvl string) {
	service := linkListService.Search(serviceName)
	agency := linkListAgency.Search(agencyName)
	if service == nil && agency == nil{
		fmt.Println(agencyName + " and " + serviceName + " not found!")
	}else if service == nil{
		fmt.Println(serviceName + " not found!")
	}else if agency == nil{
		fmt.Println(agencyName + "not found!")
	}else {
		var level int
		if lvl == "1" {
			level = 1
		} else if lvl == "2" {
			level = 2
		} else if lvl == "3" {
			level = 3
		}
		agency.Order(service,level, customer)
	}
}

func Execute(agencyName string){
	agency := linkListAgency.Search(agencyName)
	if agency == nil{
		fmt.Println(agencyName + "not found!")
	}
	agency.ListOrder()
}
func error(text string){
	text = "'"+ text[:len(text)-1] + "' is not recognized as an internal or external command,\n operable program or batch file."

	fmt.Println(text)
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

