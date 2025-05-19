package routes

import (
	"bre-api/fiber/fiberfunc/fiberfuncCOUD"
	fiberfuncCOUDaddOn "bre-api/fiber/fiberfunc/fiberfuncCOUD/fiberfuncAddOn"

	"github.com/gofiber/fiber/v2"
)

// GetAgency godoc
// @Summary ดึงข้อมูลหน่วยงานตามรหัส
// @Description ใช้สำหรับดึงข้อมูลของหน่วยงานจากรหัสที่ระบุ
// @Tags Agency
// @Accept  json
// @Produce  json
// @Param id path int true "รหัสของหน่วยงานที่ต้องการดึง"
// @Success 200 {object} []models.Agency "สำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง หรือเกิดข้อผิดพลาดจาก SQL"
// @Failure 401 {object} map[string]interface{} "ไม่ได้รับอนุญาต"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/agency/{id} [get]
func (Agency) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGeneric[Agency](c)
}

// GetAllAgency godoc
// @Summary ดึงข้อมูลหน่วยงานทั้งหมด
// @Description ดึงข้อมูลหน่วยงานทั้งหมด โดยจำกัดจำนวนตามที่กำหนด
// @Tags Agency
// @Accept json
// @Produce json
// @Param limit path string true "จำนวนรายการสูงสุดที่ต้องการดึง"
// @Success 200 {array} []models.Agency "สำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/agency/all/{limit} [get]
func (Agency) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGeneric[Agency](c)
}

// CreateAgency godoc
// @Summary สร้างหน่วยงานใหม่
// @Description เพิ่มข้อมูลหน่วยงานใหม่เข้าสู่ระบบ
// @Tags Agency
// @Accept json
// @Produce json
// @Param agency body models.Agency true "ข้อมูลของหน่วยงานที่ต้องการเพิ่ม"
// @Success 200 {object} models.Agency "สร้างสำเร็จ"
// @Failure 400 {object} map[string]interface{} "ข้อมูลไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/agency/create [post]
func (Agency) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Agency](c)
}

// UpdateAgency godoc
// @Summary แก้ไขข้อมูลหน่วยงาน
// @Description ใช้สำหรับแก้ไขข้อมูลของหน่วยงานที่มีอยู่ ตามรหัสที่ระบุ
// @Tags Agency
// @Accept json
// @Produce json
// @Param id path int true "รหัสของหน่วยงานที่ต้องการแก้ไข"
// @Param agency body models.Agency true "ข้อมูลหน่วยงานที่แก้ไขแล้ว"
// @Success 200 {object} models.Agency "อัปเดตสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบหน่วยงาน"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/agency/update/{id} [put]
func (Agency) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Agency](c)
}

// DeleteAgency godoc
// @Summary ลบข้อมูลหน่วยงาน
// @Description ใช้สำหรับลบข้อมูลหน่วยงานตามรหัสที่ระบุ
// @Tags Agency
// @Accept json
// @Produce json
// @Param id path int true "รหัสของหน่วยงานที่ต้องการลบ"
// @Success 200 {object} map[string]interface{} "ลบสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบหน่วยงาน"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/agency/delete/{id} [delete]
func (Agency) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Agency](c)
}

//---------------------------------------------------------------------------------------//

// GetAuditor godoc
// @Summary ดึงข้อมูลผู้ตรวจสอบตามรหัส
// @Description ใช้สำหรับดึงข้อมูลของผู้ตรวจสอบจากรหัสที่ระบุ
// @Tags Auditor
// @Accept json
// @Produce json
// @Param id path int true "รหัสของผู้ตรวจสอบที่ต้องการดึง"
// @Success 200 {object} models.Auditor "สำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง หรือเกิดข้อผิดพลาดจาก SQL"
// @Failure 401 {object} map[string]interface{} "ไม่ได้รับอนุญาต"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/auditor/{id} [get]
func (Auditor) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGeneric[Auditor](c)
}

// GetAllAuditor godoc
// @Summary ดึงข้อมูลผู้ตรวจสอบทั้งหมด
// @Description ดึงข้อมูลผู้ตรวจสอบทั้งหมด โดยจำกัดจำนวนตามที่กำหนด
// @Tags Auditor
// @Accept json
// @Produce json
// @Param limit path string true "จำนวนรายการสูงสุดที่ต้องการดึง"
// @Success 200 {array} []models.Auditor "สำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/auditor/all/{limit} [get]
func (Auditor) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGeneric[Auditor](c)
}

