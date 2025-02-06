package equipamentusecases

import "gym-system/src/inventory/domain/repositories/Equipment"

type CreateEquipament struct {
	db repositories.IEquipamentRepository
}

func NewCreateEquipament(db repositories.IEquipamentRepository) *CreateEquipament {
	return &CreateEquipament{db: db}
}

func (ce *CreateEquipament) Execute(cname string, category string, ccondition string){
	ce.db.Save(cname, category, ccondition)
}