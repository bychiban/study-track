package main

func main() {

}

type checkIn struct {
	stationName string
	time        int
}

type averageTime struct {
	count   int
	sumTime int
}

type UndergroundSystem struct {
	carsInTransfer map[int]*checkIn
	stationAverage map[string]map[string]*averageTime
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{make(map[int]*checkIn), make(map[string]map[string]*averageTime)}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	if _, ok := this.stationAverage[stationName]; !ok {
		this.stationAverage[stationName] = make(map[string]*averageTime)
	}
	this.carsInTransfer[id] = &checkIn{stationName, t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	startStation, _ := this.stationAverage[this.carsInTransfer[id].stationName]
	endStation, ok := startStation[stationName]
	if !ok {
		startStation[stationName] = &averageTime{}
		endStation = startStation[stationName]
	}
	endStation.count++
	endStation.sumTime += t - this.carsInTransfer[id].time
	delete(this.carsInTransfer, id)
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	averageTime := this.stationAverage[startStation][endStation]
	return float64(averageTime.sumTime) / float64(averageTime.count)
}
