// +build medium

package shared_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/web-platform-tests/wpt.fyi/shared"
	"github.com/web-platform-tests/wpt.fyi/shared/sharedtest"
)

func TestTestRunIDs_LoadTestRuns(t *testing.T) {
	testRuns := make(shared.TestRuns, 2)
	testRuns[0].BrowserName = "chrome"
	testRuns[0].BrowserVersion = "63.0"
	testRuns[0].OSName = "linux"
	testRuns[0].Revision = "1234567890"
	testRuns[0].ResultsURL = "/static/chrome-63.0-linux-summary.json.gz"

	testRuns[1].BrowserName = "firefox"
	testRuns[1].BrowserVersion = "60.0"
	testRuns[1].OSName = "linux"
	testRuns[1].Revision = "0987654321"
	testRuns[1].ResultsURL = "/static/firefox-60.0-linux-summary.json.gz"

	ctx, done, err := sharedtest.NewAEContext(true)
	assert.Nil(t, err)
	defer done()

	keys := make([]shared.Key, 0, len(testRuns))
	store := shared.NewAppEngineDatastore(ctx, false)
	for range testRuns {
		keys = append(keys, store.NewIncompleteKey("TestRun"))
	}
	keys, err = store.PutMulti(keys, testRuns)
	assert.Nil(t, err)
	for i, key := range keys {
		testRuns[i].ID = key.IntID()
	}

	trs, err := testRuns.GetTestRunIDs().LoadTestRuns(store)
	assert.Nil(t, err)
	assert.Equal(t, testRuns, trs)
}
