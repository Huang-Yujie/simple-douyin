// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package constants

const (
	DebugMode                = true
	NoteTableName            = "note"
	UserTableName            = "user"
	SecretKey                = "4631"
	IdentityKey              = "id"
	Total                    = "total"
	ApiServiceName           = "api_service"
	VideoServiceName         = "video_service"
	UserServiceName          = "user_service"
	MySQLDefaultDSN          = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress              = "127.0.0.1:2379"
	CPURateLimit     float64 = 80.0
	DefaultLimit             = 10
	FeedCount                = 5
)
