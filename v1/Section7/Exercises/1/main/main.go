package main

import "fmt"

func main(){
  for i := 1; i < 10; i++{
    test(i)
  }
}

func test(val int){
  var half = (val / 2)
  var isEven = (val % 2 == 0)
  fmt.Printf("Half of %v is %v. Is this even? %v\n", val, half, isEven)
}
