package main
import "fmt"

func main() {
  var withdraw int
  var currentBalance, finalBalance float64
  fmt.Scan(&withdraw, &currentBalance)
  finalBalance = currentBalance - float64(withdraw) - 0.50
  if withdraw % 5 != 0 || finalBalance < 0 {
    finalBalance = currentBalance
  }
  fmt.Printf("%.2f", finalBalance)
}
