# Testing in Go

This repository is designed to help developers understand the testing framework in the Go programming language. Testing is a critical part of software development, and Go provides a simple yet powerful framework to write tests for your programs.

### Introduction to Testing in Go

Go comes with a built-in testing framework that allows you to write unit tests for your code. Testing ensures that your functions behave as expected and helps catch errors before they reach production.

- The basic test function in Go looks like this:
```bash
func TestFunctionName(t *testing.T) {
    // Test logic here
}
```
- Every test function name should start with Test.
 - The t *testing.T argument provides access to testing methods that report errors or failures.

### Writing Your First Test

Let’s say you have a simple Add function:
```bash
func Add(a, b int) int {
    return a + b
}
```
To test this function, you can create a test file add_test.go:
```bash
package main

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}
```
The test checks if Add(2, 3) returns 5. If not, it will report an error using t.Errorf().

### Running Tests
To run the tests, use go test command;
```bash
go test
```
Example output:
```bash
PASS
ok      main.go 0.004s
```
If there is a failure, it will display which test failed and why.

### Table-Driven Tests

Table-driven testing is a common pattern in Go, allowing you to test multiple inputs and expected outputs efficiently.

- Here’s an example:
```bash
func TestAdd(t *testing.T) {
    testCases := []struct {
        a, b, expected int
    }{
        {2, 3, 5},
        {1, 1, 2},
        {0, 0, 0},
        {-1, -1, -2},
    }

    for _, tc := range testCases {
        result := Add(tc.a, tc.b)
        if result != tc.expected {
            t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
        }
    }
}
```
- Here, we use a slice of structs to define multiple test cases.
- The loop runs through each test case, ensuring the Add function produces the correct result for different inputs.