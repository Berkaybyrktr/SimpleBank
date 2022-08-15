package Api

import (
	"github.com/gin-gonic/gin"
	db "github/techschool/simplebank/db/sqlc"

	_ "github.com/lib/pq"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)

	router.GET("/accounts/:id", server.getAccount)

	router.GET("/accounts", server.listAccounts)

	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}