package machine

import (
	"gym-system/src/inventory/application/MachineUsecases"
	machinecontrollers "gym-system/src/inventory/infraestructure/controllers/MachineControllers"
	"gym-system/src/inventory/infraestructure/database/Machine"
	"github.com/gin-gonic/gin"

	service "gym-system/src/inventory/application/services/Machine"
	http "gym-system/src/inventory/infraestructure/http/Machine"
)

func SetupRoutesMachine(r *gin.Engine){

	dbInstance := machine.NewMySQLMachine()
	machineRepo := dbInstance

	listMachineController := machinecontrollers.NewListMachineController(*machineusecases.NewListMachine(dbInstance))
	createMachineController := machinecontrollers.NewCreateMachineController(*machineusecases.NewCreateMachine(dbInstance))
	getMachineById := machinecontrollers.NewMachineByIdController(*machineusecases.NewMachineById(dbInstance))
	updateMachine := machinecontrollers.NewUpdateMachineController(*machineusecases.NewUpdateMachine(dbInstance))
	deleteMachine := machinecontrollers.NewDeleteMachine(*machineusecases.NewDeleteMachine(dbInstance))

	machineService := service.MachineService{Repo: machineRepo}

	r.GET("/machines/status",http.GetMachineStatus(machineService))
	r.GET("/machines/available",http.WaitForAvailableMachine(machineService))

	r.GET("/machines",listMachineController.Execute)
	r.POST("/machines",createMachineController.Execute)
	r.GET("/machines/:id",getMachineById.Execute)
	r.PUT("/machines/:id",updateMachine.Execute)
	r.DELETE("/machines/:id",deleteMachine.Execute)

}