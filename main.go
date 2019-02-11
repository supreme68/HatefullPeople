package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

//prcntgToBeHated reprezents the percentage of people from a group should be hated by all particiopants
const prcntgToBeHated float32 = 40

var participants []int

func main() {
	fmt.Println("Hello friend so welcome to Hatefull People a simple CLI program that picks the best team in order to not be any kind of conflicts in our teams")
	fmt.Println("-------------------------------------------------------------------------------------------------------")
	// part := "10" //TESTING
	var part string

	fmt.Println("Input the number of participants")
	_, err := fmt.Scan(&part)
	if err != nil {
		log.Println(err)
	}
	partCout, err := strconv.Atoi(part)
	if err != nil {
		log.Fatal(err)
	}
	participants, err = genHatedPartAndExclude(partCout)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("")
	fmt.Printf("This is our team: %x", participants)
}

func genHatedPartAndExclude(numOfPart int) ([]int, error) {
	part := initPartSlice(numOfPart)
	var hatedPart = 4
	for i, n := range part {
		if randomGenerator() {
			part = append(part[:i], part[i+1:]...) //remove the person that is being hated by all
			fmt.Println("")
			fmt.Printf("Excluded participant because all hate him: %x", n)
			hatedPart--
			if hatedPart == 0 {
				return part, nil
			}
		}
	}
	return nil, errors.New("ITS 5 FUCKING AM AND I AM TOO FUCKING TIRED TO THINK OF ANY GOOD ERROR NAME")
}

func initPartSlice(part int) []int {
	participants := make([]int, part)
	for i := range participants {

		participants[i] = i

	}
	return participants
}

func randomGenerator() bool {
	if rand.Intn(2) == 1 {
		return true
	} else {
		return false
	}
}
