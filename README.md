# Hey,

## I'm Owen, I'm learning GoLang because:
> - I believe it is a better starting point then jumping straight in with the harder languages.
> - I think of it as a way to expand my knowledge for some harder to learn languages like `JavaScript` or `C#`.
> - It is an amazing and flexible language for problem solving, machine learning, network programming and such more.
> - Go has a simple and beginner friendly syntax with good-looking, straight forward code.

Since I have intermediate knowledge of basic HTML and CSS I decided to learn Go to develop back end server side code and data for future sites I end up building.

Yes, I know that learning a more upfront and modernly used language like `JavaScript` would benefit me more since there are a lot more libraries and frameworks especially when it comes to node but I thought of this as a better starting point instead of Hitting the harder stuff up first like i have been for the past 2 months


## Notes:

### Recourse:

If you’re ever stuck on something, check out:

[Codecademy's forums](https://discuss.codecademy.com/c/get-help/go/1877)
View questions and answers from learners on this site!

[Stack Overflow’s Go questions](https://stackoverflow.com/questions/tagged/go?tab=Active)
A forum of programmers answering programming questions.


[Go’s official site](https://golang.org/)
Go to the official site for documentation.


[Google](https://www.google.com/)
Roll up your sleeves and search it up!
It may help to search Golang instead of Go in certain cases.

Using the 'go doc' command is helpful.
You can find out more about a package.
Or about a function inside the package.
EXAMPLE: go doc time. Now



### Comments:

You can create comments with ```//```, or you can make a multi line comment with ```/* Comment goes here */```.



### Variables and Types:

You can create a variable by defining one with var.
You can use that variable in your code as a nickname for your values.
EXAMPLE :
```go
var funfair = "Hummingbirds' wings can beat up to 200 times a second."
fmt.Println("Did you know?")
fmt.Println(funfair)
Output :
Did you know?
Hummingbirds' wings can beat up to 200 times a second.
```

You can also make a constant. `(const)`
They Can be used to create a variable that can not be changed later on in the code. 

Variables can be assigned different types : `bool, string, int32 int64, uint32 uint64, byte // alias for uint8, rune // alias for int32, float32 float64, complex64 complex128`.
The basic types in Go : `ints, floats, complexs, and strings`.
:= is known as the short declaration operator. It is used to declare and initialize 
the variables inside and outside the functions.

When defining variables, you can define the in single line, for example, amount, unit := 10, "dollars" or you can do it like this:
var amount, unit int32
amount = 10
unit = "dollars"



### Verbs:

In go there are tags called verbs in the fmt module. Here are some basic verbs used 
in Go
The %v verb is a place holder for variables included within the fmt. Print/ln line.
The %T verb is to title the type of the included variable of the fmt. Print/ln line. 
The %% verb is a literal percent sign; consumes no value
the %t is the word true or false



### Operators:

Using an operator can decide how your if/else/else if statement will behave.  
Depending on which operator you use for your block of code, will vary on your outcome 
of the finished product.

Here are the basic operators:
> == Equal to
>
> != Not equal to
>
>  < Greater than
>
> <= Greater than or equal to
>
> &gt; Less Than
>
> &gt;= Less than or equal to

Logical operators:
> && And
>
> || Or
>
> !  Not


Go also has a math/rand library that can be used to create a random number.
This is likely to be used since if you hard caode variables, you know the outcome therefor no need for a if statement.

This is likely to be used since if you hard code variables, you know the outcome therefor no need for a if statement.

### Functions:
	
It’s important to note that the code inside the function body does not run until we call the function.
We can call a function by using it's name followed by a parentheses somewhere outside the definition of the function.
EG: After the function is created, type the functions name with a paren inside of the main function (func main())
	
### Scopes:
A function definition creates something called a scope. We’ve referenced scope before in our conditionals exercise about
scoped short declarations but it’s important to recognize how scope plays a huge role in functions and programming overall!
	
Scope is a concept that refers to where the values and functions are defined and where they can be accessed.
For instance, when a variable is defined within a function, that variable is only accessible within that function.
When we try to access that same variable from a different function, we get an error because we can’t do it.
Each function has its own specific scope.
	
There are three different scopes present in this example:
The global scope, which contains the function definitions for main() and performAddition().
performAddition() has a local scope, which defines x and y.
main() has a local scope also. It can access performAddition() because that’s defined on the same scope level as main() but can’t access the internals of performAddition‘s scope (i.e., x or y).
	
This differentiation of scope keeps the namespace, the available variables and keywords, cleaner as a result. You can only refer to variables or functions defined within a specific namespace.
	
### Returning Values from Functions:
	
While variables and their values are scoped to their functions, there are no built-in ways of passing information out of their native functions and into another namespace.
The way that information can be sent from within a function to the call site, the place where the function is called. 
This is done by returning a value — when we return a value, we pass the value to another place in our code.
A function can be given a return type, the type of a value that will be returned by the function.
At the call site, the return value can be stored within a variable of the same type as the function’s return.
A return statement tells the function to pass back a value (or multiple values) and stops the function from executing any more code, i.e. if we put more code after our return statement, it would not run!
we can access the variables we need through the return keyword!
eg.
	
``` go
func getLengthOfCentralPark() int32 {
    var lengthInBlocks int32
    lengthInBlocks = 51
    return lengthInBlocks
}
	
func main() {
	var centralParkLength int32
	centralParkLength = getLengthOfCentralPark()
	fmt.Println(centralParkLength) // Prints: 51
}
```
