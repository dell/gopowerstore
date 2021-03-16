package gopowerstore

import (
	"net/http"
	"sync"
)

// MetaDataHeader stores metadata used by other types
type MetaDataHeader struct {
	once     sync.Once // creates the metadata value once.
	metadata http.Header
}
