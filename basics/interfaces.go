package main

import "fmt"

type Hound interface{ Hunt() }

type Poodle interface{ Bark() string }

type Decorative interface{ Sleep(hours int) }

type Bloodhound struct{ name string }
type TeacupPoodle struct{ name string }
type ShihTzu struct{ name string }

func (b Bloodhound) Hunt()          { fmt.Printf("%s is hunting\n", b.name) }
func (p TeacupPoodle) Bark() string { return fmt.Sprintf("%s is barking\n", p.name) }
func (p ShihTzu) Sleep(hours int)   { fmt.Printf("%s is sleeping %d hours\n", p.name, hours) }

func zoo(dog interface{}) {
	switch dog.(type) {
	case Hound:
		dog.(Hound).Hunt()
	case Poodle:
		fmt.Printf("%s", dog.(Poodle).Bark())
	case Decorative:
		dog.(Decorative).Sleep(10)

	default:
		fmt.Printf("Unknown %T", dog)
	}
}

func main() {
	zoo(Bloodhound{"mishka"})
	zoo(TeacupPoodle{"vasya"})
	zoo(ShihTzu{"dima"})
	zoo(int(1))
}
