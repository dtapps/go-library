package goqc

import (
	"bytes"
	"context"
	"embed"
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

const (
	DefaultImgSize = 256
)

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

// QrCode 生成二维码
func QrCode(ctx context.Context, imgSize int, qcContent string, qcLevel qrcode.RecoveryLevel) (*QrCodeOperation, error) {

	// 生成二维码
	qr, err := qrcode.New(qcContent, qcLevel)
	if err != nil {
		return nil, err
	}

	// 二维码图片的大小
	codeImgSize := imgSize
	if codeImgSize == 0 {
		codeImgSize = DefaultImgSize
	}

	// 生成二维码图片
	qrImage := qr.Image(codeImgSize)

	// 调整二维码图片大小
	qrImage = resize.Resize(uint(codeImgSize), uint(codeImgSize), qrImage, resize.Lanczos3)

	// 使用 github.com/fogleman/gg 绘制文本
	dc := gg.NewContext(codeImgSize, codeImgSize)
	dc.DrawImage(qrImage, 0, 0)

	return &QrCodeOperation{ctx: ctx, dc: dc}, nil
}

// QrCodeText 生成二维码带文本
func QrCodeText(ctx context.Context, imgSize int, textContent string, qcContent string, qcLevel qrcode.RecoveryLevel) (*QrCodeOperation, error) {

	// 生成二维码
	qr, err := qrcode.New(qcContent, qcLevel)
	if err != nil {
		return nil, err
	}

	// 二维码图片的大小
	codeImgSize := imgSize
	if codeImgSize == 0 {
		codeImgSize = DefaultImgSize
	}

	// 生成二维码图片
	qrImage := qr.Image(codeImgSize)

	// 调整二维码图片大小
	qrImage = resize.Resize(uint(codeImgSize), uint(codeImgSize), qrImage, resize.Lanczos3)

	// 使用 github.com/fogleman/gg 绘制文本
	dc := gg.NewContext(codeImgSize, codeImgSize)
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
	x := (float64(codeImgSize) - textWidth) / 2
	y := (float64(codeImgSize+codeImgSize) + textHeight - textHeight) / 2

	// 绘制文本
	dc.DrawStringAnchored(textContent, x, y-6, 0, 0)

	return &QrCodeOperation{ctx: ctx, dc: dc}, nil
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
