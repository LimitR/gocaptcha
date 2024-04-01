package core

import (
	"strings"

	"github.com/google/uuid"
)

type Randomizer struct{}

func (r *Randomizer) GetString(len uint) string {
	return r.splitString(len, strings.Split(strings.Join(strings.Split(uuid.NewString(), "-"), ""), ""))
}

func (r *Randomizer) splitString(len uint, str []string) string {
	newStr := ""
	for i := 0; i < int(len); i++ {
		newStr += str[i]
	}
	return newStr
}
