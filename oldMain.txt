// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// //TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// // the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
// var dbData= []string{"id1", "id2", "id3", "id4", "id5"}
// var results = []string{}
// var wg = sync.WaitGroup{}
// var m = sync.Mutex{}
// func main() {

// 	// s := "gopher"
// 	// fmt.Printf("Hello and welcome, %s!\n", s)

// 	// for i := 1; i <= 5; i++ {
// 	// 	fmt.Println("i =", 100/i)
// 	// }

// 	t0 := time.Now()
// 	for i:= range dbData {
// 		wg.Add(1)
// 		go dbCall(i)
// 	}
// 	wg.Wait()
// 	fmt.Printf("\n Total execution time: %v", time.Since(t0))
// 	fmt.Printf("\n The results are %v", results)
// }

// func dbCall(i int){
// 	var delay float32 = rand.Float32() * 2000
// 	time.Sleep(time.Duration(delay) * time.Millisecond)
// 	fmt.Println("\n The result from the db is", dbData[i])
// 	m.Lock()
// 	results = append(results, dbData[i])
// 	m.Unlock()
// 	wg.Done()
// }

// //generics

// func sumSlice[T int | float32 | float64](slice []T) T {
// 	var sum T
// 	for _, v := range slice {
// 		sum += v
// 	}
// 	return sum
// }
