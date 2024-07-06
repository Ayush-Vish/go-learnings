package main

import "fmt"

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

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) GetDetails() {
	fmt.Println("Name : ", u.Name)
	fmt.Println("Email : ", u.Email)
	fmt.Println("Age : ", u.Age)
}

func main() {
	ayush := User{"ayush", "a@gmail.com", 20}
	ayush.GetDetails()

}
