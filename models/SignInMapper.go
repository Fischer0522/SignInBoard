package models

import (
	"SignInBoard/data"
	"time"
)

func CreateSignIn(signInParam data.SignInParam) (*SignIn,error) {

	sign := SignIn{
		Name: signInParam.Name,
		Message: signInParam.Message,
		SignInTime:time.Now(),
	}

	result := db.Create(&sign)
	err:= result.Error
	if err != nil {
		return nil,err
	}
	return &sign,nil

}

func GetSignIn(id string) (*SignIn,error) {


	var signIn SignIn

	result := db.First(&signIn,id)

	err := result.Error
	if err != nil {
		return nil,err
	}
	return &signIn,nil

}

func GetSignByName(name string) (*SignIn,error) {
	var signIn SignIn
	result := db.Where("name = ?",name).First(&signIn)
	err := result.Error
	if err != nil {
		return nil,err
	}
	return &signIn,nil

}
func GetAllSign() (*[]SignIn,error) {
	var currentSign []SignIn
	result:=db.Find(&currentSign)
	err := result.Error
	if err != nil {
		return nil,err
	}
	return &currentSign,nil;
}

func UpdateSignById(updateParam data.UpdateParam) (error) {
	signIn := SignIn{Id: updateParam.Id,Name: updateParam.Name,Message: updateParam.Message,SignInTime: time.Now()}
	result := db.Model(&signIn).Updates(SignIn{Name: updateParam.Name,Message: updateParam.Message,SignInTime: time.Now()})
	err := result.Error
	return err

}
