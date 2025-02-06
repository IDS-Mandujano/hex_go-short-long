package machineusecases

import "gym-system/src/inventory/domain/repositories/Machine"

type CreateMachine struct {
	db repositories.IMachineRepository
}

func NewCreateMachine(db repositories.IMachineRepository) *CreateMachine {
	return &CreateMachine{db: db}
}

func (createMachine *CreateMachine) Execute(cname string, ctype string, cstatus string){
	createMachine.db.Save(cname, ctype, cstatus)
}