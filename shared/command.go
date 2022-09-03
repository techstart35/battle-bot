package shared

import "os"

type CMD struct {
	Start       string
	Stop        string
	List        string
	RejectStart string
	Tanaka      string
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
			Tanaka:      "!tanaka",
		}
	default:
		return CMD{
			Start:       "b",
			Stop:        "stopb",
			List:        "listb",
			RejectStart: "rejectstartb",
			Tanaka:      "!tanaka",
		}
	}
}
