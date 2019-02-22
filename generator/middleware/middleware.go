package middleware

import "github.com/guiyomh/charlatan/contract"

type middleware func(value contract.Value) contract.Value
