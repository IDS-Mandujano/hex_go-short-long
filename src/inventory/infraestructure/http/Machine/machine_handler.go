package machine

import (

	"gym-system/src/inventory/application/services/Machine"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetMachineStatus(service machine.MachineService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		machines, err := service.GetAllMachines()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
			return
		}

		ctx.JSON(http.StatusOK,machines)
	}
}

func WaitForAvailableMachine(service machine.MachineService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		timeout := time.After(30 * time.Second)
		ticker := time.NewTicker(3 * time.Second)

		for {
			select {
			case <-timeout : 
				ctx.JSON(http.StatusRequestTimeout,gin.H{"message":"No hay cambios en las maquinas"})
				return
			case <-ticker.C :
				machines,_ := service.GetAllMachines()
				if len(machines) > 0 {
					ctx.JSON(http.StatusOK,machines)
					return
				}
			}
		}
	}
}