// CreateAuditor godoc
// @Summary สร้างผู้ตรวจสอบใหม่
// @Description เพิ่มข้อมูลผู้ตรวจสอบใหม่เข้าสู่ระบบ
// @Tags Auditor
// @Accept json
// @Produce json
// @Param auditor body models.Auditor true "ข้อมูลของผู้ตรวจสอบที่ต้องการเพิ่ม"
// @Success 200 {object} models.Auditor "สร้างสำเร็จ"
// @Failure 400 {object} map[string]interface{} "ข้อมูลไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/auditor/create [post]
func (Auditor) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Auditor](c)
}

// UpdateAuditor godoc
// @Summary แก้ไขข้อมูลผู้ตรวจสอบ
// @Description ใช้สำหรับแก้ไขข้อมูลของผู้ตรวจสอบที่มีอยู่ ตามรหัสที่ระบุ
// @Tags Auditor
// @Accept json
// @Produce json
// @Param id path int true "รหัสของผู้ตรวจสอบที่ต้องการแก้ไข"
// @Param auditor body models.Auditor true "ข้อมูลผู้ตรวจสอบที่แก้ไขแล้ว"
// @Success 200 {object} models.Auditor "อัปเดตสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบผู้ตรวจสอบ"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/auditor/update/{id} [put]
func (Auditor) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Auditor](c)
}

// DeleteAuditor godoc
// @Summary ลบข้อมูลผู้ตรวจสอบ
// @Description ใช้สำหรับลบข้อมูลผู้ตรวจสอบตามรหัสที่ระบุ
// @Tags Auditor
// @Accept json
// @Produce json
// @Param id path int true "รหัสของผู้ตรวจสอบที่ต้องการลบ"
// @Success 200 {object} map[string]interface{} "ลบสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบผู้ตรวจสอบ"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/auditor/delete/{id} [delete]
func (Auditor) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Auditor](c)
}

//---------------------------------------------------------------------------------------//

// GetCategory godoc
// @Summary ดึงข้อมูลหมวดหมู่ตามรหัส
// @Description ใช้สำหรับดึงข้อมูลหมวดหมู่ตามรหัสที่ระบุ
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "รหัสของหมวดหมู่ที่ต้องการดึง"
// @Success 200 {object} models.Category "สำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง หรือเกิดข้อผิดพลาดจาก SQL"
// @Failure 401 {object} map[string]interface{} "ไม่ได้รับอนุญาต"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/category/{id} [get]
func (Category) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGeneric[Category](c)
}

// GetAllCategory godoc
// @Summary ดึงข้อมูลหมวดหมู่ทั้งหมด
// @Description ดึงข้อมูลหมวดหมู่ทั้งหมด โดยจำกัดจำนวนตามที่กำหนด
// @Tags Category
// @Accept json
// @Produce json
// @Param limit path string true "จำนวนรายการสูงสุดที่ต้องการดึง"
// @Success 200 {array} []models.Category "สำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/category/all/{limit} [get]
func (Category) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGeneric[Category](c)
}

// CreateCategory godoc
// @Summary สร้างหมวดหมู่ใหม่
// @Description เพิ่มข้อมูลหมวดหมู่ใหม่เข้าสู่ระบบ
// @Tags Category
// @Accept json
// @Produce json
// @Param category body models.Category true "ข้อมูลของหมวดหมู่ที่ต้องการเพิ่ม"
// @Success 200 {object} models.Category "สร้างสำเร็จ"
// @Failure 400 {object} map[string]interface{} "ข้อมูลไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/category/create [post]
func (Category) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Category](c)
}

// UpdateCategory godoc
// @Summary แก้ไขข้อมูลหมวดหมู่
// @Description ใช้สำหรับแก้ไขข้อมูลหมวดหมู่ที่มีอยู่ ตามรหัสที่ระบุ
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "รหัสของหมวดหมู่ที่ต้องการแก้ไข"
// @Param category body models.Category true "ข้อมูลหมวดหมู่ที่แก้ไขแล้ว"
// @Success 200 {object} models.Category "อัปเดตสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบหมวดหมู่"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/category/update/{id} [put]
func (Category) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Category](c)
}

// UpdateCategory godoc
// @Summary แก้ไขข้อมูลหมวดหมู่
// @Description ใช้สำหรับแก้ไขข้อมูลหมวดหมู่ที่มีอยู่ ตามรหัสที่ระบุ
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "รหัสของหมวดหมู่ที่ต้องการแก้ไข"
// @Param category body models.Category true "ข้อมูลหมวดหมู่ที่แก้ไขแล้ว"
// @Success 200 {object} models.Category "อัปเดตสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบหมวดหมู่"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/category/update/{id} [put]
func (Category) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Category](c)
}

