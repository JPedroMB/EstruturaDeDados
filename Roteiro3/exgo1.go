package main
import "fmt"

func main(){
  var b float64 = 3.576
  var p2 *float64 = &b
  fmt.Println("Valor antes da modificação: ", b)
  *p2 = 2.8
  fmt.Println("Valor antes da modificação: ", b)
}
