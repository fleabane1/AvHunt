package scanners

import "AvHunt/resources"

type RisingDetection struct{}

func (w *RisingDetection) Name() string {
	return "瑞星杀毒 Defender"
}

func (w *RisingDetection) Type() resources.EDRType {
	return resources.RisingEDR
}

var RisingProcessHeuristic = []string{
	"rsmain",
	"rstray",
	"rstray64",
	"rsupdatertool",
	"rslogup",
}

var RisingServicesHeuristic = []string{
	"rsmain",
	"rstray",
	"rstray64",
	"rsupdatertool",
	"rslogup",
}

var RisingDriverHeuristic = []string{
	"rsmain",
	"rstray",
	"rstray64",
	"rsupdatertool",
	"rslogup",
}

var RisingRegistryHeuristic = []string{
	"rsmain",
	"rstray",
	"rstray64",
	"rsupdatertool",
	"rslogup",
}

func (w *RisingDetection) Detect(data resources.SystemData) (resources.EDRType, bool) {
	_, ok := data.CountMatchesAll(RisingDriverHeuristic, RisingProcessHeuristic, RisingRegistryHeuristic, RisingServicesHeuristic)
	if !ok {
		return "", false
	}

	return resources.RisingEDR, true
}
