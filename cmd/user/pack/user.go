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

package pack

import (
	"simple-douyin/cmd/user/dal/model"
	"simple-douyin/kitex_gen/userproto"
)

// User pack user info
func User(u *model.User) *userproto.UserInfo {
	if u == nil {
		return nil
	}

	return &userproto.UserInfo{UserId: int64(u.UserID), Username: u.UserName}
}

// Users pack list of user info
func Users(us []*model.User) []*userproto.UserInfo {
	users := make([]*userproto.UserInfo, 0)
	for _, u := range us {
		if user2 := User(u); user2 != nil {
			users = append(users, user2)
		}
	}
	return users
}
