# Mapper

[![PkgGoDev](https://pkg.go.dev/badge/gihub.com/tamboto2000/mapper)](https://pkg.go.dev/gihub.com/tamboto2000/mapper)
Mapper map struct to another struct with same or similiar fields.

  

### Installation

Mapper require Go version 1.14 or up

```sh
$ go get github.com/tamboto2000/mapper
```

### Example

```go
package main

import (
	"fmt"

	"github.com/tamboto2000/mapper"
)

type Struct1 struct {
	Str   string
	Num   int
	Float float64
}

type Struct2 struct {
	Str string
	Num int
}

func main() {
	struct1 := Struct1{Str: "Hello world!", Num: 1, Float: 1.5}
	struct2 := new(Struct2)
	if err := mapper.Map(struct1, struct2); err != nil {
		panic(err.Error())
	}

	fmt.Println("Str:", struct2.Str)
	fmt.Println("Num:", struct2.Num)
}

```

License
----

MIT
