package main

import "fmt"

func main() {

	// coral := [4]string{"blue coral", "staghorn coral", "pillar coral", "elkhorn coral"}
	// num := [2]int{5, 3}
	// fmt.Printf("%q\n", coral)
	// fmt.Printf("%q\n", num)
	// cuando definimos un areglo con una dimension determinada ya no pondremos el tamaño ni agregar mas elementos..
	// es por ello que debemos de usar un slice para trabajar con arreglos
	// corals := []string{"blue coral", "staghorn coral", "pillar coral", "elkhorn coral"}
	// newCoral := append(corals, "coral new agreement")
	// fmt.Printf("%q\n", corals)
	// fmt.Printf("%q\n", newCoral)

	// crear un array y despues agregar sus valores

	// animals := make([]string, 3)

	// fmt.Printf("%q\n", animals) // => ["" "" ""]

	// coral := [6]string{"blue coral", "foliose coral", "pillar coral", "elkhorn coral", "staghorn coral", "coral new agreement"}

	// coralsNews := coral[3:5] // slice de un arreglo
	// segundo valor es la suma del inicio y los elementos que queremos que se muestren.

	// fmt.Printf("%q\n", coralsNews)

	// fmt.Println(len(coral))

	goSlice()
}

func sliceMake() {

	items := make([]int, 3, 5) // 3 es el tamaño inicial y 5 es el tamaño maximo

	fmt.Printf("%d %d", len(items), cap(items))

}

func goSlice() {
	items := make([]int, 0, 0)
	for i := 0; i < 15; i++ {
		items = append(items, i)
	}
	fmt.Printf("%d %d", len(items), cap(items))
}
