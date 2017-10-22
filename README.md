# FizzBuzz implementation in Golang
Just a simple repository to implement the Fizz Buzz Challenge in Golang and practice the use of channels for synchronisation

Installation:
```
go get -u github.com/ClarksonCJ/fizzbuzz-go
```

Execution:
Uses the Cobra command configuration (will execute with the default input array of numbers from 1 - 100)
```
go run main.go run
```

Also uses Persistent flags to allow input of data as an array of integers on the commandline (will execute with input values 1 - 15)
```
go run main.go run --Input=1,2,3,4,5,6,7,8,9,10,11,12,13,14,15
```