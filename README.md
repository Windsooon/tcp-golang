### Table of content

- [Quick start](#quickstart)
- [Run test](#run-test)
- [Output](#output)
- [Requirements and Goals](#requirements-and-goals)
- [High Level System Design](#high-level-system-design)
- [The idea behind the project](#the-idea-behind-the-project)
    - [Testing](#testing)
    - [Error handling](#error-handling)
    - [Improve](#improve)



### Quick start

    go run tcp_client.go

    go run tcp_server.go

### Run test

    // inside base dir
    go test -test.v

### Output

    // Client side
    >> go run tcp_client.go
    >> Windson
    >> Hello Windson!
    >> ^[[A
    >> <name> is not printable.
    >> 
    >> <name> should not be empty.
    >> Exit
    TCP client exiting...

    // Server side
    >> go run tcp_server.go 
    This is a TCP server...
    Launching Server at 8080
    2019/09/29 18:14:29  from [::1]:51602]:51602
    2019/09/29 18:14:38  from [::1]:51602
    2019/09/29 18:14:48 Client from [::1]:51602 left..
    
### Requirements and Goals
1. The client can send `<name>` to the server, the server will log the information from the client and return "Hello, <name>!" to the client.
2. If the `<name>` is not printable or empty, the server will return error messages.
3. The client can leave by sending "Exit" to the server. 

### High Level System Design

![img](https://raw.githubusercontent.com/Windsooon/tcp-golang/master/high-level.png)

### The idea behind the project

Since it's a challenge, I choose to use Golang instead of Python. (I also completed a challenge a week ago using Golang. Please check out [fibonacci-sequence](https://github.com/Windsooon/fibonacci-sequence) and [idea behind it](https://gist.github.com/Windsooon/54e919a3c923c8f44f1d9cb239418412)). In this project, I will focus on testing and error handling. 

#### Testing
I learned how to write better tests in this project, I use table tests and subtests to achieve a better result.

#### Error handling
There are lots of ways to handle error in Golang, I learned a lot from this great post [Don’t just check errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)


#### Improve
1. If we need to support millions of TCP connections, we can use epoll instead of goroutines for saving memory usage. (From [A Million WebSockets and Go](https://www.freecodecamp.org/news/million-websockets-and-go-cc58418460bb/))
2. Add testing for `tcp_server.go` and `tcp_client.go`
