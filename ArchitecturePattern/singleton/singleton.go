package singleton

import (
	"fmt"
	"sync"

)

type singleton struct{
	value int
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{value: 42} // Example value
	})
	return instance
}

func SingletonPattern(){
	// Example usage of the singleton
	singletonInstance := GetInstance()
	fmt.Println("Singleton value:", singletonInstance.value)

	singletonInstance2 := GetInstance()
	fmt.Println("Singleton value from second instance:", singletonInstance2.value)

	if singletonInstance == singletonInstance2 {
		fmt.Println("Both instances are the same.")
	} else {
		fmt.Println("Instances are different.")
	}
}