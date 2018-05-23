using System;

namespace HelloWorld {

  public class HelloWorld {
    public static int Hello(int n) {
      Console.WriteLine("Hello is called with n = " + n.ToString());
      return 2;
    }

    public static string GetString() {
      return "mystring";
    }
  }

}