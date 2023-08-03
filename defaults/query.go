package defaults

import "fmt"

type query struct{}

var Query query

func (q query) Limit() uint64 {
	return 10
}

func (q query) LimitString() string {
	return fmt.Sprint(q.Limit())
}

func (q query) Offset() uint64 {
	return 0
}

func (q query) OffsetString() string {
	return fmt.Sprint(q.Offset())
}
