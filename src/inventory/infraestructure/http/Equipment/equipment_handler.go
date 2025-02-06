package equipment

import (

	"gym-system/src/inventory/application/services/Equipment"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetEquipmentStatus(service equipments.EquipmentService) gin.HandlerFunc {
    return func(c *gin.Context) {
        equipments, err := service.GetAllEquipments()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, equipments)
    }
}

func WaitForAvailableEquipment(service equipments.EquipmentService) gin.HandlerFunc {
    return func(c *gin.Context) {
        timeout := time.After(30 * time.Second)
        ticker := time.NewTicker(3 * time.Second)

        for {
            select {
            case <-timeout:
                c.JSON(http.StatusRequestTimeout, gin.H{"message": "No hay cambios en los equipos"})
                return
            case <-ticker.C:
                equipments, _ := service.GetAllEquipments()
                if len(equipments) > 0 {
                    c.JSON(http.StatusOK, equipments)
                    return
                }
            }
        }
    }
}