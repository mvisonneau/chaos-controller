// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2020 Datadog, Inc.

package v1beta1

import (
	chaostypes "github.com/DataDog/chaos-controller/types"
)

// CPUPressureSpec represents a cpu pressure disruption
type CPUPressureSpec struct {
}

// Validate validates args for the given disruption
func (s *CPUPressureSpec) Validate() error {
	return nil
}

// GenerateArgs generates injection or cleanup pod arguments for the given spec
func (s *CPUPressureSpec) GenerateArgs(level chaostypes.DisruptionLevel, containerID, sink string, dryRun bool) []string {
	args := []string{
		"cpu-pressure",
		"--metrics-sink",
		sink,
		"--level",
		string(level),
		"--container-id",
		containerID,
	}

	// enable dry-run mode
	if dryRun {
		args = append(args, "--dry-run")
	}

	return args
}
