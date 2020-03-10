package car

type Car struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
}

func AddElement(car *Car) string {
	return "AddElement : It's working!"
}

func DeleteElement(car *Car) string {
	return "DeleteElement : It's working!"
}
