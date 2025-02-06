package equipment

import (
	"gym-system/src/inventory/application/EquipamentUsecases"
	equipamentcontrollers "gym-system/src/inventory/infraestructure/controllers/EquipamentControllers"
	"gym-system/src/inventory/infraestructure/database/Equipment"
	"github.com/gin-gonic/gin"

	service "gym-system/src/inventory/application/services/Equipment"
	http "gym-system/src/inventory/infraestructure/http/Equipment"
)

func SetupRoutesEquipament(r *gin.Engine){

	dbInstance := equipment.NewMySQLEquipament()
	equipmentRepo := dbInstance

	listEquipamentController := equipamentcontrollers.NewListEquipmentController(*equipamentusecases.NewListEquipment(dbInstance))
	createEquipamentController := equipamentcontrollers.NewCreateEquipamentController(*equipamentusecases.NewCreateEquipament(dbInstance))
	getEquipmentById := equipamentcontrollers.NewEquipmentByIdController(*equipamentusecases.NewEquipmentById(dbInstance))
	updateEquipment := equipamentcontrollers.NewUpdateEquipmentController(*equipamentusecases.NewUpdateEquipment(dbInstance))
	deleteEquipment := equipamentcontrollers.NewDeleteEquipment(*equipamentusecases.NewDeleteEquipment(dbInstance))

	equipmentServer := service.EquipmentService{Repo: equipmentRepo}

	r.GET("/equipments/status",http.GetEquipmentStatus(equipmentServer))
	r.GET("/equipments/available",http.WaitForAvailableEquipment(equipmentServer))

	r.GET("/equipments",listEquipamentController.Execute)
	r.POST("/equipments",createEquipamentController.Execute)
	r.GET("/equipments/:id",getEquipmentById.Execute)
	r.PUT("/equipments/:id",updateEquipment.Execute)
	r.DELETE("/equipments/:id",deleteEquipment.Execute)

}