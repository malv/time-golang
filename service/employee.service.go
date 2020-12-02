package service

import (
	"log"
	"time-project/config"
	pb "time-project/proto/model"
)

// EmployeeService ...
type EmployeeService struct{}

// GetEmployees ...
func (EmployeeService) GetEmployees() (data *pb.Employees, e error) {
	defer config.CatchError(&e)
	res, err := config.EmployeeClient.GetEmployees(config.CtxEmployee, &pb.Tokens{Token: ReqToken})
	log.Print(res)
	if err != nil {
		return data, err
	}

	return res, nil
}
