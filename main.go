package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"unicode"
//	"io/ioutil" //old
//	"net/http" //old

	"golang.org/x/crypto/bcrypt"
	_ "github.com/go-sql-driver/mysql"
)
var tpl *template.Template
var db *sql.DB

func main() {

	tpl, _ = template.ParseGlob("templates/*.html")
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/loginauth", loginAuthHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/registerauth", registerAuthHandler)
	http.ListenAndServe("localhost:8080", nil)
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****loginHandler running*****")
	tpl.ExecuteTemplate(w, "login.html", nil)
}

// loginAuthHandler authenticates user login
func loginAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****loginAuthHandler running*****")
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println("username:", username, "password:", password)
	// retrieve password from db to compare (hash) with user supplied password's hash
	var hash string
	stmt := "SELECT Hash FROM bcrypt WHERE Username = ?"
	row := db.QueryRow(stmt, username)
	err := row.Scan(&hash)
	fmt.Println("hash from db:", hash)
	if err != nil {
		fmt.Println("error selecting Hash in db by Username")
		tpl.ExecuteTemplate(w, "login.html", "check username and password")
		return
	}
	// func CompareHashAndPassword(hashedPassword, password []byte) error
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	// returns nill on succcess
	if err == nil {
		fmt.Fprint(w, "You have successfully logged in :)")
		return
	}
	fmt.Println("incorrect password")
	tpl.ExecuteTemplate(w, "login.html", "check username and password")
}

// loginHandler serves form for users to login with
func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****RegisterHandler running*****")
	tpl.ExecuteTemplate(w, "login.html", nil)
}
func registerAuthHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("*****RegisterAuthHandler running*****")
	r.ParseForm()
	username := r.FormValue("username")
	// checks for username - alphaNumeric characters

	var nameAlphaNumeric = true
	for _, char := range username {
		if unicode.IsLetter(char) == false && unicode.IsNumber(char) == false {
			nameAlphaNumeric = false
		}
	}
// checking username passwordLenth
var nameLength bool
if 5 <= len(username) && len(username) <= 50 { 
nameLength = true
}

// check password criteroa
password := r.FormValue("password")
fmt.Println("password:", password, "\npswdLength:", len(password))
//variables that must password creation criteria
var pswdLowercase, pswdUppercase, pswdNumber, pswdSpecial, pswdLength, pswdNoSpaces bool
pswdNoSpaces = true
for _, char := range password {
	switch  {
	case unicode.IsLower(char):
			pswdLowercase = true
	case unicode.IsUpper(char):
		pswdUppercase = true
	case unicode.IsNumber(char):
		pswdNumber = true
	case unicode.IsPunct(char) || unicode.IsSymbol(char):
		pswdSpecial = true
	case unicode.IsSpace(int32(char)):
		pswdNoSpaces = false
	}

}
if 11 < len(password) && len(password) < 60 {
	pswdLength = true
}
fmt.Println("pswdLowercase:", pswdLowercase, "\npswdUppercase:", pswdUppercase, "\npswdNumber:", pswdNumber, "\npswdSpecial:", pswdSpecial, "\npswdLength:", pswdLength, "\npswdNoSpaces:", pswdNoSpaces, "\nnameAlphaNumeric:", nameAlphaNumeric, "\nnameLength:", nameLength)
if !pswdLowercase || !pswdUppercase || !pswdNumber || !pswdSpecial || !pswdLength || !pswdNoSpaces || !nameAlphaNumeric || !nameLength {
	tpl.ExecuteTemplate(w, "register.html", "please check username and password criteria")
	return
}
// check if username already exists for availability
stmt := "SELECT UserID FROM bcrypt WHERE username = ?"
row := db.QueryRow(stmt, username)
var uID string
err := row.Scan(&uID)
if err != sql.ErrNoRows {
	fmt.Println("username already exists, err:", err)
	tpl.ExecuteTemplate(w, "register.html", "username already taken")
	return
}
// create hash from password
var hash []byte
// func GenerateFromPassword(password []byte, cost int) ([]byte, error)
hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
if err != nil {
	fmt.Println("bcrypt err:", err)
	tpl.ExecuteTemplate(w, "register.html", "there was a problem registering account")
	return
}
fmt.Println("hash:", hash)
fmt.Println("string(hash):", string(hash))
// func (db *DB) Prepare(query string) (*Stmt, error)
var insertStmt *sql.Stmt
insertStmt, err = db.Prepare("INSERT INTO bcrypt (Username, Hash) VALUES (?, ?);")
if err != nil {
	fmt.Println("error preparing statement:", err)
	tpl.ExecuteTemplate(w, "register.html", "there was a problem registering account")
	return
}
defer insertStmt.Close()
var result sql.Result
//  func (s *Stmt) Exec(args ...interface{}) (Result, error)
result, err = insertStmt.Exec(username, hash)
rowsAff, _ := result.RowsAffected()
lastIns, _ := result.LastInsertId()
fmt.Println("rowsAff:", rowsAff)
fmt.Println("lastIns:", lastIns)
fmt.Println("err:", err)
if err != nil {
	fmt.Println("error inserting new user")
	tpl.ExecuteTemplate(w, "register.html", "there was a problem registering account")
	return
}
fmt.Fprint(w, "congrats, your account has been successfully created")

}
	/*  old 
	(forever 21 API)
	url := "https://apidojo-forever21-v1.p.rapidapi.com/products/v2/list?category=women_main&pageSize=48&pageNumber=1&sortby=0&filterColor=BLACK&filterSize=XS%2FS"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "1aa5b6fa64msh21e535afc50d574p1e0eadjsnd37ae1fe22f2")
	req.Header.Add("X-RapidAPI-Host", "apidojo-forever21-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))


	/* user authentication
	1. check user name criteria
	2. check paswworkd criteria
	3. check if username already exists in database 
	4. creat bctypt hash from password
	5.insert username and password has in database
	(email validation will be next...?)
	*/







/*

- user account (login system)
- backened database saves info (whhether they exist or prompt to )
- favorite items (store in an array)
- 

- later on:
- shopping cart/ payment page ^ arrary like favorites
- test backend 

*/

