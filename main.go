package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"os"
	"raytracing/rt"
	"time"

	"golang.org/x/image/bmp"
)

const (
	aspectRatio    = 16.0 / 9.0
	imageWidth     = 320
	imageHeight    = int(imageWidth / aspectRatio)
	samperPerPixel = 100
	maxDepth       = 50
)

var (
	threadNum = flag.Int("n", 4, "number of goroutine.")
)

type fragment struct {
	x    int
	y    int
	rgba color.RGBA
}

func rayColor(rd *rand.Rand, r rt.Ray, hitList *rt.HitableList, depth int) rt.Vector3 {
	rec := &rt.HitRecord{}
	if depth <= 0 {
		return rt.Vec3(0.0, 0.0, 0.0)
	}
	if hitList.Hit(r, 0.001, rt.INFINTE, rec) {
		if ok, attenuation, scattered := rec.Mat.Scatter(r, rec, rd); ok {
			return rt.Vec3MulVec3(rayColor(rd, scattered, hitList, depth-1), attenuation)
		}
		return rt.Vec3(0.0, 0.0, 0.0)
	}

	unitDirection := r.Dir.Normalize()
	t := 0.5 * (unitDirection.Y + 1.0)
	return rt.Vec3Add(rt.Vec3(1.0, 1.0, 1.0).MulFloat(1.0-t), rt.Vec3(0.5, 0.7, 1.0).MulFloat(t))
}

func scence() *rt.HitableList {
	world := rt.NewHitList()
	//ground
	world.Add(rt.NewSphere(rt.Vec3(0.0, -1000, -1.0), 1000, rt.NewLamertian(rt.Vec3(0.8, 0.8, 0.8))))

	world.Add(rt.NewSphere(rt.Vec3(-4.0, 1.0, -0.3), 1.0, rt.NewLamertian(rt.Vec3(0.4, 0.2, 0.1))))
	world.Add(rt.NewSphere(rt.Vec3(4.0, 1.0, -0.2), 1.0, rt.NewMetal(rt.Vec3(0.7, 0.6, 0.5), 0.2)))
	world.Add(rt.NewSphere(rt.Vec3(0.0, 1.0, 0.0), 1.0, rt.NewDielectric(1.5)))

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			// source := rand.NewSource(time.Now().UnixNano())
			// generator := rand.New(source)
			chooseMat := rt.Random()
			center := rt.Vec3(float64(a)+0.9*rt.Random(), 0.2, float64(b)+0.9*rt.Random())

			if center.Sub(rt.Vec3(4.0, 0.2, 0)).Length() > 0.9 {
				var sphereMat rt.Material
				if chooseMat < 0.8 {
					albedo := rt.Vec3Random().MulVec3(rt.Vec3Random())
					sphereMat = rt.NewLamertian(albedo)
					world.Add(rt.NewSphere(center, 0.2, sphereMat))
				} else if chooseMat < 0.95 {
					albedo := rt.Vec3RandomBetween(0.5, 1.0)
					fuzz := rt.RandomBetween(0.0, 0.5)
					sphereMat = rt.NewMetal(albedo, fuzz)
					world.Add(rt.NewSphere(center, 0.2, sphereMat))
				} else {

					sphereMat = rt.NewDielectric(1.5)
					world.Add(rt.NewSphere(center, 0.2, sphereMat))
				}
			}
		}
	}
	//return world
	nHitb := rt.NewBVH(world)
	return rt.NewHitList(nHitb)
}

func main() {
	flag.Parse()
	//defer profile.Start().Stop()
	startTime := time.Now()
	file, err := os.Create("./output.bmp")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	image := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	world := scence()
	lookFrom := rt.Vec3(13.0, 2.0, 3.0)
	//lookFrom := rt.Vec3(0.0, 0.0, 3.0)
	lookAt := rt.Vec3(0.0, 0.0, 0.0)
	camera := rt.NewCamera(lookFrom, lookAt, rt.Vec3(0.0, 1.0, 0.0), 20.0,
		aspectRatio,
		0.01,
		10.0)

	frags := make(chan fragment, *threadNum)
	stepSize := imageHeight / *threadNum
	left := imageHeight
	for left >= 1 {
		from := left - 1

		to := left - stepSize
		if to < 0 {
			to = 0
		}
		left -= (from - to + 1)
		go func(from, to int) {
			source := rand.NewSource(time.Now().UnixNano())
			generator := rand.New(source)

			//fmt.Println("from", from, "to", to)
			for y := from; y >= to; y-- {
				//fmt.Printf("%d lines remain.....\r", y)
				for x := 0; x < imageWidth; x++ {
					pixelCol := rt.Vec3(1.0, 1.0, 1.0)
					for s := 0; s < samperPerPixel; s++ {
						var u float64 = (float64(x) + rt.RandomNew(generator)) / float64(imageWidth-1.0)
						var v float64 = (float64(y) + rt.RandomNew(generator)) / float64(imageHeight-1.0)
						ray := camera.GetRay(u, v, generator)
						pixelCol = pixelCol.Add(rayColor(generator, ray, world, maxDepth))
					}
					frags <- fragment{x, imageHeight - y - 1, rt.SampleColorToRGBA(pixelCol, samperPerPixel)}

				}
			}
		}(from, to)

	}

	var count = 0
	total := imageHeight * imageWidth
	for frag := range frags {
		image.SetRGBA(frag.x, frag.y, frag.rgba)
		count++
		if count%1000 == 0 {
			fmt.Printf("%f%% finished....\r", float64(count)/float64(total)*100.0)
		}
		if count == total {
			close(frags)
		}
	}
	fmt.Print("\n")
	bmp.Encode(file, image)
	fmt.Println("Render time:", time.Now().Sub(startTime))
}
