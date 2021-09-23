package main

var students = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

var teachers = []*Teacher{}

type Teacher struct {
	Id   string
	Name string
}

func GetTeachers() []*Teacher {
	return teachers
}

func SelectTeacher(id string) *Teacher {
	for _, each := range teachers {
		if each.Id == id {
			return each
		}
	}

	return nil
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}

	return nil
}

func init() {
	students = append(students, &Student{Id: "s001", Name: "bourne", Grade: 2})
	students = append(students, &Student{Id: "s002", Name: "ethan", Grade: 2})
	students = append(students, &Student{Id: "s003", Name: "wick", Grade: 3})
	teachers = append(teachers, &Teacher{Id: "a1", Name: "Jhon"})
	teachers = append(teachers, &Teacher{Id: "a2", Name: "Iron"})
	teachers = append(teachers, &Teacher{Id: "a3", Name: "Thor"})
}
