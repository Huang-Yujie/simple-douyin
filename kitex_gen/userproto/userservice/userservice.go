// Code generated by Kitex v0.3.1. DO NOT EDIT.

package userservice

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"google.golang.org/protobuf/proto"
	"simple-douyin/kitex_gen/userproto"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*userproto.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateUser":    kitex.NewMethodInfo(createUserHandler, newCreateUserArgs, newCreateUserResult, false),
		"GetUser":       kitex.NewMethodInfo(getUserHandler, newGetUserArgs, newGetUserResult, false),
		"CheckUser":     kitex.NewMethodInfo(checkUserHandler, newCheckUserArgs, newCheckUserResult, false),
		"FollowUser":    kitex.NewMethodInfo(followUserHandler, newFollowUserArgs, newFollowUserResult, false),
		"UnFollowUser":  kitex.NewMethodInfo(unFollowUserHandler, newUnFollowUserArgs, newUnFollowUserResult, false),
		"GetFollowList": kitex.NewMethodInfo(getFollowListHandler, newGetFollowListArgs, newGetFollowListResult, false),
		"GetFanList":    kitex.NewMethodInfo(getFanListHandler, newGetFanListArgs, newGetFanListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.3.1",
		Extra:           extra,
	}
	return svcInfo
}

func createUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userproto.CreateUserReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userproto.UserService).CreateUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateUserArgs:
		success, err := handler.(userproto.UserService).CreateUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateUserResult)
		realResult.Success = success
	}
	return nil
}
func newCreateUserArgs() interface{} {
	return &CreateUserArgs{}
}

func newCreateUserResult() interface{} {
	return &CreateUserResult{}
}

type CreateUserArgs struct {
	Req *userproto.CreateUserReq
}

