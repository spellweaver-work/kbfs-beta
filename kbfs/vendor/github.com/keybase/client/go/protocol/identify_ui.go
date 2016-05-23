// Auto-generated by avdl-compiler v1.3.1 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/identify_ui.avdl

package keybase1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type ProofResult struct {
	State  ProofState  `codec:"state" json:"state"`
	Status ProofStatus `codec:"status" json:"status"`
	Desc   string      `codec:"desc" json:"desc"`
}

type IdentifyRow struct {
	RowId     int         `codec:"rowId" json:"rowId"`
	Proof     RemoteProof `codec:"proof" json:"proof"`
	TrackDiff *TrackDiff  `codec:"trackDiff,omitempty" json:"trackDiff,omitempty"`
}

type IdentifyKey struct {
	PGPFingerprint []byte     `codec:"pgpFingerprint" json:"pgpFingerprint"`
	KID            KID        `codec:"KID" json:"KID"`
	TrackDiff      *TrackDiff `codec:"trackDiff,omitempty" json:"trackDiff,omitempty"`
	BreaksTracking bool       `codec:"breaksTracking" json:"breaksTracking"`
}

type Cryptocurrency struct {
	RowId   int    `codec:"rowId" json:"rowId"`
	Pkhash  []byte `codec:"pkhash" json:"pkhash"`
	Address string `codec:"address" json:"address"`
}

type RevokedProof struct {
	Proof RemoteProof `codec:"proof" json:"proof"`
	Diff  TrackDiff   `codec:"diff" json:"diff"`
}

type Identity struct {
	Status          *Status          `codec:"status,omitempty" json:"status,omitempty"`
	WhenLastTracked Time             `codec:"whenLastTracked" json:"whenLastTracked"`
	Proofs          []IdentifyRow    `codec:"proofs" json:"proofs"`
	Cryptocurrency  []Cryptocurrency `codec:"cryptocurrency" json:"cryptocurrency"`
	Revoked         []TrackDiff      `codec:"revoked" json:"revoked"`
	RevokedDetails  []RevokedProof   `codec:"revokedDetails" json:"revokedDetails"`
	BreaksTracking  bool             `codec:"breaksTracking" json:"breaksTracking"`
}

type SigHint struct {
	RemoteId  string `codec:"remoteId" json:"remoteId"`
	HumanUrl  string `codec:"humanUrl" json:"humanUrl"`
	ApiUrl    string `codec:"apiUrl" json:"apiUrl"`
	CheckText string `codec:"checkText" json:"checkText"`
}

type CheckResultFreshness int

const (
	CheckResultFreshness_FRESH  CheckResultFreshness = 0
	CheckResultFreshness_AGED   CheckResultFreshness = 1
	CheckResultFreshness_RANCID CheckResultFreshness = 2
)

type CheckResult struct {
	ProofResult ProofResult          `codec:"proofResult" json:"proofResult"`
	Time        Time                 `codec:"time" json:"time"`
	Freshness   CheckResultFreshness `codec:"freshness" json:"freshness"`
}

type LinkCheckResult struct {
	ProofId            int          `codec:"proofId" json:"proofId"`
	ProofResult        ProofResult  `codec:"proofResult" json:"proofResult"`
	SnoozedResult      ProofResult  `codec:"snoozedResult" json:"snoozedResult"`
	TorWarning         bool         `codec:"torWarning" json:"torWarning"`
	TmpTrackExpireTime Time         `codec:"tmpTrackExpireTime" json:"tmpTrackExpireTime"`
	Cached             *CheckResult `codec:"cached,omitempty" json:"cached,omitempty"`
	Diff               *TrackDiff   `codec:"diff,omitempty" json:"diff,omitempty"`
	RemoteDiff         *TrackDiff   `codec:"remoteDiff,omitempty" json:"remoteDiff,omitempty"`
	Hint               *SigHint     `codec:"hint,omitempty" json:"hint,omitempty"`
	BreaksTracking     bool         `codec:"breaksTracking" json:"breaksTracking"`
}

