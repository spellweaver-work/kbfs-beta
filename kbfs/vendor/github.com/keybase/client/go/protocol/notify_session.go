// Auto-generated by avdl-compiler v1.1.1 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/notify_session.avdl

package keybase1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type LoggedOutArg struct {
}

type LoggedInArg struct {
	Username string `codec:"username" json:"username"`
}

type NotifySessionInterface interface {
	LoggedOut(context.Context) error
	LoggedIn(context.Context, string) error
}

func NotifySessionProtocol(i NotifySessionInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.NotifySession",
		Methods: map[string]rpc.ServeHandlerDescription{
			"loggedOut": {
				MakeArg: func() interface{} {
					ret := make([]LoggedOutArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					err = i.LoggedOut(ctx)
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"loggedIn": {
				MakeArg: func() interface{} {
					ret := make([]LoggedInArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]LoggedInArg)
					if !ok {
						err = rpc.NewTypeError((*[]LoggedInArg)(nil), args)
						return
					}
					err = i.LoggedIn(ctx, (*typedArgs)[0].Username)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type NotifySessionClient struct {
	Cli rpc.GenericClient
}

func (c NotifySessionClient) LoggedOut(ctx context.Context) (err error) {
	err = c.Cli.Notify(ctx, "keybase.1.NotifySession.loggedOut", []interface{}{LoggedOutArg{}})
	return
}

func (c NotifySessionClient) LoggedIn(ctx context.Context, username string) (err error) {
	__arg := LoggedInArg{Username: username}
	err = c.Cli.Call(ctx, "keybase.1.NotifySession.loggedIn", []interface{}{__arg}, nil)
	return
}
