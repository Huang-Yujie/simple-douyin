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

package rpc

import (
	"context"
	"time"

	// "simple-douyin/kitex_gen/notedemo"
	"simple-douyin/kitex_gen/videoproto"
	"simple-douyin/kitex_gen/videoproto/videoservice"

	// "simple-douyin/kitex_gen/notedemo/noteservice"
	"simple-douyin/pkg/constants"
	"simple-douyin/pkg/errno"
	"simple-douyin/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var videoClient videoservice.Client

func InitVideoRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := videoservice.NewClient(
		constants.NoteServiceName, // 这个没改
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

// CreateNote create note info
func CreateVideo(ctx context.Context, req *videoproto.CreateVideoReq) error {
	resp, err := videoClient.CreateVideo(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return nil
}

// QueryNotes query list of note info
func QueryVideoByUid(ctx context.Context, req *videoproto.GetVideosByUserIdReq) (*videoproto.GetVideosByUserIdResp, error) {
	resp, err := videoClient.GetVideosByUserId(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp, nil
}

func QueryVideoByTime(ctx context.Context, req *videoproto.GetVideosByTimeReq) (*videoproto.GetVideosByTimeResp, error) {
	resp, err := videoClient.GetVideosByTime(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp, nil
}

func LikeVideo(ctx context.Context, req *videoproto.LikeVideoReq) (*videoproto.LikeVideoResp, error) {
	resp, err := videoClient.LikeVideo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp, nil
}

func UnLikeVideo(ctx context.Context, req *videoproto.UnLikeVideoReq) (*videoproto.UnLikeVideoResp, error) {
	resp, err := videoClient.UnLikeVideo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp, nil
}

func CreateComment(ctx context.Context, req *videoproto.CreateCommentReq) (*videoproto.CreateCommentResp, error) {
	resp, err := videoClient.CreateComment(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp, nil
}

func DeleteComment(ctx context.Context, req *videoproto.DeleteCommentReq) (*videoproto.DeleteCommentResp, error) {
	resp, err := videoClient.DeleteComment(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp, nil
}

func GetComments(ctx context.Context, req *videoproto.GetCommentsReq) (*videoproto.GetCommentsResp, error) {
	resp, err := videoClient.GetComments(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp, nil
}

// // UpdateNote update note info
// func UpdateNote(ctx context.Context, req *notedemo.UpdateNoteRequest) error {
// 	resp, err := noteClient.UpdateNote(ctx, req)
// 	if err != nil {
// 		return err
// 	}
// 	if resp.BaseResp.StatusCode != 0 {
// 		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
// 	}
// 	return nil
// }

// // DeleteNote delete note info
// func DeleteNote(ctx context.Context, req *notedemo.DeleteNoteRequest) error {
// 	resp, err := noteClient.DeleteNote(ctx, req)
// 	if err != nil {
// 		return err
// 	}
// 	if resp.BaseResp.StatusCode != 0 {
// 		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
// 	}
// 	return nil
// }
