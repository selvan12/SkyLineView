/*
 * In a city there are N number of rectangular 2D buildings. Each buildings
 * left and right (X and Y) top corner coordinates are given. Compute the
 * outer contour points which gives SkyLineView when viewed from a distance.
 * By taking consideration that buildings 2D points may overlap.
 *
 * Programming Language : GoLang
 */

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

// constant corner value to indicate building
const build_points int = 4

// struct for each points x, y coordinates on 2D graph
type point struct {
	xData int
	yData int
}

// points which builds a building in a 2D graph
type Building2D struct {
	xLeft  int
	yLeft  int
	xRight int
	yRight int
}

// collection of skypoints which results sky outline
type skyPoints []point

// function to validate the building rectangle coordinates
// input: point1 of the rectangle array
// input: point2 of the rectangle array
// output: boolean result of points validity
func validateRectanglePoints(pt1 point, pt2 point) bool {
	result := true

	// first check if the coordinates are valid
	if (pt1.xData < 0) || (pt1.yData < 0) || (pt2.xData < 0) || (pt2.yData < 0) {
		result = false
	} else if (pt1.yData != pt2.yData) {
		// if both height are not same, it is not rectangle
		result = false
	} else if (pt1.xData == pt2.xData) {
		// if both X position same, it is not rectangle
		result = false
	}

	return result
}

// print all points from the point collection
// input: point collection to print
func printPoints(pts []point) int{
	length := len(pts)
	buildingCount := length/2

	if buildingCount == 0{
		fmt.Println("Builing count 0")
		return buildingCount
	}
	fmt.Printf("\nTotal number of building are %d and its coordinates are:\n", buildingCount)
	for i := 0; i < len(pts); i++ {
		// print 2 points followed by newline for display
		fmt.Print(pts[i])
		if i%2 != 0 {
			fmt.Println()
		}
	}

	return buildingCount
}

// construct the building array from each point coordinates
// input: point collection to made buildings
// output: building array
func constructBuildingFromCoordinates(pts []point) []Building2D {
	var building []Building2D
	length := len(pts)

	if length != 0 {
		// two point collection makes one building
		building = make([]Building2D, length/2)
		for i, j := 0, 0; i%2 == 0 && i+1 < length; i, j = i+2, j+1 {
			building[j] = Building2D{pts[i].xData, pts[i].yData,
				pts[i+1].xData, pts[i+1].yData - pts[i].yData}
		}
	}

	return building
}

// get skyPoints from building
// input: building array to convert skyPoints
// output: skyPoints from building structure
func getSkyPoints(build Building2D) skyPoints {
	var skyPts skyPoints
	skyPts = skyPoints{{build.xLeft, build.yLeft},
		{build.xRight, build.yRight}}
	return skyPts
}

// removing redundant points falling at same height
// input: skyPoints needs to be scanned for redundant points
// output: skyPoints result after cleaned
func removeRedundantPoints(input skyPoints) skyPoints {
	// skyPts which will be sent as result
	skyPtsResult := skyPoints{}
	yPrev := 0

	for i := 0; i < len(input); i++ {
		// do not append if previous and current heights are same
		if (yPrev != input[i].yData) {
			skyPtsResult = append(skyPtsResult, input[i])
		}
		// remember previous coordinate
		yPrev = input[i].yData
	}

	return skyPtsResult
}

