package routes

import (
	"bre-api/fiber/fiberfunc/fiberfuncCOUD"

	"github.com/gofiber/fiber/v2"
)

//---------------------------------------------------------------------------------------//

// GetConsider godoc
// @Summary ดึงข้อมูล Consider ตาม ID
// @Description ดึงข้อมูล Consider รายการเดียวตาม ID
// @Tags Consider
// @Accept json
// @Produce json
// @Param id path int true "รหัส Consider"
// @Success 200 {object} models.Consider "ดึงข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบ Consider ที่ระบุ"
// @Failure 500 {object} map[string]interface{} "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/consider/{id} [get]
func (Consider) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGenericWithKey[Consider](c, ConsiderKey)
}

// GetAllConsider godoc
// @Summary ดึงข้อมูล Consider ทั้งหมด
// @Description ดึงข้อมูล Consider หลายรายการพร้อมจำกัดจำนวนผลลัพธ์
// @Tags Consider
// @Accept json
// @Produce json
// @Param limit path string true "จำนวนสูงสุดของรายการที่ต้องการดึง (ใช้ค่า config limit หากเป็น string)"
// @Success 200 {array} models.Consider "รายการ Consider"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/consider/all/{limit} [get]
func (Consider) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGenericWithKey[Consider](c, ConsiderKey)
}

// CreateConsider godoc
// @Summary สร้าง Consider ใหม่
// @Description เพิ่มข้อมูล Consider ใหม่
// @Tags Consider
// @Accept json
// @Produce json
// @Param consider body models.Consider true "ข้อมูล Consider"
// @Success 201 {object} models.Consider "สร้างข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "ข้อมูลนำเข้าไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/consider/create [post]
func (Consider) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Consider](c)
}

// UpdateConsider godoc
// @Summary อัพเดต Consider
// @Description แก้ไขข้อมูล Consider ตาม ID
// @Tags Consider
// @Accept json
// @Produce json
// @Param id path int true "รหัส Consider"
// @Param consider body models.Consider true "ข้อมูล Consider ที่แก้ไข"
// @Success 200 {object} models.Consider "อัพเดตข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบ Consider ที่ระบุ"
// @Failure 500 {object} map[string]interface{} "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/consider/update/{id} [put]
func (Consider) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Consider](c)
}

// DeleteConsider godoc
// @Summary ลบ Consider
// @Description ลบข้อมูล Consider ตาม ID
// @Tags Consider
// @Accept json
// @Produce json
// @Param id path int true "รหัส Consider"
// @Success 200 {object} map[string]interface{} "ลบข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบ Consider ที่ระบุ"
// @Failure 500 {object} map[string]interface{} "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/consider/delete/{id} [delete]
func (Consider) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Consider](c)
}

//---------------------------------------------------------------------------------------//

// GetEquipment godoc
// @Summary ดึงข้อมูล Equipment ตาม ID
// @Description ดึงข้อมูล Equipment รายการเดียวตาม ID
// @Tags Equipment
// @Accept json
// @Produce json
// @Param id path int true "รหัส Equipment"
// @Success 200 {object} models.Equipment "ดึงข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบ Equipment ที่ระบุ"
// @Failure 500 {object} map[string]interface{} "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/equipment/{id} [get]
func (Equipment) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGenericWithKey[Equipment](c, EquipmentKey)
}

// GetAllEquipment godoc
// @Summary ดึงข้อมูล Equipment ทั้งหมด
// @Description ดึงข้อมูล Equipment หลายรายการพร้อมจำกัดจำนวนผลลัพธ์
// @Tags Equipment
// @Accept json
// @Produce json
// @Param limit path string true "จำนวนสูงสุดของรายการที่ต้องการดึง (ใช้ค่า config limit หากเป็น string)"
// @Success 200 {array} models.Equipment "รายการ Equipment"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/equipment/all/{limit} [get]
func (Equipment) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGenericWithKey[Equipment](c, EquipmentKey)
}

