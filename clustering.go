package main

func l2distance(v1, v2 [128]float32) (distance float32) {
	distance = 0
	for i := 0; i < 128; i++ {
		distance += (v1[i] - v2[i]) * (v1[i] - v2[i])
	}
	return distance
}

func getClasters(dts []data) (cls clasters) {
	cls = append(cls, []*data{&dts[0]})

	for i := 1; i < len(dts); i++ {
		found := false
		for cl := range cls {
			distance := l2distance(cls[cl][0].Vector, dts[i].Vector)
			if distance < DEFAULT_DISTANCE {
				cls[cl] = append(cls[cl], &dts[i])
				found = true
				break
			}
		}
		if found == false {
			cls = append(cls, []*data{&dts[i]})
		}
	}

	return cls
}
