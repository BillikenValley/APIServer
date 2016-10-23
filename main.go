package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"

	"github.com/BillikenValley/APIServer/model"
)

func main() {

	// router := NewRouter()
	//
	// log.Fatal(http.ListenAndServe(":8080", router))
	shelters := make([]model.Shelter, 50)
	for i := range shelters {
		shelters[i] = model.RandomShelter()
	}
	requesters := make([]model.Requester, 3)
	for i := range requesters {
		requesters[i] = model.RandomRequester()
	}
	binary, lookErr := exec.LookPath("python")
	if lookErr != nil {
		panic(lookErr)
	}

	sheltersJson, err := json.Marshal(shelters)
	var sheltersPretty bytes.Buffer
	checkErr(err)
	sf, err := ioutil.TempFile("", "")
	checkErr(err)
	// defer os.Remove(sf.Name()) // clean up
	sf.Write(sheltersJson)
	sf.Close()

	json.Indent(&sheltersPretty, sheltersJson, "", "\t")
	requesterJson, err := json.Marshal(requesters)
	checkErr(err)
	rf, err := ioutil.TempFile("", "")
	checkErr(err)
	// defer os.Remove(rf.Name()) // clean up
	rf.Write(requesterJson)
	rf.Close()

	locatorJson, err := json.Marshal(shelters[0].Location)
	checkErr(err)
	lf, err := ioutil.TempFile("", "")
	checkErr(err)
	// defer os.Remove(lf.Name()) // clean up
	lf.Write(locatorJson)
	lf.Close()
	args := []string{"python", "./BackendServices/plot.py", "-s", sf.Name(), "-u", rf.Name(), "-l", lf.Name()}
	env := os.Environ()

	fmt.Println(args)
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
