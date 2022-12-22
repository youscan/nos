/*
 * Copyright 2022 Nebuly.ai
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1alpha1

import (
	"errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	cfg "sigs.k8s.io/controller-runtime/pkg/config/v1alpha1"
	"time"
)

// +kubebuilder:object:root=true

type GpuPartitionerConfig struct {
	metav1.TypeMeta                        `json:",inline"`
	cfg.ControllerManagerConfigurationSpec `json:",inline"`
	SchedulerConfigFile                    string           `json:"schedulerConfigFile,omitempty"`
	KnownMigGeometriesFile                 string           `json:"knownMigGeometriesFile,omitempty"`
	BatchWindowTimeoutSeconds              time.Duration    `json:"batchWindowTimeoutSeconds"`
	BatchWindowIdleSeconds                 time.Duration    `json:"batchWindowIdleSeconds"`
	NvidiaDevicePluginConfigMap            NamespacedObject `json:"devicePluginConfigMap,omitempty"`
}

func (c *GpuPartitionerConfig) Validate() error {
	if c.BatchWindowTimeoutSeconds.Seconds() <= 0 {
		return errors.New("batchWindowTimeoutSeconds must be greater than 0")
	}
	if c.BatchWindowIdleSeconds.Seconds() <= 0 {
		return errors.New("batchWindowIdleSeconds must be greater than 0")
	}
	return nil
}

type NamespacedObject struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}
