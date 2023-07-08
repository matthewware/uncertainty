Simple package to implements the most basic functional forms of
[error propogations](https://en.wikipedia.org/wiki/Propagation_of_uncertainty).

Assumes no covariance. Each number is assumed to be an independent sample.

## Usage
```go
package main

import (
    "fmt"
    "github.com/matthewware/uncertainty"
)

func foo(f1 UFloat, f2 UFloat) UFloat{
    return f1.Multiply(f2) // f1 * f2
}

func bar(f1 UFloat, f2 UFloat) UFloat{
    return ScalarMulitply(4, f1.Divided(f2)) // 4 * (f1 / f2)
}

func main() {
    // UFloat type is {value, standard deviation}
    a := UFloat(1.2, 0.3)
    b := UFloat(4.5, 0.6)
    fmt.Println(foo(a,b))
    fmt.Println(bar(a,b))
}
```

Current functions are Add, Subtract, Multiply, Divide, ScalarMulitply.
