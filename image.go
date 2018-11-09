package main

import "image"
import "image/color"
import "golang.org/x/tour/pic"

/* 1 官方 image的结构
type Image interface {

    // 颜色模式
	ColorModel() color.Model
	
    // 图片边界
	Bounds() Rectangle
	
    // 某个点的颜色
	At(x, y int) color.Color
}
// */

// 实现 Image的3方法

// ColorModel 应当返回 color.RGBAModel
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds 应当返回一个 image.Rectangle ，例如 `image.Rect(0, 0, w, h)` 。
func (i Image) Bounds() image.Rectangle {
	// 从 i 中获取  
	return image.Rect(0, 0, 200, 200)
}

// 图片生成器的值 v 对应于此次的 `color.RGBA{v, v, 255, 255}`
func (i Image) At(x, y int) color.Color {
    v := uint8(x % (y + 1))
	return color.RGBA{v, v, uint8(255), uint8(255)}
}


type Image struct{}

func main() {
	m := Image{}
	pic.ShowImage(m)
	fmt.Println()
}
