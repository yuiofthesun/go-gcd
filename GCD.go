package main

import(
    "fmt"
 //   "math
)
//func GCD() func(int) int


func main() {
    var x,y int
    fmt.Printf("Please enter two integer: ")
    fmt.Scanf("%d %d" , &x, &y)
    if x == 0 {
        fmt.Println("The GCD is %d:", y)
    }
    if y == 0 {
        fmt.Println("The GCD is %d:", x)
    }

    for x !=y {
        if x > y {
            x = x -y
        } else {
            y = y -x
        }
    }

    fmt.Printf("The GCD is: %d ", y)
}

