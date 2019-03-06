// Copyright 2018 The WPT Dashboard Project. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package api

import (
	"net/http"

	"github.com/web-platform-tests/wpt.fyi/api/checks"
	"github.com/web-platform-tests/wpt.fyi/api/receiver"
	"github.com/web-platform-tests/wpt.fyi/shared"
)

func apiResultsUploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST is supported", http.StatusMethodNotAllowed)
		return
	}

	ctx := shared.NewAppEngineContext(r)
	a := receiver.NewAPI(ctx)
	receiver.HandleResultsUpload(a, w, r)
}

func apiResultsCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST is supported", http.StatusMethodNotAllowed)
		return
	}

	ctx := shared.NewAppEngineContext(r)
	a := receiver.NewAPI(ctx)
	s := checks.NewAPI(ctx)
	receiver.HandleResultsCreate(a, s, w, r)
}
