// Package python implements the "python" runtime.
package python

import (
	"strings"

	"github.com/apex/apex/function"
)

// Runtime for inference.
const (
	Runtime   = "python"
	Runtime27 = "python2.7"
	Runtime36 = "python3.6"
)

func init() {
	function.RegisterPlugin(Runtime, &Plugin{})
	function.RegisterPlugin(Runtime27, &Plugin{})
	function.RegisterPlugin(Runtime36, &Plugin{})
}

// Plugin implementation.
type Plugin struct{}

// Open adds python defaults.
func (p *Plugin) Open(fn *function.Function) error {
	if !strings.HasPrefix(fn.Runtime, "python") {
		return nil
	}

	// Support "python" for backwards compat.
	if fn.Runtime == "python" {
		fn.Runtime = "python2.7"
	}

	if fn.Handler == "" {
		fn.Handler = "main.handle"
	}

	return nil
}
