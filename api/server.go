package api

import (
	db "github.com/daryalazareva/microapp/db/sqlc"
	"github.com/gin-gonic/gin"
)

//to server http requests
type Server struct {
	store  *db.Store   //interacting with db
	router *gin.Engine //help us send req to the correct handler
}

//creates a new http server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//handlers need to be methods of the Server struct because we have to get access to the store object in order to save new users to the db
	router.POST("/signup", server.createUser)
	router.GET("/signin", server.authUser)
	router.PUT("/changepassword", server.updateUsersPassword)

	server.router = router
	return server
}

//runs the http server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

//error must become a key:value type
//gin.H is a shortcut for map[string]interface{}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
