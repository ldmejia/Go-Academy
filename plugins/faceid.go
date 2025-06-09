package plugins

import (
	"xfintech/interfaz"
	"context"
	"fmt"
	"time"
)

type FaceID struct{}


func (f *FaceID) Ejecutar(ctx context.Context, entrada any) error {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Simulando FaceId")

	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

func (f *FaceID) Scorable() int {
	return 95
}

func (f *FaceID) Descripcion() string {
	return "Usando FaceID"
}

func NewFaceID() interfaz.Scorable {
	return &FaceID{}
}
