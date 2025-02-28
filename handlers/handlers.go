// handlers/handlers.go
package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// DB instance
var DB *sql.DB

// Struct untuk request body
type RequestBody struct {
    Data string `json:"data"`
}

// InsertUser adalah handler untuk menambahkan user ke database
func CreateUser(c echo.Context) error {  // Make sure the function starts with an uppercase letter
    req := new(RequestBody)
    if err := c.Bind(req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
    }

    data := strings.TrimSpace(req.Data)
    if data == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Data field is required."})
    }

    arrayString := strings.Fields(data)
    if len(arrayString) < 3 {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Input not valid"})
    }

    var name, age, city string
    ageIndex := -1

    for i, word := range arrayString {
        if age == "" && strings.HasPrefix(word, "0") || strings.HasPrefix(word, "1") || strings.HasPrefix(word, "2") {
            age = word
            ageIndex = i
            continue
        }

        if ageIndex == -1 {
            name += strings.ToUpper(word) + " "
        } else if i > ageIndex && word != "th" && word != "thn" && word != "tahun" {
            city += strings.ToUpper(word) + " "
        }
    }

    // Insert data ke database
    query := `INSERT INTO users (name, age, city) VALUES ($1, $2, $3)`
    _, err := DB.Exec(query, strings.TrimSpace(name), age, strings.TrimSpace(city))
    if err != nil {
        log.Println("Error inserting data into database:", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error storing data in the database"})
    }

    return c.JSON(http.StatusCreated, map[string]string{"message": "Succeed"})
}
