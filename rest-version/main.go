// *GIN
// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type Numbers struct {
// 	Values []float64 `json:"values" binding:"required"`
// }

// type Result struct {
// 	Sum     float64 `json:"sum"`
// 	Average float64 `json:"average"`
// 	Product float64 `json:"product"`
// }

// func computeHandler(c *gin.Context) {
// 	var numbers Numbers

// 	if err := c.ShouldBindJSON(&numbers); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	sum := 0.0
// 	product := 1.0
// 	for _, v := range numbers.Values {
// 		sum += v
// 		product *= v
// 	}
// 	average := sum / float64(len(numbers.Values))

// 	result := Result{
// 		Sum:     sum,
// 		Average: average,
// 		Product: product,
// 	}

// 	c.JSON(http.StatusOK, result)
// }

// func main() {
// 	r := gin.Default()

// 	r.POST("/compute", computeHandler)

// 	r.Run(":8080")
// }

// *NATIVE
// package main

// import (
// 	"encoding/json"
// 	"net/http"
// )

// type Numbers struct {
// 	Values []float64 `json:"values"`
// }

// type Result struct {
// 	Sum     float64 `json:"sum"`
// 	Average float64 `json:"average"`
// 	Product float64 `json:"product"`
// }

// func computeHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var numbers Numbers
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&numbers); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	sum := 0.0
// 	product := 1.0
// 	for _, v := range numbers.Values {
// 		sum += v
// 		product *= v
// 	}
// 	average := sum / float64(len(numbers.Values))

// 	result := Result{
// 		Sum:     sum,
// 		Average: average,
// 		Product: product,
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	encoder := json.NewEncoder(w)
// 	encoder.Encode(result)
// }

// func main() {
// 	http.HandleFunc("/compute", computeHandler)
// 	http.ListenAndServe(":8080", nil)
// }

// *FIBER
package main

import (
	"github.com/gofiber/fiber/v2"
)

type Numbers struct {
	Values []float64 `json:"values"`
}

type Result struct {
	Sum     float64 `json:"sum"`
	Average float64 `json:"average"`
	Product float64 `json:"product"`
}

func computeHandler(c *fiber.Ctx) error {
	var numbers Numbers

	if err := c.BodyParser(&numbers); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	sum := 0.0
	product := 1.0
	for _, v := range numbers.Values {
		sum += v
		product *= v
	}
	average := sum / float64(len(numbers.Values))

	result := Result{
		Sum:     sum,
		Average: average,
		Product: product,
	}

	return c.JSON(result)
}

func main() {
	app := fiber.New()

	app.Post("/compute", computeHandler)

	app.Listen(":8080")
}
