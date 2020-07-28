/**
 * Copyright 2020 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package auth

import (
	"github.com/IBM/ibmcloud-volume-interface/config"
	"github.com/IBM/ibmcloud-volume-interface/provider/local"
	"github.com/IBM/ibmcloud-volume-vpc/common/iam"
)

// ContextCredentialsFactory ...
type ContextCredentialsFactory struct {
	softlayerConfig      *config.SoftlayerConfig
	vpcConfig            *config.VPCProviderConfig
	tokenExchangeService iam.TokenExchangeService
}

var _ local.ContextCredentialsFactory = &ContextCredentialsFactory{}

// NewContextCredentialsFactory ...
func NewContextCredentialsFactory(bluemixConfig *config.BluemixConfig, softlayerConfig *config.SoftlayerConfig, vpcConfig *config.VPCProviderConfig) (*ContextCredentialsFactory, error) {
	var tokenExchangeService iam.TokenExchangeService
	var err error
	if bluemixConfig.PrivateAPIRoute != "" {
		tokenExchangeService, err = iam.NewTokenExchangeIKSService(bluemixConfig)
	} else {
		tokenExchangeService, err = iam.NewTokenExchangeService(bluemixConfig)
	}
	if err != nil {
		return nil, err
	}

	return &ContextCredentialsFactory{
		softlayerConfig:      softlayerConfig,
		vpcConfig:            vpcConfig,
		tokenExchangeService: tokenExchangeService,
	}, nil
}
