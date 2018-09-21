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
  fmt.Printf("The adding function yields %d", add(1,1))
  fmt.Printf(" The swapping function yields %d", swap("fun","joy"))
}
