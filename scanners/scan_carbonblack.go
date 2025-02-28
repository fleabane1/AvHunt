package scanners

import "AvHunt/resources"

type CarbonBlackDetection struct{}

func (w *CarbonBlackDetection) Name() string {
	return "Carbon Black"
}

func (w *CarbonBlackDetection) Type() resources.EDRType {
	return resources.CarbonBlackEDR
}

var CarbonBlackHeuristic = []string{
	"CarbonBlack\\",
	"CbDefense\\",
	"SensorVersion",
}

func (w *CarbonBlackDetection) Detect(data resources.SystemData) (resources.EDRType, bool) {
	_, ok := data.CountMatchesAll(CarbonBlackHeuristic)
	if !ok {
		return "", false
	}

	return resources.CarbonBlackEDR, true
}
