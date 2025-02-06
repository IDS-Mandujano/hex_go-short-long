package machineusecases

import "gym-system/src/inventory/domain/repositories/Machine"

type ListMachine struct {
	db repositories.IMachineRepository
}

func NewListMachine (db repositories.IMachineRepository) *ListMachine {
	return &ListMachine{db: db}
}

func (listMachine *ListMachine) Execute () ([]map[string]interface{}, error){
	return listMachine.db.GetAll()
}