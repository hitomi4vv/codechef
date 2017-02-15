package main
import "fmt"

func main() {
  var a, b int
  fmt.Scan(&a, &b)
  if c := a - b; (c % 10) == 9 {
    fmt.Println(c - 1)
  } else {
    fmt.Println(c + 1)
  }
}
