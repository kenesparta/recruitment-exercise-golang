package factory

import (
	".main.go/assemblyspot"
	".main.go/vehicle"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type factoryUnitTestSuite struct {
	suite.Suite
	adapter *Factory
}

func (s *factoryUnitTestSuite) SetupSuite() {

	s.adapter = &Factory{}
}

func TestFactoryUnitTestSuite(t *testing.T) {
	suite.Run(t, &factoryUnitTestSuite{})
}

func (s *factoryUnitTestSuite) TestStartAssemblingProcessNotEmptyChannel() {
	s.adapter = New()
	tests := []struct {
		name  string
		input struct {
			carsAmount        int
			vehiclesAssembled chan vehicle.Car
		}
		output float64
	}{
		{
			name: "First case",
			input: struct {
				carsAmount        int
				vehiclesAssembled chan vehicle.Car
			}{
				carsAmount:        1,
				vehiclesAssembled: make(chan vehicle.Car, 1),
			},
			output: 1.01,
		},
		{
			name: "Zero case",
			input: struct {
				carsAmount        int
				vehiclesAssembled chan vehicle.Car
			}{
				carsAmount:        0,
				vehiclesAssembled: make(chan vehicle.Car, 0),
			},
			output: 0.01,
		},
		{
			name: "Default case",
			input: struct {
				carsAmount        int
				vehiclesAssembled chan vehicle.Car
			}{
				carsAmount:        9,
				vehiclesAssembled: make(chan vehicle.Car, 9),
			},
			output: 2.01,
		},
	}

	for _, tt := range tests {
		start := time.Now()
		s.adapter.StartAssemblingProcess(tt.input.carsAmount, tt.input.vehiclesAssembled)
		for car := 0; car < tt.input.carsAmount; car++ {
			v := <-tt.input.vehiclesAssembled
			s.Assert().NotEqual("", v.Id)
			s.Assert().NotEqual("", v.AssembleLog)
		}
		duration := time.Since(start)
		s.Assert().Condition(func() (success bool) {
			return duration.Seconds() <= tt.output
		}, "")
	}
}

func (s *factoryUnitTestSuite) TestAssembleVehicleLogsNotEmpty() {
	var (
		idleSpot    = assemblyspot.AssemblySpot{}
		vehicleTest = vehicle.Car{
			Id:            0,
			Chassis:       "NotSet",
			Tires:         "NotSet",
			Engine:        "NotSet",
			Electronics:   "NotSet",
			Dash:          "NotSet",
			Sits:          "NotSet",
			Windows:       "NotSet",
			EngineStarted: false,
			TestingLog:    "",
			AssembleLog:   "",
		}
	)
	idleSpot.SetVehicle(&vehicleTest)
	_, err := idleSpot.AssembleVehicle()
	if err != nil {
		s.Fail("Nil Object")
	}
	s.Assert().NotEqual("", idleSpot.GetAssembledLogs())
}
