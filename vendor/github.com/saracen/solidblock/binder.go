package solidblock

import (
	"errors"
	"fmt"
	"io"
)

var (
	// ErrInputIsUnbound is returned when an input hasn't been binded to either
	// a reader/paired without an output.
	ErrInputIsUnbound = errors.New("input is unbound")

	// ErrUnexpectedOutputCount is returned when the amount of io.Readers
	// returned from a codec handler doesn't match the amount specified when
	// adding the codec.
	ErrUnexpectedOutputCount = errors.New("unexpected output count")
)

type reader struct {
	Name string
	R    io.Reader
}

type codec struct {
	fn func([]io.Reader) ([]io.Reader, error)

	inIndexes  []int
	outIndexes []int
}

// Binder holds information regarding codecs, their inputs/outputs and how they
// join together.
type Binder struct {
	numInStreams  int
	numOutStreams int

	in  []*reader
	out []*reader

	codecs []*codec
}

// NewBinder returns a new binder.
func NewBinder() *Binder {
	return &Binder{}
}

// AddCodec adds a handler function for processing information from input(s) and
// producing output(s).
func (b *Binder) AddCodec(fn func([]io.Reader) ([]io.Reader, error), inputs, outputs int) (in, out []int) {
	c := &codec{fn: fn}
	b.in = append(b.in, make([]*reader, inputs)...)
	b.out = append(b.out, make([]*reader, outputs)...)

	for i := 0; i < inputs; i++ {
		c.inIndexes = append(c.inIndexes, b.numInStreams+i)
	}
	for i := 0; i < outputs; i++ {
		c.outIndexes = append(c.outIndexes, b.numOutStreams+i)
	}

	b.numInStreams += inputs
	b.numOutStreams += outputs

	b.codecs = append(b.codecs, c)

	return c.inIndexes, c.outIndexes
}

// Reader binds a reader to an in stream.
func (b *Binder) Reader(r io.Reader, in int) {
	if in < 0 || in >= len(b.in) {
		return
	}
	b.in[in] = &reader{fmt.Sprintf("In: %v", in), r}
}

// Pair pairs two streams, binding an in stream to an out stream.
func (b *Binder) Pair(in int, out int) {
	if in < 0 || in >= len(b.in) || out < 0 || out >= len(b.out) {
		return
	}

	if b.out[out] == nil {
		b.out[out] = &reader{fmt.Sprintf("Bind %v:%v", in, out), nil}
	}

	b.in[in] = b.out[out]
}

// Outputs returns any unbound output readers to ready from.
func (b *Binder) Outputs() ([]io.Reader, error) {
	var unbound []io.Reader

	for i := range b.codecs {
		var ins []io.Reader
		for _, num := range b.codecs[i].inIndexes {
			if b.in[num] == nil || b.in[num].R == nil {
				return unbound, ErrInputIsUnbound
			}
			ins = append(ins, b.in[num].R)
		}

		outs, err := b.codecs[i].fn(ins)
		if err != nil {
			return nil, err
		}

		if len(outs) != len(b.codecs[i].outIndexes) {
			return unbound, ErrUnexpectedOutputCount
		}
		for j, num := range b.codecs[i].outIndexes {
			if b.out[num] == nil {
				b.out[num] = &reader{fmt.Sprintf("Out %v", outs), nil}
				unbound = append(unbound, outs[j])
			}
			b.out[num].R = outs[j]
		}
	}

	return unbound, nil
}
