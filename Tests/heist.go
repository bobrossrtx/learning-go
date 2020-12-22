<<<<<<< HEAD
package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  
  var isHeistOn bool = true
  var eludedGuards int = rand.Intn(100)

  if eludedGuards <= 50 {
    fmt.Print("Looks like you've managed to make it past the guards. Good job, but remember, this is just the first step.")
  } else {
    isHeistOn = false
    fmt.Println("We've got to come up with a better plan next time.")
  }

  fmt.Println("\n***") // Added for spacing
  fmt.Println("Is the heist on? ", isHeistOn)
  fmt.Println("\n***") // Added for spacing

  var openedVault int = rand.Intn(100)

  if isHeistOn && openedVault >= 70 {

    fmt.Println("Grab all of the money and leave as fast as you can before the cops arive.")
  } else if isHeistOn {
    isHeistOn = false
    fmt.Println("The vault door cannot be opened")
  }

  var leftSafely = rand.Intn(5)
  if isHeistOn {
    switch leftSafely {
      case 0:
        isHeistOn = false
        fmt.Println("You have failed the heist.")
      case 1:
        isHeistOn = false
        fmt.Println("You have failed the heist.")
      case 2:
        isHeistOn = false
        fmt.Println("You have failed the heist.")
      case 3:
        isHeistOn = false
        fmt.Println("You have failed the heist.")
      case 4:
        isHeistOn = false
        fmt.Println("You have failed the heist.")
      default:
        isHeistOn = true
        fmt.Println("You have got out of the bank safely, get to the getaway vehicle")
    }
}

  if isHeistOn {
    var amtStolen int = 10000 + rand.Intn(1000000)
    fmt.Println("This is how much was stolen:", amtStolen)
  }

}
=======
package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  
  var isHeistOn bool = true
  var eludedGuards int = rand.Intn(100)

  if eludedGuards <= 50 {
    fmt.Print("Looks like you've managed to make it past the guards. Good job, but remember, this is just the first step.")
  } else {
    isHeistOn = false
    fmt.Println("We've got to come up with a better plan next time.")
  }

  fmt.Println("\n***") // Added for spacing
  fmt.Println("Is the heist on? ", isHeistOn)
  fmt.Println("\n***") // Added for spacing

  var openedVault int = rand.Intn(100)

  if isHeistOn && openedVault >= 70 {

    fmt.Println("Grab all of the money and leave as fast as you can before the cops arive.")
  } else if isHeistOn {
    isHeistOn = false
    fmt.Println("The vault door cannot be opened")
  }

  var leftSafely = rand.Intn(5)
  if isHeistOn {
    switch leftSafely {
      case 0:
        isHeistOn = false
        fmt.Println("You have failed the heist.")
      case 1:
        isHeistOn = false
        fmt.Println("You have failed the heist.")
      case 2:
        isHeistOn = false
        fmt.Println("You have failed the heist.")
      case 3:
        isHeistOn = false
        fmt.Println("You have failed the heist.")
      case 4:
        isHeistOn = false
        fmt.Println("You have failed the heist.")
      default:
        isHeistOn = true
        fmt.Println("You have got out of the bank safely, get to the getaway vehicle")
    }
}

  if isHeistOn {
    var amtStolen int = 10000 + rand.Intn(1000000)
    fmt.Println("This is how much was stolen:", amtStolen)
  }

}
>>>>>>> ed9344e4fa57a3a383756291d5d0b20d457815be
