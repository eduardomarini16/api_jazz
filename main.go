package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// @Summary List all albums
// @Description get all albums
// @Tags albums
// @Accept json
// @Produce json
// @Success 200 {array} album
// @Router /albums [get]
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums) // serializar a estrutura em JSON e adicioná-la a resposta
}

// @Summary Add a new album
// @Description add a new album
// @Tags albums
// @Accept json
// @Produce json
// @Param album body album true "Album to add"
// @Success 201 {object} album
// @Router /albums [post]
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Vincula o corpo JSON da solicitação à nova estrutura de álbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default() // inicializa um roteador Gin
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080") // inicia o servidor na porta 8080
}
