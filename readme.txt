gcc -shared -o libmylib.so -fPIC mylib.c

#cgo LDFLAGS: -L. -lmylib
#cgo go directive for flags to compiler or linker
LDFLAGS means we are passing linker flags
-L. loo in current directory for libraries
-lmylib link against mylib (lib prefix and .so extension are handled by linker)
