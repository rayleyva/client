// Auto-generated by avdl-compiler v1.3.23 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/emom1/ae.avdl

package emom1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type MsgType int

const (
	MsgType_CALL    MsgType = 0
	MsgType_REPLY   MsgType = 1
	MsgType_NOTIFY  MsgType = 2
	MsgType_RATCHET MsgType = 16
)

func (o MsgType) DeepCopy() MsgType { return o }

var MsgTypeMap = map[string]MsgType{
	"CALL":    0,
	"REPLY":   1,
	"NOTIFY":  2,
	"RATCHET": 16,
}

var MsgTypeRevMap = map[MsgType]string{
	0:  "CALL",
	1:  "REPLY",
	2:  "NOTIFY",
	16: "RATCHET",
}

func (e MsgType) String() string {
	if v, ok := MsgTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type AuthEnc struct {
	E []byte `codec:"e" json:"e"`
	N Seqno  `codec:"n" json:"n"`
	R Seqno  `codec:"r" json:"r"`
}

func (o AuthEnc) DeepCopy() AuthEnc {
	return AuthEnc{
		E: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte{}, x...)
		})(o.E),
		N: o.N.DeepCopy(),
		R: o.R.DeepCopy(),
	}
}

type ServerRatchet struct {
	I Seqno `codec:"i" json:"i"`
	K KID   `codec:"k" json:"k"`
}

func (o ServerRatchet) DeepCopy() ServerRatchet {
	return ServerRatchet{
		I: o.I.DeepCopy(),
		K: o.K.DeepCopy(),
	}
}

type KID []byte

func (o KID) DeepCopy() KID {
	return (func(x []byte) []byte {
		if x == nil {
			return nil
		}
		return append([]byte{}, x...)
	})(o)
}

type UID []byte

func (o UID) DeepCopy() UID {
	return (func(x []byte) []byte {
		if x == nil {
			return nil
		}
		return append([]byte{}, x...)
	})(o)
}

type Time int64

func (o Time) DeepCopy() Time {
	return o
}

type KeyGen int64

func (o KeyGen) DeepCopy() KeyGen {
	return o
}

type Seqno int64

func (o Seqno) DeepCopy() Seqno {
	return o
}

type Handshake struct {
	K KID    `codec:"k" json:"k"`
	S KeyGen `codec:"s" json:"s"`
	V int    `codec:"v" json:"v"`
}

func (o Handshake) DeepCopy() Handshake {
	return Handshake{
		K: o.K.DeepCopy(),
		S: o.S.DeepCopy(),
		V: o.V,
	}
}

type RequestPlaintext struct {
	A []byte           `codec:"a" json:"a"`
	F *SignedAuthToken `codec:"f,omitempty" json:"f,omitempty"`
	N string           `codec:"n" json:"n"`
	S *Seqno           `codec:"s,omitempty" json:"s,omitempty"`
}

func (o RequestPlaintext) DeepCopy() RequestPlaintext {
	return RequestPlaintext{
		A: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte{}, x...)
		})(o.A),
		F: (func(x *SignedAuthToken) *SignedAuthToken {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.F),
		N: o.N,
		S: (func(x *Seqno) *Seqno {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.S),
	}
}

type ResponsePlaintext struct {
	E []byte `codec:"e" json:"e"`
	R []byte `codec:"r" json:"r"`
	S Seqno  `codec:"s" json:"s"`
}

func (o ResponsePlaintext) DeepCopy() ResponsePlaintext {
	return ResponsePlaintext{
		E: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte{}, x...)
		})(o.E),
		R: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte{}, x...)
		})(o.R),
		S: o.S.DeepCopy(),
	}
}

type AuthToken struct {
	C Time `codec:"c" json:"c"`
	D KID  `codec:"d" json:"d"`
	K KID  `codec:"k" json:"k"`
	U UID  `codec:"u" json:"u"`
}

func (o AuthToken) DeepCopy() AuthToken {
	return AuthToken{
		C: o.C.DeepCopy(),
		D: o.D.DeepCopy(),
		K: o.K.DeepCopy(),
		U: o.U.DeepCopy(),
	}
}

type AuthTokenExported struct {
	C Time `codec:"c" json:"c"`
	D KID  `codec:"d" json:"d"`
	U UID  `codec:"u" json:"u"`
}

func (o AuthTokenExported) DeepCopy() AuthTokenExported {
	return AuthTokenExported{
		C: o.C.DeepCopy(),
		D: o.D.DeepCopy(),
		U: o.U.DeepCopy(),
	}
}

type SignedAuthToken struct {
	T AuthTokenExported `codec:"t" json:"t"`
	S []byte            `codec:"s" json:"s"`
}

func (o SignedAuthToken) DeepCopy() SignedAuthToken {
	return SignedAuthToken{
		T: o.T.DeepCopy(),
		S: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte{}, x...)
		})(o.S),
	}
}

type Arg struct {
	A AuthEnc    `codec:"a" json:"a"`
	H *Handshake `codec:"h,omitempty" json:"h,omitempty"`
}

func (o Arg) DeepCopy() Arg {
	return Arg{
		A: o.A.DeepCopy(),
		H: (func(x *Handshake) *Handshake {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.H),
	}
}

type Res struct {
	A AuthEnc  `codec:"a" json:"a"`
	R *AuthEnc `codec:"r,omitempty" json:"r,omitempty"`
}

func (o Res) DeepCopy() Res {
	return Res{
		A: o.A.DeepCopy(),
		R: (func(x *AuthEnc) *AuthEnc {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.R),
	}
}

type CArg struct {
	A Arg `codec:"a" json:"a"`
}

type NArg struct {
	A Arg `codec:"a" json:"a"`
}

type AeInterface interface {
	C(context.Context, Arg) (Res, error)
	N(context.Context, Arg) error
}

func AeProtocol(i AeInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "emom.1.ae",
		Methods: map[string]rpc.ServeHandlerDescription{
			"c": {
				MakeArg: func() interface{} {
					ret := make([]CArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]CArg)
					if !ok {
						err = rpc.NewTypeError((*[]CArg)(nil), args)
						return
					}
					ret, err = i.C(ctx, (*typedArgs)[0].A)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"n": {
				MakeArg: func() interface{} {
					ret := make([]NArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]NArg)
					if !ok {
						err = rpc.NewTypeError((*[]NArg)(nil), args)
						return
					}
					err = i.N(ctx, (*typedArgs)[0].A)
					return
				},
				MethodType: rpc.MethodNotify,
			},
		},
	}
}

type AeClient struct {
	Cli rpc.GenericClient
}

func (c AeClient) C(ctx context.Context, a Arg) (res Res, err error) {
	__arg := CArg{A: a}
	err = c.Cli.Call(ctx, "emom.1.ae.c", []interface{}{__arg}, &res)
	return
}

func (c AeClient) N(ctx context.Context, a Arg) (err error) {
	__arg := NArg{A: a}
	err = c.Cli.Notify(ctx, "emom.1.ae.n", []interface{}{__arg})
	return
}