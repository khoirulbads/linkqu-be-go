package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"linkqu-be-go/config"

	"github.com/labstack/echo/v4"
)

// Struct re body
type RequestBody struct {
	Data string `json:"data"`
}

func CreateUser(c echo.Context) error {
	req := new(RequestBody)

	// validasi input
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

	// Regex check number
	re := regexp.MustCompile(`^\d+`)

	for i, word := range arrayString {
		// save number as age with index
		if re.MatchString(word) {
			age = re.FindString(word)
			ageIndex = i
		}

		// save name
		if age == "" {
			name += strings.ToUpper(word) + " "
			continue
		}

		// save city
		if i > ageIndex && word != "th" && word != "thn" && word != "tahun" {
			city += strings.ToUpper(word) + " "
		}
	}

	// clear space
	name = strings.TrimSpace(name)
	city = strings.TrimSpace(city)

	// Debug log
	debug := fmt.Sprintf("Insert User - Name: %s, Age: %s, City: %s", name, age, city)
	log.Printf(debug)

	// check db connection
	sqlDB, err := config.DB.DB()
	if err != nil {
		log.Println("Error connect database:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Database connection failed"})
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Println(" Database TO:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Database is unreachable"})
	}
	log.Println("Database connented")

	query := `INSERT INTO users (name, age, city) VALUES (?, ?, ?)`
	err = config.DB.Exec(query, name, age, city).Error
	if err != nil {
		log.Println("Error insert", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error storing data in the database"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Succeed "+debug, })
}