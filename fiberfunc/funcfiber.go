package fiberfunc

import (
	"github.com/gofiber/fiber/v2"
)

func (Agency) Get(c *fiber.Ctx) error {
	return GetGeneric[Agency](c)
}
func (Agency) GetAll(c *fiber.Ctx) error {
	return GetAllGeneric[Agency](c)
}
func (Agency) Create(c *fiber.Ctx) error {
	return CreateGeneric[Agency](c)
}
func (Agency) Update(c *fiber.Ctx) error {
	return UpdateGeneric[Agency](c)
}
func (Agency) Delete(c *fiber.Ctx) error {
	return DeteleGeneric[Agency](c)
}

func (Auditor) Get(c *fiber.Ctx) error {
	return GetGeneric[Auditor](c)
}
func (Auditor) GetAll(c *fiber.Ctx) error {
	return GetAllGeneric[Auditor](c)
}
func (Auditor) Create(c *fiber.Ctx) error {
	return CreateGeneric[Auditor](c)
}
func (Auditor) Update(c *fiber.Ctx) error {
	return UpdateGeneric[Auditor](c)
}
func (Auditor) Delete(c *fiber.Ctx) error {
	return DeteleGeneric[Auditor](c)
}

func (Category) Get(c *fiber.Ctx) error {
	return GetGeneric[Category](c)
}
func (Category) GetAll(c *fiber.Ctx) error {
	return GetAllGeneric[Category](c)
}
func (Category) Create(c *fiber.Ctx) error {
	return CreateGeneric[Category](c)
}
func (Category) Update(c *fiber.Ctx) error {
	return UpdateGeneric[Category](c)
}
func (Category) Delete(c *fiber.Ctx) error {
	return DeteleGeneric[Category](c)
}

func (Money) Get(c *fiber.Ctx) error {
	return GetGeneric[Money](c)
}
func (Money) GetAll(c *fiber.Ctx) error {
	return GetAllGeneric[Money](c)
}
func (Money) Create(c *fiber.Ctx) error {
	return CreateGeneric[Money](c)
}
func (Money) Update(c *fiber.Ctx) error {
	return UpdateGeneric[Money](c)
}
func (Money) Delete(c *fiber.Ctx) error {
	return DeteleGeneric[Money](c)
}

func (Province) Get(c *fiber.Ctx) error {
	return GetGeneric[Province](c)
}
func (Province) GetAll(c *fiber.Ctx) error {
	return GetAllGeneric[Province](c)
}
func (Province) Create(c *fiber.Ctx) error {
	return CreateGeneric[Province](c)
}
func (Province) Update(c *fiber.Ctx) error {
	return UpdateGeneric[Province](c)
}
func (Province) Delete(c *fiber.Ctx) error {
	return DeteleGeneric[Province](c)
}

func (Source) Get(c *fiber.Ctx) error {
	return GetGeneric[Source](c)
}
func (Source) GetAll(c *fiber.Ctx) error {
	return GetAllGeneric[Source](c)
}
func (Source) Create(c *fiber.Ctx) error {
	return CreateGeneric[Source](c)
}
func (Source) Update(c *fiber.Ctx) error {
	return UpdateGeneric[Source](c)
}
func (Source) Delete(c *fiber.Ctx) error {
	return DeteleGeneric[Source](c)
}

func (Unit) Get(c *fiber.Ctx) error {
	return GetGeneric[Unit](c)
}
func (Unit) GetAll(c *fiber.Ctx) error {
	return GetAllGeneric[Unit](c)
}
func (Unit) Create(c *fiber.Ctx) error {
	return CreateGeneric[Unit](c)
}
func (Unit) Update(c *fiber.Ctx) error {
	return UpdateGeneric[Unit](c)
}
func (Unit) Delete(c *fiber.Ctx) error {
	return DeteleGeneric[Unit](c)
}

func (User) Get(c *fiber.Ctx) error {
	return GetGeneric[User](c)
}
func (User) GetAll(c *fiber.Ctx) error {
	return GetAllGeneric[User](c)
}
func (User) Create(c *fiber.Ctx) error {
	return CreateGeneric[User](c)
}
func (User) Update(c *fiber.Ctx) error {
	return UpdateGeneric[User](c)
}
func (User) Delete(c *fiber.Ctx) error {
	return DeteleGeneric[User](c)
}
