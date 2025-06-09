package interfaz

import "context"

type Verifier interface {
	Ejecutar(ctx context.Context, entrada any) error 
}

type Scorable interface {
	Verifier 
	Scorable() int 
	Descripcion() string 
}