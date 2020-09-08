package image

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

// ChangeImageExtension srcFileExtの拡張子からdstFileExtの拡張子に変換する
func ChangeImageExtension(imageFilePath string, srcFileExt string, dstFileExt string) error {
	imageFile, err := os.Open(imageFilePath)
	if err != nil {
		return err
	}
	defer imageFile.Close()

	img, _, err := image.Decode(imageFile)
	if err != nil {
		return err
	}

	// 変換後の画像ファイルのパスを生成
	extIndex := strings.LastIndex(imageFilePath, srcFileExt)
	dstImageFilePath := imageFilePath[:extIndex] + dstFileExt

	dstImageFile, err := os.Create(dstImageFilePath)
	if err != nil {
		return err
	}
	defer dstImageFile.Close()

	if dstFileExt == ".jpg" {
		jpeg.Encode(dstImageFile, img, nil) // jpegに変換
	} else if dstFileExt == ".png" {
		png.Encode(dstImageFile, img) // pngに変換
	}

	return nil
}
