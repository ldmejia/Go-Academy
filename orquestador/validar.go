package orquestador

import (
	"xfintech/interfaz"
	"xfintech/logger"
	"context"
	"fmt"
	"time"
)


type IdentityValidator[T interfaz.Verifier] struct {
	plugin T 
}

func NewIdentityValidator[T interfaz.Verifier](p T) *IdentityValidator[T]{
	return &IdentityValidator[T]{plugin: p}
}

func (v *IdentityValidator[T]) Ejecutar(entrada any){
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	
	defer cancel()

	/*defer func() {
		if r := recover(); r != nil {
			logger.Registrar(fmt.Sprintf("Recuperado de panic: %v", r))
		}	
	}()*/

	err := v.plugin.Ejecutar(ctx, entrada)

	if err != nil {
		panic("Paniqueo")
	}

	switch p := any(v.plugin).(type) {
	case interfaz.Scorable:
		logger.Registrar(fmt.Sprintln("Descripción", p.Descripcion()))
	default:
		logger.Registrar("No se puede")
	} 

}