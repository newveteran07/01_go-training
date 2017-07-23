package main

import "fmt"

func main() {
	myFriendsName := "Medhi"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Wassup my friends with a name of length 2")
	case myFriendsName == "Tim":
		fmt.Println("Wassup Tim")
	case myFriendsName == "Jenny":
		fmt.Println("Wassup Jenny")
	case myFriendsName == "Marcus", myFriendsName == "Medhi":
		fmt.Println("Your name is either Marcus or Medhi")
	case myFriendsName == "Julian":
		fmt.Println("Wasup Julian")
	case myFriendsName == "Sushant":
		fmt.Println("Wassup Sushant")
	default:
		fmt.Println("nothing matched; this is the default response.")

		/*expression i not needed -- if not expression is provided, go checks for the first case that evals to true -- making the switch operate like a if/if else/ else case */

	}
}