type UserCard struct {
	Following     int    `codec:"following" json:"following"`
	Followers     int    `codec:"followers" json:"followers"`
	Uid           UID    `codec:"uid" json:"uid"`
	FullName      string `codec:"fullName" json:"fullName"`
	Location      string `codec:"location" json:"location"`
	Bio           string `codec:"bio" json:"bio"`
	Website       string `codec:"website" json:"website"`
	Twitter       string `codec:"twitter" json:"twitter"`
	YouFollowThem bool   `codec:"youFollowThem" json:"youFollowThem"`
	TheyFollowYou bool   `codec:"theyFollowYou" json:"theyFollowYou"`
}

type ConfirmResult struct {
	IdentityConfirmed bool `codec:"identityConfirmed" json:"identityConfirmed"`
	RemoteConfirmed   bool `codec:"remoteConfirmed" json:"remoteConfirmed"`
	ExpiringLocal     bool `codec:"expiringLocal" json:"expiringLocal"`
}

type DismissReasonType int

const (
	DismissReasonType_NONE              DismissReasonType = 0
	DismissReasonType_HANDLED_ELSEWHERE DismissReasonType = 1
)

type DismissReason struct {
	Type     DismissReasonType `codec:"type" json:"type"`
	Reason   string            `codec:"reason" json:"reason"`
	Resource string            `codec:"resource" json:"resource"`
}

type DisplayTLFCreateWithInviteArg struct {
	SessionID  int    `codec:"sessionID" json:"sessionID"`
	FolderName string `codec:"folderName" json:"folderName"`
	IsPrivate  bool   `codec:"isPrivate" json:"isPrivate"`
	Assertion  string `codec:"assertion" json:"assertion"`
	InviteLink string `codec:"inviteLink" json:"inviteLink"`
	Throttled  bool   `codec:"throttled" json:"throttled"`
}

type DelegateIdentifyUIArg struct {
}

type StartArg struct {
	SessionID int            `codec:"sessionID" json:"sessionID"`
	Username  string         `codec:"username" json:"username"`
	Reason    IdentifyReason `codec:"reason" json:"reason"`
}

type DisplayKeyArg struct {
	SessionID int         `codec:"sessionID" json:"sessionID"`
	Key       IdentifyKey `codec:"key" json:"key"`
}

type ReportLastTrackArg struct {
	SessionID int           `codec:"sessionID" json:"sessionID"`
	Track     *TrackSummary `codec:"track,omitempty" json:"track,omitempty"`
}

type LaunchNetworkChecksArg struct {
	SessionID int      `codec:"sessionID" json:"sessionID"`
	Identity  Identity `codec:"identity" json:"identity"`
	User      User     `codec:"user" json:"user"`
}

type DisplayTrackStatementArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Stmt      string `codec:"stmt" json:"stmt"`
}

type FinishWebProofCheckArg struct {
	SessionID int             `codec:"sessionID" json:"sessionID"`
	Rp        RemoteProof     `codec:"rp" json:"rp"`
	Lcr       LinkCheckResult `codec:"lcr" json:"lcr"`
}

type FinishSocialProofCheckArg struct {
	SessionID int             `codec:"sessionID" json:"sessionID"`
	Rp        RemoteProof     `codec:"rp" json:"rp"`
	Lcr       LinkCheckResult `codec:"lcr" json:"lcr"`
}

type DisplayCryptocurrencyArg struct {
	SessionID int            `codec:"sessionID" json:"sessionID"`
	C         Cryptocurrency `codec:"c" json:"c"`
}

type ReportTrackTokenArg struct {
	SessionID  int        `codec:"sessionID" json:"sessionID"`
	TrackToken TrackToken `codec:"trackToken" json:"trackToken"`
}

type DisplayUserCardArg struct {
	SessionID int      `codec:"sessionID" json:"sessionID"`
	Card      UserCard `codec:"card" json:"card"`
}

type ConfirmArg struct {
	SessionID int             `codec:"sessionID" json:"sessionID"`
	Outcome   IdentifyOutcome `codec:"outcome" json:"outcome"`
}

type FinishArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type DismissArg struct {
	SessionID int           `codec:"sessionID" json:"sessionID"`
	Username  string        `codec:"username" json:"username"`
	Reason    DismissReason `codec:"reason" json:"reason"`
}