// CreateEquipment godoc
// @Summary สร้าง Equipment ใหม่
// @Description เพิ่มข้อมูล Equipment ใหม่
// @Tags Equipment
// @Accept json
// @Produce json
// @Param equipment body models.Equipment true "ข้อมูล Equipment"
// @Success 201 {object} models.Equipment "สร้างข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "ข้อมูลนำเข้าไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/equipment/create [post]
func (Equipment) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Equipment](c)
}

// UpdateEquipment godoc
// @Summary อัพเดต Equipment
// @Description แก้ไขข้อมูล Equipment ตาม ID
// @Tags Equipment
// @Accept json
// @Produce json
// @Param id path int true "รหัส Equipment"
// @Param equipment body models.Equipment true "ข้อมูล Equipment ที่แก้ไข"
// @Success 200 {object} models.Equipment "อัพเดตข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบ Equipment ที่ระบุ"
// @Failure 500 {object} map[string]interface{} "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/equipment/update/{id} [put]
func (Equipment) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Equipment](c)
}

// DeleteEquipment godoc
// @Summary ลบ Equipment
// @Description ลบข้อมูล Equipment ตาม ID
// @Tags Equipment
// @Accept json
// @Produce json
// @Param id path int true "รหัส Equipment"
// @Success 200 {object} map[string]interface{} "ลบข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบ Equipment ที่ระบุ"
// @Failure 500 {object} map[string]interface{} "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/equipment/delete/{id} [delete]
func (Equipment) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Equipment](c)
}

//---------------------------------------------------------------------------------------//

// GetFile godoc
// @Summary ดึงข้อมูล File ตาม ID
// @Description ดึงข้อมูล File รายการเดียวตาม ID
// @Tags File
// @Accept json
// @Produce json
// @Param id path int true "รหัส File"
// @Success 200 {object} models.File "ดึงข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบ File ที่ระบุ"
// @Failure 500 {object} map[string]interface{} "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/file/{id} [get]
func (File) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGenericWithKey[File](c, FileKey)
}

// GetAllFile godoc
// @Summary ดึงข้อมูล File ทั้งหมด
// @Description ดึงข้อมูล File หลายรายการพร้อมจำกัดจำนวนผลลัพธ์
// @Tags File
// @Accept json
// @Produce json
// @Param limit path string true "จำนวนสูงสุดของรายการที่ต้องการดึง (หรือใช้ค่า config limit)"
// @Success 200 {array} models.File "รายการ File"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/file/all/{limit} [get]
func (File) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGenericWithKey[File](c, FileKey)
}

// CreateFile godoc
// @Summary สร้าง File ใหม่
// @Description เพิ่มข้อมูล File ใหม่
// @Tags File
// @Accept json
// @Produce json
// @Param file body models.File true "ข้อมูล File"
// @Success 201 {object} models.File "สร้างข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "ข้อมูลนำเข้าไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/file/create [post]
func (File) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[File](c)
}

// UpdateFile godoc
// @Summary อัพเดต File
// @Description แก้ไขข้อมูล File ตาม ID
// @Tags File
// @Accept json
// @Produce json
// @Param id path int true "รหัส File"
// @Param file body models.File true "ข้อมูล File ที่แก้ไข"
// @Success 200 {object} models.File "อัพเดตข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบ File ที่ระบุ"
// @Failure 500 {object} map[string]interface{} "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/file/update/{id} [put]
func (File) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[File](c)
}

// DeleteFile godoc
// @Summary ลบ File
// @Description ลบข้อมูล File ตาม ID
// @Tags File
// @Accept json
// @Produce json
// @Param id path int true "รหัส File"
// @Success 200 {object} map[string]interface{} "ลบข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบ File ที่ระบุ"
// @Failure 500 {object} map[string]interface{} "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/file/delete/{id} [delete]
func (File) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[File](c)
}

//---------------------------------------------------------------------------------------//

// GetLocation godoc
// @Summary ดึงข้อมูล Location ตาม ID
// @Description ดึงข้อมูล Location รายการเดียวตาม ID
// @Tags Location
// @Accept json
// @Produce json
// @Param id path int true "รหัส Location"
// @Success 200 {object} models.Location "ดึงข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบ Location ที่ระบุ"
// @Failure 500 {object} map[string]interface{} "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/location/{id} [get]
func (Location) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGenericWithKey[Location](c, LocationKey)
}

