// $go run variadic-functions.go

package main
import "fmt"


func plus(a int, b int) int {
    return a + b
}

func plusPlus(a, b, c int) int {
    return a + b + c
}

func vals() (int, int) {
    return 3, 7
}


func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {


    res := plus(1, 2)
    fmt.Println("1+2 =", res)
    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
    
    
    // multiple return
    
    a, b := vals()
    fmt.Println(a)      // 3
    fmt.Println(b)      // 7
    
    _, c := vals()
    fmt.Println(c)         // 7 
    
    
    //////-------------- variadic --------------------------
    
    sum(1, 2)
    sum(1, 2, 3)
    
    nums := []int{1, 2, 3, 4}
    sum(nums...)
    
    
   
}
