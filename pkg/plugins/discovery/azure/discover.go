/*
Copyright 2020 The kconnect Authors.

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

package azure

import (
	"fmt"

	"github.com/fidelity/kconnect/pkg/provider"
)

func (p *aksClusterProvider) Discover(ctx *provider.Context, identity provider.Identity) (*provider.DiscoverOutput, error) {
	if err := p.setup(ctx.ConfigurationItems(), identity); err != nil {
		return nil, fmt.Errorf("setting up aks provider: %w", err)
	}

	p.logger.Info("discovering AKS clusters")

	return nil, nil
}