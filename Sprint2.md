Backend:
 * Used golang and sql package
 * Go get to import packages
 * Connected our database
 * Two handlers
 * Register
      * Once submitting takes you to 
 * Register auth
 * Abilities:
 * To create a user 
      * Username: has criteria (length and alphanumeric)   
         * Checks if its already there as well
 * Passwords with requirements
      * must include uppercase and lowercase letter as well as a special character
      * must be between 11 and 60 characters
 * Range through usernames alphabetically 
 * Places them into the database
 * Bcrypt 
 * Password (defaultCost) - increase your hardware
 * Salts are different to increase security
 * Inserts statement sql for our query 
## API Documentation
 * registerHandler(w http.ResponseWriter, r  *http.Request)
	 * 
 * registerAuthHandler(w http.ResponseWriter, r  *http.Request)
	 * takes in username and password
	 * checks that it fulfills all requirements
		 * if so, adds to database
		 * if not, returns error message

## Bugs: 
* Sql is still having trouble insert new usernames somethings
   * Not 100% capabilities
