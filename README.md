# Silent Map Access in Go: Unexpected Zero Value

This repository demonstrates a common yet subtle issue in Go: accessing a non-existent key in a map silently returns the zero value for that key's type instead of an error.  This behavior can cause bugs that are difficult to track down, as the program continues execution without indication of the issue.

## The Problem

In many other programming languages, accessing a non-existent key in a map or dictionary would throw an exception or return a special value (like null or undefined) signaling an error. However, Go's maps return the zero value for the map's value type.  For integer types, this is 0; for strings, it's an empty string; for pointers, it's `nil`; and so on. This can lead to incorrect calculations or unexpected program flow without any explicit error message.

## The Solution

The best way to handle this situation is to explicitly check for the existence of the key before accessing it using the `value, ok := map[key]` idiom.

## Files

* `bug.go`: Demonstrates the silent map access.
* `bugSolution.go`: Shows how to safely check for key existence using the idiomatic approach. 