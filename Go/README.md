# Go: CodeAcademy

https://www.codecademy.com/courses/learn-go/

## Compiling

To compile a go file

```
go build file.go
```

To run a go file

```
./file.go
```

## Running files

If we want to combine building and running, we can run this command

```
go run file.go
```

## Basic Go Structure: Packages

Every Go file will have a similar structure to this,

```go
package main

import "fmt"
```

This will import all the things we need

## Basic Go Structure: Main

If the file is using the main package, then the main function will be the first thing run.

```go
func main() {
  fmt.Println("Hello World")
}
```

## Importing Multiple Packages

Here is a list of Go Packages: https://golang.org/pkg/

We have two ways of importing multiple packages.

Option 1:

```go
import "package1"
import "package2"
```

Option 2:

```go
import (
    "package1"
    "package2"
)
```

If we want to provide an alias to a pacakage.

```go
import (
    p1 "package1"
    "package2"
)

p1.SampleFunc() // Example of using alias
```

## Comments

You can either use // or /\* \*/ to comment out code

## Go Doc Command

You can use the Go Doc command to learn about any package

```
go doc fmt
```

You can also use the command to get information about functions in a package

```
go doc fmt.println
```

## Values and Variables

Literals

```go
fmt.println(10 * 10) // Prints: 100
```

Constants (Can't be mutated)

```go
const funFact = "Hummingbirds' wings can beat up to 200 times a second."
```

Variable (Can change)

```go
var lengthOfSong uint16 = 10
```

Infering Variables

```go
// Don't need type
days := 24

// Also looks like
var days = 24
```

## Printf

Using %v (verbs) allows you to add in variables

```go
specialWord := "Test"
fmt.Printf("This value's type is %v.", specialWord)
```

## Formating keywords

- %v for string value
- %T for type of variable
- %d for number
- %f for precise (%.2f for 3.8 will print 3.80)

## Scanf

Allows for user input

```go
package main

import "fmt"

func main() {
  fmt.Println("What would you like for lunch?")

  // Add your code below:
  var food string
  fmt.Scan(&food)

  fmt.Printf("Sure, we can have %v for lunch.", food)
}
```

## Conditional Statements

Pretty simple and normal

```go
package main

import "fmt"

func main() {
	lockCombo := "2-35-19"
	robAttempt := "1-1-1"

  // Add your code below:
  if lockCombo == robAttempt {
    fmt.Println("The vault is now opened.")
  }
}
```

## Switch Statements

```go
switch clothingChoice {
case "shirt":
    fmt.Println("We have shirts in S and M only.")
case "polos":
   fmt.Println("We have polos in M, L, and XL.")
case "sweater":
    fmt.Println("We have sweaters in S, M, L, and XL.")
case "jackets":
    fmt.Println("We have jackets in all sizes.")
default:
    fmt.Println("Sorry, we don't carry that")
}
```
