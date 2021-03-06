/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package env

import (
	"strconv"
	"time"

	"knative.dev/pkg/system"

	"go.uber.org/zap"
	"knative.dev/eventing-kafka/pkg/channel/distributed/common/env"
	"knative.dev/pkg/controller"
)

// The Environment Struct
type Environment struct {

	// Eventing Configuration
	SystemNamespace string // Required

	// Metrics Configuration
	MetricsPort   int    // Required
	MetricsDomain string // Required

	// Pod information to be used by the metrics reporter
	PodName       string // Required
	ContainerName string // Required

	// Health Configuration
	HealthPort int // Required

	// Kafka Configuration
	ServiceName  string        // Required
	ResyncPeriod time.Duration // Optional

	// Kafka Authorization
	KafkaSecretName      string // Required
	KafkaSecretNamespace string // Required
}

// Get The Environment
func GetEnvironment(logger *zap.Logger) (*Environment, error) {

	// Error Reference
	var err error

	// The Environment Struct To Be Populated
	environment := &Environment{}

	// Get The Required System Namespace Config Value
	environment.SystemNamespace, err = env.GetRequiredConfigValue(logger, system.NamespaceEnvKey)
	if err != nil {
		return nil, err
	}

	// Get The Required Metrics Port Config Value & Convert To Int
	environment.MetricsPort, err = env.GetRequiredConfigInt(logger, env.MetricsPortEnvVarKey, "MetricsPort")
	if err != nil {
		return nil, err
	}

	// Get The Required Metrics Domain Config Value
	environment.MetricsDomain, err = env.GetRequiredConfigValue(logger, env.MetricsDomainEnvVarKey)
	if err != nil {
		return nil, err
	}

	// Get The Required PodName Config Value
	environment.PodName, err = env.GetRequiredConfigValue(logger, env.PodNameEnvVarKey)
	if err != nil {
		return nil, err
	}

	// Get The Required ContainerName Config Value
	environment.ContainerName, err = env.GetRequiredConfigValue(logger, env.ContainerNameEnvVarKey)
	if err != nil {
		return nil, err
	}

	// Get The Required HealthPort Port Config Value & Convert To Int
	environment.HealthPort, err = env.GetRequiredConfigInt(logger, env.HealthPortEnvVarKey, "HealthPort")
	if err != nil {
		return nil, err
	}

	// Get The Required K8S KafkaSecretNamespace Config Value
	environment.KafkaSecretNamespace, err = env.GetRequiredConfigValue(logger, env.KafkaSecretNamespaceEnvVarKey)
	if err != nil {
		return nil, err
	}

	// Get The Required K8S KafkaSecretName Config Value
	environment.KafkaSecretName, err = env.GetRequiredConfigValue(logger, env.KafkaSecretNameEnvVarKey)
	if err != nil {
		return nil, err
	}

	// Get The Required K8S ServiceName Config Value
	environment.ServiceName, err = env.GetRequiredConfigValue(logger, env.ServiceNameEnvVarKey)
	if err != nil {
		return nil, err
	}

	// Get The Optional ResyncPeriod Config Value & Convert To Duration
	resyncMinutes, err := env.GetOptionalConfigInt(
		logger,
		env.ResyncPeriodMinutesEnvVarKey,
		strconv.Itoa(int(controller.DefaultResyncPeriod/time.Minute)),
		"ResyncPeriodMinutes")
	if err != nil {
		return nil, err
	}
	environment.ResyncPeriod = time.Duration(resyncMinutes) * time.Minute

	// Log The Receiver Configuration Loaded From Environment Variables
	logger.Info("Environment Variables", zap.Any("Environment", environment))

	// Return The Populated Environment Structure
	return environment, nil
}
