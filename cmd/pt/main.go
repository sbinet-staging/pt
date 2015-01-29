package main

import (
	"github.com/fogleman/pt"
	"image"
	"image/png"
	// "image/jpeg"
	"os"
	"runtime"
)

func scene1() (pt.Scene, pt.Camera) {
	// file, err := os.Open("texture.jpg")
	// if err != nil {
	// 	panic(err)
	// }
	// im, err := jpeg.Decode(file)
	// if err != nil {
	// 	panic(err)
	// }
	scene := pt.Scene{}
	scene.AddShape(pt.NewSphere(pt.Vector{0, 0, 0}, 1, pt.HexColor(0x334D5C), pt.Material{}, nil))
	scene.AddShape(pt.NewSphere(pt.Vector{-2, 0, -2}, 1, pt.HexColor(0x45B29D), pt.Material{}, nil))
	scene.AddShape(pt.NewSphere(pt.Vector{-2, 0, 2}, 1, pt.HexColor(0xEFC94C), pt.Material{}, nil))
	scene.AddShape(pt.NewSphere(pt.Vector{2, 0, -2}, 1, pt.HexColor(0xE27A3F), pt.Material{}, nil))
	scene.AddShape(pt.NewSphere(pt.Vector{2, 0, 2}, 1, pt.HexColor(0xDF5A49), pt.Material{}, nil))
	scene.AddShape(pt.NewPlane(pt.Vector{0, -1, 0}, pt.Vector{0, 1, 0}, pt.Color{1, 1, 1}, pt.Material{}))
	scene.AddLight(pt.NewSphere(pt.Vector{-1, 3, -1}, 0.25, pt.Color{1, 1, 1}, pt.Material{}, nil))
	scene.AddLight(pt.NewSphere(pt.Vector{1, 3, -1}, 0.25, pt.Color{1, 1, 1}, pt.Material{}, nil))
	scene.AddLight(pt.NewSphere(pt.Vector{-1, 3, 1}, 0.25, pt.Color{1, 1, 1}, pt.Material{}, nil))
	scene.AddLight(pt.NewSphere(pt.Vector{1, 3, 1}, 0.25, pt.Color{1, 1, 1}, pt.Material{}, nil))
	// scene.AddLight(pt.NewCube(pt.Vector{-5, 8, -5}, pt.Vector{5, 9, 5}, pt.Color{1, 1, 1}, pt.Material{}))
	camera := pt.LookAt(pt.Vector{0, 6, -8}, pt.Vector{0, 0, -1.5}, pt.Vector{0, 1, 0}, 40)
	return scene, camera
}

func scene2() (pt.Scene, pt.Camera) {
	scene := pt.Scene{}
	for x := 0.; x < 3; x++ {
		for z := 0.; z < 3; z++ {
			scene.AddShape(pt.NewSphere(pt.Vector{x*2 - 2, 0, z*2 - 2}, 1, pt.HexColor(0xEFC94C), pt.Material{}, nil))
		}
	}
	for x := 0.; x < 2; x++ {
		for z := 0.; z < 2; z++ {
			scene.AddShape(pt.NewSphere(pt.Vector{x*2 - 1, 1.414, z*2 - 1}, 1, pt.HexColor(0xE27A3F), pt.Material{}, nil))
		}
	}
	scene.AddShape(pt.NewSphere(pt.Vector{0, 1.414 * 2, 0}, 1, pt.HexColor(0xDF5A49), pt.Material{}, nil))
	scene.AddShape(pt.NewCube(pt.Vector{7, -1, -8}, pt.Vector{8, 10, 8}, pt.Color{1, 1, 1}, pt.Material{}))
	scene.AddShape(pt.NewCube(pt.Vector{-8, -1, -8}, pt.Vector{-7, 10, 8}, pt.Color{1, 1, 1}, pt.Material{}))
	scene.AddShape(pt.NewCube(pt.Vector{-8, -1, 7}, pt.Vector{8, 10, 8}, pt.HexColor(0xE27A3F), pt.Material{}))
	scene.AddShape(pt.NewCube(pt.Vector{-8, -1, -8}, pt.Vector{8, 10, -7}, pt.HexColor(0xE27A3F), pt.Material{}))
	scene.AddShape(pt.NewPlane(pt.Vector{0, -1, 0}, pt.Vector{0, 1, 0}, pt.Color{1, 1, 1}, pt.Material{}))
	scene.AddLight(pt.NewSphere(pt.Vector{-1, 8, -1}, 0.25, pt.Color{0.6, 0.6, 0.6}, pt.Material{}, nil))
	scene.AddLight(pt.NewSphere(pt.Vector{1, 8, -1}, 0.25, pt.Color{0.6, 0.6, 0.6}, pt.Material{}, nil))
	scene.AddLight(pt.NewSphere(pt.Vector{-1, 8, 1}, 0.25, pt.Color{0.6, 0.6, 0.6}, pt.Material{}, nil))
	scene.AddLight(pt.NewSphere(pt.Vector{1, 8, 1}, 0.25, pt.Color{0.6, 0.6, 0.6}, pt.Material{}, nil))
	camera := pt.LookAt(pt.Vector{-6, 6, 0}, pt.Vector{0, 0, 0}, pt.Vector{0, 1, 0}, 55)
	return scene, camera
}

func scene3() (pt.Scene, pt.Camera) {
	scene := pt.Scene{}
	for x := 0; x < 10; x++ {
		for z := 0; z < 10; z++ {
			scene.AddShape(pt.NewSphere(pt.Vector{float64(x) - 4.5, 0, float64(z) - 4.5}, 0.45, pt.HexColor(0xEFC94C), pt.Material{}, nil))
		}
	}
	scene.AddShape(pt.NewPlane(pt.Vector{}, pt.Vector{0, 1, 0}, pt.Color{1, 1, 1}, pt.Material{}))
	// scene.AddLight(pt.NewCube(pt.Vector{-5, 8, -5}, pt.Vector{5, 9, 5}, pt.Color{1, 1, 1}, pt.Material{}))
	scene.AddLight(pt.NewSphere(pt.Vector{0, 2, 0}, 0.25, pt.Color{1, 1, 1}, pt.Material{}, nil))
	camera := pt.LookAt(pt.Vector{0, 5, -8}, pt.Vector{0, 0, -2}, pt.Vector{0, 1, 0}, 45)
	return scene, camera
}

func save(path string, im image.Image) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = png.Encode(file, im)
	if err != nil {
		panic(err)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	scene, camera := scene1()
	im := pt.Render(&scene, &camera, 2560/4, 1440/4, 4, 16, 8)
	save("out.png", im)
}
