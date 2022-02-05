# Weight
[![Go Report Card](https://goreportcard.com/badge/github.com/nothinux/weight)](https://goreportcard.com/report/github.com/nothinux/weight)  ![test status](https://github.com/nothinux/weight/actions/workflows/test.yml/badge.svg?branch=master)

This package implement weighted round-robin load balancing algorithm


## Documentation
see [pkg.go.dev](https://pkg.go.dev/github.com/nothinux/weight)


## Installation

```
$ go get github.com/nothinux/weight
```

### Getting Started
``` go
package main

import "github.com/nothinux/weight"

func main() {
    w := weight.NewWeight()
    w.AddWeight("A", 2)
    w.AddWeight("B", 3)
    w.AddWeight("C", 2)
    
    for i := 0; i < 7; i++ {
        log.Println(w.Next())
    }
}
```

## LICENSE
[MIT](https://github.com/nothinux/weight/blob/master/LICENSE)
