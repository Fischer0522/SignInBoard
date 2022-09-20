package models

func GetAllTrashData() (*[]TrashData,error) {
	var trashData []TrashData
	result:=db.Find(&trashData)
	err := result.Error
	if err != nil {
		return nil,err
	}
	return &trashData,nil
}

func DeleteTrash(id string) error  {

	result:=db.Delete(&TrashData{},id)

	if err := result.Error; err != nil {
		return err
	}
	return nil

}
func GetTrashById(id string) (*TrashData,error) {
	var trashData TrashData
	result := db.Find(&trashData,id)
	err := result.Error
	if err != nil {
		return nil,err
	}
	return &trashData,err
}

func CreateTrash() error {
	trashData:= TrashData{
		Name: "这是一条垃圾信息",
	}
	result := db.Create(&trashData)
	err := result.Error
	return err
}
