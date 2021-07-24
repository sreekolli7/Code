# Questions
## Exercises - Ninja Level 1
Contribute your code
As you work through these hands on exercises, if you create code which you would like to share with the rest of the class, tweet me the link ( https://twitter.com/Todd_McLeod ) and I will add it to our course outline.
video: 016b
1. Hands-on exercise #1
Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
42
“James Bond”
true
Now print the values stored in those variables using 
a single print statement
multiple print statements
code: here’s the solution: https://play.golang.org/p/yYXnWXIQNa 
video: 017
1. Hands-on exercise #2
Use var to DECLARE three VARIABLES. The variables should have package level scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE).
identifier “x” type int
identifier “y” type string
identifier “z” type bool
in func main
print out the values for each identifier
The compiler assigned values to the variables. What are these values called?
code: here’s the solution: https://play.golang.org/p/jzHwSlles9 
video: 018
1. Hands-on exercise #3
Using the code from the previous exercise,
At the package level scope, assign the following values to the three variables
for x assign 42
for y assign “James Bond”
for z assign true
in func main
use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
print out the value stored by variable “s”
code: here’s the solution: https://play.golang.org/p/QFctSQB_h3 
video: 019
1. Hands-on exercise #4
FYI - nice documentation and new terminology “underlying type”
https://golang.org/ref/spec#Types 
For this exercise
Create your own type. Have the underlying type be an int.
create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
in func main
print out the value of the variable “x”
print out the type of the variable “x”
assign 42 to the VARIABLE “x” using the “=” OPERATOR
print out the value of the variable “x”
code: here’s the solution: https://play.golang.org/p/snm4WuuYmG 
video: 020
1. Hands-on exercise #5
Building on the code from the previous example
at the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom TYPE “x”
eg:
	
in func main
this should already be done
print out the value of the variable “x”
print out the type of the variable “x”
assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
print out the value of the variable “x”
now do this
now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
then use the “=” operator to ASSIGN that value to “y”
print out the value stored in “y”
print out the type of “y”
code: here’s the solution: https://play.golang.org/p/cj8RrYgBOD 
video: 021
1. Hands-on exercise #6
Take this quiz.
video: 022


## Exercises - Ninja Level 2
1. Hands-on exercise #1
Write a program that prints a number in decimal, binary, and hex
solution: https://play.golang.org/p/bAQxcEuK8O 
video: 031
1. Hands-on exercise #2
Using the following operators, write expressions and assign their values to variables:
==
<=
>=
!=
<
>
Now print each of the variables. 
solution: https://play.golang.org/p/76R-poSzaY 
video: 032
1. Hands-on exercise #3
Create TYPED and UNTYPED constants. Print the values of the constants.
solution: https://play.golang.org/p/NutvJXWUx2 
video: 033
1. Hands-on exercise #4
Write a program that 
assigns an int to a variable
prints that int in decimal, binary, and hex
shifts the bits of that int over 1 position to the left, and assigns that to a variable
prints that variable in decimal, binary, and hex
solution: https://play.golang.org/p/Ms964T8SbH 
video: 034
1. Hands-on exercise #5
Create a variable of type string using a raw string literal. Print it.
solution: https://play.golang.org/p/dLy36A-V-w 
video: 035
1. Hands-on exercise #6
Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
solution: https://play.golang.org/p/MDLF3v5EGT 
video: 036
1. Hands-on exercise #7
take this quiz
video: 037
## Exercises - Ninja Level 3
1. Hands-on exercise #1
Print every number from 1 to 10,000
solution: https://play.golang.org/p/voDiuiDGGw  
video: 050
1. Hands-on exercise #2
Print every rune code point of the uppercase alphabet three times. Your output should look like this:
65
	U+0041 'A'
	U+0041 'A'
U+0041 'A'
66
	U+0042 'B'
	U+0042 'B'
	U+0042 'B' 
 … through the rest of the alphabet characters
