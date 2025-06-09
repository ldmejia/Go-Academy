package main

import (
	"xfintech/orquestador"
	"xfintech/plugins"
	"context"
)


func main(){

	pluginsD := plugins.NewFaceID()

	switch verificar := pluginsD.(type){
		case interface {
			Ejecutar(ctx context.Context, entrada any) error 
		}:
		validar := orquestador.NewIdentityValidator(verificar)
		validar.Ejecutar(nil)
	}
}