package main

import "sample-api/bootstrap"

type newStudent struct {
	Student_id       uint64 `json: "student_id" binding: "required"`
	Student_name     uint64 `json: "student_name binding: "required"`
	Student_age      uint64 `json: "student_age" binding: "required"`
	Student_address  string `json: "student_address" binding: "required"`
	Student_phone_no string `json: "student_phone_no" binding: "required"`
}

func main() {
	bootstrap.BootstrapApp()
}
