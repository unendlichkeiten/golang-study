package sample

import (
	"math/rand"
	"time"

	"github.com/google/uuid"

	"practice/pb"
)

const (
	EMPTY = ""

	NVIDIA = "Nvidia"
	AMD    = "AMD"
)

func init() {
	// set random seed
	randSeed()
}

// randomLaptopBrand
func randomLaptopBrand() string {
	laptopBrands := []string{"Apple", "Dell", "Lenovo"}
	return randomStringFromSet(laptopBrands)
}

// randomLaptopName
func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet([]string{"Macbook Air", "Macbook Pro"})
	case "Dell":
		return randomStringFromSet([]string{"Latitude", "Vostro", "XPS", "Alienware"})
	default:
		return randomStringFromSet([]string{"Thinkpad X1", "Thinkpad P1", "Thinkpad P53"})
	}
}

// randomKeyboardLayout set keyboard layout
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

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9

	resolution := &pb.Screen_Resolution{
		Width:  uint32(width),
		Height: uint32(height),
	}
	return resolution
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

// randomCPUBrand set cpu brand
func randomCPUBrand() string {
	return randomStringFromSet([]string{"Intel", "AMD"})
}

// randomGPUBrand set gpu brand
func randomGPUBrand() string {
	return randomStringFromSet([]string{"Intel", "AMD"})
}

// randomCPUName set cpu name
func randomCPUName(brand string) string {
	intelCPUNames := []string{
		"Xeon E-2286M",
		"Core i9-9980HK",
		"Core i7-9750H",
		"Core i5-9400F",
		"Core i3-1005G1",
	}

	othersCPUNames := []string{
		"Ryzen 7 PRO 2700U",
		"Ryzen 5 PRO 3500U",
		"Ryzen 3 PRO 3200GE",
	}
	if brand == "Intel" {
		return randomStringFromSet(intelCPUNames)
	}
	return randomStringFromSet(othersCPUNames)
}

// randomGPUName set gpu name
func randomGPUName(brand string) string {
	nvidiaGPUNames := []string{
		"RTX 2060",
		"RTX 2070",
		"GTX 1660-Ti",
		"GTX 1070",
	}

	amdGPUNames := []string{
		"RX 590",
		"RX 580",
		"RX 5700-XT",
		"RX Vega-56",
	}

	if brand == NVIDIA {
		return randomStringFromSet(nvidiaGPUNames)
	}

	return randomStringFromSet(amdGPUNames)
}

// randSeed set rand seed
func randSeed() {
	rand.Seed(time.Now().Unix())
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomID() string {
	return uuid.New().String()
}

func randomStringFromSet(set []string) string {
	n := len(set)
	if n == 0 {
		return EMPTY
	}

	return set[rand.Intn(n)]
}