//---------------------------------------------------------------------------------------//

// GetMoney godoc
// @Summary ดึงข้อมูลเงินตามรหัส
// @Description ใช้สำหรับดึงข้อมูลเงินตามรหัสที่ระบุ
// @Tags Money
// @Accept json
// @Produce json
// @Param id path int true "รหัสของข้อมูลเงินที่ต้องการดึง"
// @Success 200 {object} models.Money "สำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง หรือเกิดข้อผิดพลาดจาก SQL"
// @Failure 401 {object} map[string]interface{} "ไม่ได้รับอนุญาต"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/money/{id} [get]
func (Money) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGeneric[Money](c)
}

// GetAllMoney godoc
// @Summary ดึงข้อมูลเงินทั้งหมด
// @Description ดึงข้อมูลเงินทั้งหมด โดยจำกัดจำนวนตามที่กำหนด
// @Tags Money
// @Accept json
// @Produce json
// @Param limit path string true "จำนวนรายการสูงสุดที่ต้องการดึง"
// @Success 200 {array} []models.Money "สำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/money/all/{limit} [get]
func (Money) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGeneric[Money](c)
}

// CreateMoney godoc
// @Summary สร้างข้อมูลเงินใหม่
// @Description เพิ่มข้อมูลเงินใหม่เข้าสู่ระบบ
// @Tags Money
// @Accept json
// @Produce json
// @Param money body models.Money true "ข้อมูลเงินที่ต้องการเพิ่ม"
// @Success 200 {object} models.Money "สร้างสำเร็จ"
// @Failure 400 {object} map[string]interface{} "ข้อมูลไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/money/create [post]
func (Money) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Money](c)
}

// UpdateMoney godoc
// @Summary แก้ไขข้อมูลเงิน
// @Description ใช้สำหรับแก้ไขข้อมูลเงินที่มีอยู่ ตามรหัสที่ระบุ
// @Tags Money
// @Accept json
// @Produce json
// @Param id path int true "รหัสของข้อมูลเงินที่ต้องการแก้ไข"
// @Param money body models.Money true "ข้อมูลเงินที่แก้ไขแล้ว"
// @Success 200 {object} models.Money "อัปเดตสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบข้อมูลเงิน"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/money/update/{id} [put]
func (Money) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Money](c)
}

// DeleteMoney godoc
// @Summary ลบข้อมูลเงิน
// @Description ใช้สำหรับลบข้อมูลเงินตามรหัสที่ระบุ
// @Tags Money
// @Accept json
// @Produce json
// @Param id path int true "รหัสของข้อมูลเงินที่ต้องการลบ"
// @Success 200 {object} map[string]interface{} "ลบสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบข้อมูลเงิน"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/money/delete/{id} [delete]
func (Money) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Money](c)
}

//---------------------------------------------------------------------------------------//

// GetProvince godoc
// @Summary ดึงข้อมูลจังหวัดตามรหัส
// @Description ใช้สำหรับดึงข้อมูลจังหวัดตามรหัสที่ระบุ
// @Tags Province
// @Accept json
// @Produce json
// @Param id path int true "รหัสจังหวัดที่ต้องการดึงข้อมูล"
// @Success 200 {object} models.Province "สำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง หรือเกิดข้อผิดพลาดจาก SQL"
// @Failure 401 {object} map[string]interface{} "ไม่ได้รับอนุญาต"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/province/{id} [get]
func (Province) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGeneric[Province](c)
}

// GetAllProvince godoc
// @Summary ดึงข้อมูลจังหวัดทั้งหมด
// @Description ดึงข้อมูลจังหวัดทั้งหมด โดยจำกัดจำนวนตามที่กำหนด
// @Tags Province
// @Accept json
// @Produce json
// @Param limit path string true "จำนวนรายการสูงสุดที่ต้องการดึง"
// @Success 200 {array} []models.Province "สำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/province/all/{limit} [get]
func (Province) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGeneric[Province](c)
}

// CreateProvince godoc
// @Summary สร้างข้อมูลจังหวัดใหม่
// @Description เพิ่มข้อมูลจังหวัดใหม่เข้าสู่ระบบ
// @Tags Province
// @Accept json
// @Produce json
// @Param province body models.Province true "ข้อมูลจังหวัดที่ต้องการเพิ่ม"
// @Success 200 {object} models.Province "สร้างสำเร็จ"
// @Failure 400 {object} map[string]interface{} "ข้อมูลไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/province/create [post]
func (Province) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Province](c)
}

