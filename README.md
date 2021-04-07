# datetranslate

This package aims to translate format strings between 2 timestamp formats:

- `SimpleDateFormat` - https://docs.oracle.com/javase/8/docs/api/java/text/SimpleDateFormat.html
- `ctimefmt` - https://github.com/observIQ/ctimefmt/blob/4cb1bdfd4b74804fd68c5169d48a8db1e06cfe3e/ctimefmt.go#L63-L97

## How to use it

```go
package main

import (
    "fmt"
    "time"

    "github.com/observiq/ctimefmt"
    "github.com/pmalek-sumo/datetranslate"
)

func main() {
    dateFormat, err := datetranslate.SimpleDateFormat2Ctimefmt("yyyy-MM-dd HH:mm:ss.SSS")
    if err != nil {
        fmt.Printf("Couldn't translate provided format string: %v", err)
        return
    }

    out, err := ctimefmt.Format(dateFormat, time.Now())
    if err != nil {
        fmt.Printf("Couldn't format the time with ctimefmt: %v", err)
        return
    }

    fmt.Printf("The time is: %v", out)
}
```
