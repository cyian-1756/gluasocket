package gluasocket_socketcore

import (
	"github.com/yuin/gopher-lua"
)

// ----------------------------------------------------------------------------

var exports = map[string]lua.LGFunction{
	"connect": connectFn,
	"gettime": gettimeFn,
	"select":  selectFn,
	"sink":    sinkFn,
	"skip":    skipFn,
	"sleep":   sleepFn,
	"source":  sourceFn,
	"tcp":     tcpFn,
	"udp":     udpFn,
}

// ----------------------------------------------------------------------------

func Loader(l *lua.LState) int {
	mod := l.SetFuncs(l.NewTable(), exports)
	l.Push(mod)

	l.SetField(mod, "_DEBUG", lua.LBool(false))
	l.SetField(mod, "_VERSION", lua.LString("0.0.0")) // TODO

	registerClientType(l, mod)
	registerDNSType(l, mod)

	return 1
}

// ----------------------------------------------------------------------------

func registerClientType(l *lua.LState, mod *lua.LTable) {
	mt := l.NewTypeMetatable("client")
	l.SetField(mod, "client", mt)
	l.SetField(mt, "__index", l.SetFuncs(l.NewTable(), clientMethods))
}

// ----------------------------------------------------------------------------

func registerDNSType(L *lua.LState, mod *lua.LTable) {
	table := L.NewTable()
	L.SetFuncs(table, dnsMethods)
	L.SetField(mod, "dns", table)
}
