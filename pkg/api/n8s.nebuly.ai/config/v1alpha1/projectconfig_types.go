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
	"github.com/nebuly-ai/nebulnetes/pkg/constant"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	cfg "sigs.k8s.io/controller-runtime/pkg/config/v1alpha1"
)

//+kubebuilder:object:root=true

type OperatorConfig struct {
	metav1.TypeMeta                        `json:",inline"`
	cfg.ControllerManagerConfigurationSpec `json:",inline"`
	NvidiaGPUResourceMemoryGB              *int64 `json:"nvidiaGPUResourceMemoryGB,omitempty"`
}

func (c *OperatorConfig) FillDefaultValues() {
	if c.NvidiaGPUResourceMemoryGB == nil {
		var defaultValue int64 = constant.DefaultNvidiaGPUResourceMemory
		c.NvidiaGPUResourceMemoryGB = &defaultValue
	}
}

// +kubebuilder:object:root=true

type GpuPartitionerConfig struct {
	metav1.TypeMeta                        `json:",inline"`
	cfg.ControllerManagerConfigurationSpec `json:",inline"`
}

func init() {
	SchemeBuilder.Register(&OperatorConfig{})
	SchemeBuilder.Register(&GpuPartitionerConfig{})
}
