package main

import (
  "fmt"
  //"bufio"
  //"encoding/csv"
  //"flag"
  //"log"
  //"os"
  //"strings"
  //"time"
  )

func add ( x,y int) int {
  return x + y
}
func swap (h,j string) (string, string) {
  return j, h
}


func main (){
  a,b := swap("fun", "joy")
  fmt.Printf("The adding function yields %d", add(1,1))
  fmt.Printf(" The swapping function yields: ")
  fmt.Println(a,b)
}
