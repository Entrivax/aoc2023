# AOC 2023

My results for the Advent of Code 2023.

## Run specific day

To run the day `1`, run the following:
```sh
make 1
```

## Create a new day script

Take for example the day `1`
- Create a folder `1` in `cmd/` to get `cmd/1`
- Create a file `1.go` in that folder to get `cmd/1/1.go`
- Create the input file of the day (`1.txt`) in `internal/inputfile/` to get `internal/inputfile/1.txt`
- Create the embed script to go with it (`1.go`) in `internal/inputfile/` to get `internal/inputfile/1.go`
- In the embed script, use the following template:
```go
package inputfile

import _ "embed"

//go:embed 1.txt
var InputFile1 string
```