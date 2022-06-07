package rpc

import (
	"context"
	"simple-douyin/kitex_gen/videoproto"
	"simple-douyin/kitex_gen/videoproto/videoservice"
	"simple-douyin/pkg/config"
	"simple-douyin/pkg/constants"
	"simple-douyin/pkg/errno"
	"simple-douyin/pkg/middleware"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var videoClient videoservice.Client

func initVideoRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.Server.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := videoservice.NewClient(
		constants.VideoServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(time.Minute),                // rpc timeout
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

func GetVideosByUserId(ctx context.Context, req *videoproto.GetVideosByUserIdReq) ([]*videoproto.VideoInfo, error) {
	resp, err := videoClient.GetVideosByUserId(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.VideoInfos, nil
}

func GetVideosByTime(ctx context.Context, req *videoproto.GetVideosByTimeReq) ([]*videoproto.VideoInfo, int64, error) {
	resp, err := videoClient.GetVideosByTime(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.VideoInfos, resp.NextTime, nil
}

func LikeVideo(ctx context.Context, req *videoproto.LikeVideoReq) error {
	resp, err := videoClient.LikeVideo(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return nil
}

func UnLikeVideo(ctx context.Context, req *videoproto.UnLikeVideoReq) error {
	resp, err := videoClient.UnLikeVideo(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return nil
}

func GetLikeVideos(ctx context.Context, req *videoproto.GetLikeVideosReq) ([]*videoproto.VideoInfo, error) {
	resp, err := videoClient.GetLikeVideos(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.VideoInfos, nil
}

func CreateComment(ctx context.Context, req *videoproto.CreateCommentReq) (*videoproto.CommentInfo, error) {
	resp, err := videoClient.CreateComment(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.CommentInfo, nil
}

func DeleteComment(ctx context.Context, req *videoproto.DeleteCommentReq) error {
	resp, err := videoClient.DeleteComment(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return nil
}

func GetComments(ctx context.Context, req *videoproto.GetCommentsReq) ([]*videoproto.CommentInfo, error) {
	resp, err := videoClient.GetComments(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.CommentInfos, nil
}
