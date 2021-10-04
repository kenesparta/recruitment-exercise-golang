package assemblyspot

import (
	"errors"
	"fmt"
	"sync"
	"time"

	".main.go/vehicle"
)

type AssemblySpot struct {
	vehicleToAssemble *vehicle.Car
	assemblyLog       string
}

func (s *AssemblySpot) SetVehicle(v *vehicle.Car) {
	s.vehicleToAssemble = v
}

func (s *AssemblySpot) GetAssembledVehicle() *vehicle.Car {
	return s.vehicleToAssemble
}

func (s *AssemblySpot) GetAssembledLogs() string {
	return s.assemblyLog
}

//hint: improve this function to execute this process concurrenlty
func (s *AssemblySpot) AssembleVehicle() (*vehicle.Car, error) {
	if s.vehicleToAssemble == nil {
		return nil, errors.New("no vehicle set to start assembling")
	}
	s.assemblyLog = ""
	var (
		wg       sync.WaitGroup
		mu       sync.Mutex
		pipeline = []func(*sync.WaitGroup, *sync.Mutex){
			s.assembleChassis,
			s.assembleTires,
			s.assembleEngine,
			s.assembleElectronics,
			s.assembleDash,
			s.assembleSeats,
			s.assembleWindows,
		}
	)

	// This executes every process on a different go-routine
	for _, process := range pipeline {
		wg.Add(1)
		go process(&wg, &mu)
	}

	wg.Wait()
	return s.vehicleToAssemble, nil
}

func (s *AssemblySpot) assembleChassis(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	s.vehicleToAssemble.Chassis = "Assembled"
	time.Sleep(1 * time.Second)
	mu.Lock()
	s.assemblyLog += fmt.Sprintf("Chassis at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
	mu.Unlock()
}

func (s *AssemblySpot) assembleTires(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	s.vehicleToAssemble.Tires = "Assembled"
	time.Sleep(1 * time.Second)
	mu.Lock()
	s.assemblyLog += fmt.Sprintf("Tires at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
	mu.Unlock()
}

func (s *AssemblySpot) assembleEngine(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	s.vehicleToAssemble.Engine = "Assembled"
	time.Sleep(1 * time.Second)
	mu.Lock()
	s.assemblyLog += fmt.Sprintf("Engine at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
	mu.Unlock()
}

func (s *AssemblySpot) assembleElectronics(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	s.vehicleToAssemble.Electronics = "Assembled"
	time.Sleep(1 * time.Second)
	mu.Lock()
	s.assemblyLog += fmt.Sprintf("Electronics at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
	mu.Unlock()
}

func (s *AssemblySpot) assembleDash(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	s.vehicleToAssemble.Dash = "Assembled"
	time.Sleep(1 * time.Second)
	mu.Lock()
	s.assemblyLog += fmt.Sprintf("Dash at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
	mu.Unlock()
}

func (s *AssemblySpot) assembleSeats(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	s.vehicleToAssemble.Sits = "Assembled"
	time.Sleep(1 * time.Second)
	mu.Lock()
	s.assemblyLog += fmt.Sprintf("Sits at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
	mu.Unlock()
}

func (s *AssemblySpot) assembleWindows(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	s.vehicleToAssemble.Windows = "Assembled"
	time.Sleep(1 * time.Second)
	mu.Lock()
	s.assemblyLog += fmt.Sprintf("Windows at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
	mu.Unlock()
}
