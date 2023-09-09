package go_throwable

import "fmt"

type IThrowable interface {
    GetSubject() string 
    GetMessage() string
    GetValue() any
}

type TryOpts struct {
    Catch func(any) any

    Finally func() 

    Default any

    SkipWarnings bool
}

func Throwable(p any) (IThrowable, bool) {
    i, ok := interface{}(p).(IThrowable);
    return i, ok
}

func Throw(throwable IThrowable) { panic(throwable); }

func Try( fnc func() any, opts TryOpts) (fncresponse any) {

    defer func() {
        if p := recover(); p != nil {

            if !opts.SkipWarnings { fmt.Println("[WARNING] a `Try` encountered a panic.", p) }

            if v := opts.Catch ; v != nil {
                if !opts.SkipWarnings { fmt.Println("[Panic] forwarded to Catch") }
                fncresponse = v(p);

            } else {
                if !opts.SkipWarnings { fmt.Println("[Panic] responding with opts.Default value ", opts.Default) }
                fncresponse = opts.Default;

            }

        }

        if v := opts.Finally ; v != nil {
            v();
        }
    }()

    fncresponse = fnc()

    return;
}  
