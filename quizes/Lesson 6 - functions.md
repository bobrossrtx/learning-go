# Functions

## 2. What would the following block of code print?
```go
func returnsTwoThings(message string) (string, int) {
  var newMessage string
  var coolNumber int
  coolNumber = 5
  newMessage = "Surprise!" + message
  return newMessage, coolNumber
}
 
func main() {
  msg, num := returnsTwoThings("What's up?")
  fmt.Println(msg, num)
}
```

- Whats up? 5 | ✕
- ```Surprise! Whats up? 5 | ✓```

## 2. What’s printed out as a result of the following block of code?
```go
package main
import "fmt"
 
func returnsTen() int {
  return 10
}
 
func main() {
  coolVariable := returnsTen()
  fmt.Println(coolVariable + 10)
}
```

- 10 ✕
- ```20 ✓```

## 3. Which of the following use cases are good for functions?

1: ```When you need to store a variable but aren’t sure if it should be float32 or int32```

2: ```When a similar pattern of code is used multiple times but with numbers or data tweaked slightly.```

3: ```When you’re getting a lot of errors, functions will raise fewer errors.```

- ```2 ✓``` 

# 4.  Why does the following block of code fail to compile?

```go
package main
import "fmt"
 
func coolFunction() {
  var coolVariable := 5
}
 
func main() {
  coolFunction()
  fmt.Println(coolVariable)
}
```
- ``` is referenced outside of the scope it was defined in. ✓```

# 5. What does the following Go program print out?
```go
func multiplyThree(x int, y int, z int) int {
  return x * y * z
}
 
func main() {
  var product int
  product = multiplyThree(2, 2, 2)
  fmt.Println(product)
}
```
- ```8 ✓```

# 6. Which of the following defines a function in Go?
1: ```var goFunction function```

2: ```def goFunction:```

3: ```func goFunction() { }```

-  ```3 ✓```

# 7. What key word is used to call a function after the current function finishes?
1: ```async```

2: ```defer```

3: ```wait```

- ```defer ✓```

___

# Results

## ```71% / 100%```
## ```2 ✕ | 5 ✓```