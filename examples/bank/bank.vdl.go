// This file was auto-generated by the veyron vdl tool.
// Source: bank.vdl

/*
Package bank implements an application to manipulate virtual money. The client's
Veyron Identity determines the account they can manipulate. New identity's make
a new account. Clients can deposit, withdraw, transfer, or query their balance
in virtual currency.
*/
package bank

import (
	"veyron2/security"

	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_veyron2 "veyron2"
	_gen_context "veyron2/context"
	_gen_ipc "veyron2/ipc"
	_gen_naming "veyron2/naming"
	_gen_rt "veyron2/rt"
	_gen_vdlutil "veyron2/vdl/vdlutil"
	_gen_wiretype "veyron2/wiretype"
)

// Bank allows clients to store virtual money. Certain implementations can use persistent storage.
// Uses the client's Veyron Identity to determine account access.
// Bank is the interface the client binds and uses.
// Bank_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type Bank_ExcludingUniversal interface {
	// Connect causes the bank to bless a new user (string) and return their bank account number (int64). Existing users are not blessed (""), but still receive their account number (int64).
	Connect(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (newIdentity string, accountNumber int64, err error)
}
type Bank interface {
	_gen_ipc.UniversalServiceMethods
	Bank_ExcludingUniversal
}

// BankService is the interface the server implements.
type BankService interface {

	// Connect causes the bank to bless a new user (string) and return their bank account number (int64). Existing users are not blessed (""), but still receive their account number (int64).
	Connect(context _gen_ipc.ServerContext) (newIdentity string, accountNumber int64, err error)
}

// BindBank returns the client stub implementing the Bank
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindBank(name string, opts ..._gen_ipc.BindOpt) (Bank, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		client = _gen_rt.R().Client()
	case 1:
		switch o := opts[0].(type) {
		case _gen_veyron2.Runtime:
			client = o.Client()
		case _gen_ipc.Client:
			client = o
		default:
			return nil, _gen_vdlutil.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdlutil.ErrTooManyOptionsToBind
	}
	stub := &clientStubBank{client: client, name: name}

	return stub, nil
}

// NewServerBank creates a new server stub.
//
// It takes a regular server implementing the BankService
// interface, and returns a new server stub.
func NewServerBank(server BankService) interface{} {
	return &ServerStubBank{
		service: server,
	}
}

// clientStubBank implements Bank.
type clientStubBank struct {
	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubBank) Connect(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (newIdentity string, accountNumber int64, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Connect", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&newIdentity, &accountNumber, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubBank) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubBank) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubBank) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubBank wraps a server that implements
// BankService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubBank struct {
	service BankService
}

func (__gen_s *ServerStubBank) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Connect":
		return []interface{}{security.Label(2)}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubBank) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Connect"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "newIdentity", Type: 3},
			{Name: "accountNumber", Type: 37},
			{Name: "err", Type: 65},
		},
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}

func (__gen_s *ServerStubBank) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
	if unresolver, ok := __gen_s.service.(_gen_ipc.Unresolver); ok {
		return unresolver.UnresolveStep(call)
	}
	if call.Server() == nil {
		return
	}
	var published []string
	if published, err = call.Server().Published(); err != nil || published == nil {
		return
	}
	reply = make([]string, len(published))
	for i, p := range published {
		reply[i] = _gen_naming.Join(p, call.Name())
	}
	return
}

func (__gen_s *ServerStubBank) Connect(call _gen_ipc.ServerCall) (newIdentity string, accountNumber int64, err error) {
	newIdentity, accountNumber, err = __gen_s.service.Connect(call)
	return
}

// The BankAccount can only be accessed by blessed users
// BankAccount is the interface the client binds and uses.
// BankAccount_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type BankAccount_ExcludingUniversal interface {
	// Deposit adds the amount given to this account.
	Deposit(ctx _gen_context.T, amount int64, opts ..._gen_ipc.CallOpt) (err error)
	// Withdraw reduces the amount given from this account.
	Withdraw(ctx _gen_context.T, amount int64, opts ..._gen_ipc.CallOpt) (err error)
	// Transfer moves the amount given to the receiver's account.
	Transfer(ctx _gen_context.T, receiver int64, amount int64, opts ..._gen_ipc.CallOpt) (err error)
	// Balance returns the amount stored in this account.
	Balance(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply int64, err error)
}
type BankAccount interface {
	_gen_ipc.UniversalServiceMethods
	BankAccount_ExcludingUniversal
}

