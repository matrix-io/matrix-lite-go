package matrix

import (
	"fmt"
	"testing"
)

// UPDATE SENSOR
// BenchmarkSensor-4   	   10000	    133304 ns/op
// PASS
// ok  	github.com/matrix-io/matrix-lite-go	1.378s

// NEW SENSOR
// BenchmarkSensor-4   	   10000	    134916 ns/op
// PASS
// ok  	github.com/matrix-io/matrix-lite-go	1.384s

// TestMain creates a file for testing
func TestMain(m *testing.M) {
	Init()
	m.Run()
}

func BenchmarkSensorNew(b *testing.B) {

	for i := 0; i < b.N; i++ {
		//Print values
		fmt.Println("IMU", ReadIMU())
		fmt.Println("UV", ReadUV())
		fmt.Println("Humidity", ReadPressure())
		fmt.Println("Pressure", ReadHumidity())
	}
}
