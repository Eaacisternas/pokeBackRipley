package ftpservices

import (
	"fmt"
	"os"
	"time"

	"github.com/secsy/goftp"
)

/* sube los archivos por ftp a un cliente asignado*/
func SubirArchivo() error {
	config := goftp.Config{
		User:               "PruebaBancoRipley@appgrade.cl",
		Password:           "Abril201118",
		ConnectionsPerHost: 10,
		Timeout:            10 * time.Second,
		Logger:             os.Stderr,
	}

	client, err := goftp.DialConfig(config, "ftp.bpyme.cl")
	if err != nil {
		fmt.Println("fallo" + err.Error())
		return err
	}

	bigFile, err := os.Open("./documents/kanto.csv")
	if err != nil {
		return err
	}

	err = client.Store("kanto.csv", bigFile)
	if err != nil {
		return err
	}
	return nil
}
