package main

import "fmt"

const ebulicaoK = 373

func main (){

	tempFK := 373
	tempFC := 0
	tempK := ebulicaoK
	tempC := (tempK - 273)

	fmt.Printf("A temperatura de ebulição da água em K = %v (%T), temperatura de ebulição da água em °C =%v (%T) .", tempK, tempK, tempC, tempC)
	fmt.Printf("A temperatura de fusão da água em K = %v (%T), e a temperatura de fusão da água em °C é = %v (%T). ", tempFK, tempFK, tempFC, tempFC)
}