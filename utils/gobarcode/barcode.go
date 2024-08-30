package gobarcode

import (
	"bytes"
	"context"
	"embed"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"image/png"
	"io"
	"path/filepath"
)

//go:embed simhei.ttf
var fontContent embed.FS
var fontPath = "simhei.ttf"
var fontSize = 16.0

const (
	DefaultImgWidth  = 350
	DefaultImgHeight = 70
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

// Operation 二维码操作
type Operation struct {
	ctx context.Context
	dc  *gg.Context
}

// QrCode 生成条形码
func QrCode(ctx context.Context, imgWidth int, imgHeight int, qcContent string) (*Operation, error) {

	// 创建一个code128编码的 Barcode
	barcodeImg, _ := code128.Encode(qcContent)

	// 条形码图片的大小
	barcodeImgWidth := imgWidth
	if barcodeImgWidth == 0 {
		barcodeImgWidth = DefaultImgWidth
	}
	barcodeImgHeight := imgHeight
	if barcodeImgHeight == 0 {
		barcodeImgHeight = DefaultImgHeight
	}

	// 设置图片像素大小
	qrCode, _ := barcode.Scale(barcodeImg, barcodeImgWidth, barcodeImgHeight)

	// 使用 github.com/fogleman/gg 绘制文本
	dc := gg.NewContext(barcodeImgWidth, barcodeImgHeight)
	dc.DrawImage(qrCode, 0, 0)

	return &Operation{ctx: ctx, dc: dc}, nil
}

// QrCodeText 生成条形码带文本
func QrCodeText(ctx context.Context, imgWidth int, imgHeight int, textContent string, qcContent string) (*Operation, error) {

	// 创建一个code128编码的 Barcode
	barcodeImg, _ := code128.Encode(qcContent)

	// 条形码图片的大小
	barcodeImgWidth := imgWidth
	if barcodeImgWidth == 0 {
		barcodeImgWidth = DefaultImgWidth
	}
	barcodeImgHeight := imgHeight
	if barcodeImgHeight == 0 {
		barcodeImgHeight = DefaultImgHeight
	}
	ggBarcodeImgWidth := barcodeImgWidth
	ggBarcodeImgHeight := barcodeImgHeight + 25

	// 设置图片像素大小
	qrCode, _ := barcode.Scale(barcodeImg, barcodeImgWidth, barcodeImgHeight)

	// 使用 github.com/fogleman/gg 绘制文本
	dc := gg.NewContext(ggBarcodeImgWidth, ggBarcodeImgHeight)
	dc.DrawImage(qrCode, 0, 0)

	// 设置字体和文本属性
	err := loadFontEmbed(dc)
	if err != nil {
		return nil, err
	}

	// 设置文本颜色
	dc.SetRGB(0, 0, 0)

	// 计算文本尺寸
	textWidth, textHeight := dc.MeasureString(textContent)

	// 计算文本位置
	x := (float64(ggBarcodeImgWidth) - textWidth) / 2
	y := (float64(ggBarcodeImgHeight+ggBarcodeImgHeight) + textHeight - textHeight) / 2

	// 绘制文本
	dc.DrawStringAnchored(textContent, x, y-6, 0, 0)

	return &Operation{ctx: ctx, dc: dc}, nil
}

// SavePNG 保存图片
func (o *Operation) SavePNG(filePath, fileName string) error {
	if fileName == "" {
		return o.dc.SavePNG(filePath)
	} else {
		path := filepath.Join(filePath, "/", fileName)
		return o.dc.SavePNG(path)
	}
}

// Encode 返回图片字节
func (o *Operation) Encode() ([]byte, error) {
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
func (o *Operation) EncodePNG(w io.Writer) error {
	return o.dc.EncodePNG(w)
}
