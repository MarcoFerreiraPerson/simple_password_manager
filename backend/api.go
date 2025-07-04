package main

import (
	authHandlers "password_manager/api/handlers/auth"
	vaultHandlers "password_manager/api/handlers/vault"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/auth/register", authHandlers.RegisterUser)
	router.POST("/auth/login", authHandlers.Login)
	router.PUT("/auth/account/password", authHandlers.ChangePassword)
	router.DELETE("/auth/account", authHandlers.DeleteAccount)

	router.GET("/vault", vaultHandlers.GetPasswords)
	router.POST("/vault/add", vaultHandlers.AddPassword)
	router.PUT("/vault/:id", vaultHandlers.GetPasswordById)
	router.DELETE("/vault/:id", vaultHandlers.Delete)

	router.Run("localhost:8080")
}
