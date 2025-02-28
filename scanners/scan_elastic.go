package scanners

import "AvHunt/resources"

type ElasticAgentDetection struct{}

func (w *ElasticAgentDetection) Name() string {
	return "Elastic Endpoint Security"
}

func (w *ElasticAgentDetection) Type() resources.EDRType {
	return resources.ElasticAgentEDR
}

var ElasticAgentHeuristic = []string{
	"Elastic Endpoint Security",
	"Elastic Agent",
	"elastic-agent.exe",
	"elastic-endpoint.exe",
	"elastic-endpoint-driver",
	"ElasticEndpoint",
}

func (w *ElasticAgentDetection) Detect(data resources.SystemData) (resources.EDRType, bool) {
	_, ok := data.CountMatchesAll(ElasticAgentHeuristic)
	if !ok {
		return "", false
	}

	return resources.ElasticAgentEDR, true
}
