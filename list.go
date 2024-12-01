package main

import "fmt"

type ProductType struct {
	Title string `json:"title"`
	Id    string `json:"id"`
	Price string `json:"price"`
}

type CourseRatingMapType map[string]float64

func (rateMap CourseRatingMapType) OutputCourseMap() {
	fmt.Println("course rating with maps", rateMap)
}

func listAndMap() {
	//Lists
	productList := []ProductType{{Title: "First", Id: "1", Price: "23"}, {Title: "Second", Id: "2", Price: "34"}}
	priceList := []int{1, 2, 3, 4}
	discountList := []int{20, 30, 10, 40}
	priceList = append(priceList, discountList...)
	//Maps
	countMaps := map[int]string{
		0: "First one",
		1: "Second one",
		2: "Third one",
	}
	//make with arrys
	courseRatingList := make([]string, 1, 5) // First it she arrays type, second arg is inital length and last arg is maximum capacity
	courseRatingList[0] = "1.2"
	courseRatingList = append(courseRatingList, "2")

	//make with maps
	courseRatingMaps := make(CourseRatingMapType, 5) // For maps first arg is map initialization and last arg is capacity of the map
	courseRatingMaps["go"] = 4.7
	courseRatingMaps["React"] = 4.8

	fmt.Println("product list", productList)
	fmt.Println("prices and discount list", priceList)
	fmt.Println("maps list", countMaps)
	countMaps[3] = "Forth one"
	fmt.Println("modified map list", countMaps)
	fmt.Println("course rating with arrays", courseRatingList)
	courseRatingMaps.OutputCourseMap()
}
