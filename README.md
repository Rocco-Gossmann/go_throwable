# Go-Throwable

Aiming to twist GOs `panic`=> `defer recover` system into a `try { ... } catch { ... }` form

Here is a simple example on how to use it
```go
import ex "github.com/rocco-gossmann/go_throwable"

func mayCauseAPanic() int {
    panic("PAAAAANIC !!!!! * runs in circles *")

    return 20;
}


func main() {

    var response int
    var ok       bool

    response, ok = interface{}(
        ex.Try( 
            func() any {    // <- run all your panicing functions in here
                return mayCauseAPanic(); 
            },

            ex.TryOpts{ 
                Default: -1 //<- In case of a panic return -1
            },  
        ),
    ).(int)

}
```
After this the `ok` value contains informations if the returned value is  
an `int`  (`true` means, it is),  while `response` contains the returned number.



## `Catch`ing a panic with a function 

Based on the previous example, lets change the Try to have a bit more control
over the recovery from the panic

```go

    response, ok = interface{}(
        ex.Try( 
            func() any {    // <- run all your panicing functions in here
                return mayCauseAPanic(); 
            },

            ex.TryOpts{ 
                Catch: func(p) any {  // <- the value returned by this function,
                                      //    will become the new return value
                    fmt.PrintLn("Cought the Panic => Returning a new number", p)
                    return -2
                } 
            },  
        ),
    ).(int)

```



## `Finally` - Executing code, regardless of if a panic happened or not
this can be done via the TryOpts.Finally. 
A function, that takes no params and returns nothing.
Regardless of if a panic happened, it will always be executed after everything
is done

```go
    response, ok = interface{}(
        ex.Try( 
            func() any {    // <- run all your panicing functions in here
                return mayCauseAPanic(); 
            },

            ex.TryOpts{ 
                Default: 7353,

                Finally: func() {
                    fmt.PrintLn("This is printed regardless of if a panic happened")
                },
            },  
        ),
    ).(int)
```


## Disabeling the Warning output.
By Default go_throwable will make some outputs to stdout this can be disabled
by setting `TryOpts.SkipWarnings: true`

```go
    response, ok = interface{}(
        ex.Try( 
            func() any {  return mayCauseAPanic(); },

            ex.TryOpts{ 
                SkipWarnings: true,
                Default: 7353,
            },  
        ),
    ).(int)
