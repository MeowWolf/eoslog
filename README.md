# eoslog

Simple log formatter with standard (non-go) log levels

## Usage

```go
import (
	log "github.com/MeowWolf/eoslog"
)

log.Trace.Printf("...")
log.Debug.Printf("...")
log.Info.Printf("...")
log.Warn.Printf("...")
log.Error.Printf("...")
log.Critical.Printf("...")
```

The individual levels are just `Logger`s underneath so all the normal methods work, e.g. `Print`, `Printf`, `Println`, etc.
