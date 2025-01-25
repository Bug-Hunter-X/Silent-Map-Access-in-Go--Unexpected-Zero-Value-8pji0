```go
package main

import (

    "fmt"
)

func main() {
    var m = map[string]int{"a": 1, "b": 2}

    value, ok := m["c"]
    if ok {
        fmt.Println("Value of c:", value) //Prints only if "c" exists
    } else {
        fmt.Println("Key 'c' not found in the map")
    }
}
```