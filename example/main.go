package main

import (
    "fmt"
    ex "github.com/rocco-gossmann/go_throwable"
)

type drawableType byte;

const (
    DRAWABLE_NONE   drawableType = 0
    DRAWABLE_LINE   drawableType = 1
    DRAWABLE_PIXEL  drawableType = 2 
    DRAWABLE_RECT   drawableType = 3 
    DRAWABLE_CIRCLE drawableType = 4 
)

func validateType(t drawableType) (ret drawableType) {

    switch(t) {
        case DRAWABLE_LINE, 
             DRAWABLE_PIXEL,
             DRAWABLE_RECT,
             DRAWABLE_CIRCLE:

         default: if t == 5 {
             ex.Throw(drawableException{msg: "given type is not a drawable", subject: "t", value: t})
         } else {
             panic("PAAAANIC !!! *runs in circles*");
         }
    }

    ret = t;
    return 
}



func main() {

    var (
        dt drawableType
        ok bool
    )

    // By default the return of ex.Try is the same as the one of the first parameter
    // Since this function will not cause a panic, dt should be 1 after this
    //-------------------------------------------------------------------------
    dt,ok = (ex.Try( 

        func() any { return validateType(DRAWABLE_LINE) }, 
                                    //    /\-- will NOT cause a panic

        ex.TryOpts{
            Default: DRAWABLE_NONE, // <-- would be returned if a panic happens
        },
    )).(drawableType)
    fmt.Println("\nGot drawable type: ", dt, ok)




    // Reacting to a panic by defining a handler function that 
    // Will return a new Value.
    // The try function will Panic, since 5 is not a valid drawingType
    //
    // Since the panic encountered has the ex.Throwable interface
    // we can get even more details on why it happened
    //
    // After a panic is cought, the return value of the catching function 
    // Is returned instead (in this case "0")
    //-------------------------------------------------------------------------
    fmt.Print("\n\n\n");
    dt,ok = (ex.Try( 
        func() any { 
            return validateType(5) //<- causes a panic
        },

        ex.TryOpts{ Catch: func(p any) any { 
            
            throwable, isThrowable := ex.Throwable(p);

            if isThrowable {
                fmt.Println("cought a throwable!");
                fmt.Println("The value creating the issue was", throwable.GetValue()) 
            } else { 
                fmt.Println("cought none throwable panic", p);
            }
            
            return DRAWABLE_NONE;

        } },
    )).(drawableType);
    fmt.Println("\nGot drawable type: ", dt, ok)




    // Reacting to a panic by defining a default value that
    // that is returned instead. 
    // Like the previous example, this also throws a panic, but a different one.
    // If you don't care for why a panic happens, you can also define a 
    // TryOpts.Default  instead of a TryOpts.Catch.
    //
    // In that case, The Value given to TryOpts.Default is returned by ex.Try,
    // should a panic happen
    // In this case ex.Try returns DRAWABLE_CIRCLE (4) intead of 6 since 6,
    // is also not a valid drawingType
    //-------------------------------------------------------------------------
    fmt.Print("\n\n\n");
    dt,ok = (ex.Try( 

        func() any { 
            return validateType(6)  //<- causes a different panic
        }, 

        ex.TryOpts{ 
            Default: DRAWABLE_CIRCLE,  // <- will be returned if a panic happened 
        },

    )).(drawableType)

    fmt.Println("\nGot drawable type: ", dt, ok)




    // If you expect a Panic to happen, you should disable the Warning message,
    // ex.Try Prints  by setting TryOpt.SkipWarnings to true
    //-------------------------------------------------------------------------
    fmt.Print("\n\n\n");
    dt,ok = (ex.Try( 

        func() any { return validateType(10) }, 

        ex.TryOpts{
            SkipWarnings: true,
            Default: DRAWABLE_RECT, // <-- would be returned if a panic happens
        },
    )).(drawableType)
    fmt.Println("\nGot drawable type: ", dt, ok)



}
