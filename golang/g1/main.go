package main

import (
	"fmt"
	"g1/src/simple"
	"g1/src/enum_example"
	"g1/src/complex"
	"io/ioutil"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)
func Panicking (err error){
	if err !=nil {
		fmt.Println("error : ",err)
	}
}

func main () {
	// sm := doSimple()

	// fmt.Println(" \nwriting : ")
	// writeToFile("simple.bin",sm)

	// fmt.Println("\nreading :  ")
	// sm2 := &simple.SimpleMessage{}
	// readFromFile("simple.bin",sm2)
	// fmt.Println("data : ",sm2)

	// jsonDemo(sm)

	// doEnum()

	doComplex()

}

func doComplex(){
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id: 11,
			Name: "sk shahriar ahmed raka",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
			Id: 23,
			Name: "what is what",
		},
		&complexpb.DummyMessage{
			Id: 24,
			Name: "who is who",
		},
		},
	}

	fmt.Println(cm)
}



func doEnum(){
	em := enumpb.EnumMessage{
		Id : 243,
		DayOfTheWeek: enumpb.DayOfTheWeek_FRIDAY,
	}

	fmt.Println("enum message : " ,em)

	em.DayOfTheWeek =enumpb.DayOfTheWeek_SATURDAY
	fmt.Println("enum message : " ,em)
}




func jsonDemo (sm proto.Message){
	smAsString := toJSON(sm)
	fmt.Println(smAsString)

	sm2:= &simple.SimpleMessage{}
	fromJSON(smAsString,sm2)
	fmt.Println("successfully created proto struct : ",sm2)
}


func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out,err := marshaler.MarshalToString(pb)
	Panicking(err)
	return out 
}

func fromJSON (in string , pb proto.Message){
	err := jsonpb.UnmarshalString(in,pb)
	Panicking(err)

}



func writeToFile(fName string, pb proto.Message){
	out,err := proto.Marshal(pb) 
	Panicking(err)
	err = ioutil.WriteFile(fName,out,0666)
	Panicking(err)
	fmt.Println("data has written !!!")
}
func readFromFile(fName string, pb proto.Message){
	in ,err := ioutil.ReadFile(fName)
	Panicking(err)

	err= proto.Unmarshal(in,pb)
	Panicking(err)
	fmt.Println(" data has been read ")
}





func doSimple() *simple.SimpleMessage{
	a:= simple.SimpleMessage{
		Id : 12 ,
		IsSimple : true ,
		Name : "raka",
		SampleList : []int32 {2,3,4,5},
	}
	// fmt.Println(a)
	// a.Name = "ssar"

	// fmt.Println(a)

	fmt.Println("GetId() : ",a.GetId())
	fmt.Println("GetIsSimple() : ",a.GetIsSimple())
	fmt.Println("GetName() : ",a.GetName())
	fmt.Println("GetSampleList() : ",a.GetSampleList())

	return &a 
}