// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package options

import (
	"time"

	"github.com/GoogleCloudPlatform/esp-v2/src/go/util"
)

// ConfigGeneratorOptions describes the possible overrides for the service config to envoy config translation.
// Note that this rename is difficult because it will break managed api gateway team
type ConfigGeneratorOptions struct {
	CommonOptions

	// Service Management related configurations. Must be set.
	BackendProtocol string

	// Cors related configurations.
	CorsAllowCredentials bool
	CorsAllowHeaders     string
	CorsAllowMethods     string
	CorsAllowOrigin      string
	CorsAllowOriginRegex string
	CorsExposeHeaders    string
	CorsPreset           string

	// Backend routing configurations.
	BackendDnsLookupFamily string

	// Envoy specific configurations.
	ClusterConnectTimeout time.Duration

	// Network related configurations.
	ClusterAddress       string
	ListenerAddress      string
	Healthz              string
	ServiceManagementURL string
	ClusterPort          int
	ListenerPort         int
	RootCertsPath        string

	// Flags for non_gcp deployment.
	ServiceAccountKey string

	// Flags for testing purpose.
	SkipJwtAuthnFilter       bool
	SkipServiceControlFilter bool

	// Envoy configurations.
	EnvoyUseRemoteAddress  bool
	EnvoyXffNumTrustedHops int

	LogJwtPayloads            string
	LogRequestHeaders         string
	LogResponseHeaders        string
	MinStreamReportIntervalMs uint64

	SuppressEnvoyHeaders bool

	ServiceControlNetworkFailOpen bool

	JwksCacheDurationInS int

	ScCheckTimeoutMs  int
	ScQuotaTimeoutMs  int
	ScReportTimeoutMs int

	ScCheckRetries  int
	ScQuotaRetries  int
	ScReportRetries int

	ComputePlatformOverride string
}

// DefaultConfigGeneratorOptions returns ConfigGeneratorOptions with default values.
//
// The default values are expected to match the default values from the flags.
func DefaultConfigGeneratorOptions() ConfigGeneratorOptions {

	return ConfigGeneratorOptions{
		CommonOptions:                 DefaultCommonOptions(),
		BackendDnsLookupFamily:        "auto",
		BackendProtocol:               "", // Required flag with no default
		ClusterAddress:                "127.0.0.1",
		ClusterConnectTimeout:         20 * time.Second,
		ClusterPort:                   8082,
		CorsAllowCredentials:          false,
		CorsAllowHeaders:              "",
		CorsAllowMethods:              "",
		CorsAllowOrigin:               "",
		CorsAllowOriginRegex:          "",
		CorsExposeHeaders:             "",
		CorsPreset:                    "",
		EnvoyUseRemoteAddress:         false,
		EnvoyXffNumTrustedHops:        2,
		JwksCacheDurationInS:          300,
		ListenerAddress:               "0.0.0.0",
		ListenerPort:                  8080,
		RootCertsPath:                 util.DefaultRootCAPaths,
		LogJwtPayloads:                "",
		LogRequestHeaders:             "",
		LogResponseHeaders:            "",
		ServiceAccountKey:             "",
		ServiceControlNetworkFailOpen: true,
		ServiceManagementURL:          "https://servicemanagement.googleapis.com",
		ScCheckRetries:                -1,
		ScCheckTimeoutMs:              0,
		ScQuotaRetries:                -1,
		ScQuotaTimeoutMs:              0,
		ScReportRetries:               -1,
		ScReportTimeoutMs:             0,
		SkipJwtAuthnFilter:            false,
		SkipServiceControlFilter:      false,
		SuppressEnvoyHeaders:          false,
	}
}