// UpdateProvince godoc
// @Summary แก้ไขข้อมูลจังหวัด
// @Description ใช้สำหรับแก้ไขข้อมูลจังหวัดที่มีอยู่ ตามรหัสที่ระบุ
// @Tags Province
// @Accept json
// @Produce json
// @Param id path int true "รหัสจังหวัดที่ต้องการแก้ไข"
// @Param province body models.Province true "ข้อมูลจังหวัดที่แก้ไขแล้ว"
// @Success 200 {object} models.Province "อัปเดตสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบข้อมูลจังหวัด"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/province/update/{id} [put]
func (Province) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Province](c)
}

// DeleteProvince godoc
// @Summary ลบข้อมูลจังหวัด
// @Description ใช้สำหรับลบข้อมูลจังหวัดตามรหัสที่ระบุ
// @Tags Province
// @Accept json
// @Produce json
// @Param id path int true "รหัสจังหวัดที่ต้องการลบ"
// @Success 200 {object} map[string]interface{} "ลบสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบข้อมูลจังหวัด"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/province/delete/{id} [delete]
func (Province) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Province](c)
}

//---------------------------------------------------------------------------------------//

// GetSource godoc
// @Summary ดึงข้อมูลแหล่งที่มาตามรหัส
// @Description ใช้สำหรับดึงข้อมูลแหล่งที่มาตามรหัสที่ระบุ
// @Tags Source
// @Accept json
// @Produce json
// @Param id path int true "รหัสแหล่งที่มาต้องการดึงข้อมูล"
// @Success 200 {object} models.Source "สำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง หรือเกิดข้อผิดพลาดจาก SQL"
// @Failure 401 {object} map[string]interface{} "ไม่ได้รับอนุญาต"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/source/{id} [get]
func (Source) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGeneric[Source](c)
}

// GetAllSource godoc
// @Summary ดึงข้อมูลแหล่งที่มาทั้งหมด
// @Description ดึงข้อมูลแหล่งที่มาทั้งหมด โดยจำกัดจำนวนตามที่กำหนด
// @Tags Source
// @Accept json
// @Produce json
// @Param limit path string true "จำนวนสูงสุดของรายการที่ต้องการดึง"
// @Success 200 {array} []models.Source "สำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/source/all/{limit} [get]
func (Source) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGeneric[Source](c)
}

// CreateSource godoc
// @Summary สร้างข้อมูลแหล่งที่มาใหม่
// @Description เพิ่มข้อมูลแหล่งที่มาใหม่เข้าสู่ระบบ
// @Tags Source
// @Accept json
// @Produce json
// @Param source body models.Source true "ข้อมูลแหล่งที่มาที่ต้องการเพิ่ม"
// @Success 200 {object} models.Source "สร้างสำเร็จ"
// @Failure 400 {object} map[string]interface{} "ข้อมูลไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/source/create [post]
func (Source) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Source](c)
}

// UpdateSource godoc
// @Summary แก้ไขข้อมูลแหล่งที่มา
// @Description ใช้สำหรับแก้ไขข้อมูลแหล่งที่มาที่มีอยู่ ตามรหัสที่ระบุ
// @Tags Source
// @Accept json
// @Produce json
// @Param id path int true "รหัสแหล่งที่มาที่ต้องการแก้ไข"
// @Param source body models.Source true "ข้อมูลแหล่งที่มาที่แก้ไขแล้ว"
// @Success 200 {object} models.Source "อัปเดตสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบข้อมูลแหล่งที่มา"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/source/update/{id} [put]
func (Source) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Source](c)
}

// DeleteSource godoc
// @Summary ลบข้อมูลแหล่งที่มา
// @Description ใช้สำหรับลบข้อมูลแหล่งที่มาตามรหัสที่ระบุ
// @Tags Source
// @Accept json
// @Produce json
// @Param id path int true "รหัสแหล่งที่มาที่ต้องการลบ"
// @Success 200 {object} map[string]interface{} "ลบสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบข้อมูลแหล่งที่มา"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/source/delete/{id} [delete]
func (Source) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Source](c)
}

//---------------------------------------------------------------------------------------//

