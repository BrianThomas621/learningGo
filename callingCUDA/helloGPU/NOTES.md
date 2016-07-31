# CUDA, Go, and CGO! 
 
## Some CUDA Notes: CUDA Driver API / CUDA Runtime API

The code above utilizes the CUDA Driver API (NOT the CUDA Runtime API). These are distinguished as follows: 

* CUDA Driver API methods/properties begin with "cu_"	= `<include cuda.h>`
* CUDA Runtime API methods/properties start with "cuda_"  = `<include cuda_runtime.h>`

## Making it all work in Windows (10): 

In order to get this to work on my Windows 10 laptops, I had to take some additional steps as follows. For starters, we need to call the CUDA libraries with `cgo`, and this is where I ran into a problem (which apparently Windows-specific). When installing CUDA 7.5, the files and libraries were installed into the `C:\Program Files\NVIDIA GPU Computing Toolkit\CUDA\v7.5` folder. Apparently the spaces in the path to the CUDA files is a problem for `cgo` in Windows with current versions of Go. I also had set the `CUDA_PATH` environment variable to point to the location of the CUDA files, and I was hoping to use this in the CGO flags, like so:

* `#cgo CFLAGS: -I"${CUDA_PATH}"/include/` 

However, this failed, throwing a **malformed cgo argument** error. Apparently, the problem is that there are spaces in the CUDA PATH. I also got the same error when using the full PATH to the CUDA libraries and header files as well, so I was at an apparent dead end at this point. 

My fix: I made a copy of all of the CUDA 7.5 files/folders and moved these into a "cuda" folder off of the C drive, thereby eliminating the white spaces. This allowed all CUDA files/libraries to be succesfully referenced in the CGO flags. So the above CGO flag now looks like: 

* `#cgo CFLAGS: -IC:/cuda/v7.5/include` 

## One further note... 

For some reason, I had to use an earlier version of the GNU C Compiler (gcc); originally I had gcc 5.1.0-2 installed but none of this code worked. This perplexed me because this exact same code was working on my other Windows 10 ASUS laptop just fine. I discovered that, on the other laptop, the GNU C Compiler was gcc version 4.9.2 (Architecture = x86_64). When I uninstalled the later version of gcc and used this earlier version, this code worked. 

