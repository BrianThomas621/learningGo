package main

/*
#cgo LDFLAGS: -LC:/cuda/v7.5/lib/x64 -lcuda
#cgo CFLAGS: -IC:/cuda/v7.5/include
#include <cuda.h>
*/
import "C"

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Device int

func init() {
	C.cuInit(C.uint(0))
}

func DeviceGetCount() int {
	var count C.int
	C.cuDeviceGetCount(&count)
	return int(count)
}

func DeviceComputeCapability(device Device) (computeCapacity string) {
	var maj, min C.int
	C.cuDeviceComputeCapability(&maj, &min, C.CUdevice(device))
	computeCapacity = strconv.Itoa(int(maj)) + "." + strconv.Itoa(int(min))
	return
}

func handler(writer http.ResponseWriter, request *http.Request) {
	buf := C.CString(string(make([]byte, 256)))
	C.cuDeviceGetName(buf, 256, C.CUdevice(0))
	gpudevice := "GPU Device:" + C.GoString(buf)
	no_gpus := "Number of GPU Devices: " + strconv.Itoa(DeviceGetCount())
	compute_cap := "Compute Capacity: " + DeviceComputeCapability(0)
	fmt.Fprintf(writer, "%s\n %s\n %s", gpudevice, no_gpus, compute_cap)
}

func main() {
	http.HandleFunc("/gpu_device", handler)
	server := &http.Server{
		Addr: "localhost:8080",
	}
	log.Fatal(server.ListenAndServe())
}
