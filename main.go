package main

import "strconv"

func bubble_sort(sortable [10]int) [10]int {
	for i := 0; i < len(sortable); i++ {
		for j := 0; j < len(sortable); j++ {
			if sortable[j] > sortable[i] {
				temp := sortable[i]
				sortable[i] = sortable[j]
				sortable[j] = temp
			}
		}
	}
	return sortable
}

func selection_sort(sortable [10]int) [10]int {
	for i := 0; i < len(sortable); i++ {
		// Find the smallest value in the unsorted part of the array
		index_of_min := i
		for j := i + 1; j < len(sortable); j++ {
			if sortable[j] < sortable[index_of_min] {
				index_of_min = j
			}
		}

		// Place the smallest value at the front of the unsorted section of the array
		temp := sortable[index_of_min]
		sortable[index_of_min] = sortable[i]
		sortable[i] = temp
	}
	return sortable
}

func create_array_string(printable [10]int) string {
	array_string := ""
	for _, item := range printable {
		array_string += strconv.Itoa(item) + ", "
	}
	return array_string[0 : len(array_string)-2]
}

func main() {
	// Creating an array to be sorted by the sorting algoritms
	arr := [10]int{12, 5, 15, 6, 8, 4, 53, 26, 45, 2}

	// Welcome message
	welcome_message := "Welcome to the sorting algorithm party!\n" +
		"Here, we will showcase a number of sorting algorithms.\n"
	println(welcome_message)

	// Bubble sort
	var bubble_sort_message = "Lets start with Bubble Sort.\n" +
		"Bubble sort moves through the data structure multiple times, comparing each element, and switches the values if necessary\n" +
		"Bubble sort is easy to implement and works with many different types of data types, but its time complexity is O(n^2) which sucks pretty bad.\n" +
		"Bubble sort can take this array: " + create_array_string(arr) + "\n" +
		"And turn it into this sorted array: " + create_array_string(bubble_sort(arr)) + "\n\n"
	println(bubble_sort_message)

	// Selection sort
	selection_sort_message := "Next, lets look at Selection Sort.\n" +
		"Selection sort can take this array: " + create_array_string(arr) + "\n" +
		"And turn it into this sorted array: " + create_array_string(selection_sort(arr)) + "\n\n"
	println(selection_sort_message)
}