// format each start and end points for skyLine draw
// input: skyPoints needs to be formatted
// output: skyPoints result after formatted
func formatOutputPoints(inputToBeFormated skyPoints) skyPoints {
	// skyPts which will be sent as result
	skyPtsResult := skyPoints{}

	// remove redundant points before formatting the skyPoints
	var input = removeRedundantPoints(inputToBeFormated)

	xValue, yValue := 0, 0
	for i := 0; i < len(input); i++ {
		j := i+1;

		if i == 0 {
			// append starting element
			if input[i].xData != 0 {
				skyPtsResult = append(skyPtsResult, point{input[i].xData, yValue})
			}
			// do add first element from computing
			skyPtsResult = append(skyPtsResult, input[i])
		} else if (j < len(input)) && (yValue == input[i].yData) && (input[i].yData == input[j].yData) {
			// do nothing, Since last two heights are same
		} else if (xValue != input[i].xData) && (yValue == input[i].yData) {
			// append if x values are differ
			skyPtsResult = append(skyPtsResult, input[i])
		} else if (xValue == input[i].xData) && (yValue != input[i].yData) {
			// append if y values are differ
			skyPtsResult = append(skyPtsResult, input[i])
		} else if (xValue != input[i].xData) && (yValue != input[i].yData) {
			// if both x and y value are differ then
			// add bottom point wrt to previous yValue
			skyPtsResult = append(skyPtsResult, point{input[i].xData, yValue})
			// append current values
			skyPtsResult = append(skyPtsResult, input[i])
		}
		// remember X, Y values for next comparision
		xValue = input[i].xData
		yValue = input[i].yData
	}

	return skyPtsResult
}

// merge left and right skyPoints and get the output skyPoint
// input: left skyPoint collection
// input: right skyPoint collection
// output: merged skyPoint collection
func mergeTwoSkyPoints(left skyPoints, right skyPoints) skyPoints {

	// skyPts which will be sent as result
	skyPtsResult := skyPoints{}

	// To store current heights of two skylines
	lastHeight1, lastHeight2 := 0, 0
	i, j := 0, 0

	// traverse until left and right subtree are available
	for (i < len(left)) && (j < len(right)) {
		// compare x coordinates of two input skylines and put the smaller one in result
		// compute least X with max Y data from both points
		if left[i].xData < right[j].xData {
			// left X value assigned with max Y value
			x1 := left[i].xData
			// remember the left Y data
			lastHeight1 = left[i].yData
			// select height as max of two heights
			maxY := lastHeight1
			if lastHeight1 < lastHeight2 {
				maxY = lastHeight2
			}

			// append left X least value with max Y value
			skyPtsResult = append(skyPtsResult, point{x1, maxY})
			i++
		} else {
			// right X value assigned with max Y value
			x2 := right[j].xData
			// remember the right Y data
			lastHeight2 = right[j].yData
			// select height as max of two heights
			maxY := lastHeight2
			if lastHeight2 < lastHeight1 {
				maxY = lastHeight1
			}

			// append right X least value with max Y value
			skyPtsResult = append(skyPtsResult, point{x2, maxY})
			j++
		}
	}

	// append if there are skypoints to add
	for i < len(left) || j < len(right) {
		if  i < len(left) {
			// append if there are left skypoints
			skyPtsResult = append(skyPtsResult, left[i])
			i++
		} else {
			// append if there are right skypoints
			skyPtsResult = append(skyPtsResult, right[j])
			j++
		}
	}

	return skyPtsResult
}

// function will compute the skyLine points recursively
// input: Building2D array
// output: SkyPoint collection
func getSkyLineView(build []Building2D) skyPoints {
	skyPtsResult := skyPoints{}

	// get the building count
	length := len(build)

	// If only one building available return the skyPoints as its coordinates
	if length == 1 {
		skyPtsResult = getSkyPoints(build[0])
	} else {
		// breakdown buildings by two half and get the skyPoints
		skyPts1 := getSkyLineView(build[:length/2])
		skyPts2 := getSkyLineView(build[length/2:])

		//fmt.Println("skyPts1 and skyPts2 are")
		//fmt.Println(skyPts1)
		//fmt.Println(skyPts2)

		// merge two skyPoints diagonal
		skyPtsResult = mergeTwoSkyPoints(skyPts1, skyPts2)
	}

	// return the result skyPoints
	return skyPtsResult
}

