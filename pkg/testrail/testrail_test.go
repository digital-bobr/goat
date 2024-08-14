package testrail

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoadParams(t *testing.T) {
	LoadParams("user", "pass", "1", "http://example.com/")
	if username != "user" || password != "pass" || projectID != "1" || url != "http://example.com/" {
		t.Errorf("LoadParams failed to set the correct values")
	}
}

func TestUpdateTestRun(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	LoadParams("user", "pass", "1", server.URL+"/")
	testRun := &TestRun{ID: 1, Name: "Test Run"}
	err := UpdateTestRun(testRun)
	if err != nil {
		t.Errorf("UpdateTestRun failed: %v", err)
	}
}

func TestCreateTestRun(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": 1}`))
	}))
	defer server.Close()

	LoadParams("user", "pass", "1", server.URL+"/")
	testRun := &TestRun{Name: "Test Run"}
	id, err := createTestRun(testRun)
	if err != nil || id != 1 {
		t.Errorf("createTestRun failed: %v", err)
	}
}

func TestFindTestRunByName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"runs":[{"id":1,"name":"Test Run"}]}`))
	}))
	defer server.Close()

	LoadParams("user", "pass", "1", server.URL+"/")
	id, err := findTestRunByName("Test Run")
	if err != nil || id != 1 {
		t.Errorf("findTestRunByName failed: %v", err)
	}
}

func TestAddResultsToTestRun(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	LoadParams("user", "pass", "1", server.URL+"/")
	results := &Results{Results: []Result{{CaseID: 1, StatusID: 1, Comment: "Passed"}}}
	err := addResultsToTestRun(results, 1)
	if err != nil {
		t.Errorf("addResultsToTestRun failed: %v", err)
	}
}

func TestReportResults(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/get_runs/1" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"runs":[]}`))
		} else if r.URL.Path == "/add_run/1" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"id": 1}`))
		} else if r.URL.Path == "/add_results_for_cases/1" {
			w.WriteHeader(http.StatusOK)
		}
	}))
	defer server.Close()

	LoadParams("user", "pass", "1", server.URL+"/")
	testRun := &TestRun{Name: "New Run"}
	results := &Results{Results: []Result{{CaseID: 1, StatusID: 1, Comment: "Passed"}}}
	err := ReportResults(testRun, results)
	if err != nil {
		t.Errorf("ReportResults failed: %v", err)
	}
}
