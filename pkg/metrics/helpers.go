package metrics

import (
	log "github.com/sirupsen/logrus"
)

func MakeJobName(hostname string, nodeOperatorName string) string {
	jobName := hostname
	if nodeOperatorName != "" {
		jobName = nodeOperatorName + "_" + hostname
		log.Info("Using node operator name: ", nodeOperatorName)
	} else {
		log.Info("NODE_OPERATOR_NAME not set, using hostname only for metrics job name")
	}
	return jobName
}
