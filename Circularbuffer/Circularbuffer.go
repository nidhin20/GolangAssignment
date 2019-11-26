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

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type linkedlist struct {
	Value int
	next  *linkedlist
}
type circularbuffer struct {
	front       *linkedlist
	rear        *linkedlist
	size        int
	currentsize int
}

func (curlink *circularbuffer) push(pushvalue int, outputstr string) {
	// defer mutex.Unlock()
	// mutex.Lock()
	if curlink.currentsize < curlink.size {
		curlink.currentsize = curlink.currentsize + 1
		newlinkeditem := &linkedlist{pushvalue, nil}
		if curlink.front == nil && curlink.rear == nil {
			curlink.front = newlinkeditem
			curlink.rear = newlinkeditem
			newlinkeditem.next = newlinkeditem
		} else {
			prelinkeditem := curlink.front
			prelinkeditem.next = newlinkeditem
			newlinkeditem.next = curlink.rear
			curlink.front = newlinkeditem
		}
		fmt.Println(outputstr)
	} else {
		curlink.front.next.Value = pushvalue
		curlink.front = curlink.front.next
		curlink.rear = curlink.rear.next
		fmt.Println(outputstr)
	}
}

func (curlink *circularbuffer) pop(no int, seqno int) {
	// defer mutex.Unlock()
	// mutex.Lock()
	if curlink.rear != nil {
		fmt.Printf("%v Consumer seq.no %v pop element: %v \n", no+1, seqno+1, curlink.rear.Value)
		curlink.front.next = curlink.rear.next
		curlink.rear = curlink.rear.next
		curlink.currentsize = curlink.currentsize - 1
		if curlink.currentsize == 0 {
			curlink.front = nil
			curlink.rear = nil
		}
	} else {
		fmt.Printf("%v Consumer seq.no %v pop element: Curcullar buffer is empty \n", no+1, seqno+1)
	}

}

//ProducerLinked : producer function
func ProducerLinked(link *circularbuffer, no int) {

	for i := 0; i < 2; i++ {

		Produceditem := rand.Intn(100)
		outputstr := fmt.Sprintf("%v Produced sequence %v : %v \n", no, i+1, Produceditem)
		link.push(Produceditem, outputstr)
		link.Printstack()
	}
	//link.rear.Printstack()
	wg.Done()

}

//ConsumerLinked : producer function
func ConsumerLinked(link *circularbuffer, no int) {

	for i := 0; i < 2; i++ {
		link.pop(no, i)
		//link.Printstack()
		//link.Printstack()
	}

	wg.Done()

}

//Printstack
func (Link *circularbuffer) Printstack() {
	var outputstring string

	if Link.front != nil && Link.rear != nil {
		p := Link.rear
		for i := 0; i < Link.currentsize; i++ {
			outputstring = fmt.Sprintf("%v %s", p.Value, outputstring)
			p = p.next
		}
		fmt.Println(outputstring)
	}
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
	var selection, sizeofbuffer, NoProducer, NoConsumer int
	fmt.Scan(&selection)
	fmt.Print("Size of buffer:")
	fmt.Scan(&sizeofbuffer)
	fmt.Print("No. of Producer:")
	fmt.Scan(&NoProducer)
	fmt.Print("No. of Consumer:")
	fmt.Scan(&NoConsumer)
	starttime := time.Now()
	switch selection {

	case 1:
		cirbuf := circularbuffer{front: nil, rear: nil, size: sizeofbuffer, currentsize: 0}
		//link := linkedlist{0, nil}
		for i := 0; i < NoProducer; i++ {
			wg.Add(1)
			go ProducerLinked(&cirbuf, i+1)

		}
		for i := 0; i < NoConsumer; i++ {
			for i := 0; i < 10; i++ {
				wg.Add(1)
				go ConsumerLinked(&cirbuf, i)
			}
		}
		wg.Wait()
		cirbuf.Printstack()
		fmt.Printf("\n Time taken to complete : %v", time.Since(starttime))

	}

}
