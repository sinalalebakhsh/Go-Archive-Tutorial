package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"acron/animation"
)

var reader *bufio.Reader

func main() {
	animation.Animation()
	reader = bufio.NewReader(os.Stdin)
	numberOfTimesPlayed := GetNumberOfTimesPlayed()
	animation.Animation()
	NumberOfTimesPlayed(numberOfTimesPlayed)
}	

func PlayerChoice() string{
	for {
		var reader = bufio.NewReader(os.Stdin)
		fmt.Print("Please enter rock, paper, or scissors -> ")
		playerChoice, _ := reader.ReadString('\n')
		playerChoice = strings.Replace(playerChoice, "\n", "", -1)
		animation.Animation()
		if playerChoice == "rock" || playerChoice == "paper" || playerChoice == "scissors" {
			return playerChoice
		} 	
	}
}

func ComputerChoice() string{
	//create the computerChoiceSlice slice and append computerChoice to it
	computerChoiceSlice := make([]string, 0)
	computerChoiceSlice = append(computerChoiceSlice,
    "rock",
    "paper",
    "paper",)
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	computerChoice := fmt.Sprint(computerChoiceSlice[rand.Intn(len(computerChoiceSlice))])
	return computerChoice
}

func Calculator(playerChoice, computerChoice string) string {
	if playerChoice == computerChoice {
		return "The game equalised"
	} else if playerChoice == "paper" && computerChoice == "rock" {
		return "You are win"
	} else if playerChoice == "scissors" && computerChoice == "paper" {
		return "You are win"
	} else if playerChoice == "rock" && computerChoice == "scissors" {
		return "You are win"
	} else {
		return "The computer is won"
	}
}

func NumberOfTimesPlayed(number int) {
	for i := 0; i <= number-1; i++ {
		playerChoice := PlayerChoice() 
		fmt.Println("Your Choice:", playerChoice)
		//-------------------------------------------------
		computerChoice := ComputerChoice()
		fmt.Println("Computer Choice:", computerChoice)
		//-------------------------------------------------
		result := Calculator(playerChoice, computerChoice)
		fmt.Println("RESULT:", result)
	}
}

func prompt(){
	fmt.Print("-> ")
}

func GetNumberOfTimesPlayed() int {
	counter := 0
	for {
		if counter == 0 {
			fmt.Println("Number of times played?")
			counter = 1
		}
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		number, error_ := strconv.Atoi(userInput)
		if error_ != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return number
		} 
	}
}