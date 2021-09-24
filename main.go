package main

import (
	"fmt"

	"github.com/Augustin0/Types-Interfaces/Utils"
	"github.com/Augustin0/Types-Interfaces/helpers"
)

func main() {
	helpers.SetEmpresa360()
	empleado1 := helpers.GenEmpleado("Jose Hernandez", "Mala Suerte", 50, "809", 10, 50000)
	helpers.SetEmpleado("Banco BHD", empleado1)
	fmt.Println(Utils.ListEmpresas())

}
