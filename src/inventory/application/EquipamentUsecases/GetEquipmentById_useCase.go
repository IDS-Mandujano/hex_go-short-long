package equipamentusecases

import "gym-system/src/inventory/domain/repositories/Equipment"

type GetEquipmentById struct {
	db repositories.IEquipamentRepository
}

func NewEquipmentById(db repositories.IEquipamentRepository) *GetEquipmentById {
	return &GetEquipmentById{db: db}
}

func (getEquipment *GetEquipmentById) Execute(id int) ([]map[string]interface{},error){
	return getEquipment.db.GetById(id)
}