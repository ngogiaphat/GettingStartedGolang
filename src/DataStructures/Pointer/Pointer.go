package main
import "fmt"
func main() {
    var x int = 10
    var ptr *int
    //Assign the address of x to ptr
    ptr = &x
    fmt.Println("Value of x:", x)
    fmt.Println("Address of x:", &x)
    fmt.Println("Value of ptr:", ptr)
    fmt.Println("Dereferenced value of ptr:", *ptr)
    //Update the value of x through ptr
    *ptr = 20
    fmt.Println("Updated value of x:", x)
}