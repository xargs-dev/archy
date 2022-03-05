package archy

import (
	"strings"
)

type Values struct {
	KernelName string `json:"kernel_name,omitempty"`
	Machine    string `json:"machine,omitempty"`
	Source     string `json:"source,omitempty"`
}

func (v *Values) StringSlice() []string {
	// Ensure this is in order of how uname would do it.
	return []string{
		v.KernelName,
		v.Machine,
	}
}

func (v *Values) String() string {
	// Necessary to prevent extra whitespace in case one of the values is empty
	str := []string{}
	for _, val := range v.StringSlice() {
		if val != "" {
			str = append(str, val)
		}
	}

	return strings.Join(str, " ")
}