func (p *CreateUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CreateUserArgs) Unmarshal(in []byte) error {
	msg := new(userproto.CreateUserReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateUserArgs_Req_DEFAULT *userproto.CreateUserReq

func (p *CreateUserArgs) GetReq() *userproto.CreateUserReq {
	if !p.IsSetReq() {
		return CreateUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type CreateUserResult struct {
	Success *userproto.CreateUserResp
}

var CreateUserResult_Success_DEFAULT *userproto.CreateUserResp

func (p *CreateUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CreateUserResult) Unmarshal(in []byte) error {
	msg := new(userproto.CreateUserResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateUserResult) GetSuccess() *userproto.CreateUserResp {
	if !p.IsSetSuccess() {
		return CreateUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*userproto.CreateUserResp)
}

func (p *CreateUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userproto.GetUserReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userproto.UserService).GetUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetUserArgs:
		success, err := handler.(userproto.UserService).GetUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserResult)
		realResult.Success = success
	}
	return nil
}
func newGetUserArgs() interface{} {
	return &GetUserArgs{}
}

func newGetUserResult() interface{} {
	return &GetUserResult{}
}

type GetUserArgs struct {
	Req *userproto.GetUserReq
}

func (p *GetUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserArgs) Unmarshal(in []byte) error {
	msg := new(userproto.GetUserReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserArgs_Req_DEFAULT *userproto.GetUserReq

func (p *GetUserArgs) GetReq() *userproto.GetUserReq {
	if !p.IsSetReq() {
		return GetUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUserResult struct {
	Success *userproto.GetUserResp
}

var GetUserResult_Success_DEFAULT *userproto.GetUserResp

func (p *GetUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserResult) Unmarshal(in []byte) error {
	msg := new(userproto.GetUserResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserResult) GetSuccess() *userproto.GetUserResp {
	if !p.IsSetSuccess() {
		return GetUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*userproto.GetUserResp)
}

func (p *GetUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func checkUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userproto.CheckUserReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userproto.UserService).CheckUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CheckUserArgs:
		success, err := handler.(userproto.UserService).CheckUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CheckUserResult)
		realResult.Success = success
	}
	return nil
}
func newCheckUserArgs() interface{} {
	return &CheckUserArgs{}
}

func newCheckUserResult() interface{} {
	return &CheckUserResult{}
}

type CheckUserArgs struct {
	Req *userproto.CheckUserReq
}

func (p *CheckUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CheckUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CheckUserArgs) Unmarshal(in []byte) error {
	msg := new(userproto.CheckUserReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CheckUserArgs_Req_DEFAULT *userproto.CheckUserReq

func (p *CheckUserArgs) GetReq() *userproto.CheckUserReq {
	if !p.IsSetReq() {
		return CheckUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type CheckUserResult struct {
	Success *userproto.CheckUserResp
}

var CheckUserResult_Success_DEFAULT *userproto.CheckUserResp

func (p *CheckUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CheckUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CheckUserResult) Unmarshal(in []byte) error {
	msg := new(userproto.CheckUserResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckUserResult) GetSuccess() *userproto.CheckUserResp {
	if !p.IsSetSuccess() {
		return CheckUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*userproto.CheckUserResp)
}

func (p *CheckUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func followUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userproto.FollowUserReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userproto.UserService).FollowUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FollowUserArgs:
		success, err := handler.(userproto.UserService).FollowUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FollowUserResult)
		realResult.Success = success
	}
	return nil
}
func newFollowUserArgs() interface{} {
	return &FollowUserArgs{}
}

func newFollowUserResult() interface{} {
	return &FollowUserResult{}
}

type FollowUserArgs struct {
	Req *userproto.FollowUserReq
}

func (p *FollowUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FollowUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FollowUserArgs) Unmarshal(in []byte) error {
	msg := new(userproto.FollowUserReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FollowUserArgs_Req_DEFAULT *userproto.FollowUserReq

func (p *FollowUserArgs) GetReq() *userproto.FollowUserReq {
	if !p.IsSetReq() {
		return FollowUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FollowUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type FollowUserResult struct {
	Success *userproto.FollowUserResp
}

var FollowUserResult_Success_DEFAULT *userproto.FollowUserResp

func (p *FollowUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FollowUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FollowUserResult) Unmarshal(in []byte) error {
	msg := new(userproto.FollowUserResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FollowUserResult) GetSuccess() *userproto.FollowUserResp {
	if !p.IsSetSuccess() {
		return FollowUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FollowUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*userproto.FollowUserResp)
}

func (p *FollowUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func unFollowUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userproto.UnFollowUserReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userproto.UserService).UnFollowUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UnFollowUserArgs:
		success, err := handler.(userproto.UserService).UnFollowUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UnFollowUserResult)
		realResult.Success = success
	}
	return nil
}
func newUnFollowUserArgs() interface{} {
	return &UnFollowUserArgs{}
}

func newUnFollowUserResult() interface{} {
	return &UnFollowUserResult{}
}

type UnFollowUserArgs struct {
	Req *userproto.UnFollowUserReq
}

func (p *UnFollowUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UnFollowUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UnFollowUserArgs) Unmarshal(in []byte) error {
	msg := new(userproto.UnFollowUserReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UnFollowUserArgs_Req_DEFAULT *userproto.UnFollowUserReq

func (p *UnFollowUserArgs) GetReq() *userproto.UnFollowUserReq {
	if !p.IsSetReq() {
		return UnFollowUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UnFollowUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type UnFollowUserResult struct {
	Success *userproto.UnFollowUserResp
}

var UnFollowUserResult_Success_DEFAULT *userproto.UnFollowUserResp

func (p *UnFollowUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UnFollowUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UnFollowUserResult) Unmarshal(in []byte) error {
	msg := new(userproto.UnFollowUserResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UnFollowUserResult) GetSuccess() *userproto.UnFollowUserResp {
	if !p.IsSetSuccess() {
		return UnFollowUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UnFollowUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*userproto.UnFollowUserResp)
}

func (p *UnFollowUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getFollowListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userproto.GetFollowListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userproto.UserService).GetFollowList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetFollowListArgs:
		success, err := handler.(userproto.UserService).GetFollowList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetFollowListResult)
		realResult.Success = success
	}
	return nil
}
func newGetFollowListArgs() interface{} {
	return &GetFollowListArgs{}
}

func newGetFollowListResult() interface{} {
	return &GetFollowListResult{}
}

type GetFollowListArgs struct {
	Req *userproto.GetFollowListReq
}

func (p *GetFollowListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetFollowListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetFollowListArgs) Unmarshal(in []byte) error {
	msg := new(userproto.GetFollowListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetFollowListArgs_Req_DEFAULT *userproto.GetFollowListReq

func (p *GetFollowListArgs) GetReq() *userproto.GetFollowListReq {
	if !p.IsSetReq() {
		return GetFollowListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetFollowListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetFollowListResult struct {
	Success *userproto.GetFollowListResp
}

var GetFollowListResult_Success_DEFAULT *userproto.GetFollowListResp

func (p *GetFollowListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetFollowListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetFollowListResult) Unmarshal(in []byte) error {
	msg := new(userproto.GetFollowListResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetFollowListResult) GetSuccess() *userproto.GetFollowListResp {
	if !p.IsSetSuccess() {
		return GetFollowListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetFollowListResult) SetSuccess(x interface{}) {
	p.Success = x.(*userproto.GetFollowListResp)
}

func (p *GetFollowListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getFanListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userproto.GetFanListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userproto.UserService).GetFanList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetFanListArgs:
		success, err := handler.(userproto.UserService).GetFanList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetFanListResult)
		realResult.Success = success
	}
	return nil
}
func newGetFanListArgs() interface{} {
	return &GetFanListArgs{}
}

func newGetFanListResult() interface{} {
	return &GetFanListResult{}
}

type GetFanListArgs struct {
	Req *userproto.GetFanListReq
}

func (p *GetFanListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetFanListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetFanListArgs) Unmarshal(in []byte) error {
	msg := new(userproto.GetFanListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetFanListArgs_Req_DEFAULT *userproto.GetFanListReq

func (p *GetFanListArgs) GetReq() *userproto.GetFanListReq {
	if !p.IsSetReq() {
		return GetFanListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetFanListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetFanListResult struct {
	Success *userproto.GetFanListResp
}

var GetFanListResult_Success_DEFAULT *userproto.GetFanListResp

func (p *GetFanListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetFanListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetFanListResult) Unmarshal(in []byte) error {
	msg := new(userproto.GetFanListResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetFanListResult) GetSuccess() *userproto.GetFanListResp {
	if !p.IsSetSuccess() {
		return GetFanListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetFanListResult) SetSuccess(x interface{}) {
	p.Success = x.(*userproto.GetFanListResp)
}

func (p *GetFanListResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateUser(ctx context.Context, Req *userproto.CreateUserReq) (r *userproto.CreateUserResp, err error) {
	var _args CreateUserArgs
	_args.Req = Req
	var _result CreateUserResult
	if err = p.c.Call(ctx, "CreateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUser(ctx context.Context, Req *userproto.GetUserReq) (r *userproto.GetUserResp, err error) {
	var _args GetUserArgs
	_args.Req = Req
	var _result GetUserResult
	if err = p.c.Call(ctx, "GetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CheckUser(ctx context.Context, Req *userproto.CheckUserReq) (r *userproto.CheckUserResp, err error) {
	var _args CheckUserArgs
	_args.Req = Req
	var _result CheckUserResult
	if err = p.c.Call(ctx, "CheckUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowUser(ctx context.Context, Req *userproto.FollowUserReq) (r *userproto.FollowUserResp, err error) {
	var _args FollowUserArgs
	_args.Req = Req
	var _result FollowUserResult
	if err = p.c.Call(ctx, "FollowUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UnFollowUser(ctx context.Context, Req *userproto.UnFollowUserReq) (r *userproto.UnFollowUserResp, err error) {
	var _args UnFollowUserArgs
	_args.Req = Req
	var _result UnFollowUserResult
	if err = p.c.Call(ctx, "UnFollowUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFollowList(ctx context.Context, Req *userproto.GetFollowListReq) (r *userproto.GetFollowListResp, err error) {
	var _args GetFollowListArgs
	_args.Req = Req
	var _result GetFollowListResult
	if err = p.c.Call(ctx, "GetFollowList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFanList(ctx context.Context, Req *userproto.GetFanListReq) (r *userproto.GetFanListResp, err error) {
	var _args GetFanListArgs
	_args.Req = Req
	var _result GetFanListResult
	if err = p.c.Call(ctx, "GetFanList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
