package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

func wasWichtiges() {
	lock.Lock()
	defer lock.Unlock()
	//auf ressourcen zugreifen
	fmt.Println("Arbeit an gesperrter Ressource")
}

func main() {

	wasWichtiges()
	lock.Lock()
	fmt.Println("am Ende")

}
