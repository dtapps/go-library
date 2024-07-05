package goqc

import (
	"bytes"
	"context"
	"embed"
	_ "embed"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
	"image/png"
	"io"
	"path/filepath"
)

//go:embed simhei.ttf
var fontContent embed.FS
var fontPath = "simhei.ttf"
var fontSize = 16.0

func loadFontPath(dc *gg.Context) error {
	// 设置字体和文本属性
	err := dc.LoadFontFace(fontPath, fontSize)
	if err != nil {
		return err
	}
	return nil
}

func loadFontEmbed(dc *gg.Context) error {
	// 读取嵌入的字体文件
	fontBytes, err := fontContent.ReadFile(fontPath)
	if err != nil {
		return err
	}
	// 设置字体
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		return err
	}
	// 文本属性
	face := truetype.NewFace(f, &truetype.Options{
		Size: fontSize,
	})
	dc.SetFontFace(face)
	return nil
}

// QrCodeOperation 二维码操作
type QrCodeOperation struct {
	ctx context.Context
	dc  *gg.Context
}

// QrCodeText 生成二维码带文本
func QrCodeText(ctx context.Context, textContent string, qcContent string, qcLevel qrcode.RecoveryLevel) (*QrCodeOperation, error) {

	// 生成二维码
	qr, err := qrcode.New(qcContent, qcLevel)
	if err != nil {
		return nil, err
	}

	// 设置二维码图片的大小
	qrSize := 256

	// 生成二维码图片
	qrImage := qr.Image(qrSize)

	// 调整二维码图片大小
	qrImage = resize.Resize(uint(qrSize), uint(qrSize), qrImage, resize.Lanczos3)

	// 使用 github.com/fogleman/gg 绘制文本
	dc := gg.NewContext(qrSize, qrSize)
	dc.DrawImage(qrImage, 0, 0)

	// 设置字体和文本属性
	err = loadFontEmbed(dc)
	if err != nil {
		return nil, err
	}

	// 设置文本颜色
	dc.SetRGB(0, 0, 0)

	// 计算文本尺寸
	textWidth, textHeight := dc.MeasureString(textContent)

	// 计算文本位置
	x := (float64(qrSize) - textWidth) / 2
	y := (float64(qrSize+qrSize) + textHeight - textHeight) / 2

	// 绘制文本
	dc.DrawStringAnchored(textContent, x, y-6, 0, 0)

	return &QrCodeOperation{ctx: ctx, dc: dc}, err
}

// SavePNG 保存图片
func (o *QrCodeOperation) SavePNG(filePath, fileName string) error {
	if fileName == "" {
		return o.dc.SavePNG(filePath)
	} else {
		path := filepath.Join(filePath, "/", fileName)
		return o.dc.SavePNG(path)
	}
}

// Encode 返回图片字节
func (o *QrCodeOperation) Encode() ([]byte, error) {
	img := o.dc.Image()

	encoder := png.Encoder{CompressionLevel: png.BestCompression}

	var b bytes.Buffer
	err := encoder.Encode(&b, img)

	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// EncodePNG 返回图片编码
func (o *QrCodeOperation) EncodePNG(w io.Writer) error {
	return o.dc.EncodePNG(w)
}
