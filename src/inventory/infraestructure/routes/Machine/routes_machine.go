package machine

import (

	"gym-system/src/inventory/application/MachineUsecases"
	machinecontrollers "gym-system/src/inventory/infraestructure/controllers/MachineControllers"
	"gym-system/src/inventory/infraestructure/database/Machine"

	"github.com/gin-gonic/gin"
)

func SetupRoutesMachine(r *gin.Engine){

	dbInstance := machine.NewMySQLMachine()

	listMachineController := machinecontrollers.NewListMachineController(*machineusecases.NewListMachine(dbInstance))
	createMachineController := machinecontrollers.NewCreateMachineController(*machineusecases.NewCreateMachine(dbInstance))
	getMachineById := machinecontrollers.NewMachineByIdController(*machineusecases.NewMachineById(dbInstance))
	updateMachine := machinecontrollers.NewUpdateMachineController(*machineusecases.NewUpdateMachine(dbInstance))
	deleteMachine := machinecontrollers.NewDeleteMachine(*machineusecases.NewDeleteMachine(dbInstance))

	r.GET("/machines",listMachineController.Execute)
	r.POST("/machines",createMachineController.Execute)
	r.GET("/machines/:id",getMachineById.Execute)
	r.PUT("/machines/:id",updateMachine.Execute)
	r.DELETE("/machines/:id",deleteMachine.Execute)

}