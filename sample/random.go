package sample

import (
	"github.com/google/uuid"
	"math/rand"
	"pcbook/pb"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"RTX 3090",
			"Titan V",
			"RTX 2080 Ti",
			"Titan Pascal",
			"RTX 2060 SUPER",
			"GTX 1660Ti",
			"GTX 1050Ti",
		)
	}
	return randomStringFromSet(
		"Radeon RX 6800",
		"Radeon RX 5700",
		"Radeon RX 6900 XT",
		"Radeon VII",
		"Radeon RX 5700 XT",
		"Radeon RX Vega 64",
		"Radeon RX Vega 56",
		"Radeon R9 Fury X",
	)
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"i9-7980XE",
			"i7-7700K",
			"i5-7640X",
			"i3-7350K",
			"G4620",
			"G3930T",
			"i7-7920HQ",
			"i5-7440HQ",
		)
	}
	return randomStringFromSet(
		"Athlon 64 FX",
		"Athlon II",
		"Phenom II X3",
		"Ryzen 3 PRO",
		"Ryzen Threadripper",
		"Turion II Dual-Core Mobile",
		"V Series for Notebook PCs",
	)

}

func randomStringFromSet(brands ...string) string {
	n := len(brands)
	if n == 0 {
		return ""
	}
	return brands[rand.Intn(n)]
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	resolution := &pb.Screen_Resolution{
		Width:  uint32(width),
		Height: uint32(height),
	}

	return resolution
}

func randomID() string {
	return uuid.New().String()
}

func randomLaptopBrand() string {
	return randomStringFromSet(
		"Apple",
		"Dell",
		"Lenovo",
	)
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet(
			"Macbook Air",
			"Macbook Pro",
		)
	case "Dell":
		return randomStringFromSet(
			"Latitude",
			"Vostro",
			"XPS",
			"Alienware",
		)
	default:
		return randomStringFromSet(
			"Thinkpad X1",
			"Thinkpad P1",
			"Thinkpad P53",
		)
	}
}
