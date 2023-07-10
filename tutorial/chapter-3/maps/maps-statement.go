package maps

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// Creamos un map con nill zero value la key de tipo string y el value de tipo de la strucutra
var m map[string]Vertex

func MapsDeclaration() {
	//inicializaciamos el map para que quede listo para usar
	m := make(map[string]Vertex)

	//agregar un elemento al mapa
	m["Bell labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell labs"])
}

func MapsDeclarationUnderlying() {

	var m = map[string]Vertex{
		"Bell Labs": { //omiting type
			40.68433, -74.39967,
		},
		"Google": { //omiting type
			37.42202, -122.08408,
		},
	}

	fmt.Println("Maps", m)
}

func MapsMutating() {
	m := make(map[string]int) //El zero value es con respecto al tipo del elemento

	m["Answer"] = 42 //Inserting a elem
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48 //Updating
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("Tle value:", m["Answer"])

	elem, ok := m["Answer"]
	fmt.Println("The value:", elem, "Present?", ok)

}
