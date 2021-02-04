https://itnext.io/interfaces-in-go-5c6e38b81b41
https://aaronzhuo.medium.com/interfaces-in-go-3861ec1b75ae
:::::::::::::interfaces allo type assertion 

:::::::::::::Type switching 
func WriteLog(data interface{}) {
	switch result := data.(type) {
	case User:
		log.Println("Found user: ", result.Name)
	default:
		log.Println(data)

	}

}


:::::::::::::
Embedding interfaces
type UserDatabase interface {
 GetUser(username string) (*User, error)
 GetUser(email string) (*User, error)
 AddUser(username, email string) error
}

y llevarlo mejor a 

type UserDatabase interface {
	UserRetriever
	AddUser(username, email string) error
}

// UserRetriver is a broken out interface that will be embedded by UserDatabase
type UserRetriever interface {
	GetUser(username string) (*User, error)
	GetUserByEmail(email string) (*User, error)
}



