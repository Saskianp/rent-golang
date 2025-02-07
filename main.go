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
	models.MigrateDriver(database.DB)   // Migrate untuk Driver
	models.MigrateBookingType(database.DB)   // Migrate untuk BookingType
	models.MigrateMembership(database.DB)   // Migrate untuk Membership
	models.MigrateDriverIncentive(database.DB)   // Migrate untuk DriverIncentive

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

	// Initialize repository, service, and controller untuk Driver
	repoDriver := repository.NewDriverRepository(database.DB)
	serviceDriver := service.NewDriverService(repoDriver)
	controllerDriver := controller.NewDriverController(serviceDriver)

	// Initialize repository, service, and controller untuk BookingType
	repoBookingType := repository.NewBookingTypeRepository(database.DB)
	serviceBookingType := service.NewBookingTypeService(repoBookingType)
	controllerBookingType := controller.NewBookingTypeController(serviceBookingType)

	// Initialize repository, service, and controller untuk Membership
	repoMembership := repository.NewMembershipRepository(database.DB)
	serviceMembership := service.NewMembershipService(repoMembership)
	controllerMembership := controller.NewMembershipController(serviceMembership)

	// Initialize repository, service, and controller untuk DriverIncentive
	repoDriverIncentive := repository.NewDriverIncentiveRepository(database.DB)
	serviceDriverIncentive := service.NewDriverIncentiveService(repoDriverIncentive)
	controllerDriverIncentive := controller.NewDriverIncentiveController(serviceDriverIncentive)


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

	// Routes untuk Driver
	router.GET("/drivers", controllerDriver.GetAll)
	router.GET("/driver", controllerDriver.GetByID)
	router.POST("/driver/create", controllerDriver.Create)
	router.PUT("/driver/update", controllerDriver.Update)
	router.DELETE("/driver/delete", controllerDriver.Delete)

	// Routes untuk BookingType
	router.GET("/bookingTypes", controllerBookingType.GetAll)
	router.GET("/bookingType", controllerBookingType.GetByID)
	router.POST("/bookingType/create", controllerBookingType.Create)
	router.PUT("/bookingType/update", controllerBookingType.Update)
	router.DELETE("/bookingType/delete", controllerBookingType.Delete)

	// Routes untuk Membership
	router.GET("/memberships", controllerMembership.GetAll)
	router.GET("/membership", controllerMembership.GetByID)
	router.POST("/membership/create", controllerMembership.Create)
	router.PUT("/membership/update", controllerMembership.Update)
	router.DELETE("/membership/delete", controllerMembership.Delete)

	// Routes untuk DriverIncentive
	router.GET("/driverIncentives", controllerDriverIncentive.GetAll)
	router.GET("/driverIncentive", controllerDriverIncentive.GetByID)
	router.POST("/driverIncentive/create", controllerDriverIncentive.Create)
	router.PUT("/driverIncentive/update", controllerDriverIncentive.Update)
	router.DELETE("/driverIncentive/delete", controllerDriverIncentive.Delete)

	// Jalankan server di port 7070
	router.Run(":7070")
}
