//pnoise is a package for generating perlin noise specifically
//for the purpose of my gork game (thus I do not waste computation
//on smoothing the results as I do not need it to be smooth).
package pnoise

import (
	"math"
	"math/rand"
)

var persistence = 1500
var numOctaves = 2

//3 prime numbers for noise generator
var (
	p1 int32 = 15731
	p2 int32 = 789221
	p3 int32 = 1376312589
)

//SetPersistence sets the persistence in the perlin noise algorithm
//basically the amplitude
func SetPersistence(p int) {
	persistence = p
}

//SetNumOctaves sets the number of octaves to sum up.
//Higher means more noise (and more detail).
func SetNumOctaves(n int) {
	numOctaves = n
}

func isPrime(n int) bool {
	end := int(math.Sqrt(float64(n)))
	for i := 2; i <= end; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//Randomize sets new seeds randomly
func Randomize() {
	//We want our prime numbers to be greater than this:
	startnum := 1000000

	var primes [3]int32

	for i := 0; i < 3; i++ {
		//how much to increment the starting number by each time
		inc := rand.Intn(800000) + 5000
		startnum += inc
		for n := startnum; ; n++ {
			if isPrime(n) {
				primes[i] = int32(n)
				break
			}
		}
	}
	p1, p2, p3 = primes[0], primes[1], primes[2]
}

//Seed sets the seeds for the noise generator. It uses 3 prime numbers.
func Seed(prime1, prime2, prime3 int) {
	p1, p2, p3 = int32(prime1), int32(prime2), int32(prime3)
}

func noise(x, y int32) float64 {
	n := x + y*57
	n = (n << 13) ^ n
	return float64(1.0 - float64((n*(n*n*p1+p2)+p3)&0x7fffffff)/1073741824.0)
}

func fnoise(x, y float64) float64 {
	n := int32(x + y*57)
	n = (n << 13) ^ n
	return float64(1.0 - float64((n*(n*n*p1+p2)+p3)&0x7fffffff)/1073741824.0)
}

func smoothNoise(x, y int32) float64 {
	corners := (noise(x-1, y-1) + noise(x+1, y-1) + noise(x-1, y+1) + noise(x+1, y+1)) / 16
	sides := (noise(x-1, y) + noise(x+1, y) + noise(x, y-1) + noise(x, y+1)) / 8
	center := noise(x, y) / 4
	return corners + sides + center
}
func fsmoothNoise(x, y float64) float64 {
	corners := (fnoise(x-1, y-1) + fnoise(x+1, y-1) + fnoise(x-1, y+1) + fnoise(x+1, y+1)) / 16
	sides := (fnoise(x-1, y) + fnoise(x+1, y) + fnoise(x, y-1) + fnoise(x, y+1)) / 8
	center := fnoise(x, y) / 4
	return corners + sides + center
}

//interpolate is a linear interpolation function
func interpolate(a, b, x float64) float64 {
	return a*(1-x) + b*x
}

func interpolatedNoise(x, y float64) float64 {
	intx := int32(x)
	fracx := x - float64(intx)

	inty := int32(y)
	fracy := y - float64(inty)

	v1 := float64(smoothNoise(intx, inty))
	v2 := float64(smoothNoise(intx+1, inty))
	v3 := float64(smoothNoise(intx, inty+1))
	v4 := float64(smoothNoise(intx+1, inty+1))

	i1 := interpolate(v1, v2, fracx)
	i2 := interpolate(v3, v4, fracx)

	return interpolate(i1, i2, fracy)
}

//Noise returns the perlin noise at the given x,y coordinate
func Noise(x, y float64) float64 {
	total := float64(0)
	for i := 0; i < numOctaves; i++ {
		freq := float64(2 ^ i)
		amp := float64(persistence ^ i)
		total = total + interpolatedNoise(x*freq, y*freq)*amp
	}
	return total
}

//IntNoise uses simpler calculations to find integer values
//of the noise at the given coordinate.
func IntNoise(x, y int) int {
	total := 0
	for i := 0; i < numOctaves; i++ {
		freq := 2 ^ i
		amp := persistence ^ i
		total = total + int(smoothNoise(int32(x*freq), int32(y*freq))*float64(amp))
	}
	return total
}
