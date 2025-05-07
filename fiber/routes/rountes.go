package routes

import (
	"bre-api/fiber/fiberfunc/fiberfuncCOUD"
	fiberfuncCOUDaddOn "bre-api/fiber/fiberfunc/fiberfuncCOUD/fiberfuncAddOn"

	"github.com/gofiber/fiber/v2"
)

func (Agency) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGeneric[Agency](c)
}
func (Agency) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGeneric[Agency](c)
}
func (Agency) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Agency](c)
}
func (Agency) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Agency](c)
}
func (Agency) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Agency](c)
}

func (Auditor) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGeneric[Auditor](c)
}
func (Auditor) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGeneric[Auditor](c)
}
func (Auditor) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Auditor](c)
}
func (Auditor) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Auditor](c)
}
func (Auditor) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Auditor](c)
}

func (Category) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGeneric[Category](c)
}
func (Category) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGeneric[Category](c)
}
func (Category) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Category](c)
}
func (Category) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Category](c)
}
func (Category) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Category](c)
}

func (Money) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGeneric[Money](c)
}
func (Money) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGeneric[Money](c)
}
func (Money) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Money](c)
}
func (Money) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Money](c)
}
func (Money) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Money](c)
}

func (Province) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGeneric[Province](c)
}
func (Province) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGeneric[Province](c)
}
func (Province) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Province](c)
}
func (Province) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Province](c)
}
func (Province) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Province](c)
}

func (Source) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGeneric[Source](c)
}
func (Source) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGeneric[Source](c)
}
func (Source) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Source](c)
}
func (Source) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Source](c)
}
func (Source) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Source](c)
}

func (Unit) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGeneric[Unit](c)
}
func (Unit) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGeneric[Unit](c)
}
func (Unit) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Unit](c)
}
func (Unit) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Unit](c)
}
func (Unit) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Unit](c)
}

func (User) Get(c *fiber.Ctx) error {
	return c.JSON("this func dont support")
}
func (User) GetAll(c *fiber.Ctx) error {
	return c.JSON("this func dont support")
}
func (User) Create(c *fiber.Ctx) error {
	return fiberfuncCOUDaddOn.CreateUser(c)
}
func (User) Login(c *fiber.Ctx) error {
	return fiberfuncCOUDaddOn.LoginUser(c)
}
func (User) Update(c *fiber.Ctx) error {
	return c.JSON("this func dont support")
}
func (User) Delete(c *fiber.Ctx) error {
	return c.JSON("this func dont support")
}

// ----------------------------------------
func BlockFunc(c *fiber.Ctx) error {
	return fiberfuncCOUDaddOn.BlockFunc(c)
}