solution: https://play.golang.org/p/1OjnCX1D5H 
video: 051
1. Hands-on exercise #3
Create a for loop using this syntax
for condition { }
Have it print out the years you have been alive.
solution: https://play.golang.org/p/tnyqBPJ-i5 
video: 052
1. Hands-on exercise #4
Create a for loop using this syntax
for { }
Have it print out the years you have been alive.
solution: https://play.golang.org/p/9VpnB-I1Pz 
video: 053
1. Hands-on exercise #5
Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.
solution: https://play.golang.org/p/ohfJOW9euy 
video: 054
1. Hands-on exercise #6
Create a program that shows the “if statement” in action.
solution: https://play.golang.org/p/DpZ_FLfn5s 
video: 055
1. Hands-on exercise #7
Building on the previous 1. Hands-on exercise, create a program that uses “else if” and “else”.
solution: https://play.golang.org/p/IDnrJpE7vT 
video: 056
1. Hands-on exercise #8
Create a program that uses a switch statement with no switch expression specified.
solution: https://play.golang.org/p/YpPgkWGqKY 
video: 057
1. Hands-on exercise #9
Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
solution: https://play.golang.org/p/h-8FnjpzWB 
video: 058
1. Hands-on exercise #10
Write down what these print:
fmt.Println(true && true) 
fmt.Println(true && false) 
fmt.Println(true || true) 
fmt.Println(true || false) 
fmt.Println(!true)
solution: 
video: 059
## Exercises - Ninja Level 4
1. Hands-on exercise #1
Using a COMPOSITE LITERAL: 
create an ARRAY which holds 5 VALUES of TYPE int
assign VALUES to each index position. 
Range over the array and print the values out. 
Using format printing
print out the TYPE of the array
solution: https://play.golang.org/p/tD0SzV3hdf 
video: 071
1. Hands-on exercise #2
Using a COMPOSITE LITERAL: 
create a SLICE of TYPE int
assign 10 VALUES 
Range over the slice and print the values out. 
Using format printing
print out the TYPE of the slice
solution: https://play.golang.org/p/sAQeFB7DIs 
video: 072 
1. Hands-on exercise #3
Using the code from the previous example, use SLICING to create the following new slices which are then printed:
[42 43 44 45 46]
[47 48 49 50 51]
[44 45 46 47 48]
[43 44 45 46 47]
solution: https://play.golang.org/p/SGfiULXzAB 
video: 073 
1. Hands-on exercise #4
Follow these steps:
start with this slice
x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
append to that slice this value
52
print out the slice
in ONE STATEMENT append to that slice these values
53
54
55
print out the slice
append to the slice this slice
y := []int{56, 57, 58, 59, 60}
print out the slice
solution: https://play.golang.org/p/QUYhtSBaDS 
video: 074 
1. Hands-on exercise #5
To DELETE from a slice, we use APPEND along with SLICING. For this 1. Hands-on exercise, follow these steps:
start with this slice
x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
use APPEND & SLICING to get these values here which you should ASSIGN to a variable “y” and then print:
[42, 43, 44, 48, 49, 50, 51]
solution: https://play.golang.org/p/u8zpHLfgSE 
just for fun: 
https://goo.gl/frz786 
https://www.youtube.com/channel/UCJ8YIwWQCO7hMiqpOw2ZLFw 
video: 075
1. Hands-on exercise #6
Create a slice to store the names of all of the states in the United States of America. What is the length of your slice? What is the capacity? Print out all of the values, along with their index position in the slice, without using the range clause. Here is a list of the states:

` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`, 

solution:
the solution shown in the video was incorrect - this one is incorrect 
https://play.golang.org/p/tRKQDQuQCE 
I used assignment and a composite literal instead of append
here is the correct solution
https://play.golang.org/p/dzxZh4lhEpT 
video: 076
1. Hands-on exercise #7
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "Helloooooo, James."

