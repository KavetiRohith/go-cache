package iomultiplexer

import "errors"

// ErrInvalidMaxClients is returned when the maxClients is less than 0
var ErrInvalidMaxClients = errors.New("invalid max clients")
