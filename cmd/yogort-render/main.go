package main

import (
	"github.com/martinplaner/yogort"
	"fmt"
	"image"
	"os"
	"image/png"
)

func main() {

	scene := yogort.Scene{
		Objects: []yogort.Object{
			yogort.Sphere{
				Position:yogort.Vec3{3, 1, 10},
				Radius: 0.5,
			},
			yogort.Sphere{
				Position:yogort.Vec3{7, 5, 0},
				Radius: 1,
			},
			yogort.Sphere{
				Position:yogort.Vec3{2, 5, 10},
				Radius: 1.5,
			},
			yogort.Sphere{
				Position:yogort.Vec3{6, 3, 20},
				Radius: 2,
			},
			//yogort.Plane{
			//	Position:yogort.Vec3{0, 0, 0},
			//	Normal:yogort.Vec3{0, 1, 0},
			//},
		},
		Lights:[]yogort.Light{
			yogort.Light{
				Position:yogort.Vec3{2, 1, 3},
				Brightness:0.8,
			},
		},
	}
	camera := yogort.Camera{
		Position:yogort.Vec3{4, 3, -30},
	}

	img := image.NewRGBA(image.Rect(0, 0, 1600, 1200))

	camera.Render(scene, img)

	err := saveImage(img, "out.png")
	if err != nil {
		fmt.Println("Could not save image:", err)
	}
}

func saveImage(img image.Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create image file: %v", err)
	}

	err = png.Encode(file, img)
	if err != nil {
		return fmt.Errorf("could not encode image: %v", err)
	}

	return nil
}
