package main

import ("fmt"
        "strings")

func WelcomeMessage(customer string) string {
    return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return strings.Repeat("*", numStarsPerLine) + "\n"+ fmt.Sprintln(welcomeMsg) + strings.Repeat("*", numStarsPerLine)
}

func CleanupMessage(oldMsg string) string {
    return strings.TrimSpace(" \t\n" + strings.ReplaceAll(oldMsg, "*", "") + " \n\t\r\n")
}

func main(){
   //fmt.Println(WelcomeMessage("Juan"))
   //fmt.Println(AddBorder(WelcomeMessage("Juan"), 10))
   fmt.Println(CleanupMessage(WelcomeMessage("Juan")))
}