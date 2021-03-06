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

package identity

import (
	"go.uber.org/zap"

	"github.com/fidelity/kconnect/pkg/config"
	"github.com/fidelity/kconnect/pkg/provider"
)

func init() {
	if err := provider.RegisterIdentityProviderPlugin("empty", newEmptyProvider()); err != nil {
		// TODO: handle fatal error
		zap.S().Fatalw("failed to register Empty identity provider plugin", "error", err)
	}
}

func newEmptyProvider() *emptyIdentityProvider {
	return &emptyIdentityProvider{}
}

// EmptyIdentityProvider is an identity provider that does nothing. Its used for testing
type emptyIdentityProvider struct {
}

// Name returns the name of the plugin
func (p *emptyIdentityProvider) Name() string {
	return "Empty Identity Provider"
}

// Flags will return the flags for this plugin
func (p *emptyIdentityProvider) ConfigurationItems(clusterProviderName string) (config.ConfigurationSet, error) {
	return nil, nil
}

// Authenticate will authenticate a user and returns their identity
func (p *emptyIdentityProvider) Authenticate(ctx *provider.Context, clusterProvider string) (provider.Identity, error) {
	return &EmptyIdentity{}, nil
}

// Usage returns a description for use in the help/usage
func (p *emptyIdentityProvider) Usage(clusterProvider string) (string, error) {
	return "", nil
}

// EmptyIdentity is a empty empty returned from the EmptyIdentityProvider
type EmptyIdentity struct {
}

func (i *EmptyIdentity) Type() string {
	return "empty"
}

func (i *EmptyIdentity) Name() string {
	return "empty"
}

func (i *EmptyIdentity) IsExpired() bool {
	return false
}
