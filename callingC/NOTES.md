# Calling C from Go! 

Yet another feature to love about this language: the ability to interoperate with libraries written in other languages such as C. 
To begin to understand how Go packages can call C code, start with this: [C? Go? Cgo!](https://blog.golang.org/c-go-cgo) 

##  A brief note on CGO and CGO Flags:

`cgo` creates Go bindings for C libraries and system calls, which lets Go packages call C code. The `#cgo` directives are used to provide flags for the compiler and linker when building the C parts of the package. The most common `#cgo` directives are: 

* `CFLAGS`: these are debug and optimization options/switches passed to the C compiler for compiling, telling the compiler where it can find C header files. 
* `CXXFLAGS`: same as `CFLAGS` but for C++ files! 
* `LDFLAGS`: these contain options that are passed to the linker so that library and archive files may be referenced. 

These are followed by the following flags: 

* `-I <directory>`: adds a directory to the list of places searched by the compiler for a file named on an `#include` line: these are directories searched for header files. 
* `-L <directory>`: is the path to the directories containing library files (whose file extensions are usually `.a`, `.so`, `.dll`, or `.lib`). `.a` files are _static libraries_ (Windows equivalent = `.lib` files), which means that the code in them is embedded into your binary, and thus if any changes occur to these, you have to recompile your code again. `.so` are _dynamic libraries_ (Windows equivalent = .`dll`), which means that the code stored in them will not be embedded into your own binary and thus can be simply linked at runtime. 
* `-l <library name>`: is the specific name of the library you want to link to.

All of these `#cgo` directives, along with any `#include` directives that reference C header files, must be included in the same block of code at the beginning of your .go files, which will then be followed by the pseudo-package "C", which will be placed above all other imported libraries.
 
## The `unsafe` package

In addition to the "pseudo package" C, which cgo interprets as a reference to C's name space, another package almost always used when working with C code is the `unsafe` package. 

When Go compiles packages, type checking attempts to find any attempts to apply an invalid operation to a value that is inappropriate for its type. Also, dynamic checking looks for any forbidden operations that would occur during execution. However, sometimes we need to sidestep these rules in order to interoperate with libraries written in other languages, notably C. This is where the `unsafe` package comes in. This package is most unusual in that it is implemented by the compiler itself and exposes details of Go's memory layout that are not accessible to other packages. It is used extensively witin low-level packages such as `os`, `runtime`, `syscall`, and `net`, which interact with the operating system. (See page 354 of Donovan/Kernighan's _The Go Programming Language_). 




