package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}
var wg sync.WaitGroup

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

func (Link *linkedlist) pop(no int) {
	defer mutex.Unlock()
	mutex.Lock()
	if Link.prev != nil {

		fmt.Printf("%v Consumer pop element : %v \n", no, Link.Value)
		newlink := Link.prev

		//Link = Link.prev

		Link.Value = newlink.Value
		Link.prev = newlink.prev

	} else {
		fmt.Printf("%v Consumer : Stack is empty \n", no)
	}

}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

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
		link.pop(no)
		//link.Printstack()
		link.Printstack()
	}

	wg.Done()

}

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
	starttime := time.Now()
	fmt.Printf("Linked list started at: %v", starttime)
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
				go ConsumerLinked(&link, i+1)
			}
		}
		wg.Wait()
		fmt.Printf("Final stack")
		link.Printstack()
		fmt.Printf("Time taken to complete : %v \n", time.Since(starttime))

	}

}