// main function which drives the code
func main() {

	// input building coordinate points
	var coordinates []point

	// test case 1 (which given for discussion)
	//coordinates = []point{{0,4},{1,4},  {1,1},{3,1},  {2,2},{4,2},  {3,3},{5,3},  {6,1},{8,1}}
	// test case2:
	//coordinates = []point{{2,10},{9,10},  {3,15},{6,15},  {5,12},{12,12},  {13,10},{16,10}, {15,5},{17,5}}
	// test case 3: (corner case)
	// coordinates = []point{{3,8},{8,8},  {3,10},{10,10}}
	//test case 4: (corner case)
	// coordinates = []point{{2,10},{10,10},  {4,8},{10,8}}

	/******Removed user input for each value
	// initialize values
	count := 0
	point1, point2 := point{}, point{}

	fmt.Print("Enter building count of the city: ")
enter:
	fmt.Scanf("%d", &count)
	if count < 1 {
		fmt.Println("Building count is not valid. Enter again")
		goto enter
	}

	fmt.Println("Building coordinates has to be given by user")
	for i := 0; i < count; i++ {
		fmt.Printf("Enter building %d points as below\n", i+1)
		// get all four inputs of the building
		for j := 0; j < build_points; j++ {
			switch j {
			case 0:
				fmt.Print(" X1:")
				fmt.Scanf("%d", &point1.xData)
			case 1:
				fmt.Print(" Y1:")
				fmt.Scanf("%d", &point1.yData)
			case 2:
				fmt.Print(" X2:")
				fmt.Scanf("%d", &point2.xData)
			case 3:
				fmt.Print(" Y2:")
				fmt.Scanf("%d", &point2.yData)
			}
		}

		// check if building coordinates are valid
		if !validateRectanglePoints(point1, point2) {
			fmt.Println("Invalid rectangle coordinates. Enter again")
			// reinitialize the values
			point1, point2 = point{0, 0}, point{0, 0}
			// decrement the counter since coordinates are not valid
			i--;
		} else {
			// building coordinates are valid and append to array
			coordinates = append(coordinates, point1)
			coordinates = append(coordinates, point2)
		}
	}
	***************/

	// get the input tuple string from console until new line
	fmt.Println("Enter building array of tuples - example: ((0,4),(1,4)) ::")
	reenter:
	reader := bufio.NewReader(os.Stdin)
	tupleString, _ := reader.ReadString('\n')
	//fmt.Println(tupleString)

	// function to filter string values from tuple
	splitFunc := func(r rune) bool {
		ret := true
		// filter only numeric char
		if r >= '0' && r <= '9' {
			ret = false
		}
		return ret
	}
	var tupleArray []string
	tupleArray = strings.FieldsFunc(tupleString, splitFunc)
	tupleLen := len(tupleArray)

	// print tuple split array and its length
	fmt.Println("\nTuple array: ", tupleArray, "\nTuple length: ", tupleLen)
	// building point count
	if tupleLen < build_points {
		fmt.Println("Less than 4 inputs. Enter building coordinates again")
		goto reenter
	} else if tupleLen%build_points != 0 {
		fmt.Println("Input points are not balanced and buildings considered are", tupleLen/build_points)
	}
	balancedCorner := tupleLen - (tupleLen%build_points)
	point1, point2 := point{}, point{}

	// get the all integer coordinates
	for i := 0; i < balancedCorner; {
		for j := 0; j < build_points; j++ {
			// string to int conversion
			val, err := strconv.Atoi(tupleArray[i+j])
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
			}
			// collect all 4 values
			switch j {
				case 0:
					point1.xData = val
				case 1:
					point1.yData =val
				case 2:
					point2.xData = val
				case 3:
					point2.yData = val
			}
		}
		// check if building coordinates are valid
		if !validateRectanglePoints(point1, point2) {
			fmt.Println("Invalid rectangle coordinates.")
		} else {
			// append valid coordinates to building
			coordinates = append(coordinates, point1)
			coordinates = append(coordinates, point2)
			//fmt.Println(point1)
			//fmt.Println(point2)
		}
		// reinitialize the values
		point1, point2 = point{0, 0}, point{0, 0}
		// increment to next set of tuple
		i = i+build_points
	}

	// print all the input skylines in format
	buildCount := printPoints(coordinates)
	// atleast one building required
	if buildCount == 0 {
		fmt.Println("Enter coordinates again")
		goto reenter
	}
	// construct the building structure from each pair of points
	building := constructBuildingFromCoordinates(coordinates)

	// calculate the skyLine points
	var skyLinePts skyPoints
	skyLinePts = getSkyLineView(building)

	// display output points
	fmt.Println("\nOutput:")
	// format the points before display
	fmt.Println(formatOutputPoints(skyLinePts))
}
