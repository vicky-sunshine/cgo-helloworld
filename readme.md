# cgo-helloworld

Self-practice for cgo.

## helloworld-01

use basic C function

```sh
cd helloworld-01
go run main.go
```

## helloworld-02-func

call customize C function

```sh
cd helloworld-02-func
go run main.go
```

## helloworld-03-gcg

main.go call C function which call Go function

Go -> C -> Go

```sh
cd helloworld-03-gcg
go run main.go
```

## helloworld-04-c-source

import independent `.c` file

```sh
cd helloworld-04-c-source
go run main.go
```

## helloworld-05-static-link-lib

use static link library

```sh
cd helloworld-05-static-link-lib

cd helloworld
gcc -c helloworld.c -o helloworld.o
ar -crs libhelloworld.a helloworld.o

cd ..
go run main.go
```

## helloworld-06-dynamic-link-lib

use dynamic link library

```sh
cd helloworld-06-dynamic-link-lib

# build dynamic link
cd src
gcc -shared -o libhelloworld.so helloworld.c

# move .so to lib foldeer
mv libhelloworld.so ../lib
# [in MacOS] set DYLD_LIBRARY_PATH
export DYLD_LIBRARY_PATH=<absolutie_path_of_cgo-101>/helloworld-06-dynamic-link-lib/lib

# back to helloworld-06-dynamic-link-lib folder
cd ..
go run main.go
```
