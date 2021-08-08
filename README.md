# SENTENCEPIECEGO
This is a golang interface of [sentencepiece](https://github.com/google/sentencepiece) for serving. For more information, please follow google sentencepiece.

# Usages
1. Download shared library
Please check sentencepiece library first with pip and spm-train. You can download v0.1.96 and v0.1.90 in releases.
```bash
sudo wget https://github.com/evan176/sentencepiecego/releases/download/v0.1.96-x86-64/libsentencepiecego.so -P /usr/local/lib/
ldconfig
```
2. Export CGO variable
```
export CGO_CXXFLAGS=-std=c++11
```
3. Go get
```
go get github.com/evan176/sentencepiecego
```

```go
package main

import (
        "fmt"

        "github.com/evan176/sentencepiecego"
)

func main() {
        // Load pre-trained spm model
        m, err := sentencepiecego.Load("spm.model")
        if err != nil {
                panic(err)
        }
        // Encode text to ids([]int) with spm model
        ids, err := m.Encode("test")
        if err != nil {
                panic(err)
        }
        fmt.Printf("%+v\n", ids)
        // Release model before exit
        m.Free()
}
~
```

# References
https://github.com/google/sentencepiece
