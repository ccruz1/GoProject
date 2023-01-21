package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*			DELCARATIONS 			*/

//Map initialization of "updatedTestCases()"
var m = make(map[string]bool)

//String used when we add new unique values in the map.
var a = []string{}

//Map initialization of "updatedTestCases()"
var n = make(map[string]bool)

//String used when we add new unique values in the map.
var b = []string{}

/*			EXECUTION		*/

//Main execution
func main() {

	updatedFiles()
	keyToStringFiles()

	updatedTestCases()
	keyToStringTestCases()

}

//Function to retrieve all updated files
func updatedFiles() {

	//Define the file path
	file, err := os.Open("/parserPRUpdate/parserPRUpdate/mydiff.txt")

	//Error handling
	if err != nil {
		log.Fatal(err)
	}

	//Close the file
	defer file.Close()

	//Initialize the scanner
	scanner := bufio.NewScanner(file)

	//Print title
	fmt.Print("\nUpdated & Created FILES:\n\n")

	//Read file line by line and get the ones that contain "diff"
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "diff") {

			//Initialize a varibale that stores the line
			classLine := scanner.Text()
			//From every line get the final string after the last '/' path symbol.
			last := classLine[strings.LastIndex(classLine, "/")+1:]

			//Print classes
			addUniqueFiles(last)
		}
	}

	//DIVISION LINE BETWEEN FUNCTIONS
	fmt.Print("\n")

	//Handles scanner error in case of file problems
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

//Function to retrieve all updated files
func updatedTestCases() {

	//Define the file path
	file, err := os.Open("/parserPRUpdate/parserPRUpdate/mydiff.txt")

	//Error handling
	if err != nil {
		log.Fatal(err)
	}

	//Close the file
	defer file.Close()

	//Initialize the scanner
	scanner := bufio.NewScanner(file)

	//Print title
	fmt.Print("\n\nUpdated & Created TEST CASES\n")

	//Read file line by line and get the ones that contain "void Test"
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "@@") && strings.Contains(scanner.Text(), "void Test") && strings.Contains(scanner.Text(), "txProduct") {

			//Initialize a varibale that stores the line
			classLine := scanner.Text()
			//From every line get the final string after the last '/' path symbol.

			out := classLine[strings.LastIndex(classLine, "Test"):]
			last := out[:strings.IndexByte(out, '(')]

			//Print the
			addUniqueTC(last)
		}
	}

	//DIVISION LINE BETWEEN FUNCTIONS
	fmt.Print("\n\n")

	//Handles scanner error in case of file problems
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

//Function to add unique values to a map(The values of the modified -TEST CASES-)
func addUniqueTC(s string) {
	if m[s] {
		return // Already in the map
	}
	a = append(a, s)
	m[s] = true
}

//Function to add unique values to a map(The values of the modified -FILES-)
func addUniqueFiles(t string) {
	if n[t] {
		return // Already in the map
	}
	b = append(a, t)
	n[t] = true
}

//Get keys in map & print a string (-TEST CASES-)
func keyToStringTestCases() {
	var s string
	for key := range m {

		// Convert each key/value pair in m to a string
		s = fmt.Sprintf("%s", key)

		// Do whatever you want to do with the string;
		// in this example I just print out each of them.
		fmt.Print(" - " + s)
	}
}

//Get keys in map & print a string (-FILES-)
func keyToStringFiles() {
	var s string
	for key := range n {

		// Convert each key/value pair in m to a string
		s = fmt.Sprintf("%s", key)

		// Do whatever you want to do with the string;
		// in this example I just print out each of them.
		fmt.Print(" - " + s)
	}
}
