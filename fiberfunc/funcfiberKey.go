package fiberfunc

import "github.com/gofiber/fiber/v2"

func (Consider) Get(c *fiber.Ctx) error {
	return GetGenericWithKey[Consider](c, ConsiderKey)
}
func (Consider) GetAll(c *fiber.Ctx) error {
	return GetAllGenericWithKey[Consider](c, ConsiderKey)
}
func (Consider) Create(c *fiber.Ctx) error {
	return CreateGeneric[Consider](c)
}
func (Consider) Update(c *fiber.Ctx) error {
	return UpdateGeneric[Consider](c)
}
func (Consider) Delete(c *fiber.Ctx) error {
	return DeteleGeneric[Consider](c)
}

func (Equipment) Get(c *fiber.Ctx) error {
	return GetGenericWithKey[Equipment](c, EquipmentKey)
}
func (Equipment) GetAll(c *fiber.Ctx) error {
	return GetAllGenericWithKey[Equipment](c, EquipmentKey)
}
func (Equipment) Create(c *fiber.Ctx) error {
	return CreateGeneric[Equipment](c)
}
func (Equipment) Update(c *fiber.Ctx) error {
	return UpdateGeneric[Equipment](c)
}
func (Equipment) Delete(c *fiber.Ctx) error {
	return DeteleGeneric[Equipment](c)
}

func (File) Get(c *fiber.Ctx) error {
	return GetGenericWithKey[File](c, FileKey)
}
func (File) GetAll(c *fiber.Ctx) error {
	return GetAllGenericWithKey[File](c, FileKey)
}
func (File) Create(c *fiber.Ctx) error {
	return UploadfileGeneric[File](c, FileModelsUpload)
}
func (File) Update(c *fiber.Ctx) error {
	return UpdateFile[File](c, FileModelsUpload)
}
func (File) Delete(c *fiber.Ctx) error {
	return DeteleGeneric[File](c)
}

func (Location) Get(c *fiber.Ctx) error {
	return GetGenericWithKey[Location](c, LocationKey)
}
func (Location) GetAll(c *fiber.Ctx) error {
	return GetAllGenericWithKey[Location](c, LocationKey)
}
func (Location) Create(c *fiber.Ctx) error {
	return CreateGeneric[Location](c)
}
func (Location) Update(c *fiber.Ctx) error {
	return UpdateGeneric[Location](c)
}
func (Location) Delete(c *fiber.Ctx) error {
	return DeteleGeneric[Location](c)
}

func (Necessary) Get(c *fiber.Ctx) error {
	return GetGenericWithKey[Necessary](c, NecessaryKey)
}
func (Necessary) GetAll(c *fiber.Ctx) error {
	return GetAllGenericWithKey[Necessary](c, NecessaryKey)
}
func (Necessary) Create(c *fiber.Ctx) error {
	return CreateGeneric[Necessary](c)
}
func (Necessary) Update(c *fiber.Ctx) error {
	return UpdateGeneric[Necessary](c)
}
func (Necessary) Delete(c *fiber.Ctx) error {
	return DeteleGeneric[Necessary](c)
}

func (Price) Get(c *fiber.Ctx) error {
	return GetGenericWithKey[Price](c, PriceKey)
}
func (Price) GetAll(c *fiber.Ctx) error {
	return GetAllGenericWithKey[Price](c, PriceKey)
}
func (Price) Create(c *fiber.Ctx) error {
	return CreateGeneric[Price](c)
}
func (Price) Update(c *fiber.Ctx) error {
	return UpdateGeneric[Price](c)
}
func (Price) Delete(c *fiber.Ctx) error {
	return DeteleGeneric[Price](c)
}

func (Reqlacement) Get(c *fiber.Ctx) error {
	return GetGenericWithKey[Reqlacement](c, ReqlacementKey)
}
func (Reqlacement) GetAll(c *fiber.Ctx) error {
	return GetAllGenericWithKey[Reqlacement](c, ReqlacementKey)
}
func (Reqlacement) Create(c *fiber.Ctx) error {
	return CreateGeneric[Reqlacement](c)
}
func (Reqlacement) Update(c *fiber.Ctx) error {
	return UpdateGeneric[Reqlacement](c)
}
func (Reqlacement) Delete(c *fiber.Ctx) error {
	return DeteleGeneric[Reqlacement](c)
}
