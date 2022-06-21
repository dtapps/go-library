# Solidblock

Solidblock is a Go library providing `io.Reader`s for solid compression and
codec binding/chaining.

## Solid Compression Reader

Wrapped around a compressed solid block of concatenated files, it provides
sequential access to the files:
```
// file contents
files := [][]byte{
    []byte("file 1\n"),
    []byte("file 2\n"),
}

// file metadata
var metadata struct {
    sizes []uint64
    crcs  []uint32
}
metadata.sizes = []uint64{
    uint64(len(files[0])),
    uint64(len(files[1])),
}
metadata.crcs = []uint32{
    crc32.ChecksumIEEE(files[0]),
    crc32.ChecksumIEEE(files[1]),
}

// Concatenate files to compressed block
block := new(bytes.Buffer)
w := gzip.NewWriter(block)
w.Write(files[0])
w.Write(files[1])
w.Close()

// Open gzip reader to compressed block
r, err := gzip.NewReader(block)
if err != nil {
    panic(err)
}

// Create a new solidblock reader
s := solidblock.New(r, metadata.sizes, metadata.crcs)

for {
    err := s.Next()
    if err == io.EOF {
        break
    }
    if err != nil {
        panic(err)
    }

    io.Copy(os.Stdout, s)
}
```

## Codec Binding

To improve compression, some codecs (such as BCJ2), split data up into multiple
streams that compress better individually. `solidblock.Binder` provides a simple
way to pair together the inputs and outputs of various codecs/readers.

For example:

```
func BCJ2Decoder(inputs []io.Reader) ([]io.Reader, error) {
    // 1. take 4 input readers
    // 2. do magic
    // 3. return 1 reader
}

func GzipDecoder(inputs []io.Reader) ([]io.Reader, error) {
    if len(inputs) != 1 {
        panic("unsupported input configuration")
    }
    r, err := gzip.NewReader(inputs[0])
    return []io.Reader{r}, nil
}

file, err := os.Open("file")
if err != nil {
    panic(err)
}

// Assume file has 4 concatenated streams. 3 of the streams are from a BCJ2 
// encoder, compressed to gzip streams. 1 is the 4th stream of the BCJ2 encoder,
// but left uncompressed.
streams := make([]io.Reader, 4)
streams[0] = io.NewSectionReader(file, 0, 100)
streams[1] = io.NewSectionReader(file, 101, 200)
streams[2] = io.NewSectionReader(file, 201, 300)
streams[3] = io.NewSectionReader(file, 301, 400)

// Create a new binder
binder := solidblock.NewBinder()

// Create gzip decompressors for the 4 initial input streams.
gzip0InputIDs, gzip0OutputIDs := binder.AddCodec(GzipDecoder, 1, 1)
gzip1InputIDs, gzip1OutputIDs := binder.AddCodec(GzipDecoder, 1, 1)
gzip2InputIDs, gzip2OutputIDs := binder.AddCodec(GzipDecoder, 1, 1)

// Create BCJ2 decoder for the 4 gzip decoded streams.
bcj2InputIDs, bcj2outputIDs := binder.AddCodec(BCJ2Decoder, 4, 1)

// Connect initial streams to gzip decoders
binder.Reader(streams[0], gzip0InputIDs[0])
binder.Reader(streams[1], gzip1InputIDs[0])
binder.Reader(streams[2], gzip2InputIDs[0])

// Connect 4th initial stream straight to 4th input of BCJ2 decoder.
binder.Reader(streams[3], bcj2InputIDs[3])

// Pair the 3 gzip output streams to the 1st, 2nd, 3rd input of BCJ2 decoder.
binder.Pair(gzip0OutputIDs[0], bcj2InputIDs[0])
binder.Pair(gzip1OutputIDs[0], bcj2InputIDs[1])
binder.Pair(gzip2OutputIDs[0], bcj2InputIDs[2])

// Create single output to read from
outputs, err := binder.Outputs()
if err != nil {
    panic(err)
}
if len(outputs) != 1 {
    panic("output should only contain one stream")
}

io.Copy(os.Stdout, outputs[0])
```

A picture says 60 lines of code...
```      
                                        +------------+
   concatenated file                    |bcj2 decoder+--->io.Reader
+--------------------+                  +-+--+--+--+-+
|                    |                    ^  ^  ^  ^
|  +--------------+  |   +------------+   |  |  |  |
|  |gzipped stream+------>gzip decoder+---+  |  |  |
|  +--------------+  |   +------------+      |  |  |
|                    |                       |  |  |
|  +--------------+  |   +------------+      |  |  |
|  |gzipped stream+------>gzip decoder+------+  |  |
|  +--------------+  |   +------------+         |  |
|                    |                          |  |
|  +--------------+  |   +------------+         |  |
|  |gzipped stream+------>gzip decoder+---------+  |
|  +--------------+  |   +------------+            |
|                    |                             |
|  +--------------+  |                             |
|  | uncompressed +--------------------------------+
|  |    stream    |  |
|  +--------------+  |
|                    |
+--------------------+

```

