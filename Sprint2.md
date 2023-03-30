Backend:
## Tasks completed
 Used golang and sql package
 Go get to import packages
  Two handlers
 Register
    		 Once submitting takes you to 
 Register auth
 Abilities:
 To create a user 
      Username: has criteria (length and alphanumeric)   
         Checks if its already there as well
 Passwords with requirements
      must include uppercase and lowercase letter as well as a special character
      must be between 11 and 60 characters
 Range through usernames alphabetically 
 Places them into the database
 Bcrypt 
 Password (defaultCost) - increase your hardware
 Salts are different to increase security
 Inserts statement sql for our query 

## API Documentation
database/sql
a. sql.DB
i. sql.Open()
ii. DB.QueryRow()
iii. DB.Close()
iv. Stmt
b. sql.Stmt
i. Stmt.Exec()
Fmt
a. fmt.Println()
html/template
a. template.ParseGlob()
b. template.ExecuteTemplate()
net/http
a. http.HandleFunc()
b. http.ListenAndServe()
Unicode
a. unicode.IsLetter()
b. unicode.IsNumber()
c. unicode.IsLower()
d. unicode.IsUpper()
e. unicode.IsPunct()
f. unicode.IsSymbol()
g. unicode.IsSpace()
golang.org/x/crypto/bcrypt
a. bcrypt.GenerateFromPassword()

## API Description:

database/sql: This package provides a generic interface for interacting with various SQL databases. The APIs used from this package include:
a. sql.DB: This is a database handle representing a pool of zero or more underlying connections. It exposes various methods to interact with the database, including opening and closing the connection, and executing queries.
i. sql.Open(): This function opens a new database connection and returns a pointer to a sql.DB object that represents a pool of zero or more connections to the database. It takes two parameters - the name of the driver to use and a string specifying the connection details, such as the username, password, and database name.
ii. DB.QueryRow(): This function executes a database query that is expected to return at most one row and scans the result into the provided destination pointers. It takes a query string and any number of arguments to be substituted into placeholders in the query string.
iii. DB.Close(): This function closes the database connection and releases any associated resources.
iv. Stmt: This type represents a prepared statement. It exposes various methods for executing the prepared statement with different arguments.

b. sql.Stmt: This type represents a prepared statement. It exposes various methods for executing the prepared statement with different arguments.
i. Stmt.Exec(): This method executes the prepared statement with the given arguments and returns a Result summarizing the effect of the statement.

fmt: This package implements formatted I/O with functions analogous to C's printf and scanf. The API used from this package is:
a. fmt.Println(): This function prints the given arguments to the standard output.

html/template: This package implements data-driven templates for generating HTML and other text-based formats. The APIs used from this package are:
a. template.ParseGlob(): This function parses all the files matching the given pattern and returns a pointer to a Template object.
b. template.ExecuteTemplate(): This method executes the named template with the given data object and writes the output to the given writer.

net/http: This package provides a set of functions and types for HTTP clients and servers. The APIs used from this package are:
a. http.HandleFunc(): This function registers the given function as a request handler for the given pattern on the default ServeMux.
b. http.ListenAndServe(): This function listens on the given address and serves requests using the default ServeMux.

unicode: This package provides functions for testing rune values against various character classes. 




## Bugs: 
Sql is still having trouble insert new usernames sometimes,
When working for the login section after registration it sometimes cannot recognize already existing users


## Testing:
Username: sweteam
Password: Password1234! 
Allows the user to create the user name
Fulfills password requirements

Username: 123
Password: nonumbers
Does not allow user to create an account



