Simple package to implements the most basic functional forms of
[error propogation](https://en.wikipedia.org/wiki/Propagation_of_uncertainty).

Assumes no covariance. Each number is assumed to be an independent sample.

## Usage
```go
package main

import (
    "fmt"

    "github.com/matthewware/uncertainty"
)

func foo(f1 uncertainty.UFloat, f2 uncertainty.UFloat) uncertainty.UFloat{
    return f1.Multiply(f2) // f1 * f2
}

func bar(f1 uncertainty.UFloat, f2 uncertainty.UFloat) uncertainty.UFloat{
    uf := f1.Divide(f2))
    return uf.MultiplyByFloat(2.5) // 4 * (f1 / f2)
}

func main() {
    // UFloat type is {value, standard deviation}
    a := uncertainty.UFloat(Value: 1.2, Uncertianty: 0.3)
    b := uncertainty.UFloat(Value: 4.5, Uncertainty: 0.6)
    fmt.Println(foo(a,b))
    fmt.Println(bar(a,b))
}
```

Current functions are Add, Subtract, Multiply, Divide, MultiplyByFloat.
