// cmd/micronetwork/main.go
package main

import (
"flag"
"log"
"os"

"micronetwork/internal/micronetwork"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := micronetwork.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
