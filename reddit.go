package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var taskId int = 0

type Course struct {
	CourseId    string `json:"courseid"`
	CourseName  string
	CoursePrice int
	Author      *Author
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

// middleware helper
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	r := mux.NewRouter()

	author1 := &Author{
		Fullname: "John Doe",
		Website:  "https://www.johndoe.com",
	}

	author2 := &Author{
		Fullname: "Jane Smith",
		Website:  "https://www.janesmith.com",
	}

	author3 := &Author{
		Fullname: "Bob Johnson",
		Website:  "https://www.bobjohnson.com",
	}

	// Create three different Courses with different Authors
	course1 := Course{
		CourseId:    "123",
		CourseName:  "Go Programming",
		CoursePrice: 49,
		Author:      author1,
	}

	course2 := Course{
		CourseId:    "456",
		CourseName:  "Python Programming",
		CoursePrice: 39,
		Author:      author2,
	}

	course3 := Course{
		CourseId:    "789",
		CourseName:  "JavaScript Programming",
		CoursePrice: 29,
		Author:      author3,
	}

	courses = append(courses, course1, course2, course3)
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")
	//listening to a port

	log.Fatal(http.ListenAndServe(":4000", r))

}

//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to my house </h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//grab id from request
	params := mux.Vars(r)

	for _, val := range courses {
		if val.CourseId == params["id"] {
			json.NewEncoder(w).Encode(val)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found ")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// nil body
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send data")
		return
	}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send data")
		return
	}
	taskId++
	course.CourseId = strconv.Itoa(taskId)
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, val := range courses {
		if val.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	//loop id, remove, add with my id

}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, val := range courses {
		if val.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("DELETED")
			break
		}
	}
}
