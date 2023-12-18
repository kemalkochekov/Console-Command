# **Console Commands**

## **Introduction**
Implemented console commands functionality in the project in a way that requires minimal changes when adding a new command.

## Installation
Clone this repository
  ```bash
    git clone https://github.com/kemalkochekov/Console-Commands.git
```
## How to Run
```$
  bar@foo$: go run cmd/cli/main.go
```
## **Command Types**
- help command
- spell command
- gofmt command

## **Command Overview**
  ### **Help Command**
    Displays a list of available commands
  
  ### **Spell Command**
    Accepts a word as input and outputs all the letters of that word separated by spaces in the console based on the results of the operation.
  
  ### **Gofmt Command**
    Accepts a *.txt file as input and generates a modified text file with alterations such as inserting a tab before each paragraph and adding a period (.) at the end of sentences.

## **After running the project:**
  ```
  Available commands:

  help: Displays a list of available commands
  
  spell: spell --word [string] | Accepts a word as input and outputs all the letters of that word separated by spaces in the console based on the results of the operation.
  
  gofmt: gofmt --file [string] | Accepts a *.txt file as input and generates a modified text file with alterations such as inserting a tab before each paragraph and adding a period (.) at the end of sentences.
  ```

## **Run unit tests**
```
  go test ./...
```

## **Run unit tests with coverage**
```
go test ./... -cover
```

## **Future Enhancements**
  We can add several new commands with minimum changes in the whole project.
