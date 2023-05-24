package cmd

import "fmt"

type FlagParsingError struct {
	Err error
}

func (e *FlagParsingError) Error() string {
	return fmt.Sprintf("Error parsing flag:  %v", e.Err)
}

type JSONParsingError struct {
	Err error
}

func (e *JSONParsingError) Error() string {
	return fmt.Sprintf("Error parsing JSON:  %v", e.Err)
}
