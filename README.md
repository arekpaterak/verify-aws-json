# verify-aws-json

## What is it?

It is a repository for a solution to some recruitment task.

The goal was to implement a function that verfies if a given JSON representing AWS::IAM::Role Policy does not contain `'*'` in the `Resource` field. If so the function should return `false`, otherwise `true`.

It was preferred to use Go for the implementation so I have chosen it despite the fact that it was my first time using it. I have written unit tests and tried to cover as many edge cases as possible.

## Instructions

### How to run it?

1. Install Go.
2. Run `go run main.go` in the root directory of the project.

### How to test it?

1. Run `go test` in the root directory of the project.