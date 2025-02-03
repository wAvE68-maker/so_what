package main

/*
#cgo LDFLAGS: -L. -lmylib
#include <stdint.h>

void set_timeout(int seconds);
void blocking_function(uint8_t* buffer);
*/
import "C"
import (
	"fmt"
	"sync"
	"unsafe"
)

func callBlockingFunction(wg *sync.WaitGroup, buffer []C.uint8_t, size int) {
	defer wg.Done()

	// Call the C function (it will block for the specified timeout)
	C.blocking_function((*C.uint8_t)(unsafe.Pointer(&buffer[0])))

	// Convert the buffer to a Go slice and print it
	result := C.GoBytes(unsafe.Pointer(&buffer[0]), C.int(size))
	fmt.Println("Returned data:", result)
}

func main() {
	// Set the timeout to 2 seconds (you can modify this as needed)
	C.set_timeout(2)

	var wg sync.WaitGroup
	const bufferSize = 20

	// Run 5 goroutines concurrently, each with its own buffer
	for i := 0; i < 5; i++ {
		// Allocate a unique buffer for each goroutine as a Go slice
		buffer := make([]C.uint8_t, bufferSize)

		// Add a goroutine to the wait group
		wg.Add(1)

		fmt.Printf("run %v\n", i)

		// Launch a goroutine
		go callBlockingFunction(&wg, buffer, bufferSize)

		goSlice := unsafe.Slice(&buffer[0], len(buffer))
		uint32Value := *(*uint32)(unsafe.Pointer(&goSlice[0]))
		fmt.Printf("First 4 bytes as uint32: %v\n", uint32Value)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
