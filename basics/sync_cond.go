package main

import (
	"fmt"
	"sync"
	"time"
)

type Animal struct{ name string }

func (animal *Animal) Eat(food *AnimalFood) {
	food.Lock()
	food.cond.Wait()
	food.foodAmount--
	food.Unlock()
}

type AnimalFood struct {
	sync.Mutex
	foodAmount int

	cond *sync.Cond
}

func NewAnimalFood(foodAmount int) *AnimalFood {
	r := AnimalFood{foodAmount: foodAmount}
	r.cond = sync.NewCond(&r)
	return &r
}

func main() {
	wg := &sync.WaitGroup{}
	food := NewAnimalFood(4)

	for _, animal := range []Animal{{"Vasya"}, {"Dima"}, {"Aleksey"}, {"Kolya"}} {
		wg.Add(1)
		go func(animal Animal) {
			defer wg.Done()
			animal.Eat(food)
		}(animal)
	}

	fmt.Println("Waiting for food to arrive...")
	time.Sleep(1 * time.Second)

	food.cond.Broadcast()

	wg.Wait()
	fmt.Printf("Food left: %d\n", food.foodAmount)
}
