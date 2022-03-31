// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

var previous_questions string

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		questions <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)
	// TODO: Answer questions.
	// TODO: Make prophecies.
	// TODO: Print answers.
	go handle_questions(questions, answers)
	go predictions(answers)
	go print_prediction(answers)

	return questions
}

func handle_questions(questions <-chan string, answers chan<- string) {
	for q := range questions {
		go prophecy(q, answers)
	}
}

func predictions(predictions chan string) {
	words := []string{
		"The forces of nature",
		"Spirits",
		"Swords",
		"Winds of change",
		"Winds of magic",
		"The tides of war",
		"The might of the heavens",
		"Folk",
		"",
	}
	time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)
	prophecy(words[rand.Intn(len(words))], predictions)
}

func print_prediction(answers <-chan string) {
	for answer := range answers {
		fmt.Printf("%s: ", star)
		for _, char := range answer {
			fmt.Printf("%c", char)
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		}
		fmt.Printf("\n")
	}
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)

	previous_questions += question + "\n"
	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}
	if strings.Contains(question, "previous questions") {
		answer <- "Your previous questions are?" + "..." + previous_questions
	} else {
		// Cook up some pointless nonsense.
		nonsense := []string{
			"The moon is dark.",
			"The sun is bright.",
			"The rain is wet",
			"fire is hot",
			"Viola can't rätta i tid",
		}
		answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))]

	}
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
