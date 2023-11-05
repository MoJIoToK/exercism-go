package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	welcomeMsg := fmt.Sprint("Welcome to my party, ", name, "!")
	return welcomeMsg
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	happyBirthdayMsg := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
	return happyBirthdayMsg
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	assignedTableMsg := fmt.Sprintf("\nYou have been assigned to table %03d."+
		" Your table is %s, exactly %.1f meters from here.", table, direction, distance)
	neighborMsg := fmt.Sprintf("\nYou will be sitting next to %s.", neighbor)
	return Welcome(name) + assignedTableMsg + neighborMsg
}
