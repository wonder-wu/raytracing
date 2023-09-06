package rt

type HitRecord struct {
	P         Vector3
	Normal    Vector3
	T         float64
	FrontFace bool
	Mat       Material
}

func (h *HitRecord) SetFaceNoraml(r Ray, outwardNormal Vector3) {
	h.FrontFace = (Vec3Dot(r.Dir, outwardNormal) < 0)
	if h.FrontFace {
		h.Normal = outwardNormal
	} else {
		h.Normal = outwardNormal.MulFloat(-1.0)
	}
}

type Hittable interface {
	Hit(r Ray, tMin, tMax float64, hitRecord *HitRecord) bool
	BoundingBox() AABB
}

func HitableBoxCompare(a, b Hittable) bool {

	boxA := a.BoundingBox()
	boxB := b.BoundingBox()
	return BoxCompare(boxA, boxB, 1)
}

type HitableList struct {
	objects    hitableSlice
	compareIdx int //axis x=0 y=1 z=2
}

func NewHitList(objects ...Hittable) *HitableList {
	if len(objects) == 0 {
		return &HitableList{}
	}
	hl := HitableList{}
	hl.objects = append(hl.objects, objects...)
	return &hl
}

func NewHitListWithCompareIdx(idx int, objects ...Hittable) *HitableList {
	hl := NewHitList(objects...)
	hl.compareIdx = idx
	return hl
}

func (hl *HitableList) Add(object Hittable) {
	hl.objects = append(hl.objects, object)
}
func (hl *HitableList) Clear() {
	hl.objects = nil
}
func (hl *HitableList) Hit(r Ray, tMin, tMax float64, hitRecord *HitRecord) bool {
	tempRecord := &HitRecord{}
	var hitAnything = false
	var closestSoFar = tMax

	for i := 0; i < len(hl.objects); i++ {
		if hl.objects[i].Hit(r, tMin, closestSoFar, tempRecord) {
			hitAnything = true
			closestSoFar = tempRecord.T
			*hitRecord = *tempRecord

		}
	}
	//fmt.Println(hitRecord.Normal)
	return hitAnything
}

func (hl *HitableList) BoundingBox(outputBox AABB) (bool, AABB) {
	if len(hl.objects) == 0 {
		return false, AABB{}
	}

	tempBox := AABB{}
	first := true

	var finalBox AABB

	for _, hb := range hl.objects {
		tempBox = hb.BoundingBox()
		if first {
			finalBox = tempBox
			first = false
		} else {
			finalBox = SurroundingBox(outputBox, tempBox)
		}
	}

	return true, finalBox

}

//hitablelist sortable
type hitableSlice []Hittable

func (hl hitableSlice) Len() int {
	return len(hl)
}

func (hl hitableSlice) Swap(i, j int) {
	hl[i], hl[j] = hl[j], hl[i]
}

func (hl hitableSlice) Less(i, j int) bool {
	var boxA AABB
	var boxB AABB
	boxA = hl[i].BoundingBox()
	boxB = hl[j].BoundingBox()
	return BoxCompare(boxA, boxB, 1)
}
