package routes

import (
	"bre-api/fiber/fiberfunc/fiberfuncCOUD"

	"github.com/gofiber/fiber/v2"
)

func (Consider) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGenericWithKey[Consider](c, ConsiderKey)
}
func (Consider) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGenericWithKey[Consider](c, ConsiderKey)
}
func (Consider) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Consider](c)
}
func (Consider) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Consider](c)
}
func (Consider) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Consider](c)
}

func (Equipment) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGenericWithKey[Equipment](c, EquipmentKey)
}
func (Equipment) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGenericWithKey[Equipment](c, EquipmentKey)
}
func (Equipment) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Equipment](c)
}
func (Equipment) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Equipment](c)
}
func (Equipment) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Equipment](c)
}

func (File) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGenericWithKey[File](c, FileKey)
}
func (File) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGenericWithKey[File](c, FileKey)
}
func (File) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[File](c)
}
func (File) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[File](c)
}
func (File) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[File](c)
}

func (Location) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGenericWithKey[Location](c, LocationKey)
}
func (Location) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGenericWithKey[Location](c, LocationKey)
}
func (Location) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Location](c)
}
func (Location) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Location](c)
}
func (Location) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Location](c)
}

func (Necessary) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGenericWithKey[Necessary](c, NecessaryKey)
}
func (Necessary) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGenericWithKey[Necessary](c, NecessaryKey)
}
func (Necessary) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Necessary](c)
}
func (Necessary) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Necessary](c)
}
func (Necessary) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Necessary](c)
}

func (Price) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGenericWithKey[Price](c, PriceKey)
}
func (Price) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGenericWithKey[Price](c, PriceKey)
}
func (Price) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Price](c)
}
func (Price) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Price](c)
}
func (Price) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Price](c)
}

func (Reqlacement) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGenericWithKey[Reqlacement](c, ReqlacementKey)
}
func (Reqlacement) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGenericWithKey[Reqlacement](c, ReqlacementKey)
}
func (Reqlacement) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Reqlacement](c)
}
func (Reqlacement) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Reqlacement](c)
}
func (Reqlacement) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Reqlacement](c)
}
