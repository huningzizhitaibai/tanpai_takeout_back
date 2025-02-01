package util

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func ImgSaver(dir string, img image.Image) (string, error) {
	//获取项目运行根目录
	parentPath, err := os.Getwd()
	fmt.Println("parentPath:", parentPath)
	finalPath := parentPath + dir + ".png"
	file, _ := os.Create(finalPath)
	err = png.Encode(file, img)
	return finalPath, err
}
