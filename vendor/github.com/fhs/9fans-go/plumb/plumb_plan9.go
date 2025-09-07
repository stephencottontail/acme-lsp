package plumb

import (
	"github.com/fhs/9fans-go/plan9/client"
)

func mountPlumb() {
	fsys = &client.Fsys{Mtpt: "/mnt/plumb"}
	fsysErr = nil
}
