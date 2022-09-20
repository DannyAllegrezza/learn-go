
Personal repository to document my journey learning the Go programming language.


- [Resources for getting started](#resources-for-getting-started)
    - [1. The official Go website](#1-the-official-go-website)
    - [2. 📚 The Go Programming Language](#2--the-go-programming-language)
- [📓 Journal](#-journal)
  - [Day 1](#day-1)
  - [Day 2](#day-2)
    - [TIL](#til)
    - [TIL - formatting](#til---formatting)
  - [Day 3](#day-3)
    - [Concurrent programming](#concurrent-programming)

# Resources for getting started

### 1. The official [Go website](https://go.dev/)

No surprises here. The official website has a bunch of helpful resources for getting started. Quite a few resources, actually...

1. [Documentation page](https://go.dev/doc/)
   - Has a ton of individual tutorials, as well as actual formal documentation on the language itself
2. [Tour of Go](https://go.dev/tour/)
3. [Effective Go](https://go.dev/doc/effective_go) 
   - Was originally written for Go's release in 2009 and has not been updated significantly since. Reader should understand it is far from a complete guide. 
   - **I think I'll check this out after reviewing more basic tutorials**

### 2. 📚 The Go Programming Language
# 📓 Journal 

## Day 1
  - [x] Setup this repository
  - [x] Setup learning plan and collect resources
  - [x] Create a hello world program in `go` and play around with running vs building

**Q: What does it take to print ["hello world"](go-by-example/1-hello-world.go) in go?**

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}
```

Easy enough. But what is the significance of `package main`? Does `func main()` have significance, like in a lot of C-derived languages?

```csharp
using System;

namespace MyApp
{
   class Program
   {
      static void Main(string[] args)
      {
         Console.WriteLine("hello world");
      }
   }
}
```

As it turns out, yes, there is significance. For more info, see the [program execution](https://go.dev/ref/spec#Program_execution) section of the language spec.

* A complete program is created by linking a single, unimported package called the main package with all the packages it imports, transitively.
* The main package must have package name main and declare a function main that takes no arguments and returns no value.
  * `package main`
  * `func main() { ... }`

Program execution begins by initializing the `main` package and then invoking the function `main()`. When that function invocation returns, the program exits. It does not wait for other (non-main) goroutines to complete.


## Day 2 

I've been finding the official Go resources to be the most informative. Their website covers beginner to advanced topics. Overall, learning Go has been a mostly painless transition based on my existing language experience.

Right now, I'm juggling between the official Go documentation, https://gobyexample.com and "The Go Programming Language" book.

### TIL

**Slices** are like a dynamically sized sequence `s` of array elements where individual elements can be accessed as `s[i]` and a contiguous subsequence as `s[m:n]`. The number of elements is given by `len(s)`.

In [echo3.go](go-programming-language/ch1/echo3.go), we built a little clone of the Unix `echo` command, working our way up to the final implementation in `echo3`. One thing I found interesting was this syntax:

`os.Args[1:]`

* `os.Args` is a _`slice`_ of strings
* The first element of this particular slice is always the name of the command itself, and the other elements are the command line arguments that were passed to the program when it started execution
  
This is shorthand for `os.Args[1:len(os.Args)]`. If `m` or `n` is omitted, it defaults to 0 or len(s) respectively, so we can abbreviate the desired slice as `os.Args[1:]`.

> We used this as a way to skip the 1st element in the slice and just loop or range through the rest of the elements.



### TIL - formatting

In [dup1.go](go-programming-language/ch1/dup1.go) we explore working with `stdin` and formatting strings in Go.

The [fmt](https://pkg.go.dev/fmt) package implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.

There a ton of verbs, but the most common ones are as follows:

```
%v	the value in a default format
	when printing structs, the plus flag (%+v) adds field names
%#v	a Go-syntax representation of the value
%T	a Go-syntax representation of the type of the value
%%	a literal percent sign; consumes no value
```

The most common use of formatting is to print to a terminal window through `stdout`...

```go
func Print(a ...interface{}) (n int, err error)
func Printf(format string, a ...interface{}) (n int, err error)
func Println(a ...interface{}) (n int, err error)
```

* [Print()](https://pkg.go.dev/fmt#Print) formats using the default formats for its operands and writes to standard output. Spaces are added between operands when neither is a string. It returns the number of bytes written and any write error encountered.

* [Printf()](https://pkg.go.dev/fmt#Printf) formats according to a format specifier and writes to standard output. 

```go
// In dup1.go
fmt.Printf("%d\t%s\n", n, line)
```

* [Println()](https://pkg.go.dev/fmt#Println) formats using the** default formats for its operands** and writes to standard output. Spaces are **always added between operands** and a **newline** is appended. It also returns the number of bytes written and any write error that was encountered.


```go
// In dup1.go
fmt.Println(n, line)
```

## Day 3

At this point, I've switched my learning plan to following exercises listed in "The Go Programming Language" book. Similar to the K&R book on `C`, it has a nice progression and flow that I am enjoying.

Exercises are listed in [go-programming-language](go-programming-language/) directory.

### Concurrent programming

In [fetch.go](go-programming-language/ch1/fetch.go), we write a CLI app that fetches the HTML for a given URL. In [fetchall.go](go-programming-language/ch1/fetchall.go) we get a light introduction to concurrency in go.

`fetchall.go` fetches n number of URLs at once by spinning up separate `goroutines`. A `goroutine` is a concurrent function execution. We learned that a `channel` is a communication mechanism that allows one goroutine to pass values of a specified type to another goroutine. In all of our programs thus far, even our `main` function itself runs in a goroutine, and the `go` statement creates additional goroutines ie: 

```go
// before
fetch(url)
// after  
go fetch(url, channel)
```

When one goroutine attempts to send or receive on a channel, it blocks until antoher goroutine attempts the corresponding `receive` or `send` operation, at which point the value is **transferred** and both goroutines proceed. In [fetchall](go-programming-language/ch1/fetchall.go), each `fetch()` sends a value (`ch` <- _expression_) on the channel `ch`, and `main()` receives all of them (`<-ch`). Having `main()` do all of the printing ensures that output from each goroutine is processed as a unit.