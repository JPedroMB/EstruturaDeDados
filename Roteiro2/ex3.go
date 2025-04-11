package main

import "fmt"

func main() {
  var ra [6]int
  for i := 0; i < 6; i++ {
    fmt.Printf("Insira o dígito %d do seu RA: ", i+1)
    fmt.Scan(&ra[i])
  }
  fmt.Println("R.A: ", ra)
}
