package repl

import (
	"time"
)

type Tracer struct {
	Index      int
	Epoch_time int64
	Datetime   time.Time
	Body       []byte
}
