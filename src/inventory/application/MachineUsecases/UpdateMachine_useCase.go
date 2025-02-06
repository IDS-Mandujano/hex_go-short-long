package machineusecases

import "gym-system/src/inventory/domain/repositories/Machine"

type UpdateMachine struct {
	db repositories.IMachineRepository
}

func NewUpdateMachine(db repositories.IMachineRepository) *UpdateMachine {
	return &UpdateMachine{db: db}
}

func (updateMachine *UpdateMachine) Execute (id int,cname string, ctype string, cstatus string){
	updateMachine.db.Update(id,cname,ctype,cstatus)
}