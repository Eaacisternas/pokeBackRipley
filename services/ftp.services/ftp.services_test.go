package ftpservices_test

import (
	"testing"

	ftpservices "github.com/Eaacisternas/pokeBackRipley/services/ftp.services"
)

/*TestCreate, test unitario para SubirArchivo de ftpservice*/
func TestSubirArchivo(t *testing.T) {
	err := ftpservices.SubirArchivo()
	if err != nil {
		t.Error("La prueba de subir el archivo ha fallado")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito")
	}
}
