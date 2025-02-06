package equipamentusecases

import "gym-system/src/inventory/domain/repositories/Equipment"

type UpdateEquipment struct {
	db repositories.IEquipamentRepository
}

func NewUpdateEquipment(db repositories.IEquipamentRepository) *UpdateEquipment {
	return &UpdateEquipment{db: db}
}

func (updateEquipment *UpdateEquipment) Execute (id int,cname string, category string, ccondition string){
	updateEquipment.db.Update(id,cname,category,ccondition)
}