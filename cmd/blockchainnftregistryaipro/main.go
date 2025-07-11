// cmd/blockchainnftregistryaipro/main.go
package main

import (
"flag"
"log"
"os"

"blockchainnftregistryaipro/internal/blockchainnftregistryaipro"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainnftregistryaipro.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
