package rt

import (
	"sort"
)

type BVHNode struct {
	box   AABB
	left  Hittable
	right Hittable
}

func (b *BVHNode) Hit(r Ray, tMin, tMax float64, hitRecord *HitRecord) bool {
	if b.box.Hit(r, tMin, tMax) {
		//continue to check left and right
		hitLeft := b.left.Hit(r, tMin, tMax, hitRecord)
		var hitRight bool
		if hitLeft {
			hitRight = b.right.Hit(r, tMin, hitRecord.T, hitRecord)

		} else {
			hitRight = b.right.Hit(r, tMin, tMax, hitRecord)
		}

		return hitRight || hitLeft
	}
	return false
}

func (b *BVHNode) BoundingBox() AABB {
	return b.box
}

func NewBVH(hl *HitableList) *BVHNode {
	return newBVH(hl.objects, 0, len(hl.objects)-1)
}

var i = 0

func newBVH(objects hitableSlice, start, end int) *BVHNode {
	i++
	//fmt.Println("new loop")
	var node BVHNode
	span := end - start + 1
	if span == 0 {
		panic("fuuuuc")
	} else if span == 1 {
		//fmt.Println("span = 1")
		node.left = objects[start]
		node.right = objects[start]

	} else if span == 2 {
		//fmt.Println("2 span:", start, end)
		if HitableBoxCompare(objects[start], objects[end]) {
			node.left = objects[start]
			node.right = objects[start+1]
		} else {
			node.right = objects[start]
			node.left = objects[start+1]
		}

	} else {
		//bigger than 3
		//fmt.Println("3 span:", start, end)
		sort.Sort(objects[start:end])
		mid := start + span/2

		node.left = newBVH(objects, start, mid)
		node.right = newBVH(objects, mid, end)

	}
	//fmt.Println("level", i)
	//wrap all
	boxLeft := node.left.BoundingBox()
	boxRight := node.right.BoundingBox()
	node.box = SurroundingBox(boxLeft, boxRight)
	//fmt.Println("left box", boxLeft, "right box", boxRight, "final box:", node.box)
	return &node

}
