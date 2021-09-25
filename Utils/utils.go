package Utils

import "fmt"

type Empleado struct {
	Nombre   string
	Apellido string
	Edad     int
	Tel      string
	Anos     int
	Sueldo   int
}

type Empleados map[int]Empleado

type Empresa struct {
	Name          string
	Logo          string
	Address       string
	Tel           string
	TotalEmpleado int
	EmpleadoInfo  Empleados
}

type Empresas map[int]Empresa
type EmpresaName string

type EmpresaManager interface {
	SetEmpleado(e *Empresa, empleadoInfo Empleado) bool
	getEmpleados(empresaName string, epleadoName string) Empleado
	listEmpleados(empresaName string) Empleados
	getAumento(Empresa string, empleado string) int
}

var listaEmpresas = make(Empresas)

func SetEmpresa(empresaInfo Empresa) bool {
	positions := len(listaEmpresas)
	empresaInfo.EmpleadoInfo = make(Empleados)
	listaEmpresas[positions+1] = empresaInfo
	return true
}

func (e *Empresa) SetEmpleado(empleadoInfo Empleado) bool {
	positionEmpleado := len(e.EmpleadoInfo)
	e.TotalEmpleado = positionEmpleado
	fmt.Println(positionEmpleado)
	e.EmpleadoInfo[positionEmpleado+1] = empleadoInfo
	return true
}

func GetEmpresa(empresaName string) Empresa {
	var empresa Empresa
	for _, empresa_ := range listaEmpresas {
		if empresa_.Name == empresaName {
			empresa = empresa_
		}
	}
	return empresa
}

func (e *Empresa) getEmpleados(empresaName string, epleadoName string) Empleado {
	var empleado Empleado
	if e != nil {
		for _, empleado_ := range e.EmpleadoInfo {
			empleado = empleado_
		}
	} else {
		for _, empresa := range listaEmpresas {
			if empresa.Name == empresaName {
				for _, empleado_ := range empresa.EmpleadoInfo {
					empleado = empleado_
				}
			}
		}
	}

	return empleado
}

func ListEmpresas() Empresas {
	return listaEmpresas
}

func (e *Empresa) listEmpleados(empresaName string) Empleados {
	var empleados map[int]Empleado = make(map[int]Empleado)
	if e != nil {
		empleados = e.EmpleadoInfo
	} else {

		for _, empresa := range listaEmpresas {
			if empresa.Name == empresaName {
				empleados = empresa.EmpleadoInfo
			}
		}
	}
	return empleados
}
