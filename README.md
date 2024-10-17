# Notes App by Group Task 16

The Notes App is a simple command-line application built with Go, designed to help you manage collections of notes with a password-protection feature. Each collection of notes is saved in its own database (a text file), and you can add, view, or delete notes with ease. The application uses the ROT13 cipher to encrypt and decrypt your notes and passwords security.
## Features

- Add, view, and delete notes from different databases (files).
- Notes can be tagged and timestamped for easy reference.
- Password protection for each notes database.
- Simple encryption with ROT13 for notes and passwords.
- Clean and easy-to-use command-line interface.
## Installation

   ```bash
   git clone https://github.com/yourusername/notestool.git
   cd notestool

   go build -o notestool main.go
   ```
   This will create an executable file named notestool.
   
## Usage

```go
./notestool [TAG]

```

Replace [TAG] with the tag of the note collection you want to work with. The tool will create a new file for the collection if it doesn't exist.


### Example Commands
Start the tool with a specific collection:
```go
./notestool coding_ideas
```
This will open the notes collection stored in coding_ideas.txt.

Follow the prompts in the menu to:
- View all notes
- Add a new note
- Delete a note (coming soon)
- Exit the program

### Password Protection

- When you first open a notes database, you will be prompted to set a password (you can leave it blank if no password is required).
- If a password is set, you have to enter it each time you open the that collection.
- Passwords are encrypted and stored in password.txt, using ROT13 encryption.

### Note Format
- Each note consists of a message, a tag, and a timestamp.
- Example note: My first note - [important] - 15:04:05 - 02/01/2024

## File Structure
- main.go: The main Go file containing all the program logic.
- password.txt: Stores the password associated with each notes file.
- *.txt: Each notes collection is stored in its own .txt file based on the tag you use (e.g., personal.txt, work.txt).

## Tutorial
![](https://gitea.koodsisu.fi/rinaldiadam/notes/raw/branch/main/Tutorial.gif)

# Authors

- [@rinaldiadam](https://gitea.koodsisu.fi/rinaldiadam)
- [@villekivivuori](https://gitea.koodsisu.fi/villekivivuori)
- [@habeebabdulkareem2](https://gitea.koodsisu.fi/habeebabdulkareem2)