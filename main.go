package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


// Students
var students = []Student{
	{Id: 1, Name: "John", Class: "2-A", Teacher: "Smith"},
    {Id: 2, Name: "Emily", Class: "8-B", Teacher: "Johnson"},
    {Id: 3, Name: "Michael", Class: "5-C", Teacher: "Williams"},
    {Id: 4, Name: "Sophia", Class: "1-B", Teacher: "Brown"},
    {Id: 5, Name: "William", Class: "9-A", Teacher: "Jones"},
}

type Student struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Class   string `json:"class"`
	Teacher string `json:"teacher"`
}

// Students
func listStudents(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, students)
}

// Create Student
func createStudent(context *gin.Context) {
	var studentByUser Student
	if err := context.BindJSON(&studentByUser); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON Data"})
		return
	}
	if studentByUser.Id != 0 && studentByUser.Name != "" && studentByUser.Class != "" && studentByUser.Teacher != "" {
		students = append(students, studentByUser)
		context.IndentedJSON(http.StatusCreated, gin.H{"message": "Student has been created!", "Student ID": studentByUser.Id})
		return
	} else {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Student cannot be created!", "Student ID": studentByUser.Id})
		return
	}
}


// Find Student
func getStudentByID(int_id int) (*Student, error) {
	for i, s := range students {
		if s.Id == int_id {
			return &students[i], nil
		}
	}
	return nil, errors.New("Student cannot be found!")
}

// Get Student
func getStudent(context *gin.Context) {
	str_id := context.Param("id")
	int_id, err := strconv.Atoi(str_id)
	if err != nil {
		panic(err)
	}

	student, err := getStudentByID(int_id)
	if err == nil {
		context.IndentedJSON(http.StatusOK, student)
	} else {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Student cannot found!"})
	}

}

func main() {
	router := gin.Default()
	router.GET("/students", listStudents)
	router.POST("/students", createStudent)
	router.GET("/students/:id", getStudent)
	router.Run("localhost:8080")
}
