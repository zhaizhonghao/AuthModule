package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zhaizhonghao/auth/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/user/all", controllers.GetAllUsers)
	app.Post("/api/user/delete", controllers.DeleteUser)

	app.Post("/api/acl/add", controllers.AddACLEntry)
	app.Post("/api/acl/delete", controllers.DeleteACLEntry)
	app.Get("/api/acl/all", controllers.GetAllEntries)
}
