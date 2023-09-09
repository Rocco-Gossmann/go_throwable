package main


type drawableException struct {
    msg string
    subject string
    value any
}

func (e drawableException) GetMessage() string { return e.msg; }
func (e drawableException) GetSubject() string { return e.subject; }
func (e drawableException) GetValue() any { return e.value; }
