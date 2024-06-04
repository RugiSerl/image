package filter

import "github.com/RugiSerl/image/app/image"

type Filter func(image.Image, ...float64) image.Image
