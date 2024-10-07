package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()
	router.POST("find-pairs", findingPairs)
	router.Run("localhost:9090")

}

func findPairs(numbers []int, target int) [][]int {
	output := [][]int{}

	for i, v := range numbers {
		for j, w := range numbers {
			if v == w {
				// same number cannot be added
				continue
			} else if v+w == target {
				output = append(output, []int{i, j})
			}
		}

	}
	return output
}

type Request struct {
	numbers []int `json:"numbers"`
	target  int   `json:"target"`
}

type Response struct {
	solutions  [][]int
	statusCode int
}

func findingPairs(context *gin.Context) {
	newRequest := Request{
		numbers: []int{1, 2, 3, 4, 5, 6},
		target:  7,
	}

	if err := context.BindJSON(&newRequest); err != nil {
		return
	}

	output := findPairs(newRequest.numbers, newRequest.target)

	context.IndentedJSON(http.StatusOK, output)
}
