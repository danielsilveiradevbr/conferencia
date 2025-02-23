package statusService

type Status struct {
	Status  string
	Version string
	Uptime  string
}

func GetStatus() *Status {
	return &Status{
		Status:  "running",
		Version: "1.0.0",
		Uptime:  "24h",
	}

}