// BankAccountService is the interface the server implements.
type BankAccountService interface {

	// Deposit adds the amount given to this account.
	Deposit(context _gen_ipc.ServerContext, amount int64) (err error)
	// Withdraw reduces the amount given from this account.
	Withdraw(context _gen_ipc.ServerContext, amount int64) (err error)
	// Transfer moves the amount given to the receiver's account.
	Transfer(context _gen_ipc.ServerContext, receiver int64, amount int64) (err error)
	// Balance returns the amount stored in this account.
	Balance(context _gen_ipc.ServerContext) (reply int64, err error)
}

// BindBankAccount returns the client stub implementing the BankAccount
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindBankAccount(name string, opts ..._gen_ipc.BindOpt) (BankAccount, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		client = _gen_rt.R().Client()
	case 1:
		switch o := opts[0].(type) {
		case _gen_veyron2.Runtime:
			client = o.Client()
		case _gen_ipc.Client:
			client = o
		default:
			return nil, _gen_vdlutil.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdlutil.ErrTooManyOptionsToBind
	}
	stub := &clientStubBankAccount{client: client, name: name}

	return stub, nil
}

// NewServerBankAccount creates a new server stub.
//
// It takes a regular server implementing the BankAccountService
// interface, and returns a new server stub.
func NewServerBankAccount(server BankAccountService) interface{} {
	return &ServerStubBankAccount{
		service: server,
	}
}

// clientStubBankAccount implements BankAccount.
type clientStubBankAccount struct {
	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubBankAccount) Deposit(ctx _gen_context.T, amount int64, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Deposit", []interface{}{amount}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubBankAccount) Withdraw(ctx _gen_context.T, amount int64, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Withdraw", []interface{}{amount}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubBankAccount) Transfer(ctx _gen_context.T, receiver int64, amount int64, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Transfer", []interface{}{receiver, amount}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubBankAccount) Balance(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply int64, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Balance", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubBankAccount) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubBankAccount) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubBankAccount) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubBankAccount wraps a server that implements
// BankAccountService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubBankAccount struct {
	service BankAccountService
}

func (__gen_s *ServerStubBankAccount) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Deposit":
		return []interface{}{security.Label(2)}, nil
	case "Withdraw":
		return []interface{}{security.Label(2)}, nil
	case "Transfer":
		return []interface{}{security.Label(2)}, nil
	case "Balance":
		return []interface{}{security.Label(1)}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubBankAccount) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Balance"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 37},
			{Name: "", Type: 65},
		},
	}
	result.Methods["Deposit"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "amount", Type: 37},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["Transfer"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "receiver", Type: 37},
			{Name: "amount", Type: 37},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["Withdraw"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "amount", Type: 37},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}

func (__gen_s *ServerStubBankAccount) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
	if unresolver, ok := __gen_s.service.(_gen_ipc.Unresolver); ok {
		return unresolver.UnresolveStep(call)
	}
	if call.Server() == nil {
		return
	}
	var published []string
	if published, err = call.Server().Published(); err != nil || published == nil {
		return
	}
	reply = make([]string, len(published))
	for i, p := range published {
		reply[i] = _gen_naming.Join(p, call.Name())
	}
	return
}

func (__gen_s *ServerStubBankAccount) Deposit(call _gen_ipc.ServerCall, amount int64) (err error) {
	err = __gen_s.service.Deposit(call, amount)
	return
}

func (__gen_s *ServerStubBankAccount) Withdraw(call _gen_ipc.ServerCall, amount int64) (err error) {
	err = __gen_s.service.Withdraw(call, amount)
	return
}

func (__gen_s *ServerStubBankAccount) Transfer(call _gen_ipc.ServerCall, receiver int64, amount int64) (err error) {
	err = __gen_s.service.Transfer(call, receiver, amount)
	return
}

func (__gen_s *ServerStubBankAccount) Balance(call _gen_ipc.ServerCall) (reply int64, err error) {
	reply, err = __gen_s.service.Balance(call)
	return
}
