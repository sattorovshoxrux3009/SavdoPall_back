package main

import (
	"context"
	"fmt"
	"log"

	"GitHub.com/sattorovshoxrux3009/SavdoPall_back/config"
	"GitHub.com/sattorovshoxrux3009/SavdoPall_back/server"
	"GitHub.com/sattorovshoxrux3009/SavdoPall_back/storage"
	"GitHub.com/sattorovshoxrux3009/SavdoPall_back/storage/repo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {
	cfg := config.Load(".")
	// fmt.Println(cfg)

	mysqlUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Mysql.User,     // Foydalanuvchi nomi
		cfg.Mysql.Password, // Parol
		cfg.Mysql.Host,     // Host (masalan, "localhost")
		cfg.Mysql.Port,     // Port (masalan, "3306")
		cfg.Mysql.Database, // Ma'lumotlar bazasi nomi
	)

	// GORM bilan ulanish
	mysqlConn, err := gorm.Open(mysql.Open(mysqlUrl), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// SingularTable: true, // Jadvallarni ko‘plik shaklida yaratmasin
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})
	_ = mysqlConn
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}

	log.Println("Connection success!")
	err = mysqlConn.AutoMigrate(
		&repo.Product{},
	)
	if err != nil {
		log.Fatal("Migrationda xatolik:", err)
	}

	fmt.Println("Migration muvaffaqiyatli yakunlandi!")

	strg := storage.NewStorage(mysqlConn)
	product, err := strg.Product().Create(context.TODO(), &repo.Product{
		ImgUrl:      "https://images.uzum.uz/cp3o4bnfrr80f2gllh0g/original.jpg",
		Name:        "Televizor Roison 32",
		Price:       1454000,
		Height:      55,
		Width:       70,
		Depth:       12,
		Quantity:    3,
		Left:        4,
		Description: "Televizor Roison Smart LED HD TV RE 32-060,43-430 BL, Android 12, ovozli pulti bilan. Dasturiy ta'minot: Netflix, Youtube, Google Play",
	})
	if err != nil {
		fmt.Println("Error creating product:", err)
		return
	}
	fmt.Println(product)
	router := server.NewServer(&server.Options{
		Strg: strg,
	})

	if err := router.Listen(cfg.Port); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
