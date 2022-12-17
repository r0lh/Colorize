# Colorize - go module for terminal colors

go module for terminal colors

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
  fmt.Println(color.Ize(color.Blue, "colorize module works!"))
}
```

Possible colors:
- White
- Red
- Green
- Yellow
- Blue
- Purple
- Cyan
- Gray
