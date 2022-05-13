package template

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kamil5b/backend-template/database"
	"github.com/kamil5b/backend-template/models"
	"github.com/kamil5b/backend-template/routes"
)

func SetupTemplate(server_url, db_url, user, password, protocol, db string) {

	database.Connect(
		db_url, user, password, protocol, db,
		&models.User{},
		&models.Barang{},
		&models.Customer{},
		&models.Suplier{},
		&models.FakturPembelian{},
		&models.FakturPenjualan{},
		&models.Hutang{},
		&models.Pengajuan{},
		&models.Piutang{},
		&models.Hutang{},
		&models.Pengajuan{},
		&models.Sales{},
		&models.Stock{},
		&models.Penjualan{},
		&models.Pembelian{},
		&models.KasKecil{},
		&models.Giro{},
		&models.Bank{},
	)
	app := fiber.New()

	//origin := utilities.GoDotEnvVariable("VIEW_URL") //ganti view url ini di .env
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		//AllowOrigins:     []string{origin},
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	routes.Setup(app)

	err := app.Listen(server_url)
	if err != nil {
		fmt.Println(err)
		fmt.Scan(&err)
	}
}
