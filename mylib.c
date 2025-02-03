// mylib.c
#include <unistd.h>
#include <stdint.h>

static int timeout_seconds = 1;  // Default timeout (1 second)

void set_timeout(int seconds) {
    timeout_seconds = seconds;
}

void blocking_function(uint8_t* buffer) {
    sleep(timeout_seconds);  // Simulate a blocking call with configurable timeout
    
    // Fill the buffer with some data (just a simple example)
    for (int i = 0; i < 20; i++) {
        buffer[i] = i;
    }
}
