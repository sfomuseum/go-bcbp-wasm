package bcbp

import (
	"io"
)

type Barcode interface {
	Encode(*BCBP, io.Writer) error
	Decode(io.Reader) (*BCBP, error)
}
