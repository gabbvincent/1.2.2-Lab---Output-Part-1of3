package main

import "fmt"

func main() {
  var week string = "Tuesday the Second."
  var month string = "Febuary,"
  var day , month2 , year int
  day = 2
  month2 = 4
  year = 2020
  var town string = "Chico"
  var today float64 = 64.3
  var tomorrow float64 = 78.8

  fmt.Println("Hello!, Today's date is" , month , week)

  fmt.Println(day , "/" , month2 , "/" , year)

  fmt.Println("We're here in" , town , "California where it is" , today , "degrees outside. Tomorrow we are looking at a significant rise in tempature with a high of" , tomorrow , "degrees, so go outside and enjoy the sun!")

// end.
}
               