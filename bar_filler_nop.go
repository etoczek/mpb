package mpb

import (
	"io"

	"github.com/etoczek/mpb/v8/decor"
)

// NopStyle provides BarFillerBuilder which builds NOP BarFiller.
func NopStyle() BarFillerBuilder {
	return BarFillerBuilderFunc(func() BarFiller {
		return BarFillerFunc(func(io.Writer, decor.Statistics) {})
	})
}
