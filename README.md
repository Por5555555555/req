# นี่คือ project req ส่วน api
นี่เป็น project เรียกคำขอร้อง นี่เป็นส่วน backEnd หรือ api

## มีอะไรในนั้นบ้าง
  - มีระบบ login
  - มีระบบ router ทั้้งหมดที่เราค่าดว่าควรมี
  - สามารถเชื่อม DataBase
  - Suport MySql and  Sqlite (Sqlite Suport for linux/max)
  - #### [Dont click me](https://www.youtube.com/watch?v=dQw4w9WgXcQ) 

## โครงสร้าง file 
  - addOn  => **ส่วนเสริม**
     - colortext => **นี่เป็นส่วนเสริม show color**
  - config
    - config
  
  - filber
    - fiberfunc
        - fiberfuncConfig
          - fiberConfig.go  => **ส่วนเสริมืี่ใช้ใน config fiberfunc**
          - fiberFuncAddor.go => **ส่วนเสริมืี่ใช้ใน fiberfunc**
        
        - fiberfuncCOUD
          - fiberfuncAddOn
            - fiberfuncBlock.go => **ใช้ตรวจสอบมาเข้าไปฝช้ได้ไหม**
            - fiberfuncUserAddOn.go => **นี่ใช้สร้าง User,login,logout และ cookie**
          
          - fiberCreate.go => **ใช้ในการ สร้าง ข้อมูลใหม่**
          - fiberDetele.go => **ใช้ในการ ลบ ข้อมูล**
          - fiberGet.go => **ใช้ในการ ขอ ข้อมูลใหม่**
          - fiberUpdate.go => **ใช้ในการ อัปเดต ข้อมูลใหม่**

    - fiberfuncFile
       - fiberFuncDeleteFile.go => **func เก่าลบไฟล์ ไม่ได้ถูกใช้จริง**
       - fiberFuncFileUpdate.go => **func เก่าอัปเดตไฟล์ ไม่ได้ถูกใช้จริง**
       - fiberFuncFileUpload.go => **func เก่าอัปโหลดไฟล์ ไม่ได้ถูกใช้จริง**

    - router => **เป็นตัวเก็บ func rounter**
      - config.go => **ใช้ใน router เท่านั้น**
      - routeKeys.go => **เป็นที่เก็บ func rounter ที่มีการเชื่อม หลาย Table** 
      - rountes.go => **เป็นที่เก็บ func rounter**

    - openserver.go => **ใช้เป็ด server เเละบอกตำเเหล่ง rounter** 

  - gorm
    - database
      - testdb
        - dbreq.db <h3> => ***ใช้ทดสอบไฟล์ database เป็น SQLite suport linux/max ถ้า windown ต้องโหลดส่วนเสริม***</h3>

      - DataBaseConnect.go => **เป็นจุดเปิดไฟล์ Sql**
      - dbtest.go => **ใช้สร้าง Table ไหม่**
      - install.go => **ใช้สร้าง User ใหท่**
      
    - handler
      - COUD.go => **เป็นจุด เก็บคำสั้งที่ เชือมกับ database พื้นฐาน** 

    - models
      - struct.go => **file models**
  
  - go.mod 
  - go.sum
  - main.go
  - README.md