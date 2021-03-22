package erlang

import (
	"math"
)

// Calculate traffic intensity (a.k.a. workload)
// arrivalRate: Number of arrival Calls per interval
// avgHandleTime: Average handling time in seconds
// intervalLength: Length of an interval in minutes
// return: Traffic intensity in Erlangs
func Intensity(arrivalRate int, avgHandleTime int, intervalLength int) float64 {
	var erlangunit float64 = (float64(arrivalRate) / (60 * float64(intervalLength))) * float64(avgHandleTime)
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
//intervalLength: Length of an interval in minutes
//return: Average waiting time in second
func AvgWaitTime(numberOfAgents int, arrivalRate int, avgHandleTime int, intervalLength int) float64 {
	var intensity float64 = Intensity(arrivalRate, avgHandleTime, intervalLength)
	var awt float64 = (ErlangC(numberOfAgents, intensity) * float64(avgHandleTime)) / (float64(numberOfAgents) - intensity)
	return awt
}

//Calculate the servicelevel
//Calculates the percentage of calls that are answered within the acceptable waiting time
//
//numberOfAgents: Number of available agents
//arrivalRate: Number of arrivals per interval
//avgHandleTime: Average handling time in seconds
//intervalLength: Length of an interval in minutes
//waitTime: Acceptable waiting time
//return: Service level (% of calls answered within acceptable waiting time)
func ServiceLevel(numberOfAgents int, arrivalRate int, avgHandleTime int, intervalLength int, waitTime int) float64 {
	var a float64 = Intensity(arrivalRate, avgHandleTime, intervalLength)
	var sl float64 = 1 - ErlangC(numberOfAgents, a)*math.Exp(-(float64(numberOfAgents)-a)*(float64(waitTime)/float64(avgHandleTime)))
	return sl
}

// Calculate the number of needed agents for SL goal
// Calculates the number of agents that are needed to achieve a required service level. Currently only calculates a whole (integer) number of agents.
//
//arrivalRate: Number of arrivals per interval
//avgHandleTime: Average handling time in seconds
//intervalLength: Length of an interval in minutes
//waitTime: Acceptable waiting time (i.e. 20 seconds in a 80/20 SL goal)
//serviceLevelGoal: Service level goal, the percentage of calls answered within the acceptable waiting time
//return: Number of agents needed to achieve service level
func NumberOfAgentsForSl(arrivalRate int, avgHandleTime int, intervalLength int, waitTime int, serviceLevelGoal float64) int {
	var intensity float64 = Intensity(arrivalRate, avgHandleTime, intervalLength)
	var agents int = int(math.Ceil(intensity))
	for ServiceLevel(agents, arrivalRate, avgHandleTime, intervalLength, waitTime) < serviceLevelGoal {
		agents = agents + 1
	}
	return agents
}

//Calculate the number of needed agents to achieve an ASA goal
//Calculates the number of agents that are needed to achieve a required average speed of answer. Currently only calculates a whole (integer) number of agents.
//
//arrivalRate: Number of arrivals per interval
//avgHandleTime: Average handling time in seconds
//intervalLength: Length of an interval in minutes
//waitTime: Waiting time goal in seconds
//return: Number of agents needed to achieve ASA
func NumberOfAgentsForAsa(arrivalRate int, avgHandleTime int, intervalLength int, waitTime int) int {
	var intensity float64 = Intensity(arrivalRate, avgHandleTime, intervalLength)
	var agents int = int(math.Ceil(intensity))
	for AvgWaitTime(agents, arrivalRate, avgHandleTime, intervalLength) > float64(waitTime) {
		agents = agents + 1
	}
	return agents
}
