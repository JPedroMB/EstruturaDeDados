package main

import "fmt"

func main() {
  var ra [6]int
  for i := range ra {
    fmt.Printf("Insira o dígito %d do seu RA: ", i+1)
    fmt.Scan(&ra[i])
  }
  fmt.Println("R.A: ", ra)
}
