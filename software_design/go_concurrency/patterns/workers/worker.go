package workers

//Useful to control the number of goRoutines in an execution
// we CAN NOT allow the operativesystem creates
// an unlimited number or goRoutines it could blow up !
// Goroutines are light, but the work they perform could be very heavy. A workers pool helps
// us to solve this problem. It hels to:
//Control access to shared resources using quotas
//Create a limited amount of Goroutines per app
//Provide more parallelism capabilities to other concurrent structures


