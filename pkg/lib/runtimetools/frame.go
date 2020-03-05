// Copyright (c) Facebook, Inc. and its affiliates.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package runtimetools

import (
	"runtime"
)

// Frame returns information about a frame at depth "skipFrames"
// of the call stack. It could be handy for logs (to see what
// line of the source code caused the message).
func Frame(skipFrames uint) runtime.Frame {
	// exclude runtime.Callers() and Frame() from the result:
	skipFrames += 2

	pcs := make([]uintptr, skipFrames+2)
	n := runtime.Callers(0, pcs)

	frames := runtime.CallersFrames(pcs[:n])
	for i := uint(0); i <= skipFrames; i++ {
		frame, more := frames.Next()
		if !more {
			break
		}
		if i == skipFrames {
			return frame
		}
	}

	return runtime.Frame{}
}