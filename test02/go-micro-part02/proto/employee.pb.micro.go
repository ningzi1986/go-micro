// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: employee.proto

package go_svr_proto_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for EmployeeService service

type EmployeeService interface {
	GetEmployees(ctx context.Context, in *GetEmployeeRequest, opts ...client.CallOption) (*GetEmployeeResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...client.CallOption) (*LogoutResponse, error)
}

type employeeService struct {
	c    client.Client
	name string
}

func NewEmployeeService(name string, c client.Client) EmployeeService {
	return &employeeService{
		c:    c,
		name: name,
	}
}

func (c *employeeService) GetEmployees(ctx context.Context, in *GetEmployeeRequest, opts ...client.CallOption) (*GetEmployeeResponse, error) {
	req := c.c.NewRequest(c.name, "EmployeeService.GetEmployees", in)
	out := new(GetEmployeeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "EmployeeService.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeService) Logout(ctx context.Context, in *LogoutRequest, opts ...client.CallOption) (*LogoutResponse, error) {
	req := c.c.NewRequest(c.name, "EmployeeService.Logout", in)
	out := new(LogoutResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EmployeeService service

type EmployeeServiceHandler interface {
	GetEmployees(context.Context, *GetEmployeeRequest, *GetEmployeeResponse) error
	Login(context.Context, *LoginRequest, *LoginResponse) error
	Logout(context.Context, *LogoutRequest, *LogoutResponse) error
}

func RegisterEmployeeServiceHandler(s server.Server, hdlr EmployeeServiceHandler, opts ...server.HandlerOption) error {
	type employeeService interface {
		GetEmployees(ctx context.Context, in *GetEmployeeRequest, out *GetEmployeeResponse) error
		Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error
		Logout(ctx context.Context, in *LogoutRequest, out *LogoutResponse) error
	}
	type EmployeeService struct {
		employeeService
	}
	h := &employeeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&EmployeeService{h}, opts...))
}

type employeeServiceHandler struct {
	EmployeeServiceHandler
}

func (h *employeeServiceHandler) GetEmployees(ctx context.Context, in *GetEmployeeRequest, out *GetEmployeeResponse) error {
	return h.EmployeeServiceHandler.GetEmployees(ctx, in, out)
}

func (h *employeeServiceHandler) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.EmployeeServiceHandler.Login(ctx, in, out)
}

func (h *employeeServiceHandler) Logout(ctx context.Context, in *LogoutRequest, out *LogoutResponse) error {
	return h.EmployeeServiceHandler.Logout(ctx, in, out)
}