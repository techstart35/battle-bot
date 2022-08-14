package shared

import "os"

type CMD struct {
	Start       string
	Stop        string
	Process     string
	RejectStart string
}

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
	default:
		return CMD{
			Start:       "b",
			Stop:        "stopb",
			Process:     "processb",
			RejectStart: "rejectstartb",
		}
	}
}
