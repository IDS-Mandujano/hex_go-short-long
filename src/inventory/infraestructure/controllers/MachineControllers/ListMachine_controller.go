package machinecontrollers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gym-system/src/inventory/application/MachineUsecases"
)


type ListMachineController struct {
	useCase machineusecases.ListMachine
}

func NewListMachineController(useCase machineusecases.ListMachine) *ListMachineController {
	return &ListMachineController{useCase: useCase}
}

func (listMachine_controller *ListMachineController) Execute(g *gin.Context){
	machines, err := listMachine_controller.useCase.Execute()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"Error" : "Error al obtener los maquinas"})
		return
	}

	g.JSON(http.StatusOK, machines)
}