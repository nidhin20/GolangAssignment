package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

//Stack interface
type Stack interface {
	push(int)
	pop()
	Printstack()
}

func (Link *linkedlist) push(pushvalue int) {
	existinglink := &linkedlist{Link.Value, Link.prev}
	mutex.Lock()
	Link.Value = pushvalue
	Link.prev = existinglink
	mutex.Unlock()
}

func (Link *linkedlist) pop() {
	mutex.Lock()
	if Link.prev != nil {
		fmt.Printf("Consumer pop element: %v \n", Link.Value)
		newlink := Link.prev

		Link.Value = newlink.Value
		Link.prev = newlink.prev

	}
	mutex.Unlock()
}

func init() {
	rand.Seed(1)
}

func push(pushvalue int, top linkedlist) linkedlist {
	newlink := linkedlist{pushvalue, &top}
	return newlink

}

var mutex = &sync.Mutex{}
var wg sync.WaitGroup

type linkedlist struct {
	Value int
	prev  *linkedlist
}

//ProducerLinked : producer function
func ProducerLinked(link *linkedlist, no int) {

	for i := 0; i < 30; i++ {

		Produceditem := rand.Intn(100)
		fmt.Printf("%v Produced sequence %v : %v \n", no, i+1, Produceditem)
		link.push(Produceditem)
		//link.Printstack()
	}
	link.Printstack()
	wg.Done()

}

//ConsumerLinked : producer function
func ConsumerLinked(link *linkedlist, no int) {

	for i := 0; i < 10; i++ {
		link.pop()
		//link.Printstack()
	}
	link.Printstack()
	wg.Done()

}

//Producer function
// func Producer(item chan linkedlist, signal chan bool) {
// 	defer wg.Done()
// 	Produceditem := strconv.Itoa(rand.Intn(100))
// 	var linkedar linkedlist
// 	if item != nil {
// 		linkedar = push(Produceditem, linkedar)
// 	} else {
// 		select {
// 		case linkedar = <-item:
// 			linkedar = push(Produceditem, linkedar)
//
// 		}
// 	}
// 	item <- linkedar
// }

//Printstack
func (Link *linkedlist) Printstack() {
	var outputstring string
	for p := Link; p.prev != nil; p = p.prev {
		outputstring = fmt.Sprintf("%v %s", p.Value, outputstring)
	}
	fmt.Println(outputstring)
}

// Consumer function
func Consumer(item chan []int) {
	for i := 0; i < 1; i++ {
		s := <-item
		slen := len(s)
		item <- s[:slen-1]
		fmt.Println("Consumer : " + strconv.Itoa(s[slen]))
	}

}
func main() {
	fmt.Print("Choose from following\n")
	fmt.Print("1. Linked List - Non Parallel\n")
	var selection, NoProducer, NoConsumer int
	fmt.Scan(&selection)
	fmt.Print("No. of Producer:")
	fmt.Scan(&NoProducer)
	fmt.Print("No. of Consumer:")
	fmt.Scan(&NoConsumer)

	switch selection {

	case 1:
		link := linkedlist{0, nil}
		for i := 0; i < NoProducer; i++ {
			wg.Add(1)
			go ProducerLinked(&link, i+1)

		}
		for i := 0; i < NoConsumer; i++ {
			for i := 0; i < 10; i++ {
				wg.Add(1)
				go ConsumerLinked(&link, i)
			}
		}
		wg.Wait()
		link.Printstack()

	}

	// Produceditembyproducer := make(chan linkedlist)
	// signal := make(chan bool)
	// for i := 0; i < 3; i++ {
	// 	wg.Add(1)
	// 	go Producer(Produceditembyproducer, signal)
	//
	// }

	//wg.Wait()

}
