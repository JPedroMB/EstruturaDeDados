package main
import "fmt"
func swap() {
  var x int = 3
  var y int = 4
  var p *int = &x
  var q *int = &y
  fmt.Printf("Valores antes da troca: %d %d\n", x, y)
  *p, *q = *q, *p
  fmt.Printf("Valores depois da troca: %d %d", *p, *q)
}
func main(){
  swap()
}
