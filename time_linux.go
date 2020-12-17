// +build linux

package fasttime

import (
	_ "unsafe" // for runtime.walltime
)

//go:noescape
//go:linkname walltime runtime.walltime
func walltime() (sec int64, nsec int32)
