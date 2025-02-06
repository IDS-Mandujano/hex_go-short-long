package machineusecases

import "gym-system/src/inventory/domain/repositories/Machine"

type GetMachineById struct {
	db repositories.IMachineRepository
}

func NewMachineById(db repositories.IMachineRepository) *GetMachineById {
	return &GetMachineById{db: db}
}

func (getMachine *GetMachineById) Execute(id int) ([]map[string]interface{},error){
	return getMachine.db.GetById(id)
}