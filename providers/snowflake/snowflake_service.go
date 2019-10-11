// Copyright 2019 The Terraformer Authors.
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

package snowflake

import (
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/snowflakedb/gosnowflake"

	"database/sql"
)

type SnowflakeService struct {
	terraform_utils.Service
}

func (s *SnowflakeService) generateService() (*client, error) {
	account := s.Args["account"].(string)
	username := s.Args["username"].(string)
	region := s.Args["region"].(string)
	role := s.Args["role"].(string)
	dsn, err := gosnowflake.DSN(&gosnowflake.Config{
		Account:       account,
		User:          username,
		Region:        region,
		Role:          role,
		Authenticator: gosnowflake.AuthTypeExternalBrowser,
	})
	db, err := sql.Open("snowflake-provider", dsn)
	return &client{db: db}, err
}
