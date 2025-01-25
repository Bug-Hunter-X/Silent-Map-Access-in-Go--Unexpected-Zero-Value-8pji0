```go
func main() {
    var m = map[string]int{"a": 1, "b": 2}
    fmt.Println(m["c"]) // This will print 0, not an error
}
```