package utils

import "fmt"

// Global variables
var RRenum = "RoundRobin"
var LCenum = "LeastConnections"

// if an unknown error occurs here then panic will stop the control flow, if checkNilErr has to be removed then adequate return should be added
func CheckNilErr(err error, st string) {

	if err != nil {
		fmt.Println(st)
		panic(err) // can later be replaced with safe handles
	}
}
