package repositories

type IMachineRepository interface {

	

	Save(cname string, ctype string, cstatus string)
	GetAll() ([]map[string]interface{}, error)
	GetById(id int) ([]map[string]interface{}, error)
	Update(id int, cname string, ctype string, cstatus string)
	Delete(id int)
}