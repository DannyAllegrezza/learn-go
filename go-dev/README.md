# Official `go.dev` website

| Name          	| URL                      	| Description                                                                      	|
|---------------	|--------------------------	|----------------------------------------------------------------------------------	|
| Tutorial: Get started with Go | https://go.dev/doc/tutorial/getting-started 	| Getting started tutorial. Call code in an external package. 	|


## Imports

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
```

This code groups the imports into a parenthesized, "factored" import statement.

You can also write multiple import statements, like:

```go
import "fmt"
import "math"
```

But it is good style to use the factored import statement.

## Exported names
In Go, a name is exported if it begins with a capital letter. For example, `Pizza` is an exported name, as is `Pi`, which is exported from the `math` package.

`pizza` and `pi` do not start with a capital letter, so they are not exported.

> When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

