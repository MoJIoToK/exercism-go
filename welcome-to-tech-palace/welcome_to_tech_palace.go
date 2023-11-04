package techpalace

import (
	"strings"
)

// Welcome is constant for the first part of greeting.
const welcome string = "Welcome to the Tech Palace,"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return welcome + " " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	stars := strings.Repeat("*", numStarsPerLine)
	return stars + "\n" + welcomeMsg + "\n" + stars
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	newMsg := strings.ReplaceAll(oldMsg, "*", " ")
	return strings.TrimSpace(newMsg)
}
