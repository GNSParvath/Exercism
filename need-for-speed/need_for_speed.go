package speed

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	car := Car{speed: speed, batteryDrain: batteryDrain, battery: 100, distance: 0}
	return car
}

type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	track := Track{distance: distance}
	return track
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery-car.batteryDrain < 0 {
		return car
	} else {
		car.distance += car.speed
		car.battery -= car.batteryDrain
		return car
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	var canDrive int = car.battery / car.batteryDrain
	if canDrive*car.speed >= track.distance {
		return true
	}
	return false
}