using System;

public class Program {
  public static void Main() {
    while(true) {
      int num = int.Parse(Console.ReadLine());
      if(num == 42) {
        break;
      } else {
        Console.WriteLine(num);
      }
    }
  }
}
