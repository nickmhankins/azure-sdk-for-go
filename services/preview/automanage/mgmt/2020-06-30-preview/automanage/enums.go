package automanage

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ConfigurationProfile enumerates the values for configuration profile.
type ConfigurationProfile string

const (
	// AzurevirtualmachinebestpracticesDevTest ...
	AzurevirtualmachinebestpracticesDevTest ConfigurationProfile = "Azure virtual machine best practices – Dev/Test"
	// AzurevirtualmachinebestpracticesProduction ...
	AzurevirtualmachinebestpracticesProduction ConfigurationProfile = "Azure virtual machine best practices – Production"
)

// PossibleConfigurationProfileValues returns an array of possible values for the ConfigurationProfile const type.
func PossibleConfigurationProfileValues() []ConfigurationProfile {
	return []ConfigurationProfile{AzurevirtualmachinebestpracticesDevTest, AzurevirtualmachinebestpracticesProduction}
}

// EnableRealTimeProtection enumerates the values for enable real time protection.
type EnableRealTimeProtection string

const (
	// False ...
	False EnableRealTimeProtection = "False"
	// True ...
	True EnableRealTimeProtection = "True"
)

// PossibleEnableRealTimeProtectionValues returns an array of possible values for the EnableRealTimeProtection const type.
func PossibleEnableRealTimeProtectionValues() []EnableRealTimeProtection {
	return []EnableRealTimeProtection{False, True}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Created ...
	Created ProvisioningState = "Created"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Created, Failed, Succeeded}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// None ...
	None ResourceIdentityType = "None"
	// SystemAssigned ...
	SystemAssigned ResourceIdentityType = "SystemAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{None, SystemAssigned}
}

// RunScheduledScan enumerates the values for run scheduled scan.
type RunScheduledScan string

const (
	// RunScheduledScanFalse ...
	RunScheduledScanFalse RunScheduledScan = "False"
	// RunScheduledScanTrue ...
	RunScheduledScanTrue RunScheduledScan = "True"
)

// PossibleRunScheduledScanValues returns an array of possible values for the RunScheduledScan const type.
func PossibleRunScheduledScanValues() []RunScheduledScan {
	return []RunScheduledScan{RunScheduledScanFalse, RunScheduledScanTrue}
}

// ScanType enumerates the values for scan type.
type ScanType string

const (
	// Full ...
	Full ScanType = "Full"
	// Quick ...
	Quick ScanType = "Quick"
)

// PossibleScanTypeValues returns an array of possible values for the ScanType const type.
func PossibleScanTypeValues() []ScanType {
	return []ScanType{Full, Quick}
}

// UpdateStatus enumerates the values for update status.
type UpdateStatus string

const (
	// UpdateStatusCreated ...
	UpdateStatusCreated UpdateStatus = "Created"
	// UpdateStatusFailed ...
	UpdateStatusFailed UpdateStatus = "Failed"
	// UpdateStatusSucceeded ...
	UpdateStatusSucceeded UpdateStatus = "Succeeded"
)

// PossibleUpdateStatusValues returns an array of possible values for the UpdateStatus const type.
func PossibleUpdateStatusValues() []UpdateStatus {
	return []UpdateStatus{UpdateStatusCreated, UpdateStatusFailed, UpdateStatusSucceeded}
}
