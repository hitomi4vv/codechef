package main
import "fmt"
func main() {
  var t int
  fmt.Scanln(&t)
  for ; t > 0; t-- {
    p := 1
    c := 0
    var n int
    fmt.Scanln(&n)
    for p <= n {
      p *= 5
      c += n / p
    }
    fmt.Println(c)
  }
}
