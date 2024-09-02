package main

import (
    "html/template"
    "log"
    "net/http"
    "os"
    "reakgo/models"
    "reakgo/utility"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
     // Load environment variables from .env file
	 if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Database connection
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    if dbUser == "" || dbPassword == "" || dbName == "" {
        log.Fatal("Database environment variables are not set")
    }

    var err error
    utility.Db, err = sqlx.Open("mysql", dbUser+":"+dbPassword+"@/"+dbName)
    if err != nil {
        log.Fatal("Error opening database connection:", err)
    }

    // Initialize templates
    utility.View = template.Must(template.ParseGlob("templates/*.html"))

    // Routes
    http.HandleFunc("/signup", signupHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/home", homeHandler)

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// Define the User struct
type User struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Address  string `json:"address"`
    Password string `json:"password"`
}

// Signup handler
// Signup handler
func signupHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        // Parse form data
        err := r.ParseForm()
        if err != nil {
            log.Println("Error parsing form:", err)
            http.Error(w, "Unable to parse form", http.StatusBadRequest)
            return
        }

        // Extract form values
        name := r.FormValue("name")
        email := r.FormValue("email")
        address := r.FormValue("address")
        password := r.FormValue("password")

        // Check for missing fields
        if name == "" || email == "" || address == "" || password == "" {
            http.Error(w, "Missing required fields", http.StatusBadRequest)
            return
        }

        // Insert user into the database
        err = utility.InsertUser(utility.Db, name, email, address, password)
        if err != nil {
            log.Println("Error inserting user:", err)
            http.Error(w, "Failed to save user", http.StatusInternalServerError)
            return
        }

        // Redirect to login page
        http.Redirect(w, r, "/login", http.StatusSeeOther)
    } else {
        // Serve signup page
        utility.View.ExecuteTemplate(w, "signup.html", nil)
    }
}


// Login handler
// Login handler
func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        email := r.FormValue("email")
        password := r.FormValue("password")

        var user models.User
        err := utility.Db.Get(&user, "SELECT * FROM users WHERE email = ? AND password = ?", email, password)
        if err != nil {
            http.Error(w, "Invalid credentials", http.StatusUnauthorized)
            log.Println("Login error:", err) // Log the error for debugging
            return
        }

        http.Redirect(w, r, "/home?email="+email, http.StatusFound)
    } else {
        // Serve login page
        utility.View.ExecuteTemplate(w, "login.html", nil)
    }
}


// Home handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
    email := r.URL.Query().Get("email")

    var user models.User
    err := utility.Db.Get(&user, "SELECT * FROM users WHERE email = ?", email)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    utility.View.ExecuteTemplate(w, "home.html", user)
}
