package controller

import (
	"fmt"
	"golang-distributed-parallel-image-processing/models"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sonyarouje/simdb/db"
	"go.nanomsg.org/mangos"

	// register transports
	"go.nanomsg.org/mangos/protocol/surveyor"
	_ "go.nanomsg.org/mangos/transport/all"
)

func die(format string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Start(controllerAddress string, currentWorkers map[string]models.Worker, db *db.Driver) {
	errorMessage := "[ERR] Controller -> "
	socket, err := surveyor.NewSocket()
	if err != nil {
		die(errorMessage+"Couldn't get a socket connection -> %+v", err.Error())
	}

	err = socket.Listen(controllerAddress)
	if err != nil {
		die(errorMessage+"Couldn't listen on \""+controllerAddress+"\" -> %+v", err.Error())
	}
	err = socket.SetOption(mangos.OptionSurveyTime, time.Second/2)
	if err != nil {
		die(errorMessage+"SetOption(): %+v", err.Error())
	}
	for {
		err = socket.Send([]byte("Is anyone there?"))
		if err != nil {
			die(errorMessage+"The process for looking for workers failed! -> %+v", err.Error())
		}

		for {
			msg, err := socket.Recv()
			if err != nil {
				break
			}
			worker := ParseResponse(string(msg))
			currentWorkers[worker.Name] = worker
		}
	}
}

func ParseResponse(msg string) models.Worker {
	worker := models.Worker{}
	data := strings.Split(msg, "|")
	worker.Name = data[0]
	worker.Status = data[1]
	usage, _ := strconv.Atoi(data[2])
	worker.Usage = usage
	worker.Tags = data[3]
	port, _ := strconv.Atoi(data[4])
	worker.Port = port
	jobsDone, _ := strconv.Atoi(data[5])
	worker.JobsDone = jobsDone
	worker.Active = true
	worker.URL = "localhost:" + data[4]
	return worker
}
