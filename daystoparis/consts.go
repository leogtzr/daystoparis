package daystoparis

const (
	requiredNumberOfArgs           = 3
	dateFormat                     = "2006-01-02"
	errorBuildingAppConfigFromArgs = 2
	errorRunningApplication        = 1
)

var cmdArgs = [6]string{"-f", "smblock", "-w", "900", "-F", "border"}