type IdentifyUiInterface interface {
	DisplayTLFCreateWithInvite(context.Context, DisplayTLFCreateWithInviteArg) error
	DelegateIdentifyUI(context.Context) (int, error)
	Start(context.Context, StartArg) error
	DisplayKey(context.Context, DisplayKeyArg) error
	ReportLastTrack(context.Context, ReportLastTrackArg) error
	LaunchNetworkChecks(context.Context, LaunchNetworkChecksArg) error
	DisplayTrackStatement(context.Context, DisplayTrackStatementArg) error
	FinishWebProofCheck(context.Context, FinishWebProofCheckArg) error
	FinishSocialProofCheck(context.Context, FinishSocialProofCheckArg) error
	DisplayCryptocurrency(context.Context, DisplayCryptocurrencyArg) error
	ReportTrackToken(context.Context, ReportTrackTokenArg) error
	DisplayUserCard(context.Context, DisplayUserCardArg) error
	Confirm(context.Context, ConfirmArg) (ConfirmResult, error)
	Finish(context.Context, int) error
	Dismiss(context.Context, DismissArg) error
}

func IdentifyUiProtocol(i IdentifyUiInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.identifyUi",
		Methods: map[string]rpc.ServeHandlerDescription{
			"displayTLFCreateWithInvite": {
				MakeArg: func() interface{} {
					ret := make([]DisplayTLFCreateWithInviteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DisplayTLFCreateWithInviteArg)
					if !ok {
						err = rpc.NewTypeError((*[]DisplayTLFCreateWithInviteArg)(nil), args)
						return
					}
					err = i.DisplayTLFCreateWithInvite(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"delegateIdentifyUI": {
				MakeArg: func() interface{} {
					ret := make([]DelegateIdentifyUIArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.DelegateIdentifyUI(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"start": {
				MakeArg: func() interface{} {
					ret := make([]StartArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]StartArg)
					if !ok {
						err = rpc.NewTypeError((*[]StartArg)(nil), args)
						return
					}
					err = i.Start(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"displayKey": {
				MakeArg: func() interface{} {
					ret := make([]DisplayKeyArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DisplayKeyArg)
					if !ok {
						err = rpc.NewTypeError((*[]DisplayKeyArg)(nil), args)
						return
					}
					err = i.DisplayKey(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"reportLastTrack": {
				MakeArg: func() interface{} {
					ret := make([]ReportLastTrackArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ReportLastTrackArg)
					if !ok {
						err = rpc.NewTypeError((*[]ReportLastTrackArg)(nil), args)
						return
					}
					err = i.ReportLastTrack(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"launchNetworkChecks": {
				MakeArg: func() interface{} {
					ret := make([]LaunchNetworkChecksArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]LaunchNetworkChecksArg)
					if !ok {
						err = rpc.NewTypeError((*[]LaunchNetworkChecksArg)(nil), args)
						return
					}
					err = i.LaunchNetworkChecks(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"displayTrackStatement": {
				MakeArg: func() interface{} {
					ret := make([]DisplayTrackStatementArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DisplayTrackStatementArg)
					if !ok {
						err = rpc.NewTypeError((*[]DisplayTrackStatementArg)(nil), args)
						return
					}
					err = i.DisplayTrackStatement(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"finishWebProofCheck": {
				MakeArg: func() interface{} {
					ret := make([]FinishWebProofCheckArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]FinishWebProofCheckArg)
					if !ok {
						err = rpc.NewTypeError((*[]FinishWebProofCheckArg)(nil), args)
						return
					}
					err = i.FinishWebProofCheck(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"finishSocialProofCheck": {
				MakeArg: func() interface{} {
					ret := make([]FinishSocialProofCheckArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]FinishSocialProofCheckArg)
					if !ok {
						err = rpc.NewTypeError((*[]FinishSocialProofCheckArg)(nil), args)
						return
					}
					err = i.FinishSocialProofCheck(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"displayCryptocurrency": {
				MakeArg: func() interface{} {
					ret := make([]DisplayCryptocurrencyArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DisplayCryptocurrencyArg)
					if !ok {
						err = rpc.NewTypeError((*[]DisplayCryptocurrencyArg)(nil), args)
						return
					}
					err = i.DisplayCryptocurrency(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"reportTrackToken": {
				MakeArg: func() interface{} {
					ret := make([]ReportTrackTokenArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ReportTrackTokenArg)
					if !ok {
						err = rpc.NewTypeError((*[]ReportTrackTokenArg)(nil), args)
						return
					}
					err = i.ReportTrackToken(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"displayUserCard": {
				MakeArg: func() interface{} {
					ret := make([]DisplayUserCardArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DisplayUserCardArg)
					if !ok {
						err = rpc.NewTypeError((*[]DisplayUserCardArg)(nil), args)
						return
					}
					err = i.DisplayUserCard(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"confirm": {
				MakeArg: func() interface{} {
					ret := make([]ConfirmArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ConfirmArg)
					if !ok {
						err = rpc.NewTypeError((*[]ConfirmArg)(nil), args)
						return
					}
					ret, err = i.Confirm(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"finish": {
				MakeArg: func() interface{} {
					ret := make([]FinishArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]FinishArg)
					if !ok {
						err = rpc.NewTypeError((*[]FinishArg)(nil), args)
						return
					}
					err = i.Finish(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"dismiss": {
				MakeArg: func() interface{} {
					ret := make([]DismissArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DismissArg)
					if !ok {
						err = rpc.NewTypeError((*[]DismissArg)(nil), args)
						return
					}
					err = i.Dismiss(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type IdentifyUiClient struct {
	Cli rpc.GenericClient
}

func (c IdentifyUiClient) DisplayTLFCreateWithInvite(ctx context.Context, __arg DisplayTLFCreateWithInviteArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.identifyUi.displayTLFCreateWithInvite", []interface{}{__arg}, nil)
	return
}

func (c IdentifyUiClient) DelegateIdentifyUI(ctx context.Context) (res int, err error) {
	err = c.Cli.Call(ctx, "keybase.1.identifyUi.delegateIdentifyUI", []interface{}{DelegateIdentifyUIArg{}}, &res)
	return
}

func (c IdentifyUiClient) Start(ctx context.Context, __arg StartArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.identifyUi.start", []interface{}{__arg}, nil)
	return
}

func (c IdentifyUiClient) DisplayKey(ctx context.Context, __arg DisplayKeyArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.identifyUi.displayKey", []interface{}{__arg}, nil)
	return
}

func (c IdentifyUiClient) ReportLastTrack(ctx context.Context, __arg ReportLastTrackArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.identifyUi.reportLastTrack", []interface{}{__arg}, nil)
	return
}

func (c IdentifyUiClient) LaunchNetworkChecks(ctx context.Context, __arg LaunchNetworkChecksArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.identifyUi.launchNetworkChecks", []interface{}{__arg}, nil)
	return
}

func (c IdentifyUiClient) DisplayTrackStatement(ctx context.Context, __arg DisplayTrackStatementArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.identifyUi.displayTrackStatement", []interface{}{__arg}, nil)
	return
}

func (c IdentifyUiClient) FinishWebProofCheck(ctx context.Context, __arg FinishWebProofCheckArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.identifyUi.finishWebProofCheck", []interface{}{__arg}, nil)
	return
}

func (c IdentifyUiClient) FinishSocialProofCheck(ctx context.Context, __arg FinishSocialProofCheckArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.identifyUi.finishSocialProofCheck", []interface{}{__arg}, nil)
	return
}

func (c IdentifyUiClient) DisplayCryptocurrency(ctx context.Context, __arg DisplayCryptocurrencyArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.identifyUi.displayCryptocurrency", []interface{}{__arg}, nil)
	return
}

func (c IdentifyUiClient) ReportTrackToken(ctx context.Context, __arg ReportTrackTokenArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.identifyUi.reportTrackToken", []interface{}{__arg}, nil)
	return
}

func (c IdentifyUiClient) DisplayUserCard(ctx context.Context, __arg DisplayUserCardArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.identifyUi.displayUserCard", []interface{}{__arg}, nil)
	return
}

func (c IdentifyUiClient) Confirm(ctx context.Context, __arg ConfirmArg) (res ConfirmResult, err error) {
	err = c.Cli.Call(ctx, "keybase.1.identifyUi.confirm", []interface{}{__arg}, &res)
	return
}

func (c IdentifyUiClient) Finish(ctx context.Context, sessionID int) (err error) {
	__arg := FinishArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.identifyUi.finish", []interface{}{__arg}, nil)
	return
}

func (c IdentifyUiClient) Dismiss(ctx context.Context, __arg DismissArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.identifyUi.dismiss", []interface{}{__arg}, nil)
	return
}
