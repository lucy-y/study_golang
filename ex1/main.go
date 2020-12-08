package main

import (
  "fmt"
  "strings"
)

func study1_init(){
  // const 상수 (js:const)
  const name1 string= "a"
  // 상수 값 변경 불가능
  // name1= "b"
  fmt.Println(name1);

  // var 변수 (js: let)
  // var 선언방식은 name2 혹은 name3 처럼 선언가능.
  var name2 string= "b"
  name2 = "c"
  fmt.Println(name2)

  name3 := "b"
  name3 = "c"
  fmt.Println(name3)

  return
}

func multiply(a, b int) int {
  return a * b
}

func study2_function(name string) (length int){
  length = len(name)  
  return
}

func lenAndUpper(name string) (length int, uppercase string){
  defer fmt.Println("I'm done.")
  length = len(name)
  uppercase = strings.ToUpper(name)
  return
}

func superAdd(numbers ...int) int {
  //fmt.Println(numbers)
  //return 1

  total := 0

  //for i:=0; i<len(numbers); i++ {
  //  total += numbers[i]
  //}

  for _, number := range numbers {
    total += number
  }

  return total
}

func canI(age int) bool {
  //if age > 18 {
  //  return true
  //} 
  if koreaAge := age+2; koreaAge >= 18 {
    return true
  }
  return false
}

type person struct{
  name    string
  age     int
  favFood []string
}


func main() {

  //study1_init()

  //fmt.Println(multiply(2,2))

  //const name string = "lucy"
  //fmt.Println(study2_function(name))

  //_, upperName := lenAndUpper(name)
  //fmt.Println(upperName)

  //result := superAdd(1,2,3,4,5,6)
  //fmt.Println(result)

  //fmt.Println(canI(20))
  //fmt.Println(canI(16))

  favFood := []string{"kimchi", "ramen"}
  nico := person{name: "nico", age: 18, favFood: favFood}
  fmt.Println(nico)
  fmt.Println(nico.name)


}
