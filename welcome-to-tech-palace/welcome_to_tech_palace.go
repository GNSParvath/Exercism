package techpalace

import (
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	one := "Welcome to the Tech Palace, "
	return one + "" + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	two := strings.Repeat("*", numStarsPerLine)
	return two + "\n" + welcomeMsg + "\n" + two
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	three := strings.ReplaceAll(oldMsg, "*", "")
	return strings.TrimSpace(three)
}
