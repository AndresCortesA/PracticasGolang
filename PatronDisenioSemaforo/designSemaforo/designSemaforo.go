package designSemaforo

import (
	"errors"
	"time"
)

/*
	Declarar atributos o funciones publicas se hace con letra mayus de primera
	para declararla privada es en minusculas.
*/

// creamos unos errores para el diseño
var (
	ErrorNoEntrada  = errors.New("semaforo: no puede adquirir un semaforo")
	ErrorDeRelacion = errors.New("semaforo: no puede existir una relacion del semaforo sin adquirirlo primero")
)

// Interface que contiene los comportamientos del adquirir y/o relacionar los semaforos
type Interface interface {
	Adquirir() error
	Relacionar() error
}

type implementacion struct {
	sem         chan struct{}
	tiempoFuera time.Duration
}

// Adquirir
func (s *implementacion) Adquirir() error {
	select {
	case s.sem <- struct{}{}:
		return nil
	case <-time.After(s.tiempoFuera):
		return ErrorNoEntrada
	}
}

// Relacion
func (s *implementacion) Relacionar() error {
	select {
	case _ = <-s.sem:
		return nil
	case <-time.After(s.tiempoFuera):
		return ErrorDeRelacion
	}
}

func Nuevo(entradas int, fuera time.Duration) Interface {
	return &implementacion{
		sem:         make(chan struct{}, entradas),
		tiempoFuera: fuera,
	}
}
