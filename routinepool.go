package libs

import (
	"github.com/panjf2000/ants/v2"
)

func CreatePool() *ants.Pool {
	p, _ := ants.NewPool(100)
	return p
}
