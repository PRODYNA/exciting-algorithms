package main

import (
	"fmt"
	"math"
)

func main() {

	var int float64 = Intensity(200, 180)
	fmt.Printf("Intensity:  %f \n", int)

	var occupancy float64 = Occupancy(23, 20)
	fmt.Printf("Occupancy:  %f  \n", occupancy)

	var erlangB float64 = ErlangB(200, 180)
	fmt.Printf("Erlang B:  %f  \n", erlangB)

	var erlangC float64 = ErlangC(25, 20)
	fmt.Printf("Erlang C:  %f \n", erlangC)

	var serviceLevel float64 = ServiceLevel(25, 200, 180, 60)
	fmt.Printf("Service Level:  %f \n", serviceLevel)

	var avgWaitTime float64 = AvgWaitTime(25, 200, 180)
	fmt.Printf("Avg Wait Time:  %f  \n", avgWaitTime)

	var numberOfAgentsForSl = NumberOfAgentsForSl(200, 180, 60, 0.8)
	fmt.Printf("Number of Agents for SL:  %d  \n", numberOfAgentsForSl)

	var numberOfAgentsForAsa = NumberOfAgentsForAsa(200, 180, 120)
	fmt.Printf("Number of Agents for SL:  %d \n", numberOfAgentsForAsa)
}

// Calculate traffic intensity (a.k.a. workload)
// arrivalRate: Number of arrival Calls per interval
// avgHandleTime: Average handling time in seconds
// return: Traffic intensity in Erlangs
func Intensity(arrivalRate int, avgHandleTime int) float64 {
	var erlangunit float64 = (float64(arrivalRate) * float64(avgHandleTime)) / 60
	return erlangunit
}

// Calculate agent occupancy
// numberOfAgents: Number of available agents
// intensity: Traffic intensity in Erlangs
// return: Occupancy
func Occupancy(numberOfAgents int, intensity int) float64 {
	return float64(intensity) / float64(numberOfAgents)
}

// Calculate the chance of a blocked call (Erlang B function)
// numberOfServers : Number of available agents
// intensity: Traffic intensity in Erlang Units
// return: Chance of blocking
func ErlangB(numberOfServers int, intensity float64) float64 {
	var InvBlock float64 = 1

	for i := 1; i < numberOfServers; i++ {
		InvBlock = float64(i)/intensity*InvBlock + 1
	}
	return 1 / InvBlock
}

// Calculate the chance of a queued call (Erlang C function
// numberOfAgents: Number of available agents
// intensity: Traffic intensity in Erlangs
// return: Chance of queueing
func ErlangC(numberOfAgents int, intensity float64) float64 {
	var ec float64 = float64(numberOfAgents) * ErlangB(numberOfAgents, intensity) / (float64(numberOfAgents) - intensity*(1-ErlangB(numberOfAgents, intensity)))
	return ec
}

// Calculate the average waiting time
//
//numberOfAgents: Number of available agents
//arrivalRate: Number of arrivals per interval
//avgHandleTime: Average handling time in seconds
//return: Average waiting time in second
func AvgWaitTime(numberOfAgents int, arrivalRate int, avgHandleTime int) float64 {
	var intensity float64 = Intensity(arrivalRate, avgHandleTime)
	var awt float64 = (ErlangC(numberOfAgents, intensity) * float64(avgHandleTime)) / (float64(numberOfAgents) - intensity)
	return awt
}

//Calculate the servicelevel
//Calculates the percentage of calls that are answered within the acceptable waiting time
//
//numberOfAgents: Number of available agents
//arrivalRate: Number of arrivals per interval
//avgHandleTime: Average handling time in seconds
//waitTime: Acceptable waiting time
//return: Service level (% of calls answered within acceptable waiting time)
func ServiceLevel(numberOfAgents int, arrivalRate int, avgHandleTime int, waitTime int) float64 {
	var a float64 = Intensity(arrivalRate, avgHandleTime)
	var sl float64 = 1 - ErlangC(numberOfAgents, a)*math.Exp(-(float64(numberOfAgents)-a)*(float64(waitTime)/float64(avgHandleTime)))
	return sl
}

// Calculate the number of needed agents for SL goal
// Calculates the number of agents that are needed to achieve a required service level. Currently only calculates a whole (integer) number of agents.
//
//arrivalRate: Number of arrivals per interval
//avgHandleTime: Average handling time in seconds
//waitTime: Acceptable waiting time (i.e. 20 seconds in a 80/20 SL goal)
//serviceLevelGoal: Service level goal, the percentage of calls answered within the acceptable waiting time
//return: Number of agents needed to achieve service level
func NumberOfAgentsForSl(arrivalRate int, avgHandleTime int, waitTime int, serviceLevelGoal float64) int {
	var intensity float64 = Intensity(arrivalRate, avgHandleTime)
	var agents int = int(math.Ceil(intensity))
	for ServiceLevel(agents, arrivalRate, avgHandleTime, waitTime) < serviceLevelGoal {
		agents = agents + 1
	}
	return agents
}

//Calculate the number of needed agents to achieve an ASA goal
//Calculates the number of agents that are needed to achieve a required average speed of answer. Currently only calculates a whole (integer) number of agents.
//
//arrivalRate: Number of arrivals per interval
//avgHandleTime: Average handling time in seconds
//waitTime: Waiting time goal in seconds
//return: Number of agents needed to achieve ASA
func NumberOfAgentsForAsa(arrivalRate int, avgHandleTime int, waitTime int) int {
	var intensity float64 = Intensity(arrivalRate, avgHandleTime)
	var agents int = int(math.Ceil(intensity))
	for AvgWaitTime(agents, arrivalRate, avgHandleTime) > float64(waitTime) {
		agents = agents + 1
	}
	return agents
}
