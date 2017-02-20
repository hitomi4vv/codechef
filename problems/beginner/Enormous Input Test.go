package main
import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func main() {
  sc := bufio.NewScanner(os.Stdin)
  sc.Split(bufio.ScanWords)

  sc.Scan()
  var n, _ = strconv.Atoi(sc.Text())
  sc.Scan()
  var k, _ = strconv.Atoi(sc.Text())

  for sc.Scan() {
    num, _ := strconv.Atoi(sc.Text())
    if num%k != 0 {
      n--
    }
  }

  fmt.Println(n)
}
