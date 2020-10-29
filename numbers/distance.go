package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

func main() {
	firstCity := setCityCoordinate()
	secondCity := setCityCoordinate()

	distance := firstCity.calculateDistance(secondCity)

	fmt.Println("choose unit of distance")
	fmt.Println("'m' is statute miles")
	fmt.Println("'k' is kilometers")
	fmt.Println("'n' is nautical miles")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch scanner.Text() {
		case "m":
			fmt.Println(distance)
		case "k":
			fmt.Println(distance * 1.609344)
		case "n":
			fmt.Println(distance * 0.8684)
		case "e":
			os.Exit(0)
		default:
			fmt.Println("wrong input")
		}
	}
	if scanner.Err() != nil {
		panic("an error occurred when read input")
	}
}

func setCityCoordinate() Coordinate {
	fmt.Println("Enter city's latitude")
	var latitude float64
	_, err := fmt.Scanf("%f", &latitude)
	if err != nil {
		panic(err)
	}

	fmt.Println("Enter city's longitude")
	var longitude float64
	_, err = fmt.Scanf("%f", &longitude)
	if err != nil {
		panic(err)
	}

	return Coordinate{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

func (origin Coordinate) calculateDistance(destination Coordinate) float64 {
	radlatOrigin := degrees2radians(origin.Latitude)
	radlatDest := degrees2radians(destination.Latitude)

	theta := origin.Longitude - destination.Longitude
	radtheta := degrees2radians(theta)

	dist := math.Sin(radlatOrigin)*math.Sin(radlatDest) +
		math.Cos(radlatOrigin)*math.Cos(radlatDest)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515
	return dist
}

func degrees2radians(degree float64) float64 {
	return degree * math.Pi / 180
}
