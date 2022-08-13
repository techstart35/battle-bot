package shared

import "os"

type CMD struct {
	Start       string
	Stop        string
	Process     string
	RejectStart string
}

//const (
//	Command            = "b"
//	StopCommand        = "stopb"
//	ProcessCommand     = "processb"
//	RejectStartCommand = "rejectstartb"
//)

func Command() CMD {
	env := os.Getenv("ENV")

	switch env {
	case "dev":
		return CMD{
			Start:       "!b",
			Stop:        "!stopb",
			Process:     "!processb",
			RejectStart: "!rejectstartb",
		}
	case "prd":
		return CMD{
			Start:       "b",
			Stop:        "stopb",
			Process:     "processb",
			RejectStart: "rejectstartb",
		}
	}

	return CMD{}
}
