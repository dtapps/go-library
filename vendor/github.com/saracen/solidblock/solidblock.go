package solidblock

import (
	"errors"
	"hash"
	"hash/crc32"
	"io"
	"io/ioutil"
)

var (
	// ErrChecksumMismatch is returned when a file's crc check fails.
	ErrChecksumMismatch = errors.New("checksum mismatch")
)

// Solidblock provides sequential access to files that have been concatenated
// into a single compressed data block.
type Solidblock struct {
	sizes []uint64
	crcs  []uint32

	base io.Reader
	file io.Reader
	crc  hash.Hash32

	target int
	index  int
}

// New returns a new solidblock reader.
func New(r io.Reader, sizes []uint64, crcs []uint32) *Solidblock {
	if len(sizes) != len(crcs) {
		panic("crcs slice needs to be the same length as sizes slice")
	}

	return &Solidblock{
		sizes:  sizes,
		crcs:   crcs,
		target: -1,
		base:   r,
	}
}

// Next advances to the next file entry in solid block.
//
// Calling Next without reading the current file is supported. Only when Read
// is called will decompression occur for current file. Any skipped files will
// still need to be decompressed, but their contents is discarded.
//
// io.EOF is returned at the end of the input.
func (fr *Solidblock) Next() error {
	if fr.target < len(fr.sizes)-1 {
		fr.target++
		return nil
	}
	return io.EOF
}

// Read reads from the current file in solid block.
// It returns (0, io.EOF) when it reaches the end of that file,
// until Next is called to advance to the next file.
func (fr *Solidblock) Read(p []byte) (int, error) {
	if fr.file != nil && fr.index != fr.target {
		// drain current fileReader
		_, err := io.Copy(ioutil.Discard, fr.file)
		if err != nil {
			return 0, err
		}
	}

	if fr.file == nil || fr.index != fr.target {
		// discard until we're at the position we want to be at
		for i := fr.index + 1; i < fr.target; i++ {
			_, err := io.CopyN(ioutil.Discard, fr.base, int64(fr.sizes[i]))
			if err != nil {
				return 0, err
			}
		}

		fr.crc = crc32.NewIEEE()
		fr.file = io.TeeReader(io.LimitReader(fr.base, int64(fr.sizes[fr.target])), fr.crc)
		fr.index = fr.target
	}

	n, err := fr.file.Read(p)
	if err == io.EOF {
		if fr.crc.Sum32() != fr.crcs[fr.index] {
			return n, ErrChecksumMismatch
		}
	}

	return n, err
}

func (fr *Solidblock) Size() int64 {
	if fr.target < 0 {
		return 0
	}
	return int64(fr.sizes[fr.target])
}