package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

//Stack interface
type Stack interface {
	push(int)
	pop()
	Printstack()
}

func (binnode *binarynode) findelementbyvalue(value int, prevnode *binarynode) *binarynode {

	if binnode.value == value {
		return prevnode
	} else if binnode.value > value {

		return binnode.Left.findelementbyvalue(value, binnode)
	} else {
		return binnode.Right.findelementbyvalue(value, binnode)
	}
}
func (binnode *binarynode) findsecondlastlegtoswap(isleftleg bool) *binarynode {
	if isleftleg == true {
		if binnode.Left.Left != nil {
			return binnode.Left.findsecondlastlegtoswap(true)
		} else {
			return binnode
		}
	} else {
		if binnode.Right.Right != nil {
			return binnode.Right.findsecondlastlegtoswap(false)
		} else {
			return binnode
		}
	}
}
func (binnode *binarynode) pop() {
	if len(Listitem) > 0 {
		popelement := Listitem[len(Listitem)-1]
		fmt.Printf("Consumer comsumed : %v \n", popelement)
		Listitem = Listitem[:len(Listitem)-1]
		elementtopop := binnode.findelementbyvalue(popelement, binnode)
		if elementtopop.Left != nil {
			if elementtopop.Left.value == popelement {
				elementtopop.Left = nil
			}
		}
		if elementtopop.Right != nil {
			if elementtopop.Right.value == popelement {
				elementtopop.Right = nil
			}
		}
	}

}
func (binnode *binarynode) push(value int) {
	if binnode.value == value {
		fmt.Printf("Same number exist in binary tree : %v \n", value)
	} else if binnode.value > value {
		if binnode.Left == nil {
			binnode.Left = &binarynode{value, nil, nil}
			Listitem = append(Listitem, value)
		} else {
			binnode.Left.push(value)
		}
	} else {
		if binnode.Right == nil {
			binnode.Right = &binarynode{value, nil, nil}
			Listitem = append(Listitem, value)
		} else {
			binnode.Right.push(value)
		}
	}
}
func (bin *binarytree) push(pushvalue int) {
	if bin.root == nil {
		bin.root = &binarynode{pushvalue, nil, nil}
		Listitem = append(Listitem, pushvalue)
	} else {
		bin.root.push(pushvalue)
	}
}

func ProducerLinked(bin *binarytree) {
	for i := 0; i < 50; i++ {
		Produceditem := rand.Intn(100)
		fmt.Printf("%v Produced number : %v \n", i, Produceditem)

		bin.push(Produceditem)
		bin.root.Printstack()
	}
}

var mutex = &sync.Mutex{}
var wg sync.WaitGroup

type binarynode struct {
	value int
	Left  *binarynode
	Right *binarynode
}

var Listitem []int

type binarytree struct {
	root *binarynode
}

func getelements(ele *binarynode) string {
	var out string = ""
	if ele == nil {
		out = ""
	} else {
		out = fmt.Sprintf("%s %v %s", getelements(ele.Left), ele.value, getelements(ele.Right))
	}
	return out
}
func Consumer(bin *binarytree) {
	for i := 0; i < 10; i++ {
		bin.root.pop()
		bin.root.Printstack()

	}
}

//Printstack
func (bin *binarynode) Printstack() {
	var outputstring string
	outputstring = fmt.Sprintf("%s", getelements(bin))

	fmt.Println(outputstring)
}

func main() {
	var bin binarytree
	ProducerLinked(&bin)
	bin.root.Printstack()
	Consumer(&bin)

}
