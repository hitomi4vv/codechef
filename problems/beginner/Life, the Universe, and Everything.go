package main
import (
  "bufio"
  "fmt"
  "os"
)
func main() {
  sc := bufio.NewScanner(os.Stdin)
  var line = ""
  for sc.Scan() {
    line = sc.Text()
    if line == "42" {
      break
    }
    fmt.Println(line)
  }
}
