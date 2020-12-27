# Conditionals

## 1. What is the purpose of a seed value in Go?
> A seed value provides a starting point for Go to generate random numbers. 

## 2. Given the following code snippet, what will print to the terminal?
```weekend := true
dayOfWeek := 7
 
if !weekend {
  switch day {
  case 1: 
    fmt.Println("It's Monday already?")
  case 2: 
    fmt.Println("Tuesday... still better than Monday.")
  case 3:
    fmt.Println("Wednesday, it's hump day!")
  case 4:
    fmt.Println("Thursday, almost done.")
  case 5:  
    fmt.Println("Friday! TGIF!")
  default:
    fmt.Println("What day of the week is it?")
  }
} else {
  fmt.Println("Take the day off!")
}
```
> Take the day off!

## 3. Given the following code snippet, what is the value of mixedOperators?
```
var mixedOperators bool 
 
mixedOperators = true || false
mixedOperators = mixedOperators && true 
mixedOperators = !mixedOperators
```
> false

## 4. Complete the code snippet below to have a working conditional.
``` 
 (Answer 1) randNum := rand.Intn(10) ; randNum < 7 {
  fmt.Println("Less than 7.")
} (Answer 2) randNum > 5 {
  fmt.Println("Less than 5.")
} (Answer 3) {
  fmt.Println("Greater than 7.")
}
```
> 1. if
> 2. else if
> 3. else

## 5. What is the purpose of using conditionals?
> Conditionals allow different blocks of code to be run depending on the values passed to it. 

## 6. Given the following code snippet, what will print to the terminal?
```
gradeAvg := 89
 
if gradeAvg >= 90 {
  fmt.Println("You've earned an A!")
} else if gradeAvg >= 80 {
  fmt.Println("You've earned a B.")
} else if gradeAvg >= 70 {
  fmt.Println("You've earned a C?") 
} else if gradeAvg >= 60 { 
  fmt.Println("You've earned a D...")
} else {
  fmt.Println("An F???")
}
```
> You've earned a B.

## 7. What is the purpose of a seed value in Go?
> A seed value provides a starting point for Go to generate random numbers. 

## 8. Given the following code snippet, what is printed to the terminal?
```
walksLikeDuck := true
talksLikeDuck := false
 
if walksLikeDuck && talksLikeDuck {
  fmt.Println("It's a duck!")  
} else if walksLikeDuck || talksLikeDuck {
  fmt.Println("It may or may not be a duck.")  
} else {
  fmt.Println("It's not a duck!")
}
```
> It may or may not be a duck.

## 9. When is an else statement used?
> else statements check a provided condition and runs code if the condition is true. `(Wrong)`
> Inlucing an else statement provides a default block of code to run when none of the previous conditions are true. `(Correct answer)`

## 10. Given the following code snippet, what will print to the terminal?
```
weekend := true
dayOfWeek := 7
 
if !weekend {
  switch day {
  case 1: 
    fmt.Println("It's Monday already?")
  case 2: 
    fmt.Println("Tuesday... still better than Monday.")
  case 3:
    fmt.Println("Wednesday, it's hump day!")
  case 4:
    fmt.Println("Thursday, almost done.")
  case 5:  
    fmt.Println("Friday! TGIF!")
  default:
    fmt.Println("What day of the week is it?")
  }
} else {
  fmt.Println("Take the day off!")
}
```
> Take the day off!

___

# Results

## 90% / 100%
## 1 ✗ | 9 ✓
