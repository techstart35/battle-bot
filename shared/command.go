package shared

import "os"

type CMD struct {
	Start       string
	Stop        string
	List        string
	RejectStart string
	Start5Min   string
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
			Start5Min:   "!b5",
		}
	default:
		return CMD{
			Start:       "b",
			Stop:        "stopb",
			List:        "listb",
			RejectStart: "rejectstartb",
			Start5Min:   "b5",
		}
	}
}
