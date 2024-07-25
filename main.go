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

// "bufio"

// "os"
// "strconv"
// "strings"

// func main ( ) {
// 	fmt.Println("Hello World ");
// 	var username string = "ayush " ;
// 	fmt.Println(username);
// 	var flag bool = true ;
// 	fmt.Println(flag);
// 	var age int = 20 ;
// 	fmt.Println(age);
// 	var age1 float32 = 20.5 ;
// 	fmt.Println(age1);
// 	var age2 float64 = 20.5 ;
// 	fmt.Println(age2);
// 	var age3 complex64 = 20.5 ;
// 	fmt.Println(age3);
// 	var age4 complex128 = 20.5 ;
// 	fmt.Println(age4);
// 	var age5 byte = 20 ;
// 	fmt.Println(age5);
// 	var age6 rune = 20 ;
// 	fmt.Println(age6);
// 	user := "ayush";
// 	fmt.Println(user);
// 	user = "ayush1";
// }

// func main()  {
// 	 reader := bufio.NewReader(os.Stdin);
// 	 fmt.Println("Enter the text : ");
// 	 text , error :=  reader.ReadString('\n');
// 	 fmt.Println(text);
// 	 fmt.Println(error);
// }

// Conversions

// func main () {
// 	fmt.Println("Enter the rating ")
// 	reader:= bufio.NewReader(os.Stdin);
// 	rating ,_ := reader.ReadString('\n');
// 	fmt.Println(rating);
// 	// rating

// 	numRating, error :=strconv.ParseFloat(strings.TrimSpace(rating) ,64);

// 	if error != nil {
// 		fmt.Println(error);

// 	} else {
// 		fmt.Println(numRating  + 1 );

// 	}

// }

// Time

// func main()  {
// 	present := time.Now();
// 	fmt.Println(present);
// 	fmt.Println(present.Format("01-02-2006 15:04:05 Monday"));
// }

// Memory Management
//  new() => Memory is allocated but not initiazed  -> 0 data i.e => We cannot put up data initially

// make() => Memory is allocated and init also -> 1 based indexing => We can put up the data initally

// Pointers

// func main()  {
// 	var ptr *int
// 	fmt.Println(ptr)
// 	ayush:= "ayush";
// 	var pointer =  &ayush;
// 	fmt.Println(pointer);
// 	fmt.Println(*pointer);
// }

// // Arrays-> are not used mostly  in the golang
// func  main()  {
// 	var array [5]string;
// 	array[0] = "ayush";
// 	array[1] = "Niharika";
// 	array[2] = "Poorvi" ;
// 	fmt.Println(array);
// }

//  Slices

// func main ()  {
// 	// var slice = []string{"a",  "b", "c", "d"};
// 	// slice = append(slice, "ayush");
// 	// fmt.Println(len(slice))
// 	// slice2 := slice[1:3];
// 	// fmt.Println(slice);
// 	// fmt.Println(slice2);

// 	// highScore := make([]int,4 ) ;
// 	// highScore[0]=2;
// 	// highScore[1]=7;

// 	// highScore[2]=3;

// 	// highScore[3]=10 ;
// 	// highScore = append(highScore, 3 ,4,56,67,)
// 	// sort.Ints(highScore);
// 	// fmt.Println(highScore);

// 	// Remove the element from the slice
// 	var arr = []int{1,2,3,4,5,6,7,8,9,10};
// 	fmt.Println(arr);

// 	index := 2;

// 	arr = append(arr[:index],arr[index +1 : ]... );
// 	fmt.Println(arr);

// }

//  map

// func main ()  {
// 	lang := make(map[string]string)
// 	lang["ayush"] = "golang";
// 	lang["niharika"] = "python";
// 	lang["poorvi"] = "java";
// 	fmt.Println(lang);
// 	delete(lang,"poorvi")
// 	fmt.Println(lang);

// 	for key , value:= range lang  {
// 		fmt.Println("Key : ", key , "Value : ", value);
// 	}
// }

// Structs -> classes in the golang

// capital letter -> public
// small letter -> private

// func main() {
// 	ayush := User{"ayush", 20, "d", "d"}
// 	fmt.Println(ayush)
// 	fmt.Printf("%+v", ayush);
// 	fmt.Println()
// }

// type User struct {
// 	Name     string
// 	Age      int
// 	Email    string
// 	location string
// }

//  Conditionals if else

// func main() {
// 	a := 1010
// 	if a < 10 {
// 		fmt.Println("a is less than 10")
// 	} else {
// 		fmt.Println("a is greater than 10")
// 	}
// }

// Swith case

