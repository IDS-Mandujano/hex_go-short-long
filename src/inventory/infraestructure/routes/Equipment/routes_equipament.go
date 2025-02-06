package equipment

import (
	"gym-system/src/inventory/application/EquipamentUsecases"
	equipamentcontrollers "gym-system/src/inventory/infraestructure/controllers/EquipamentControllers"
	"gym-system/src/inventory/infraestructure/database/Equipment"

	"github.com/gin-gonic/gin"
)

func SetupRoutesEquipament(r *gin.Engine){

	dbInstance := equipment.NewMySQLEquipament()

	listEquipamentController := equipamentcontrollers.NewListEquipmentController(*equipamentusecases.NewListEquipment(dbInstance))
	createEquipamentController := equipamentcontrollers.NewCreateEquipamentController(*equipamentusecases.NewCreateEquipament(dbInstance))
	getEquipmentById := equipamentcontrollers.NewEquipmentByIdController(*equipamentusecases.NewEquipmentById(dbInstance))
	updateEquipment := equipamentcontrollers.NewUpdateEquipmentController(*equipamentusecases.NewUpdateEquipment(dbInstance))
	deleteEquipment := equipamentcontrollers.NewDeleteEquipment(*equipamentusecases.NewDeleteEquipment(dbInstance))

	r.GET("/equipments",listEquipamentController.Execute)
	r.POST("/equipments",createEquipamentController.Execute)
	r.GET("/equipments/:id",getEquipmentById.Execute)
	r.PUT("/equipments/:id",updateEquipment.Execute)
	r.DELETE("/equipments/:id",deleteEquipment.Execute)

}