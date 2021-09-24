package helpers

import (
	"github.com/Augustin0/Types-Interfaces/Utils"
	"github.com/Augustin0/Types-Interfaces/db"
)

func SetEmpleado(empresaName string, empleado Utils.Empleado) bool {
	BancoHhd := Utils.GetEmpresa(empresaName)
	BancoHhd.SetEmpleado(empleado)
	return true
}

func GenEmpleado(Nombre string, Apellido string, Edad int, Tel string, Anos int, Sueldo int) Utils.Empleado {
	return Utils.Empleado{Nombre: Nombre, Apellido: Apellido, Edad: Edad, Tel: Tel, Anos: Anos, Sueldo: Sueldo}
}

func SetEmpresa360() {
	for _, empresa := range db.DBEmpresas {
		Utils.SetEmpresa(empresa)
	}
}
