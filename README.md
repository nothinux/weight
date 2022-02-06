# Weight
[![Go Report Card](https://goreportcard.com/badge/github.com/nothinux/weight)](https://goreportcard.com/report/github.com/nothinux/weight)  ![test status](https://github.com/nothinux/weight/actions/workflows/test.yml/badge.svg?branch=master)

This package implement weighted round-robin load balancing algorithm, inspired from [nginx](https://github.com/phusion/nginx/commit/27e94984486058d73157038f7950a0a36ecc6e35).

It's works by increase all current_weight by their respective weight, then choose backend which has the largest current_weight, after that the previously selected backend will reduce its current_weight by total_weight, and this process will repeat every `Next` method called.

In case you have 3 backend server A, B, C and its weight 2, 3, 2 the selected server will -> { B, A, C, B, A, C, B }


## Documentation
see [pkg.go.dev](https://pkg.go.dev/github.com/nothinux/weight)


## Installation

```
$ go get github.com/nothinux/weight
```

### Getting Started
``` go
package main

import (
    "github.com/nothinux/weight"
    "log"
)

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
