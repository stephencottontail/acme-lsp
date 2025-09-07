// +build !plan9

package plumb

import (
	"github.com/fhs/9fans-go/plan9/client"
)

func mountPlumb() {
	fsys, fsysErr = client.MountService("plumb")
}
