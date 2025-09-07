// +build !plan9

package acme

import "github.com/fhs/9fans-go/plan9/client"

func mountAcme() {
	fs, err := client.MountService("acme")
	fsys = fs
	fsysErr = err
}
