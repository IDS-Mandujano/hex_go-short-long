package equipamentusecases

import "gym-system/src/inventory/domain/repositories/Equipment"

type ListEquipament struct {
	db repositories.IEquipamentRepository
}

func NewListEquipment (db repositories.IEquipamentRepository) *ListEquipament {
	return &ListEquipament{db: db}
}

func (listEquipment *ListEquipament) Execute () ([]map[string]interface{}, error){
	return listEquipment.db.GetAll()
}