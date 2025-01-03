package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	// Initialize a Gin router using the default settings
	router := gin.Default()

	// Define the route for getting all albums
	router.GET("/albums", getAlbums)

	// Define the route for getting a specific album by ID
	router.GET("/albums/:id", getAlbumByID)

	// Define the route for adding a new album
	router.POST("/albums", postAlbums)

	// Run the server on localhost at port 8080
	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	// Respond with the list of albums in JSON format
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)

	// Respond with the newly created album in JSON format
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			// Respond with the found album in JSON format
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	// If no album is found, respond with a "not found" message
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
