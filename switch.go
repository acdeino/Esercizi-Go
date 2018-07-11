package main

import "fmt"
import "time"

func main() {

  i := 2
  fmt.Println("Write ", i, " as ")
  switch i {
  case 1:
    fmt.Println("One")
  case 2:
    fmt.Println("Two")
  case 3: fmt.Println("three")
  }

  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("It's the Weekend !")
  default:
    fmt.Println("It's a Weekday")
  }

  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("It's before noon")
  case t.Hour() > 12:
    fmt.Println("It's after noon")
  case t.Hour() > 23:
    fmt.Println("It's night")
  default:
    fmt:Println("It's morning!")
  }

  whatamI := func(i interface{}) {
    switch t := i {
    case bool:
      fmt.Println("I'm a boole, named after George Boole")
    case int
      fmt.Println("I'm an integer")
    default:
      fmt.Println("Don't know type %T\n", t)
    }
  }

  whatamI(true)
  whatamI(1)
  whatamI("Hey!")
}
