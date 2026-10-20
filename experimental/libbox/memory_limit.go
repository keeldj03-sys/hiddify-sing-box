package libbox

import "runtime/debug"

// SetMemoryLimit toggles the Go runtime soft memory limit (GC pacing).
// hiddify-core calls this on iOS or when the caller does not explicitly
// disable it; -1 removes the limit.
func SetMemoryLimit(enabled bool) {
	if enabled {
		debug.SetMemoryLimit(1024 * 1024 * 1024)
	} else {
		debug.SetMemoryLimit(-1)
	}
}
