package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Recipe struct {
	Id                 int      `json:"id"`
	Name               string   `json:"name"`
	Ingredients        []string `json:"ingredients"`
	Instructions       []string `json:"instructions"`
	PrepTimeMinutes    int      `json:"prepTimeMinutes"`
	CookTimeMinutes    int      `json:"cookTimeMinutes"`
	Servings           int      `json:"servings"`
	Difficulty         string   `json:"difficulty"`
	Cuisine            string   `json:"cuisine"`
	CaloriesPerServing int      `json:"caloriesPerServing"`
	Tags               []string `json:"tags"`
	UserId             int      `json:"userId"`
	Image              string   `json:"image"`
	Rating             float64  `json:"rating"`
	ReviewCount        int      `json:"reviewCount"`
	MealType           []string `json:"mealType"`
}

type RecipeResponse struct {
	Recipes []Recipe `json:"recipes"`
}

func RegisterAPIRoutes(server *gin.Engine) {
	server.GET("/api", func(c *gin.Context) {
		url := "https://dummyjson.com/recipes"
		// method := "GET"

		res, err := http.Get(url)

		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(body))
		var response RecipeResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.JSON(200, gin.H{
			"message": "hello world",
			"data":    response,
		})
	})
}
