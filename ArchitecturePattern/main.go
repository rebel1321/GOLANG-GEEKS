package main

import (
	"arch/factory"
	"arch/observer"
	"arch/singleton"
	"arch/decorator"
	"fmt"
	"time"
)

func main() {
	singleton.SingletonPattern()

	payment := factory.GetPaymentMethod("credit")
	if payment != nil {
		println(payment.Pay(100))
	} else {
		println("Invalid payment method")
	}

	newsAgency := &observer.NewsAgency{}

	// Create observers
	observer1 := &observer.Observer{
		Name:                "The Daily Times",
		NotificationChannel: make(chan string),
	}
	observer2 := &observer.Observer{
		Name:                "The Morning Herald",
		NotificationChannel: make(chan string),
	}

	// Register observers
	newsAgency.Register(observer1)
	newsAgency.Register(observer2)

	go func() {
		for message := range observer1.NotificationChannel {
			fmt.Printf("%s received: %s\n", observer1.Name, message)
		}
	}()

	go func() {
		for message := range observer2.NotificationChannel {
			fmt.Printf("%s received: %s\n", observer2.Name, message)
		}
	}()

	// Notify all observers
	go func() {
		newsAgency.NotifyAll("Breaking News: GoLang 1.21 Released!")
	}()

	// Wait to process notifications
	time.Sleep(1 * time.Second)
	fmt.Println("Deregistering observer: The Daily Times")
	newsAgency.Deregister(observer1)

	// Notify all observers again
	go func() {
		newsAgency.NotifyAll("Update: GoLang is awesome!")
	}()

	// Wait to process notifications
	time.Sleep(1 * time.Second)

	// Close channels to clean up resources
	close(observer1.NotificationChannel)
	close(observer2.NotificationChannel)

	var myCoffee decorator.Coffee = &decorator.BasicCoffee{}
	fmt.Println("Cost:", myCoffee.Cost(), "Ingredients:", myCoffee.Ingredients())

	// Add milk
	myCoffee = &decorator.WithMilk{Coffee: myCoffee}
	fmt.Println("Cost:", myCoffee.Cost(), "Ingredients:", myCoffee.Ingredients())


    // Add sugar
    myCoffee = &decorator.WithSugar{Coffee: myCoffee}
    fmt.Println("Cost:", myCoffee.Cost(), "Ingredients:", myCoffee.Ingredients())
}
