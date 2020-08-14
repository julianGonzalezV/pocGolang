package main

import "fmt"

//Use maps para operaciones set, get and delete data con mejor rendimiento
// Map al igual que slice son un constructor especial en go
// No garantiza el orden de los elementos, ni por key ni value
//  it's good practice to avoid defining a map using v para evitar panic message
// en caso de que no esté inicializado

func getUsers() map[string]string {
	users := map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
	}
	users["073"] = "Tracy"
	return users
}

func getUser(id string) (string, bool) {
	user, exists := getUsers()[id]
	if !exists {
		fmt.Printf("Passed user ID (%v) not found.\nUsers:\n", id)
		for key, value := range getUsers() {
			fmt.Println("  ID:", key, "Name:", value)
		}
	}
	return user, exists
}

func main() {
	fmt.Println("::::::::.Create a Map::::::")
	fmt.Println("Users:", getUsers())

	fmt.Println("::::::::.Reading from Map::::::")
	fmt.Println(getUser("073"))
	fmt.Println(getUser("12"))

	fmt.Println("::::::::Deleting Elements from a Map::::::")
	delete(getUsers(), "073")
	fmt.Println("Users:", getUsers())

}
