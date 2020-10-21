package main

import "fmt"

type Tester struct {
	Name     string
	Alter    int
	Sonstwas string
}

type StringTester struct {
	Name     string
	Alter    int
	Sonstwas string
}

func (st *StringTester) String() string {
	return fmt.Sprintf("StringTester %v, Alter: %v", st.Name, st.Alter)
}

func main() {

	t := &Tester{Name: "Dingo", Alter: 3, Sonstwas: "Wert"}
	fmt.Printf("%+v\n", t)

	st := &StringTester{Name: "Dingo", Alter: 3, Sonstwas: "Wert"}
	fmt.Printf("%+v\n", st)

}
