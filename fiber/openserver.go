package fiberopen

import (
	"bre-api/fiber/routes"

	_ "bre-api/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

var (
	Agency      = new(routes.Agency)
	Auditor     = new(routes.Auditor)
	Category    = new(routes.Category)
	Consider    = new(routes.Consider)
	Equipment   = new(routes.Equipment)
	File        = new(routes.File)
	Location    = new(routes.Location)
	Money       = new(routes.Money)
	Necessary   = new(routes.Necessary)
	Price       = new(routes.Price)
	Province    = new(routes.Province)
	Reqlacement = new(routes.Reqlacement)
	Requests    = new(routes.Requests)
	Source      = new(routes.Source)
	Unit        = new(routes.Unit)
	User        = new(routes.User)
)

func OpenServer() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{}))

	// Swagger rounter
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Use("/home", routes.BlockFunc)
	app.Post("/create", User.Create)
	app.Post("/login", User.Login)
	app.Post("/logout", routes.Logout)
	app.Get("/check", routes.CheckJwt)

	app.Get("/home/agency/:id", Agency.Get)

	app.Get("/home/agency/all/:limit", Agency.GetAll)
	app.Post("/home/agency/create", Agency.Create)
	app.Put("/home/agency/update/:id", Agency.Update)
	app.Delete("/home/agency/delete/:id", Agency.Delete)

	app.Get("/home/auditor/:id", Auditor.Get)
	app.Get("/home/auditor/all/:limit", Auditor.GetAll)
	app.Post("/home/auditor/create", Auditor.Create)
	app.Put("/home/auditor/update/:id", Auditor.Update)
	app.Delete("/home/auditor/delete/:id", Auditor.Delete)

	app.Get("/home/category/:id", Category.Get)
	app.Get("/home/category/all/:limit", Category.GetAll)
	app.Post("/home/category/create", Category.Create)
	app.Put("/home/category/update/:id", Category.Update)
	app.Delete("/home/category/delete/:id", Category.Delete)

	app.Get("/home/consider/:id", Consider.Get)
	app.Get("/home/consider/all/:limit", Consider.GetAll)
	app.Post("/home/consider/create", Consider.Create)
	app.Put("/home/consider/update/:id", Consider.Update)
	app.Delete("/home/consider/delete/:id", Consider.Delete)

	app.Get("/home/equipment/:id", Equipment.Get)
	app.Get("/home/equipment/all/:limit", Equipment.GetAll)
	app.Post("/home/equipment/create", Equipment.Create)
	app.Put("/home/equipment/update/:id", Equipment.Update)
	app.Delete("/home/equipment/delete/:id", Equipment.Delete)

	app.Get("/home/file/:id", File.Get)
	app.Get("/home/file/all/:limit", File.GetAll)
	app.Post("/home/file/create", File.Create)
	app.Put("/home/file/update/:id", File.Update)
	app.Delete("/home/file/delete/:id", File.Delete)

	app.Get("/home/location/:id", Location.Get)
	app.Get("/home/location/all/:limit", Location.GetAll)
	app.Post("/home/location/create", Location.Create)
	app.Put("/home/location/update/:id", Location.Update)
	app.Delete("/home/location/delete/:id", Location.Delete)

	app.Get("/home/money/:id", Money.Get)
	app.Get("/home/money/all/:limit", Money.GetAll)
	app.Post("/home/money/create", Money.Create)
	app.Put("/home/money/update/:id", Money.Update)
	app.Delete("/home/money/delete/:id", Money.Delete)

	app.Get("/home/necessary/:id", Necessary.Get)
	app.Get("/home/necessary/all/:limit", Necessary.GetAll)
	app.Post("/home/necessary/create", Necessary.Create)
	app.Put("/home/necessary/update/:id", Necessary.Update)
	app.Delete("/home/necessary/delete/:id", Necessary.Delete)

	app.Get("/home/price/:id", Price.Get)
	app.Get("/home/price/all/:limit", Price.GetAll)
	app.Post("/home/price/create", Price.Create)
	app.Put("/home/price/update/:id", Price.Update)
	app.Delete("/home/price/delete/:id", Price.Delete)

	app.Get("/home/province/:id", Province.Get)
	app.Get("/home/province/all/:limit", Province.GetAll)
	app.Post("/home/province/create", Province.Create)
	app.Put("/home/province/update/:id", Province.Update)
	app.Delete("/home/province/delete/:id", Province.Delete)

	app.Get("/home/reqlacement/:id", Reqlacement.Get)
	app.Get("/home/reqlacement/all/:limit", Reqlacement.GetAll)
	app.Post("/home/reqlacement/create", Reqlacement.Create)
	app.Put("/home/reqlacement/update/:id", Reqlacement.Update)
	app.Delete("/home/reqlacement/delete/:id", Reqlacement.Delete)

	app.Get("/home/Requests/:id", Requests.Get)
	app.Get("/home/Requests/all/:limit", Requests.GetAll)
	app.Post("/home/Requests/create", Requests.Create)
	app.Put("/home/Requests/update/:id", Requests.Update)
	app.Delete("/home/Requests/delete/:id", Requests.Delete)

	app.Get("/home/source/:id", Source.Get)
	app.Get("/home/source/all/:limit", Source.GetAll)
	app.Post("/home/source/create", Source.Create)
	app.Put("/home/source/update/:id", Source.Update)
	app.Delete("/home/source/delete/:id", Source.Delete)

	app.Get("/home/unit/:id", Unit.Get)
	app.Get("/home/unit/all/:limit", Unit.GetAll)
	app.Post("/home/unit/create", Unit.Create)
	app.Put("/home/unit/update/:id", Unit.Update)
	app.Delete("/home/unit/delete/:id", Unit.Delete)

	app.Get("/home/user/:id", User.Get)
	app.Get("/home/user/all/:limit", User.GetAll)
	app.Put("/home/user/update/:id", User.Update)
	app.Delete("/home/user/delete/:id", User.Delete)

	app.Listen(":8000")
}
