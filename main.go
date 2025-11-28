package main

import (
	"net/http"

	_ "github.com/eduardomarini16/api_jazz/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API de Álbuns
// @version 1.0
// @description Esta API retorna e adiciona álbuns musicais.
// @host localhost:8080
// @BasePath /
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

// @Summary Get an album by ID
// @Description get album by ID
// @Tags albums
// @Accept json
// @Produce json
// @Param id path string true "Album ID"
// @Success 200 {object} album
// @Router /albums/{id} [get]
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default() // inicializa um roteador Gin
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	// rota swageer
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run("localhost:8080") // inicia o servidor na porta 8080
}