// GetAllLocation godoc
// @Summary ดึงข้อมูล Location ทั้งหมด
// @Description ดึงข้อมูล Location หลายรายการพร้อมจำกัดจำนวนผลลัพธ์
// @Tags Location
// @Accept json
// @Produce json
// @Param limit path string true "จำนวนสูงสุดของรายการที่ต้องการดึง (หรือใช้ค่า config limit)"
// @Success 200 {array} models.Location "รายการ Location"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/location/all/{limit} [get]
func (Location) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGenericWithKey[Location](c, LocationKey)
}

// CreateLocation godoc
// @Summary สร้าง Location ใหม่
// @Description เพิ่มข้อมูล Location ใหม่
// @Tags Location
// @Accept json
// @Produce json
// @Param location body models.Location true "ข้อมูล Location"
// @Success 201 {object} models.Location "สร้างข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "ข้อมูลนำเข้าไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/location/create [post]
func (Location) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Location](c)
}

// UpdateLocation godoc
// @Summary อัพเดต Location
// @Description แก้ไขข้อมูล Location ตาม ID
// @Tags Location
// @Accept json
// @Produce json
// @Param id path int true "รหัส Location"
// @Param location body models.Location true "ข้อมูล Location ที่แก้ไข"
// @Success 200 {object} models.Location "อัพเดตข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบ Location ที่ระบุ"
// @Failure 500 {object} map[string]interface{} "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/location/update/{id} [put]
func (Location) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Location](c)
}

// DeleteLocation godoc
// @Summary ลบ Location
// @Description ลบข้อมูล Location ตาม ID
// @Tags Location
// @Accept json
// @Produce json
// @Param id path int true "รหัส Location"
// @Success 200 {object} map[string]interface{} "ลบข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบ Location ที่ระบุ"
// @Failure 500 {object} map[string]interface{} "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/location/delete/{id} [delete]
func (Location) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Location](c)
}

//---------------------------------------------------------------------------------------//

// GetNecessary godoc
// @Summary ดึงข้อมูล Necessary ตาม ID
// @Description ดึงข้อมูล Necessary รายการเดียวโดยใช้ ID
// @Tags Necessary
// @Accept json
// @Produce json
// @Param id path int true "รหัส Necessary"
// @Success 200 {object} models.Necessary "ดึงข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบข้อมูล Necessary"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/necessary/{id} [get]
func (Necessary) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGenericWithKey[Necessary](c, NecessaryKey)
}

// GetAllNecessary godoc
// @Summary ดึงข้อมูล Necessary ทั้งหมด
// @Description ดึงข้อมูล Necessary หลายรายการโดยจำกัดจำนวน
// @Tags Necessary
// @Accept json
// @Produce json
// @Param limit path string true "จำนวนสูงสุดของข้อมูลที่จะดึง"
// @Success 200 {array} models.Necessary "รายการ Necessary"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/necessary/all/{limit} [get]
func (Necessary) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGenericWithKey[Necessary](c, NecessaryKey)
}

// CreateNecessary godoc
// @Summary สร้างข้อมูล Necessary ใหม่
// @Description เพิ่มข้อมูล Necessary ใหม่
// @Tags Necessary
// @Accept json
// @Produce json
// @Param necessary body models.Necessary true "ข้อมูล Necessary"
// @Success 201 {object} models.Necessary "สร้างข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "ข้อมูลส่งมาไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/necessary/create [post]
func (Necessary) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Necessary](c)
}

// UpdateNecessary godoc
// @Summary อัปเดตข้อมูล Necessary ตาม ID
// @Description อัปเดตข้อมูล Necessary ที่มีอยู่
// @Tags Necessary
// @Accept json
// @Produce json
// @Param id path int true "รหัส Necessary"
// @Param necessary body models.Necessary true "ข้อมูล Necessary ที่ต้องการอัปเดต"
// @Success 200 {object} models.Necessary "อัปเดตข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบข้อมูล Necessary"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/necessary/update/{id} [put]
func (Necessary) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Necessary](c)
}

