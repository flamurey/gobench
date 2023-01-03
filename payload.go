package main

const LightN = 100
const HardN = 1000_000

func LighPayload() {
	fibonacci(LightN, 0, 1)
}

func HardPayload() {
	fibonacci(HardN, 0, 1)
}

func fibonacci(n int64, a int64, b int64) int64 {
	for n > 1 {
		c := a + b
		a = b
		b = c
		n--
	}
	return b
}
