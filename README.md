# VALEDICTORY

盲猜是出自 Young Sheldon

```go
package main

import (
    "fmt"

    "github.com/tonicbupt/valedictory"
)

type Milk struct {
    Name  string `url:"name,default:xxx"`
    Query string `url:"q,default:yyy"`
    Page  int    `url:"p,default:0"`
    Limit int    `url:"limit,default:20"`
    More  bool   `url:"more,default:true"`
}

func main() {
    m := Milk{}
    v := url.Values{
        "name": {"tonic"},
        "q":    {"iterm"},
        "p":    {"10"},
        "more": {"false"},
    }

    valedictory.Decode(&m, v)
    fmt.Println(m)
}
```