// DeleteNecessary godoc
// @Summary ลบข้อมูล Necessary ตาม ID
// @Description ลบข้อมูล Necessary โดยใช้ ID
// @Tags Necessary
// @Accept json
// @Produce json
// @Param id path int true "รหัส Necessary"
// @Success 200 {object} map[string]interface{} "ลบข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบข้อมูล Necessary"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/necessary/delete/{id} [delete]
func (Necessary) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Necessary](c)
}

//---------------------------------------------------------------------------------------//

// GetPrice godoc
// @Summary ดึงข้อมูล Price ตาม ID
// @Description ดึงข้อมูล Price รายการเดียวโดยใช้ ID
// @Tags Price
// @Accept json
// @Produce json
// @Param id path int true "รหัส Price"
// @Success 200 {object} models.Price "ดึงข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบข้อมูล Price"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/price/{id} [get]
func (Price) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGenericWithKey[Price](c, PriceKey)
}

// GetAllPrice godoc
// @Summary ดึงข้อมูล Price ทั้งหมด
// @Description ดึงข้อมูล Price หลายรายการโดยจำกัดจำนวน
// @Tags Price
// @Accept json
// @Produce json
// @Param limit path string true "จำนวนสูงสุดของข้อมูลที่จะดึง"
// @Success 200 {array} models.Price "รายการ Price"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/price/all/{limit} [get]
func (Price) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGenericWithKey[Price](c, PriceKey)
}

// CreatePrice godoc
// @Summary สร้างข้อมูล Price ใหม่
// @Description เพิ่มข้อมูล Price ใหม่
// @Tags Price
// @Accept json
// @Produce json
// @Param price body models.Price true "ข้อมูล Price"
// @Success 201 {object} models.Price "สร้างข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "ข้อมูลส่งมาไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/price/create [post]
func (Price) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Price](c)
}

// UpdatePrice godoc
// @Summary อัปเดตข้อมูล Price ตาม ID
// @Description อัปเดตข้อมูล Price ที่มีอยู่
// @Tags Price
// @Accept json
// @Produce json
// @Param id path int true "รหัส Price"
// @Param price body models.Price true "ข้อมูล Price ที่ต้องการอัปเดต"
// @Success 200 {object} models.Price "อัปเดตข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบข้อมูล Price"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/price/update/{id} [put]
func (Price) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Price](c)
}

// DeletePrice godoc
// @Summary ลบข้อมูล Price ตาม ID
// @Description ลบข้อมูล Price โดยใช้ ID
// @Tags Price
// @Accept json
// @Produce json
// @Param id path int true "รหัส Price"
// @Success 200 {object} map[string]interface{} "ลบข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบข้อมูล Price"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/price/delete/{id} [delete]
func (Price) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Price](c)
}

//---------------------------------------------------------------------------------------//

// GetReqlacement godoc
// @Summary ดึงข้อมูล Reqlacement ตาม ID
// @Description ดึงข้อมูล Reqlacement รายการเดียวโดยใช้ ID
// @Tags Reqlacement
// @Accept json
// @Produce json
// @Param id path int true "รหัส Reqlacement"
// @Success 200 {object} models.Reqlacement "ดึงข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบข้อมูล Reqlacement"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/Reqlacement/{id} [get]
func (Reqlacement) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGenericWithKey[Reqlacement](c, ReqlacementKey)
}

// GetAllReqlacement godoc
// @Summary ดึงข้อมูล Reqlacement ทั้งหมด
// @Description ดึงข้อมูล Reqlacement หลายรายการโดยจำกัดจำนวน
// @Tags Reqlacement
// @Accept json
// @Produce json
// @Param limit path string true "จำนวนสูงสุดของข้อมูลที่จะดึง"
// @Success 200 {array} models.Reqlacement "รายการ Reqlacement"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/Reqlacement/all/{limit} [get]
func (Reqlacement) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGenericWithKey[Reqlacement](c, ReqlacementKey)
}

