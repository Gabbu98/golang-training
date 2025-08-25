package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/Gabbu98.golang-training-phase-two/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func httpClient() *http.Client {
	return &http.Client{}
}

func requestNew(ctx context.Context, httpMethod, uri string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, httpMethod, uri, body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return req, err
	}

	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

func execRequest(request *http.Request, bodyInterface interface{}) (status int, err error) {
	resp, err := httpClient().Do(request)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	status = resp.StatusCode
	defer resp.Body.Close()

	// check for success status
	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	// Decode json
	decoder := json.NewDecoder(resp.Body)
	if bodyInterface != nil {
		if err = decoder.Decode(bodyInterface); err != nil {
			return status, fmt.Errorf("error decoding response body: %v", err)
		}
	}

	return
}

func GetPokemonDetails(ctx context.Context, pokemonName string) (response *models.PokemonDetails, status int, err error) {
	req, err := requestNew(ctx, http.MethodGet, fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", strings.ToLower(pokemonName)), nil)
	if err != nil {
		return response, http.StatusInternalServerError, err
	}

	status, err = execRequest(req, &response)
	if err != nil {
		return
	}

	return
}

func testHttpClient() {
	ctx := context.Background()
	resData, status, err := GetPokemonDetails(ctx, "metapod")
	if err != nil {
		logrus.Error(err)
		return
	}

	switch status {
	case http.StatusOK:
		jsonjson, _ := json.Marshal(resData)
		logrus.Info(string(jsonjson))
	default:
		err = fmt.Errorf("response api %v", status)
		logrus.Error(err)
		return
	}
}

func routingTraining() {
	var students = []int{1,2,3,4,5}
	// Create a router with default middleware: logger & recovery
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"response": "pong",
		})
	})

	r.POST("/submit", func(c *gin.Context){
		c.JSON(200, gin.H{"status": "submitted"})
	})

	r.PUT("/update", func(c *gin.Context){
		c.JSON(200, gin.H{"status": "updated"})
	})

	r.DELETE("/delete", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "deleted"})
	})

	studentRoutes:=r.Group("/students") 
	{
		// Parameters
		studentRoutes.GET("/:id", func (c *gin.Context)  {
			id , err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(400, gin.H{"error":"Bad Request"})
				return
			}
			if slices.Contains(students, id) {
				c.JSON(200, gin.H{"status": "Student found!"})
				return
			}

			c.JSON(404, gin.H{"status": "Student not found!"})
		})

		studentRoutes.GET("/", func (c *gin.Context)  {
			q := c.Query("largerThan")
			var query int
			var err error
			if q != "" {
				query, err = strconv.Atoi(q) // Get query param ?largerThan=
				if err != nil {
					c.JSON(400, gin.H{"error":"Bad Request"})
					return
				}
			}

			var students_filtered = []int{}
			for _, res := range students {
				if res > query {
					students_filtered = append(students_filtered, res)
				}
			}
			c.JSON(200, gin.H{"response": students_filtered})
		})
	}
	
	r.Run(":8080")
}

func main() {
	

}
