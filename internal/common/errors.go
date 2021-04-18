package common

import "fmt"

type Errors []error

func (c *Errors) Add(e error) { *c = append(*c, e) }

func (c *Errors) Error() (err string) {
	err = "Errors:\n"
	for i, e := range *c {
		err += fmt.Sprintf("\tError %d: %s\n", i, e.Error())
	}

	return err
}

// AsError returns as error, returning nil if errors is empty
func (c *Errors) AsError() error {
	if len(*c) == 0 {
		return nil
	}
	return c
}