Range over the records, then range over the data in each record.
solution: https://play.golang.org/p/FMM4c2PhZg 
video:  077
1. Hands-on exercise #8
Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things. Store three records in your map. Print out all of the values, along with their index position in the slice.

	`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
`no_dr`, `Being evil`, `Ice cream`, `Sunsets`

solution: https://play.golang.org/p/nTzSlRa9_A 
video: 078
1. Hands-on exercise #9
Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop
solution: https://play.golang.org/p/_CkxAhRrDJ 
video: 079
1. Hands-on exercise #10
Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop
solution: https://play.golang.org/p/TYl5EbjoeC 
video: 080



## Exercises - Ninja Level 5
1. Hands-on exercise #1
Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors. 
solution:
https://play.golang.org/p/ouRHmH_POg 
video: 086
1. Hands-on exercise #2
Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.
solution: https://play.golang.org/p/RvvLyAMvGo 
video: 087
1. Hands-on exercise #3
Create a new type: vehicle. 
The underlying type is a struct. 
The fields: 
doors
color 
Create two new types: truck & sedan. 
The underlying type of each of these new types is a struct. 
Embed the “vehicle” type in both truck & sedan. 
Give truck the field “fourWheel” which will be set to bool. 
Give sedan the field “luxury” which will be set to bool. solution 
Using the vehicle, truck, and sedan structs: 
using a composite literal, create a value of type truck and assign values to the fields; 
using a composite literal, create a value of type sedan and assign values to the fields. 
Print out each of these values. 
Print out a single field from each of these values.
solution: https://play.golang.org/p/PrTtTv_vVO 
video: 088 
1. Hands-on exercise #4
Create and use an anonymous struct.
solution: https://play.golang.org/p/otBHFs-lPp 
video: 089




## Exercises - Ninja Level 6
1. Hands-on exercise #1
Review
functions
purpose of functions
abstract code
code reusability
DRY - Don’t Repeat Yourself
func, receiver, identifier, params, returns
parameters vs arguments
variadic funcs
multiple “variadic” params
multiple “variadic” args
returns
multiple returns
named returns - yuck!
func expressions
assigning a func to a variable
callbacks
passing a func into another func as an argument
closure 
one scope enclosing another
variables declared in the outer scope are accessible in inner scopes
closure helps us limit the scope of variables
recursion
a func that calls itself 
factorial
life philosophy
focus on what’s important; not upon what’s urgent
Hands on exercise
create a func with the identifier foo that returns an int
create a func with the identifier bar that returns an int and a string
call both funcs
print out their results
code: https://play.golang.org/p/owEJNF5WAD 
video: 102 
1. Hands-on exercise #2
create a func with the identifier foo that 
takes in a variadic parameter of type int
pass in a value of type []int into your func (unfurl the []int)
returns the sum of all values of type int passed in
create a func with the identifier bar that 
takes in a parameter of type []int
returns the sum of all values of type int passed in
code: https://play.golang.org/p/B0yRxtBQPD 
video: 103 
1. Hands-on exercise #3
Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
code: https://play.golang.org/p/XluEuUD0Nw 
video: 104 
1. Hands-on exercise #4
Create a user defined struct with 
the identifier “person”
the fields:
first
last
age
attach a method to type person with
the identifier “speak”
the method should have the person say their name and age
create a value of type person
call the method from the value of type person
code: https://play.golang.org/p/NnXyWdqbbw 
video: 105
1. Hands-on exercise #5
create a type SQUARE
create a type CIRCLE
attach a method to each that calculates AREA and returns it
circle area= π r 2
square area = L * W
create a type SHAPE that defines an interface as anything that has the AREA method
create a func INFO which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
use func info to print the area of circle
code: https://play.golang.org/p/NGGikHNvHv 
video: 106 
1. Hands-on exercise #6
Build and use an anonymous func 
code: https://play.golang.org/p/DQX3xEIcRe  
video: 107
1. Hands-on exercise #7
Assign a func to a variable, then call that func
code: https://play.golang.org/p/_Qu7ZAyFDH 
video: 108
1. Hands-on exercise #8
Create a func which returns a func
assign the returned func to a variable
call the returned func
code: https://play.golang.org/p/c2HwqVE1Rd 
video: 109
1. Hands-on exercise #9
A “callback” is when we pass a func into a func as an argument. For this exercise, 
pass a func into a func as an argument 
code: https://play.golang.org/p/0yGYPKh1y7 
video: 110
1. Hands-on exercise #10
Closure is when we have “enclosed” the scope of a variable in some code block. For this 1. Hands-on exercise, create a func which “encloses” the scope of a variable:
code: https://play.golang.org/p/a56uWtoFSL 
video: 111 
1. Hands-on exercise #11
The best way to learn is to teach. For this 1. Hands-on exercise, 
choose one of the above exercises, or use the recursion example of factorial
download, install, and get it running 
https://obsproject.com/ 
record a video of YOU teaching the topic
upload the video to youtube
share the video on twitter and tag me in it ( https://twitter.com/Todd_McLeod ) so that I can see it!
video: 112



## Exercises - Ninja Level 7
1. Hands-on exercise #1
Create a value and assign it to a variable. 
Print the address of that value.
code: https://play.golang.org/p/57kW8xd0qT
video: 116
1. Hands-on exercise #2
create a person struct
create a func called “changeMe” with a *person as a parameter
change the value stored at the *person address
important
to dereference a struct, use (*value).field 
p1.first
(*p1).first
“As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method), x.f is shorthand for (*x).f.”
https://golang.org/ref/spec#Selectors 
create a value of type person
print out the value
in func main
call “changeMe”
in func main
print out the value
code: https://play.golang.org/p/AehM662HKS 
video: 117 


## Exercises - Ninja Level 8
1. Hands-on exercise #1
Starting with this code, marshal the []user to JSON. There is a little bit of a curve ball in this 1. Hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a package.
solution: https://play.golang.org/p/8BK6PXj3aG 
video: 125 
1. Hands-on exercise #2
Starting with this code, unmarshal the JSON into a Go data structure. Hint: https://mholt.github.io/json-to-go/ 
code:
code setup (just fyi, not needed for exercise): 
https://play.golang.org/p/nWPP5Z2Q4e 
solution:
https://play.golang.org/p/r8oSG8DIPR 
video: 126 
1. Hands-on exercise #3
Starting with this code, encode to JSON the []user sending the results to Stdout. Hint: you will need to use json.NewEncoder(os.Stdout).encode(v interface{})
solution: https://play.golang.org/p/ql90D1OQqd 
video: 127
1. Hands-on exercise #4
Starting with this code, sort the []int and []string for each person.
solution: https://play.golang.org/p/jz_llY1aPp 
video: 128 
1. Hands-on exercise #5
Starting with this code, sort the []user by 
age
last
Also sort each []string “Sayings” for each user
print everything in a way that is pleasant
solution: https://play.golang.org/p/8RKkdtLl6w 
video: 129 

## Exercises - Ninja Level 9
1. Hands-on exercise #1
in addition to the main goroutine, launch two additional goroutines
each additional goroutine should print something out
use waitgroups to make sure each goroutine finishes before your program exists
code: https://github.com/GoesToEleven/go-programming 
video: 148 
1. Hands-on exercise #2
This exercise will reinforce our understanding of method sets:
create a type person struct
attach a method speak to type person using a pointer receiver
*person
create a type human interface
to implicitly implement the interface, a human must have the speak method
create func “saySomething” 
have it take in a human as a parameter
have it call the speak method
show the following in your code 
you CAN pass a value of type *person into saySomething
you CANNOT pass a value of type person into saySomething
here is a hint if you need some help
https://play.golang.org/p/FAwcQbNtMG 

Receivers       Values
-----------------------------------------------
(t T)           T and *T
(t *T)          *T

code: https://github.com/GoesToEleven/go-programming 
video: 149
1. Hands-on exercise #3
Using goroutines, create an incrementer program
have a variable to hold the incrementer value
launch a bunch of goroutines
each goroutine should
read the incrementer value
store it in a new variable
yield the processor with runtime.Gosched()
increment the new variable
write the value in the new variable back to the incrementer variable
use waitgroups to wait for all of your goroutines to finish
the above will create a race condition. 
Prove that it is a race condition by using the -race flag
if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej 
code: https://github.com/GoesToEleven/go-programming 
video: 150
1. Hands-on exercise #4
Fix the race condition you created in the previous exercise by using a mutex
it makes sense to remove runtime.Gosched()
code: https://github.com/GoesToEleven/go-programming 
video: 151
1. Hands-on exercise #5
Fix the race condition you created in exercise #4 by using package atomic
code: https://github.com/GoesToEleven/go-programming 
video: 152
1. Hands-on exercise #6
Create a program that prints out your OS and ARCH. Use the following commands to run it
go run
go build
go install
code: https://github.com/GoesToEleven/go-programming 
video: 153
1. Hands-on exercise #7
Download OBS. Create a screen recording teaching some aspect of the Go programming language. Some topics you might teach:
Why Go
installing go
GOROOT & GOPATH
environment variables
hello world
go commands
help
variables
Short declaration operator
constants
loops
init, cond, post
break
continue
functions
func (receiver) identifier(params) (returns) { code }
methods
interfaces
method sets
type
conversion
concurrency vs parallelism
goroutines
waitgroups
mutex
atomic
After creating your recording, upload your video to youtube. Then tweet a link to your video and tag me in the tweet so that I can see it.
code: https://github.com/GoesToEleven/go-programming 
video: 154


## Exercises - Ninja Level 10
1. Hands-on exercise #1
get this code working:
using func literal, aka, anonymous self-executing func
solution: https://play.golang.org/p/SHr3lpX4so 
a buffered channel
solution: https://play.golang.org/p/Y0Hx6IZc3U 
video: 164
1. Hands-on exercise #2
Get this code running:
https://play.golang.org/p/oB-p3KMiH6 
solution: https://play.golang.org/p/isnJ8hMMKg 
https://play.golang.org/p/_DBRueImEq 
solution: https://play.golang.org/p/mgw750EPp4 
video: 165
1. Hands-on exercise #3
Starting with this code, pull the values off the channel using a for range loop
solution: https://play.golang.org/p/D3N4Tq54SN 
video: 166
1. Hands-on exercise #4
Starting with this code, pull the values off the channel using a select statement
solution: https://play.golang.org/p/FulKBY5JNj  
video: 167
1. Hands-on exercise #5
Show the comma ok idiom starting with this code.
solution: https://play.golang.org/p/qh2ywLB5OG 
video: 168
1. Hands-on exercise #6
write a program that
puts 100 numbers to a channel
pull the numbers off the channel and print them
solution: https://play.golang.org/p/096Lk1BR7o 
video: 169
1. Hands-on exercise #7
write a program that
launches 10 goroutines
each goroutine adds 10 numbers to a channel
pull the numbers off the channel and print them
solutions: 
https://play.golang.org/p/R-zqsKS03P
https://play.golang.org/p/quWnlwzs2z 
student solutions:
https://twitter.com/mannion
https://gist.github.com/mannion007/3c8899913974c1027ef6f13ec37b2b3f
video: 170

## Exercises - Ninja Level 11
1. Hands-on exercise #1
Start with this code. Instead of using the blank identifier, make sure the code is checking and handling the error.
solution: 
https://play.golang.org/p/tn8oiuL1Yn 
video: 176
1. Hands-on exercise #2
Start with this code. Create a custom error message using “fmt.Errorf”.
solution: 
https://play.golang.org/p/HugU4HJEEO 
https://play.golang.org/p/NII-lmGasj 
https://play.golang.org/p/Vo5kIoR-sG 
video: 177
1. Hands-on exercise #3
Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that has a value of type error as a parameter. Create a value of type “customErr” and pass it into “foo”. If you need a hint, here is one.
solution: 
https://play.golang.org/p/ixeowY2fd2 
assertion
https://play.golang.org/p/pbl2kCYsM0 
conversion
https://play.golang.org/p/1ldiBdkdzA 
video: 178
1. Hands-on exercise #4
Starting with this code, use the sqrt.Error struct as a value of type error. If you would like, use these numbers for your 
lat "50.2289 N"
long "99.4656 W"
solution: 
https://play.golang.org/p/nsRxbDfkCh 
video: 179
1. Hands-on exercise #5
We are going to learn about testing next. For this exercise, take a moment and see how much you can figure out about testing by reading the testing documentation & also Caleb Doxsey’s article on testing. See if you can get a basic example of testing working. 
video: 180

## Exercises - Ninja Level 12
1. Hands-on exercise #1
Create a dog package. The dog package should have an exported func “Years” which takes human years and turns them into dog years (1 human year = 7 dog years). Document  your code with comments. Use this code in func main. 
run your program and make sure it works
run a local server with godoc and look at your documentation.
solution: https://github.com/GoesToEleven/go-programming 
video: 180-f
1. Hands-on exercise #2
Push the code to github. Get your documentation on godoc.org and take a screenshot. Delete your code from github. Refresh godoc.org so that it no longer has your code. Tweet me (https://twitter.com/Todd_McLeod) your screenshot.
solution: https://github.com/GoesToEleven/go-programming 
video: no video
1. Hands-on exercise #3
Use godoc at the command line to look at the documentation for:
fmt
fmt Print
strings
strconv
solution: https://github.com/GoesToEleven/go-programming 
video: no video


## Exercises - Ninja Level 13
1. Hands-on exercise #1
Start with this code. Get the code ready to BET on the code (add benchmarks, examples, tests). Run the following in this order:
tests
benchmarks
coverage
coverage shown in web browser
examples shown in documentation in a web browser
solution: 
https://github.com/GoesToEleven/go-programming 
video: 189
1. Hands-on exercise #2
Start with this code. Get the code ready to BET on (add benchmarks, examples, tests) however do not write an example for the func that returns a map; and only write a test for it as an extra challenge. Add documentation to the code. Run the following in this order:
tests
benchmarks
coverage
coverage shown in web browser
examples shown in documentation in a web browser
solution: 
https://github.com/GoesToEleven/go-programming 
video: 190
1. Hands-on exercise #3
Start with this code. Get the code ready to BET on (add benchmarks, examples, tests). Write a table test. Add documentation to the code. Run the following in this order:
tests
benchmarks
coverage
coverage shown in web browser
examples shown in documentation in a web browser
helpful to know: 
https://play.golang.org/p/4GUqs1HMpp 
https://play.golang.org/p/P9unTIFeOq 
solution: 
https://github.com/GoesToEleven/go-programming 
video: 191