// func main() {
// 	rand.Seed(time.Now().UnixNano())
// 	num := rand.Intn(6) + 1
// 	fmt.Println("Num => ", num)
// 	switch num {
// 	case 1:
// 		fmt.Println("onw ")
// 	default:
// 		fmt.Println("Anything ")
// 	}
// }

// Loops
// func main() {
// 	days := []string{"Sunday", "Tuesday", "Wednesday"}

// 	for i := 0; i < len(days); i++ {
// 		fmt.Println(days[i])
// 	}

// 	for i, day := range days {
// 		fmt.Println(i, day)
// 	}

// }

// func main() {
// 	fmt.Println("Hello World")
// 	f(1, "dfhgsjhdf")
// 	fd(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
// }

// func f(i int, s string) string {
// 	fmt.Println("waesdasasa World", i, s)
// 	return s
// }

// func fd(val ...int) {
// 	fmt.Println(val)
// 	for i := range val {
// 		fmt.Println(val[i])
// 	}

// }

// type User struct {
// 	Name  string
// 	Email string
// 	Age   int
// }

// func (u User) GetDetails() {
// 	fmt.Println("Name : ", u.Name)
// 	fmt.Println("Email : ", u.Email)
// 	fmt.Println("Age : ", u.Age)
// }

// func main() {
// 	ayush := User{"ayush", "a@gmail.com", 20}
// 	ayush.GetDetails()
// 	root := Node{data: 0, left: nil, right: nil}
// 	root = *makeTree(&root)
// 	fmt.Println(root)

// }

// func makeTree(root *Node) *Node {

// 	dataReader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Enter the data : ")
// 	val, _ := dataReader.ReadString('\n')
// 	val = strings.TrimSpace(val)
// 	intval, err := strconv.Atoi(val)
// 	if err != nil {
// 		fmt.Println("Error in the input")
// 		return root;

// 	}
// 	if intval == -1 {
// 		return nil
// 	}
// 	root = &Node{data: intval, left: nil, right: nil}

// 	fmt.Println("Enter the left child of the node : " , root.data);
// 	root.left = makeTree(root.left);
// 	fmt.Println("Enter the right child of the node : " , root.data);
// 	root.right = makeTree(root.right);
// 	return root;

// }

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// }

// Files
// func main() {
// 	content := "Hellosdasdasdcxzcxzsadasdasd World"
// 	file, err := os.Create("test.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	length, err := io.WriteString(file, content)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Length of the file : ", length)
// 	defer file.Close()
// 	readFile("test.txt")
// }
// func readFile(filename string) {
// 	dataByte, _ := ioutil.ReadFile(filename)
// 	fmt.Println(string(dataByte))
// }

// WEb Reqwuest

// const url = "https://jsonplaceholder.typicode.com/posts"

// func main() {
// 	response, err := http.Get(url)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer response.Body.Close()
// 	fmt.Println("Status Code : ", response)
// 	defer response.Body.Close()
// 	data, _ := ioutil.ReadAll(response.Body)
// 	fmt.Println(string(data))

// }

// GO MOD ->

// func main() {
// 	fmt.Println("Hello ")
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", home).Methods("GET")
// 	log.Fatal(http.ListenAndServe(":4000", r))

// }
// func home(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("<h1> Hello Eodsljknfgsdxlk </h1>"))
// }

// APis

//  Models

type Course struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Author *Author `json:"author"`
	Price  int     `json:"price"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:website`
}

// fake database
var courses []Course

// Middleware
func (c *Course) isEmpty() bool {
	return c.Name == ""
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range courses {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the id ")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("No data found")
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data found")
		return

	}

	// Generate the id
	rand.Seed(time.Now().UnixNano())
	course.ID = strconv.Itoa(rand.Intn(1000))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(courses)

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, course := range courses {
		if course.ID == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.ID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course with Id Found ")
	return

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, course := range courses {
		if params["id"] == course.ID {
			courses = append(courses[:i], courses[i+1:]...)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course with  the Id Found ! ")

}

func main() {
	fmt.Println("Hello World")
	r := mux.NewRouter()

	courses = append(courses, Course{ID: "2", Name: "Ayush", Author: &Author{Fullname: "sjfdnbjskand", Website: "dsfkjbsdkjfnb"}, Price: 30000})
	courses = append(courses, Course{ID: "3", Name: "Ayush", Author: &Author{Fullname: "sdsd", Website: "dsds"}, Price: 303423000})

	r.HandleFunc("/api/courses", getAllCourses).Methods("GET")	
	r.HandleFunc("/api/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/api/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/api/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/api/courses/{id}", deleteOneCourse).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":4000", r))

}