// GetUnit godoc
// @Summary ดึงข้อมูลหน่วยตามรหัส
// @Description ใช้สำหรับดึงข้อมูลหน่วยตามรหัสที่ระบุ
// @Tags Unit
// @Accept json
// @Produce json
// @Param id path int true "รหัสหน่วยที่ต้องการดึงข้อมูล"
// @Success 200 {object} models.Unit "สำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง หรือเกิดข้อผิดพลาดจาก SQL"
// @Failure 401 {object} map[string]interface{} "ไม่ได้รับอนุญาต"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/unit/{id} [get]/
func (Unit) Get(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetGeneric[Unit](c)
}

// GetAllUnit godoc
// @Summary ดึงข้อมูลหน่วยทั้งหมด
// @Description ดึงข้อมูลหน่วยทั้งหมดโดยจำกัดจำนวนตามที่กำหนด
// @Tags Unit
// @Accept json
// @Produce json
// @Param limit path string true "จำนวนสูงสุดของรายการที่ต้องการดึง"
// @Success 200 {array} []models.Unit "สำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/unit/all/{limit} [get]
func (Unit) GetAll(c *fiber.Ctx) error {
	return fiberfuncCOUD.GetAllGeneric[Unit](c)
}

// CreateUnit godoc
// @Summary สร้างข้อมูลหน่วยใหม่
// @Description เพิ่มข้อมูลหน่วยใหม่เข้าสู่ระบบ
// @Tags Unit
// @Accept json
// @Produce json
// @Param unit body models.Unit true "ข้อมูลหน่วยที่ต้องการเพิ่ม"
// @Success 200 {object} models.Unit "สร้างสำเร็จ"
// @Failure 400 {object} map[string]interface{} "ข้อมูลไม่ถูกต้อง"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/unit/create [post]
// @Router /home/unit/create [post]
func (Unit) Create(c *fiber.Ctx) error {
	return fiberfuncCOUD.CreateGeneric[Unit](c)
}

// UpdateUnit godoc
// @Summary Update a unit
// @Description Update an existing unit by ID
// @Tags Unit
// @Accept json
// @Produce json
// @Param id path int true "Unit ID"
// @Param unit body models.Unit true "Updated unit data"
// @Success 200 {object} models.Unit "Successfully updated"
// @Failure 400 {object} map[string]interface{} "Invalid parameters"
// @Failure 404 {object} map[string]interface{} "Unit not found"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /home/unit/update/{id} [put]
func (Unit) Update(c *fiber.Ctx) error {
	return fiberfuncCOUD.UpdateGeneric[Unit](c)
}

// DeleteUnit godoc
// @Summary ลบข้อมูลหน่วย
// @Description ใช้สำหรับลบข้อมูลหน่วยตามรหัสที่ระบุ
// @Tags Unit
// @Accept json
// @Produce json
// @Param id path int true "รหัสหน่วยที่ต้องการลบ"
// @Success 200 {object} map[string]interface{} "ลบสำเร็จ"
// @Failure 400 {object} map[string]interface{} "พารามิเตอร์ไม่ถูกต้อง"
// @Failure 404 {object} map[string]interface{} "ไม่พบข้อมูลหน่วย"
// @Failure 500 {object} map[string]interface{} "ข้อผิดพลาดจากเซิร์ฟเวอร์"
// @Router /home/unit/delete/{id} [delete]
func (Unit) Delete(c *fiber.Ctx) error {
	return fiberfuncCOUD.DeteleGeneric[Unit](c)
}

//---------------------------------------------------------------------------------------//

func (User) Get(c *fiber.Ctx) error {
	return c.JSON("this func dont support")
}
func (User) GetAll(c *fiber.Ctx) error {
	return c.JSON("this func dont support")
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a user by providing user details
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body models.User true "User data"
// @Success 200 {object} models.User "Created"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /create [post]
func (User) Create(c *fiber.Ctx) error {
	return fiberfuncCOUDaddOn.CreateUser(c)
}

// LoginUser gpdoc
// @Summary Login User For jwt
// @Description login User
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 200 {object} models.User "Created"
// @Router /login [post]
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

// LoginUser gpdoc
// @Summary Logout User
// @Description logout User
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "logout success"
// @Router /logout [post]
func Logout(c *fiber.Ctx) error {
	return fiberfuncCOUDaddOn.Logout(c)
}

// CheckJwt gpdoc
// @Summary ตรวจสอบ JWT
// @Description ใช้ดูว่า ่ jwt ยังใช้ได้หรือไม่
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Check success"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /check [get]
func CheckJwt(c *fiber.Ctx) error {
	return fiberfuncCOUDaddOn.CheckJwt(c)
}
