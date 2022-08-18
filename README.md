# learn-go

Personal repository to document my journey learning the Go programming language.

## Resources for getting started

### 1. The official [Go website](https://go.dev/)

No surprises here. The official website has a bunch of helpful resources for getting started. Quite a few resources, actually...

1. [Documentation page](https://go.dev/doc/)
   - Has a ton of individual tutorials, as well as actual formal documentation on the language itself
2. [Tour of Go](https://go.dev/tour/)
3. [Effective Go](https://go.dev/doc/effective_go) 
   - Was originally written for Go's release in 2009 and has not been updated significantly since. Reader should understand it is far from a complete guide. 
   - **I think I'll check this out after reviewing more basic tutorials**

## 📓 Journal 

#### Day 1 7/23/22
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


### Day 2 

I've been finding the official Go resources to be the most informative. Their website covers beginner to advanced topics. Overall, learning Go has been a mostly painless transition based on my existing language experience.
