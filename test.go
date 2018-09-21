package main

import (
  "fmt"
  "bufio"
  "encoding/csv"
  "flag"
  "log"
  "os"
  "strings"
  "time"
  )

func add ( x int, y int) int {
  return x + y
}

func main (){
  fmt.Println(add(1,1))
}
