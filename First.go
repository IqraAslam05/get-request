package main

import (
	"fmt"
	
)

// func main() {

// 	fmt.Println(  calculatePrice(14.5, 3))
// }

// func calculatePrice(rate float32,night int)float32{
// 	return rate * float32 (night)
// }

//------------------second-----------------------//

// func totalPrice(rate float32,nights int)float32{
// 	return rate*float32(nights)//same data type is required to perform multiplication so int is converted into float//

// }

// func main(){
// Ali:=totalPrice(4.5,4)
// Ahmad:=totalPrice(3.5,7)
// total:=Ali+Ahmad
// fmt.Printf("Total Price is  Rs %0.2f ",total)
// }
//-------------------3rd--------------------------//
// func main() {
//     johnPrice := calculatePrice(145.90, 3)
//     fmt.Println("John:", johnPrice, "rate:", rate)//rate must be defined in main function
// }

// func calculatePrice(rate float32, nights int) float32 {
//     return rate * float32(nights)
// }

//-----------------4th-------------------------//
// func calculatePrice(rate float32,nights int )float32{
//    n:=float32(nights)
//    return rate*n

// }
// func main(){

//    fmt.Println(calculatePrice(14.5,3))

// }
//--------------5th----------------------//

//result param should be in parenthesis with its type
// func FindPrice(rate float32,nights int) (price float32){
// price=rate* float32(nights)
// return price  //this price var is optional,can be ommitted.
// }
// func main(){

//     fmt.Println("Price is", FindPrice(5.7,8))
// }
//-------------6th-----------------//
// func Calculateprice( rate float32,night int)(price float32){
//   p:=rate*float32(night)
//   p=p*2
//   return        //bcz missing p after return ,and given output is according to Default value of (price float 32) that is 0.00

// }
// func main(){
//     a:=Calculateprice(334.14,8)
//     fmt.Printf("%0.2f",a)
// }
//---------------7th---------------------//

// //import
// //"time"
// import
// "math/rand"
// func EmptyRooms()int {
//    //[ rand.Seed(time.Now().UTC().UnixNano())]
//    return rand.Intn(20)
// }
// func main(){
//     v:=EmptyRooms()
//     fmt.Println(v)
// }

//-----------------------8th  without any param and result----------------------------//

// func ShowMessage(){
//     fmt.Println("Hello World")
//     fmt.Println("Welcome to Pakistan")
// }
// func main(){
//     ShowMessage()
// }

//-----------------------9th------------------//
// func main(){
// const a="Pakistan"
//         Print()
// }
// func Print(){

//     fmt.Println("This is my country",a)  //undefined a ,a is not in scope of Print func
//     fmt.Println("Error")
// }
//------------------10th--------------------------//


import
"strings"
               //Define your ASCII space characters here
var asciiSpace = [256]byte{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}
func TrimSpace(input string) string {
	// Searching for the first non-ASCII non-space byte
	startingPos := 0
	for ; startingPos < len(input); startingPos++ {
		char := input[startingPos]
		if asciiSpace[char] == 0 {
			break
		}
	}
	// Scanning backwards for the last non-ASCII non-space byte from the end
	endingPos := len(input)
	for ; endingPos > startingPos; endingPos-- {
		char := input[endingPos-1]
		if asciiSpace[char] == 0 {
			break
		}
	}

	// At this point, input[startingPos:endingPos] consists of non-space bytes, completing the process.
	return input[startingPos:endingPos]
}
func main() {
	input := 	"Hello,World"
	   
	result := strings.TrimSpace(input)
	fmt.Printf("Original: '%s'\nTrimmed: '%s'\n", input, result)
}

//------------------------11--------------------------//

//  import
//  "strings"

// func main() {
// 	input := "	Hello,World   "
// 	result := strings.TrimSpace(input)
// 	fmt.Printf("Original: '%s'\nTrimmed: '%s'\n", input, result)
// }
//-----------------------------
// import
// "strings"
// func main(){
// a:="@qwertyui@"
// fmt.Println(strings.Trim(a,"@"))
//}


// import
// "strings"

// func main(){
// s:= "    qwertyui   "
// fmt.Println(s,"before Trim")
// fmt.Println(strings.TrimSpace(s))
// }