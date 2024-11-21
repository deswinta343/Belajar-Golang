package main
import "fmt"

func main() {
//--------------------------------------------//

// get text
fmt.Println("Hello world")
  
//--------------------------------------------//
  
// get variabel initial

//soal a
  var student1 string = "Deswinta"
  var student2 = "Susanto"
  x := 2
  
  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)

//soal b

    // var a string
    // var b int
    // var c bool
    
    // fmt.Println(a)
    // fmt.Println(b)
    // fmt.Println(c)
    
    //soal c
    
    var studentke1 string
    student1 = "Desss"
    fmt.Println(studentke1)
    
//--------------------------------------------//
    
//multiple variable

//soal a

var a,b,c,d int = 1,3,5,7
fmt.Println(a)
fmt.Println(b)
fmt.Println(c)
fmt.Println(d)

//soal b

// var a, b = 6, "Hello"
// c, d := 7, "World"

// fmt.Println(a)
// fmt.Println(b)
// fmt.Println(c)
// fmt.Println(d)

//--------------------------------------------//
    
//constants

const PI = 3.14
fmt.Println(PI)

const name = "Dessss"
fmt.Println(name)

const firstname = "Deswinta"
const lastname = "Susanto"
const fullname = firstname + lastname

// fmt.Println(firstname)
// fmt.Println(lastname)
fmt.Println(fullname)


// --------------------------------------------//
    
//array

var arrayke1 = [...]int{1,2,3}
arrayke2 := [...]int{4,5,6,7,8,9}

fmt.Println(arrayke1)
fmt.Println(arrayke2)

var array1 = [3]int{1,2,3}
array2 := [5]int{4,5,6,7,8}
//jika initial sebelum deret angka lebih dari total angka di deret maka akan diisi 0 0 dst

fmt.Println(array1)
fmt.Println(array2)

var merk = [...]string{"honda", "bmw", "mazda", "toyota"}
fmt.Println(merk)

arr1 := [5]int{}
arr2 := [5]int{1,2}
arr3 := [5]int{1,2,3,4,5}

fmt.Println(arr1)
fmt.Println(arr2)
fmt.Println(arr3)

}