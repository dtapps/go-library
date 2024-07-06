//go:build jsoniter

package json

import jsoniter "github.com/json-iterator/go"

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
	// Marshal 封装了 gojson/json 包的 Marshal 函数。
	Marshal = json.Marshal
	// Unmarshal 封装了 gojson/json 包的 Unmarshal 函数。
	Unmarshal = json.Unmarshal
	// MarshalIndent 封装了 gojson/json 包的 MarshalIndent 函数。
	MarshalIndent = json.MarshalIndent
	// NewDecoder 是一个函数，返回一个新的自定义 Decoder 实例。
	NewDecoder = func(r io.Reader) *Decoder {
		return &Decoder{decoder: json.NewDecoder(r)}
	}
	// NewEncoder 是一个函数，返回一个新的自定义 Encoder 实例。
	NewEncoder = func(w io.Writer) *Encoder {
		return &Encoder{encoder: json.NewEncoder(w)}
	}
)

// Decoder 是一个用于解码 JSON 的自定义结构体。
type Decoder struct {
	decoder *json.Decoder
}

// Decode 读取下一个 JSON 编码值并将其存储在由 v 指向的值中。
func (d *Decoder) Decode(v interface{}) error {
	return d.decoder.Decode(v)
}

// Encoder 是一个用于编码 JSON 的自定义结构体。
type Encoder struct {
	encoder *json.Encoder
}

// Encode 将 v 的 JSON 编码写入流中。
func (e *Encoder) Encode(v interface{}) error {
	return e.encoder.Encode(v)
}
