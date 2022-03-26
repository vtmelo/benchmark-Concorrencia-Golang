package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"runtime"
	"strconv"
	"sync"
)

const NUM_SAPOS int = 5
const DISTANCIA int = 500
const PULO_MAXIMO int64 = 50

var COLOCACAO = 0

type frog struct {
	name              string
	distanceRace      int
	totalDistanceRace int
	distanceJump      int
	totalJumps        int
	FinalPosition     int
	maxJump           int64
}

func newFrog(name string) frog {
	return frog{
		name:              name,        // nome do sapo
		distanceRace:      0,           // distancia total corrida pelo sapo
		totalDistanceRace: DISTANCIA,   // distancia total da corrida
		distanceJump:      0,           // tamanho do pulo do sapo em cm
		totalJumps:        0,           // quantidades de pulos dados na corrida
		FinalPosition:     0,           // colocação do sapo ao final da corrida
		maxJump:           PULO_MAXIMO, // pulo máximo em cm que um sapo pode dar
	}
}

//increment the jump of the frog
func (f *frog) jump() {
	f.totalJumps++

	n, _ := rand.Int(rand.Reader, big.NewInt(f.maxJump))
	f.distanceJump = int((n.Int64() + 1))

	f.distanceRace += f.distanceJump
	if f.distanceRace > f.totalDistanceRace {
		f.distanceRace = f.totalDistanceRace
	}
}

func (f *frog) printActualSituation() {
	fmt.Printf("O sapo %s pulou %dcm e ja percorreu %dcm\n", f.name, f.distanceJump, f.distanceRace)
}

func (f *frog) restFrog() {
	runtime.Gosched()
}

func (f *frog) printFinalPosition() {
	rw := &sync.RWMutex{}
	rw.Lock()
	COLOCACAO++
	fmt.Printf("*****O sapo %s chegou na posição %d com %d pulos*****\n", f.name, COLOCACAO, f.totalJumps)
	rw.Unlock()
}
func (f *frog) run(wg *sync.WaitGroup) {
	for f.distanceRace < f.totalDistanceRace {
		f.jump()
		f.printActualSituation()
		f.restFrog()
	}
	f.printFinalPosition()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(NUM_SAPOS)
	for i := 1; i <= NUM_SAPOS; i++ {
		frog := newFrog("SAPO_" + strconv.Itoa(i))
		go frog.run(&wg)
	}
	wg.Wait()
}
