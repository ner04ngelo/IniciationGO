package File

type Empleado struct {
	Persona
	Puesto string
}

func (emp Empleado) ImprimirPuesto () string {
	return emp.Puesto
}


func (emp *Empleado) SobreEscribirPuesto(cadena string) {
	 emp.Puesto = cadena
}