package main

import "log"

func importantStuff() {
	log.Println("rein")
	defer log.Println("raus")

	log.Println("drin")

}

func main() {
	importantStuff()
}
