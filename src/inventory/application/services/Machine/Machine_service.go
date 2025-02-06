package machine

import (
	"gym-system/src/inventory/domain/repositories/Machine"
)

type MachineService struct {
	Repo repositories.IMachineRepository
}

func (service *MachineService) GetAllMachines() ([]map[string]interface{}, error){
	return service.Repo.GetAll()
}

func (service *MachineService) GetMachineById(id int) ([]map[string]interface{},error){
	return service.Repo.GetById(id)
}

func (service *MachineService) UpdateMachine(id int, canme string, ctype string, cstatus string) {
	service.Repo.Update(id,canme,ctype,cstatus)
}