// CreateReqlacement godoc
// @Summary สร้างข้อมูล Reqlacement ใหม่
// @Description เพิ่มข้อมูล Reqlacement ใหม่
// @Tags Reqlacement
// @Accept json
// @Produce json
// @Param Reqlacement body models.Reqlacement true "ข้อมูล Reqlacement"
// @Success 201 {object} models.Reqlacement "สร้างข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "ข้อมูลส่งมาไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/Reqlacement/create [post]
func (Reqlacement) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Reqlacement](c)
}

// UpdateReqlacement godoc
// @Summary อัปเดตข้อมูล Reqlacement ตาม ID
// @Description อัปเดตข้อมูล Reqlacement ที่มีอยู่
// @Tags Reqlacement
// @Accept json
// @Produce json
// @Param id path int true "รหัส Reqlacement"
// @Param Reqlacement body models.Reqlacement true "ข้อมูล Reqlacement ที่ต้องการอัปเดต"
// @Success 200 {object} models.Reqlacement "อัปเดตข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบข้อมูล Reqlacement"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/Reqlacement/update/{id} [put]
func (Reqlacement) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Reqlacement](c)
}

// DeleteReqlacement godoc
// @Summary ลบข้อมูล Reqlacement ตาม ID
// @Description ลบข้อมูล Reqlacement โดยใช้ ID
// @Tags Reqlacement
// @Accept json
// @Produce json
// @Param id path int true "รหัส Reqlacement"
// @Success 200 {object} map[string]interface{} "ลบข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบข้อมูล Reqlacement"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/Reqlacement/delete/{id} [delete]
func (Reqlacement) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Reqlacement](c)
}

//---------------------------------------------------------------------------------------//

// GetRequest godoc
// @Summary ดึงข้อมูล Request ตาม ID
// @Description ดึงข้อมูล Request รายการเดียวโดยใช้ ID
// @Tags Requests
// @Accept json
// @Produce json
// @Param id path int true "รหัส Request"
// @Success 200 {object} models.Requests "ดึงข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบข้อมูล Request"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/requests/{id} [get]
func (Requests) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGenericWithKey[Requests](c, RequestKey)
}

// GetAllRequests godoc
// @Summary ดึงข้อมูล Request ทั้งหมด
// @Description ดึงข้อมูล Request หลายรายการโดยจำกัดจำนวน
// @Tags Requests
// @Accept json
// @Produce json
// @Param limit path string true "จำนวนสูงสุดของข้อมูลที่จะดึง"
// @Success 200 {array} models.Requests "รายการ Request"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/requests/all/{limit} [get]
func (Requests) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGenericWithKey[Requests](c, RequestKey)
}

// CreateRequest godoc
// @Summary สร้างข้อมูล Request ใหม่
// @Description เพิ่มข้อมูล Request ใหม่
// @Tags Requests
// @Accept json
// @Produce json
// @Param request body models.Requests true "ข้อมูล Request"
// @Success 201 {object} models.Requests "สร้างข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "ข้อมูลส่งมาไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/requests/create [post]
func (Requests) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Requests](c)
}

// UpdateRequest godoc
// @Summary อัปเดตข้อมูล Request ตาม ID
// @Description อัปเดตข้อมูล Request ที่มีอยู่
// @Tags Requests
// @Accept json
// @Produce json
// @Param id path int true "รหัส Request"
// @Param request body models.Requests true "ข้อมูล Request ที่ต้องการอัปเดต"
// @Success 200 {object} models.Requests "อัปเดตข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบข้อมูล Request"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/requests/update/{id} [put]
func (Requests) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Requests](c)
}

// DeleteRequest godoc
// @Summary ลบข้อมูล Request ตาม ID
// @Description ลบข้อมูล Request โดยใช้ ID
// @Tags Requests
// @Accept json
// @Produce json
// @Param id path int true "รหัส Request"
// @Success 200 {object} map[string]interface{} "ลบข้อมูลสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบข้อมูล Request"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /home/requests/delete/{id} [delete]
func (Requests) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Requests](c)
}

//---------------------------------------------------------------------------------------//
