package gluasocket_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/BixData/gluasocket"
	"github.com/BixData/gluasocket/http"
	"github.com/BixData/gluasocket/mime"
	"github.com/BixData/gluasocket/smtp"
	"github.com/BixData/gluasocket/url"
	"github.com/stretchr/testify/assert"
	"github.com/yuin/gopher-lua"
)

func TestHttp(t *testing.T) {
	doTest("httptest.lua", t)
}

func TestMime(t *testing.T) {
	doTest("mimetest.lua", t)
}

func TestSmtp(t *testing.T) {
	doTest("smtptest.lua", t)
}

func TestUrl(t *testing.T) {
	doTest("urltest.lua", t)
}

// ----------------------------------------------------------------------------
func doTest(testScript string, t *testing.T) {
	assert := assert.New(t)

	// Bring up a GopherLua VM
	luaState := lua.NewState()
	defer luaState.Close()

	// Register the Gluasocket module
	luaState.PreloadModule("mime", gluasocket_mime.Loader)
	luaState.PreloadModule("socket", gluasocket.Loader)
	luaState.PreloadModule("socket.http", gluasocket_http.Loader)
	luaState.PreloadModule("socket.smtp", gluasocket_smtp.Loader)
	luaState.PreloadModule("socket.url", gluasocket_url.Loader)

	// Change working directory to where scripts are, so that nested scripts are found
	os.Chdir("testdata/luasocket-test")

	// Run Lua test script
	fmt.Printf("Running test %s\n", testScript)
	assert.NoError(luaState.DoFile(testScript))
}
