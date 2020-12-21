# Hey,

## I'm owen, I'm learning GoLang because:
> - I believe it is a better starting point then jumping straight in with the harder languages.
> - I think of it as w way to expand my knowledge for some harder to learn languages like `JavaScript` or `C#`.
> - It is an amazing and flexible language for problem solving, machine learning, network programming and buch more.
> - Go has a simple and beginner friendly syntax with good looking, straight forward code.

Since I have intermit knowledge of basic HTML and CSS I decided to learn Go to develop backend server side code and data for future sites I end up building.

Yes, I know that learning a more upfront and modernly used language like `JavaScript` would benefit me more since there are alot more libraries and frameworks especially when it comes to node but I thought of ths as a better starting point instead of Hitting the harder stuff up first like i have been for the past 2 months


## Notes:

### Recourses:

If you’re ever stuck on something, check out:

[<span style="color: #0C9EE8;">Codecademy's forums](https://discuss.codecademy.com/c/get-help/go/1877)<br>
View questions and answers from learners on this site!

[<span style="color: #0C9EE8;">Stack Overflow’s Go questions](https://stackoverflow.com/questions/tagged/go?tab=Active)<br>
A forum of programmers answering programming questions.


[<span style="color: #0C9EE8;">Go’s official site](https://golang.org/)<br>
Go to the official site for documentation.


[<span style="color: #0C9EE8;">Google](https://www.google.com/)<br>
Roll up your sleeves and search it up!
It may help to search Golang instead of Go in certain cases.

<span style="color: #0C9EE8;">Using the 'go doc' command is helpful.</span><br>
You can find out more about a package.
Or about a function inside the package.
EXAMPLE: go doc time.Now



### Comments:

You can create comments with //, or you can make a multi line comment with /* Comment goes here */.
<br><br>


### Variables and Types:

You can create a variable by defining one with var.
You can use that variable in your code as a nickname for your values.
EXAMPLE :
var funFact = "Hummingbirds' wings can beat up to 200 times a second."
fmt.Println("Did you know?")
fmt.Println(funFact)
Output :
Did you know?
Hummingbirds' wings can beat up to 200 times a second.

You can also make a constant. (const)
They Can be used to create a variable that can not be changed later on in the code. 

Variables can be assigned different types : bool, string, int32  int64, uint32 uint64, byte // alias for uint8, rune // alias for int32, float32 float64, complex64 complex128.
The basic types in Go : ints, floats, complexs, and strings.
:= is known as the short declaration operator. It is used to declare and initialize 
the variables inside and outside the functions.

When defining variables, you can defnie the in single line, for example, ammount, unit := 10, "dollers" or you can do it like this:
var ammount, unit int32
ammount = 10
unit = "dollers"
<br><br>


### Verbs:

In go there are tags called verbs in the fmt module. Here are some basic verbs used 
in Go
The %v verb is a is a place holder for variables included within the fmt.Print/ln line.
The %T verb is to title the type of the included variable of the fmt.Print/ln line. 
The %% verb is a literal percent sign; consumes no value
the %t is the word true or false
<br><br>


### Operators:

Using an operator can decide how your if/else/else if statement will behave.        
Depending on wich operator you use for you block of code, will vary on your outcome 
of the finnished product.

Here are the basic operators:
> ==       Equal to
>
> !=       Not equal to
>
>  <       Less than
>
>  &lt;       Greater than
>
> <=       Less than or equal to
>
> &lt;=       Greater than or equal to

Logical operators:
> &&      And
>
> ||      Or
>
>  !      Not


Go also has a math/rand library that can be used to create a random number.<br>
This is likely to be used since if you hard caode variables, you know the outcome therefor no need for a if statement.
