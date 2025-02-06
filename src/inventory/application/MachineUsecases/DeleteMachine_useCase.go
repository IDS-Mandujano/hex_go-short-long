package machineusecases

import "gym-system/src/inventory/domain/repositories/Machine"

type DeleteMachine struct {
	db repositories.IMachineRepository
}

func NewDeleteMachine (db repositories.IMachineRepository) *DeleteMachine {
	return &DeleteMachine{db: db}
}

func (deleteMachine *DeleteMachine) Execute(id int){
	deleteMachine.db.Delete(id)
}