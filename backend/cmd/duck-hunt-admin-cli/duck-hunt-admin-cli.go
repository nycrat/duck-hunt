package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/nycrat/duck-hunt/backend/internal/repository"

	_ "github.com/joho/godotenv/autoload"
)

func addNewParticipant(name string, db *sql.DB) {
	participantRepository := repository.NewParticipantRepository(db)
	participantRepository.AddNewParticipant(name)

	fmt.Printf("Inserted new participant: %s\n", name)
}

func addNewPasscode(id int, passcode string, db *sql.DB, pepper []byte) {
	authRepository := repository.NewAuthRepo(db, pepper)
	authRepository.AddNewLoginInfo(int(id), passcode)

	fmt.Printf("Inserted new passcode for user: %d\n", id)
}

func printHelpMessage() {
	fmt.Println(`Usage: duck-hunt-admin-cli <command> [arguments]

The commands are:

	add participant [name]           adds new participant
	add passcode    [id] [passcode]  assigns new passcode to participant by id`)
}

func getCommand(db *sql.DB, pepper []byte) func() {
	if len(os.Args) == 1 {
		return printHelpMessage
	}

	if os.Args[1] == "add" {
		switch os.Args[2] {
		case "participant":
			name := os.Args[3]
			return func() {
				addNewParticipant(name, db)
			}
		case "passcode":
			id, err := strconv.ParseInt(os.Args[3], 10, 32)
			if err != nil {
				fmt.Println(err)
			}
			passcode := os.Args[4]

			return func() {
				addNewPasscode(int(id), passcode, db, pepper)
			}
		default:
			return printHelpMessage
		}
	} else {
		return printHelpMessage
	}
}

func main() {
	pepper := []byte(os.Getenv("DH_PEPPER"))
	dbConnUrl := os.Getenv("DH_DATABASE_URL")

	if pepper == nil || dbConnUrl == "" {
		fmt.Println("FATAL: Environment variables not set: DH_PEPPER DH_DATABASE_URL")
		return
	}

	db, err := sql.Open("postgres", dbConnUrl)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	getCommand(db, pepper)()
}
