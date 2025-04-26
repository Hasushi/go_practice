package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/internal/router"
)
func TestCreateTodo(t *testing.T) {
	e := router.NewRouter()

	ts := httptest.NewServer(e)
	defer ts.Close()

	tests := []struct {
		name    string
		todo    map[string]string
		wantResponse string
		wantStatus int
		wantError bool
	}{
		{
			name: "title is empty",
			todo: map[string]string{
				"title": "",
				"description": "This is a test todo",
			},
			wantResponse: "validation error",
			wantStatus: http.StatusBadRequest,
			wantError: true,
		},
		{
			name: "valid todo",
			todo: map[string]string{
				"title": "Test Todo",
				"description": "This is a test todo",
			},
			wantResponse: "success",
			wantStatus: http.StatusOK,
			wantError: false,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			todoJSON, err := json.Marshal(tt.todo)
			if err != nil {
				t.Fatalf("failed to marshal todo: %v", err)
			}

			req, err := http.NewRequest(http.MethodPost, ts.URL+"/todos", bytes.NewBuffer(todoJSON))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("failed to send request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.wantStatus {
				t.Errorf("expected status %d, got %d", tt.wantStatus, resp.StatusCode)
			}

			if tt.wantError {
				var responseBody map[string]string
				json.NewDecoder(resp.Body).Decode(&responseBody)
				if responseBody["error"] != tt.wantResponse {
					t.Errorf("expected error message %q, got %q", tt.wantResponse, responseBody["error"])
				}
			}
		})
	}

}

func TestGetTodo(t *testing.T) {

}

func TestGetTodos(t *testing.T) {

}

func TestUpdateTodo(t *testing.T) {

}

func TestDeleteTodo(t *testing.T) {

}

