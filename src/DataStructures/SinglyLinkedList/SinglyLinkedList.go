package main
import "fmt"
type Node struct {
    data int
    next *Node
}
type LinkedList struct {
    head *Node
}
func (ll *LinkedList) addNode(data int) {
    newNode := &Node{data, nil}
    if ll.head == nil {
        ll.head = newNode
    } else {
        current := ll.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
}
func (ll *LinkedList) printList() {
    current := ll.head
    for current != nil {
        fmt.Printf("%d ", current.data)
        current = current.next
    }
    fmt.Printf("\n")
}
func main() {
    ll := &LinkedList{}
    ll.addNode(1)
    ll.addNode(2)
    ll.addNode(3)
    ll.addNode(4)
    ll.printList()
}