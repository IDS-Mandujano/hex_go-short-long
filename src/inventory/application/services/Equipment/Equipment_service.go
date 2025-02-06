package equipments

import "gym-system/src/inventory/domain/repositories/Equipment"

type EquipmentService struct {
	Repo repositories.IEquipamentRepository
}

func (service *EquipmentService) GetAllEquipments() ([]map[string]interface{}, error) {
    return service.Repo.GetAll()
}

func (service *EquipmentService) GetEquipmentById(id int) ([]map[string]interface{}, error) {
    return service.Repo.GetById(id)
}

func (service *EquipmentService) UpdateEquipment(id int, cname string, category string, ccondition string) {
    service.Repo.Update(id, cname, category, ccondition)
}