package structural

import "fmt"

type Image interface {
	Display()
}

type RealImage struct {
	fileName string
}

func (r *RealImage) Load() {
	fmt.Println("Loading", r.fileName)
}

func (r *RealImage) Display() {
	fmt.Println("Displaying", r.fileName)
}

type ProxyImage struct {
	realImage *RealImage
	fileName  string
}

func (p *ProxyImage) Display() {
	if p.realImage == nil {
		p.realImage = &RealImage{fileName: p.fileName}
		p.realImage.Load()
	}
	p.realImage.Display()
}

func Proxy() {
	image := &ProxyImage{fileName: "test_image.jpg"}
	image.Display()
	image.Display()
}
