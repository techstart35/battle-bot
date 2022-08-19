package shared

import "os"

type CMD struct {
	Start       string
	Stop        string
	List        string
	RejectStart string
}

func Command() CMD {
	env := os.Getenv("ENV")

	switch env {
	case "dev":
		return CMD{
			Start:       "!b",
			Stop:        "!stopb",
			List:        "!listb",
			RejectStart: "!rejectstartb",
		}
	default:
		return CMD{
			Start:       "b",
			Stop:        "stopb",
			List:        "listb",
			RejectStart: "rejectstartb",
		}
	}
}
