package main

import "fmt"

type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string{
	strFormat := ` Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0`

	return fmt.Sprint(strFormat,de.dividee)
}

func Divide(varDividee int,varDivider int)(result int,errorMsg string)  {

}

func main()  {

}



