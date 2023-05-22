package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaforo struct {
	color    string
	duracion time.Duration
}

func cambiarSemaforo(semaforo *Semaforo, ch chan<- *Semaforo) {
	fmt.Printf("%s semaforo encendido\n", semaforo.color)
	time.Sleep(semaforo.duracion)

	fmt.Printf("%s semaforo apagado\n", semaforo.color)
	time.Sleep(time.Second)

	ch <- semaforo // Enviar semáforo actual al canal
}

func controladorSemaforos(semaforos []*Semaforo, ch <-chan *Semaforo, cambio <-chan struct{}) {
	var siguiente *Semaforo

	for {
		semaforo := <-ch // Recibir semáforo actual del canal

		// Encontrar el siguiente semáforo en la lista
		for i, s := range semaforos {
			if s == semaforo {
				siguiente = semaforos[(i+1)%len(semaforos)]
				break
			}
		}

		time.Sleep(time.Second) // Esperar antes de cambiar al siguiente semáforo

		fmt.Printf("Cambiando de %s a %s semaforo\n", semaforo.color, siguiente.color)

		<-cambio // Esperar señal de la goroutine anónima
	}
}

func main() {
	semaforo1 := &Semaforo{color: "Rojo", duracion: 5 * time.Second}
	semaforo2 := &Semaforo{color: "Amarillo", duracion: 2 * time.Second}
	semaforo3 := &Semaforo{color: "Verde", duracion: 3 * time.Second}

	semaforos := []*Semaforo{semaforo1, semaforo2, semaforo3}

	ch := make(chan *Semaforo, 2)
	cambio := make(chan struct{})

	go controladorSemaforos(semaforos, ch, cambio)

	var wg sync.WaitGroup
	wg.Add(len(semaforos))

	for _, semaforo := range semaforos {
		go func(s *Semaforo) {
			cambiarSemaforo(s, ch)
			cambio <- struct{}{} // Señalar cambio de semáforo a la goroutine controladorSemaforos
			wg.Done()
		}(semaforo)
	}

	wg.Wait()
}
