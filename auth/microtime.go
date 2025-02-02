package auth // import "github.com/Buccaneer69/wc-api-go/auth"

import (
	"strconv"
	"time"
)

// MicroTimer ...
type MicroTimer struct {
}

// Get current micro time
func (m *MicroTimer) Get() string {
	mc := time.Now().UnixNano()
	return strconv.FormatInt(mc, 10)
}
