package uid

import (
	"github.com/segmentio/ksuid"
)

type Account string
type Decision string

const (
	AccountPrefix  = "acc"
	DecisionPrefix = "dec"
)

func MakeAccount() Account {
	return Account(prefixedKSUID(AccountPrefix))
}

func (uid Account) String() string {
	return string(uid)
}

func MakeDecision() Decision {
	return Decision(prefixedKSUID(DecisionPrefix))
}

func (uid Decision) String() string {
	return string(uid)
}

func prefixedKSUID(prefix string) string {
	return prefix + "_" + ksuid.New().String()
}
