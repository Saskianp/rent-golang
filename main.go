package main

import (
	"testBeGo/database"
	"testBeGo/models"
	"testBeGo/repository"
	"testBeGo/service"
	"testBeGo/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect ke database
	database.ConnectDB()
	models.MigrateCust(database.DB)  // Migrate untuk Customer
	models.MigrateCar(database.DB)   // Migrate untuk Car
	models.MigrateBooking(database.DB)   // Migrate untuk Booking

	// Initialize repository, service, and controller untuk Customer
	repoCust := repository.NewCustomerRepository(database.DB)
	serviceCust := service.NewCustomerService(repoCust)
	controllerCust := controller.NewCustomerController(serviceCust)

	// Initialize repository, service, and controller untuk Car
	repoCar := repository.NewCarRepository(database.DB)
	serviceCar := service.NewCarService(repoCar)
	controllerCar := controller.NewCarController(serviceCar)

	// Initialize repository, service, and controller untuk Booking
	repoBooking := repository.NewBookingRepository(database.DB)
	serviceBooking := service.NewBookingService(repoBooking)
	controllerBooking := controller.NewBookingController(serviceBooking)

	// Inisialisasi router
	router := gin.Default()

	// Routes untuk Customer
	router.GET("/customers", controllerCust.GetAll)
	router.GET("/customer", controllerCust.GetByID)
	router.POST("/customer/create", controllerCust.Create)
	router.PUT("/customer/update", controllerCust.Update)
	router.DELETE("/customer/delete", controllerCust.Delete)

	// Routes untuk Car
	router.GET("/cars", controllerCar.GetAll)
	router.GET("/car", controllerCar.GetByID)
	router.POST("/car/create", controllerCar.Create)
	router.PUT("/car/update", controllerCar.Update)
	router.DELETE("/car/delete", controllerCar.Delete)

	// Routes untuk Booking
	router.GET("/bookings", controllerBooking.GetAllBookings)
	router.GET("/booking/", controllerBooking.GetBookingByID)	
	router.POST("/booking/create", controllerBooking.CreateBooking)
	router.PUT("/booking/update", controllerBooking.UpdateBooking)
	router.DELETE("/booking/delete", controllerBooking.DeleteBooking)

	// Jalankan server di port 7070
	router.Run(":7070")
}
