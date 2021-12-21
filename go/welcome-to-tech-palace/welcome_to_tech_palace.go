package techpalace

import (
	"fmt"
	"regexp"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s\n", border))
	sb.WriteString(fmt.Sprintf("%s\n", welcomeMsg))
	sb.WriteString(fmt.Sprintf("%s", border))
	return sb.String()
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	r := regexp.MustCompile("[*\n]")
	return strings.TrimSpace(r.ReplaceAllString(oldMsg, ""))
}
