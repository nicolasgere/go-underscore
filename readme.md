
# go-underscore

go-underscore is a utility-belt library for Golang that provides support for the usual functional suspects (each, map, reduce, filter...) without creating special container. This is largely inspeired by javascript package underscore.
It rely on the new golang generic feature available with the 1.18 (still in development). Read more here about go generic.


## Usage/Examples

```go
package main
import (
  underscore "github.com/nicolasgere/go-underscore"
)

func Main() {
  list := []int{1,2,3,4,5}
  listFiltered := underscore.filter(list, func(e int) bool{
    return e > 2
  })
  fmt.println(listFiltered) // Print: [3,4,5]
}
```


## Available function

- [x] Each
- [x] Map
- [x] Filter
- [x] Every
- [x] Some
- [x] Reduce
- [x] Find
- [x] Sort

