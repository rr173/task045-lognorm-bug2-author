package httpapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"task045-lognorm/internal/lognorm"
)

func TestProbe_IngestSuccessReturnsEmptyErrorsArray(t *testing.T) {
	body := bytes.NewBufferString(`{"lines":["hello"]}`)
	req := httptest.NewRequest(http.MethodPost, "/ingest", body)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	New(lognorm.NewService()).Handler().ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status=%d want 200; body=%s", rec.Code, rec.Body.String())
	}
	var out map[string]json.RawMessage
	if err := json.Unmarshal(rec.Body.Bytes(), &out); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	var errors []json.RawMessage
	if err := json.Unmarshal(out["errors"], &errors); err != nil {
		t.Fatalf("decode errors: %v", err)
	}
	if errors == nil {
		t.Fatal("errors must be an empty JSON array, not null")
	}
	if len(errors) != 0 {
		t.Fatalf("errors=%d want 0", len(errors))
	}
}
