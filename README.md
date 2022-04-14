# Colorize - go module for terminal colors

## install

go get github.com/r0lh/color

## usage

Example:

```
package main

import (
   "fmt"
   "github.com/r0lh/color"
)

func main() {
  fmt.Println(color.Colorize(color.Success, "colorize module works!"))
}
```
