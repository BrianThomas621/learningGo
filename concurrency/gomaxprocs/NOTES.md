# A Note about `GOMAXPROCS`

`GOMAXPROCS` is a Go environment variable that limits the number of OS threads that can execute Go code simultaneously. In Go versions 1.4 and earlier, the default was 1, which meant that only one OS thread was used for goroutines. However, starting with Go version 1.5, Go's default behavior was changed to set the number of OS thread execution contexts from 1 to the number of CPU's returned by the `runtime.NumCPU()` function. 

However, there seems to be a misconception that `GOMAXPROCS` represents the number of CPUs Go will use to run go routines. This is not the case, and the number of OS threads that can execute Go code simultaneously can be adjusted by calling the `runtime.GOMAXPROCS()` function. This can be adjusted to a number far more than the number of your CPUs, i.e. up to 256. 

Note that if you change `GOMAXPROCS` to a number *larger* than 256, no error will occur...it will just set `GOMAXPROCS` to 256, as shown here: 

```
fmt.Println("Default GOMAXPROCS:\t", runtime.GOMAXPROCS(-1))    // Prints Default GOMAXPROCS:   8
fmt.Println("Number of CPUs:\t\t", runtime.NumCPU())            // Prints Number of CPUs:       8
runtime.GOMAXPROCS(300)
fmt.Println("Maximum GOMAXPROCS:\t", runtime.GOMAXPROCS(-1))    // Prints Maximum GOMAXPROCS:	256
```



