package gojieba

import (
	"github.com/CharlotteFTD/gojieba/deps/cppjieba"
	"github.com/CharlotteFTD/gojieba/deps/limonp"
	"github.com/CharlotteFTD/gojieba/dict"
)

func init() {
	dict.Init()
	limonp.Init()
	cppjieba.Init()
}
