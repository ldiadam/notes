package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// Define the Varitable Globally
var database string
var collection string
var notes []string

// Define The Color Globally
var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func main() {

	if len(os.Args) != 2 {
		ShowHelp()
		return
	}

	database = os.Args[1] + ".txt"
	collection = os.Args[1]

	if !LoadNotes() {
		return
	}
	ClearScreen()
	PrintColor(Blue, "Welcome to the notes tool")
	fmt.Println()
	DisplayMenu()
}

func ShowHelp() {
	PrintColor(Yellow, "Usage: ./notestool [TAG]")
}

func DisplayMenu() {
	reader := bufio.NewReader(os.Stdin)

	for {
		PrintColor(Cyan, "Select Operation:")
		fmt.Println("1. Show notes.")
		fmt.Println("2. Add a note.")
		fmt.Println("3. Delete a note.")
		fmt.Println("4. Exit.")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		ClearScreen()

		switch choice {
		case "1":
			ShowNotes()
			SaveNotes()
		case "2":
			AddNote()
			SaveNotes()
		case "3":
			DeleteNote()
			SaveNotes()
		case "4":
			PrintColor(Red, "Exiting the program.....")
			return
		default:
			PrintColor(Red, "Invalid choice. Please select again.\n ")
		}
	}
}

func Auth() bool {
	ClearScreen()

	// Read from password.txt
	passwordFile := "password.txt"
	passwordData, err := os.ReadFile(passwordFile)
	if err != nil {
		if os.IsNotExist(err) {
			_, _ = os.Create(passwordFile)
		} else {
			fmt.Println("Error loading notes:", err)
			return false
		}
	}

	// Find password
	var storedPassword string
	lines := strings.Split(string(passwordData), "\n")
	for _, line := range lines { // Loop Every Password Line
		parts := strings.Split(line, "|")            // Split Every Passowrd line with separator "|""
		if len(parts) == 2 && parts[0] == database { // Check wether password is define for that character
			storedPassword = strings.TrimSpace(parts[1]) // Store password to variable storedPassword
			break
		}
	}

	// Check wether password is empty or not, if empty, ask user for password
	if storedPassword == "" {
		PrintColor(Blue, "Would you like to set a password (leave blank for no password): ")
		reader := bufio.NewReader(os.Stdin)
		newPassword, _ := reader.ReadString('\n')
		newPassword = rot13(strings.TrimSpace(newPassword)) // Encrypt the password using rot13

		file, err := os.OpenFile(passwordFile, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return false
		}
		defer file.Close()

		if newPassword != "" {
			_, err = file.WriteString(fmt.Sprintf("%s|%s\n", database, newPassword))
			if err != nil {
				fmt.Println("Something went wrong:", err)
				return false
			}
		}

		storedPassword = newPassword
		ClearScreen()
	}

	// If password defined, Ask user the password
	if storedPassword != "" {
		var attemps int
		var loginState bool
		maxLogin := 3
		// Looping with condition if attemps < maxLogin
		for attemps < maxLogin {
			PrintColor(Blue, "Enter password for "+collection+": ")
			reader := bufio.NewReader(os.Stdin)
			inputPassword, _ := reader.ReadString('\n')
			inputPassword = rot13(strings.TrimSpace(inputPassword))

			// Compare the password, if incorrect
			if inputPassword != storedPassword {
				// Add one Attemps every incorrect password
				attemps++
				text := fmt.Sprintf("Incorrect password (%d/%d)", attemps, maxLogin)
				PrintColor(Red, text)
				// If attemps reached maxLogin, break the loop with loginState = false
				if attemps == maxLogin {
					loginState = false
					break
				}
				// If not, continue the loop
				continue
			} else {
				// If Password Correct, break the loop and set true on loginState
				loginState = true
				break
			}
		}

		// Check wheter login is false, return flase is login state false with error
		if !loginState {
			text := fmt.Sprintf("Access denied after %d attempts \n", maxLogin)
			PrintColor(Yellow, text)
			return false
		}
	}
	//Return true if correct
	return true

}

func LoadNotes() bool {
	if !Auth() {
		return false
	}

	fi, err := os.ReadFile(database)
	if err != nil {
		if os.IsNotExist(err) {
			_, _ = os.Create(database)
		} else {
			fmt.Println("Error loading notes:", err)
			return false
		}
	}

	var temp string
	for _, value := range fi {
		if value != 10 {
			temp += string(value)
		} else {
			notes = append(notes, temp)
			temp = ""
		}
	}

	return true
}

func ShowNotes() {
	if len(notes) == 0 { // if no notes, then no notes :)
		PrintColor(Red, "No notes available. \n ")
		return
	}
	PrintColor(Green, "Notes:") // if there are notes, showing them

	for i, encryptedNote := range notes { // Decrypt the note using ROT13
		decryptedNote := rot13(encryptedNote)
		fmt.Printf("%03d - %s\n", i+1, decryptedNote)
	}

	fmt.Println()
}

func DeleteNote() {
	if len(notes) == 0 {
		PrintColor(Red, "No notes to delete. \n ")
		return
	}

	ShowNotes()
	PrintColor(White, "Enter the number of note you want to delete:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	ClearScreen()
	// Convert input to an index
	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > len(notes) {

		PrintColor(Red, "Invalid input.\nPlease enter a valid note number.\n")
		return
	}

	// Remove note from the array.
	notes = append(notes[:index-1], notes[index:]...)
	SaveNotes()
	PrintColor(Green, "Note deleted successfully.\n")
}

func SaveNotes() bool {
	file, err := os.Create(database)
	if err != nil {
		fmt.Println("Error saving notes:", err)
		return false
	}
	defer file.Close()
	for _, note := range notes {
		_, _ = file.WriteString(note + "\n")
	}

	return true
}

func AddNote() {
	reader := bufio.NewReader(os.Stdin) // creates reader for efficient reading

	fmt.Println("Enter the note text:")
	note, err := reader.ReadString('\n')
	if err != nil { // error handling
		fmt.Println("Error reading input:", err)
	}
	note = strings.TrimSpace(note) // removing whitespaces

	if len(note) == 0 {
		ClearScreen()
		PrintColor(Red, "Error reading input, Input was empty\n")
		return
	}

	fmt.Println("Enter a tag for the note:")
	tag, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
	}
	tag = strings.TrimSpace(tag) // remove any extra whitespace from the tag

	timestamp := time.Now().Format("15:04:05 - 02/01/2006") // Get the current timestamp

	fullNote := fmt.Sprintf("%s - [%s] - %s", note, tag, timestamp) // Combine the note, tag, and timestamp

	encryptedNote := rot13(fullNote)     // Encrypt the note using ROT13
	notes = append(notes, encryptedNote) // Add the encrypted note to the collection
	SaveNotes()                          // Write to Database
	ClearScreen()                        // Clear Screen
}

// Clearscreen helper
func ClearScreen() {
	cmd := exec.Command("clear") // Use "cls" on Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Encrypt helper
func rot13(input string) string {
	var result strings.Builder
	for _, char := range input {
		switch {
		case char >= 'a' && char <= 'z':
			result.WriteRune('a' + (char-'a'+13)%26)
		case char >= 'A' && char <= 'Z':
			result.WriteRune('A' + (char-'A'+13)%26)
		default:
			result.WriteRune(char) // non-alphabet characters remain the same
		}
	}
	return result.String()
}

// Colour Text helper
func PrintColor(color string, text string) {
	fmt.Println(color + text + Reset)
}
