# Pokemon Challenge Premium Minds

This test was developed in Go, that is the language that I am learning and working in my actual company.
The installation of Go and use of the Go commands and environment is not detailed in this document.


## Instructions for compilation

 - Install Go distribution according your Operational System following the instructions here: [Go Installation](https://golang.org/doc/install).
 - Clone this repository to the GOPATH/src directory.
 - Inside the pm project directory, run this following command:

```
$ go build -o main
```

## Instructions to run the program

 - According to your OS, you can run the program like this:

Linux/Osx
```
$ ./main -movs=NSEO
```
Windows
```
c:\...\pm\main.exe -movs=NSEO
```
There is a little help to check the command:
```
$ ./main -h
```
## Instructions to run the tests

 - Inside the pm project directory, run this following command:

```
$ go test
```