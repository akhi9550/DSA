package main

import (
    "fmt"
)

func makeFancyString(s string) string {
    arr := []byte(s)
    for i := 1; i < len(arr)-1; i++ {
        if arr[i-1] == arr[i] && arr[i] == arr[i+1] {
            arr = append(arr[:i], arr[i+1:]...)
            i-- 
        }
    }
    return string(arr)
}

func main() {
    var s string
    fmt.Print("Enter a string: ")
    fmt.Scanln(&s)

    result := makeFancyString(s)
    fmt.Println("Fancy string:", result)
}
