package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//model for course - file

type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

// face db
var courses []Course

// middleware or helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

type Author struct {
	FullName string `json:"full_name"`
	Website  string `json:"website"`
}

func main() {
	fmt.Println("API")
	r := mux.NewRouter()

	//seeding of data in courses
	courses = append(courses, Course{
		CourseId:    "1",
		CourseName:  "React js",
		CoursePrice: 100,
		Author: &Author{
			FullName: "Jane Doe",
			Website:  "https://jane.doe.com",
		},
	},
		Course{
			CourseId:    "2",
			CourseName:  "React js",
			CoursePrice: 100,
			Author: &Author{
				FullName: "Jane Doe",
				Website:  "https://jane.doe.com",
			},
		},
		Course{
			CourseId:    "3",
			CourseName:  "React js",
			CoursePrice: 100,
			Author: &Author{
				FullName: "Jane Doe",
				Website:  "https://jane.doe.com",
			},
		})

	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/courses", GetAllCourses).Methods("GET")
	r.HandleFunc("/get-one/{id}", GetCourseById).Methods("GET")
	r.HandleFunc("/create", CreateOneCourse).Methods("POST")
	r.HandleFunc("/update/{id}", UpdateOneCourse).Methods("PUT")
	r.HandleFunc("/delete/{id}", DeleteOneCourse).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers - file

// serve home route

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Luci</h1>"))
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func GetCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get course by id")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop through courses and find matching ID and return the response
	for _, item := range courses {
		if item.CourseId == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course Found with given id")
	return
}

func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please enter a valid JSON body")
	}

	//what if: {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please enter a valid JSON body")
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	_ = json.NewEncoder(w).Encode(courses)
	return
}
func UpdateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first: grab id from request
	params := mux.Vars(r)

	// loop, id, remove, add with same ID
	for index, item := range courses {
		if item.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			if course.IsEmpty() {
				_ = json.NewEncoder(w).Encode("Please enter a valid JSON body")
				return
			}
			rand.Seed(time.Now().UnixNano())
			course.CourseId = params["id"]
			courses = append(courses, course)
			_ = json.NewEncoder(w).Encode(courses)
			return
		}
	}
}

func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range courses {
		if item.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			_ = json.NewEncoder(w).Encode("Course Deleted")
			return
		}
	}
}
