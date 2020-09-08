package image

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

// ChangeImageExtension srcFileExtの拡張子からdstFileExtの拡張子に変換する
func ChangeImageExtension(imageFilePath string, srcFileExt string, dstFileExt string) error {

	// 入力ファイルの拡張子がjpgとpng以外の場合, エラーを返す
	if srcFileExt != ".jpg" && srcFileExt != ".png" {
		err := fmt.Errorf("Invalid source Extension: %s", srcFileExt)
		return err
	}

	// 出力ファイルの拡張子がjpgとpng以外の場合, エラーを返す
	if dstFileExt != ".jpg" && dstFileExt != ".png" {
		err := fmt.Errorf("Invalid dest Extension: %s", dstFileExt)
		return err
	}

	// 入力ファイルと出力ファイルの拡張子が一致したら, エラーを返す
	if srcFileExt == dstFileExt {
		err := fmt.Errorf("Same extension: %s", srcFileExt)
		return err
	}

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
