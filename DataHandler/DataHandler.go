package dataHandler

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}
type Data []Point

const coordinatesSeparator string = ";"
const pointsSeparator string = "\r\n"

func ReadDataAsStringFromFile(path string) (string, error) {
	arByte, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(arByte[:]), nil
}
func (slice Data) Sort() {
	sort.Slice(slice, func(i, j int) bool {
		if slice[i].x < slice[j].x {
			return true
		}
		if slice[i].x > slice[j].x {
			return false
		}
		return slice[i].y < slice[j].y

	})
}
func (slice Data) ToString() string {
	var builder strings.Builder

	for i := range slice {
		builder.WriteString(fmt.Sprintf("%d;%d\r\n", slice[i].x, slice[i].y))
	}
	return builder.String()
}
func ParseDataFromString(stringData string) (Data, error) {

	_points := strings.Split(stringData, pointsSeparator)
	slicePoints := make([]Point, 0, len(_points))

	for _, p := range _points {

		_point := strings.Split(p, coordinatesSeparator)

		if len(_point) != 2 {
			return nil, fmt.Errorf("Unsuccessful attempt to recognize the point")
		}
		x, err := strconv.Atoi(_point[0])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(_point[1])
		if err != nil {
			return nil, err
		}
		slicePoints = append(slicePoints, Point{x, y})
	}
	return slicePoints, nil
}
