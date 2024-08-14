package testrail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var (
	url       string
	username  string
	password  string
	projectID string
)

type Results struct {
	Results []Result `json:"results"`
}

type Result struct {
	CaseID   int    `json:"case_id"`
	StatusID int    `json:"status_id"` // 1 - Passed; 5 - Failed
	Comment  string `json:"comment"`
}

type TestRun struct {
	ID          int    `json:"id,omitempty"`
	SuiteID     int    `json:"suite_id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	MilestoneID int    `json:"milestone_id,omitempty"`
	AssignedTo  int    `json:"assignedto_id,omitempty"`
	IncludeAll  bool   `json:"include_all"`
	CaseIDs     []int  `json:"case_ids,omitempty"`
}

func LoadParams(testRailUsername, testRailPassword, testRailProjectID string, testRailUrl string) {
	username = testRailUsername
	password = testRailPassword
	projectID = testRailProjectID
	url = testRailUrl // Ex.: "https://bobr.testrail.net/index.php?/api/v2/"
}

func ReportResults(run *TestRun, results *Results) error {
	testRunId, err := createIfDoesNotExistTestRun(run)
	if err != nil {
		return err
	}
	return addResultsToTestRun(results, testRunId)
}

func UpdateTestRun(testRun *TestRun) error {
	url := fmt.Sprintf("%supdate_run/%d", url, testRun.ID)
	body, err := json.Marshal(testRun)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("TestRail API error: %s", respBody)
	}
	fmt.Printf("Test Run '%s' successfully updated", testRun.Name)

	return nil
}

func createIfDoesNotExistTestRun(testRun *TestRun) (int, error) {
	testRunId, err := findTestRunByName(testRun.Name)
	if err != nil {
		if testRunId > 0 {
			return testRunId, nil
		}
		return testRunId, err
	}
	if testRunId > 0 {
		fmt.Printf("\nTestRailTestRun '%s' with ID '%d' already exists.\n", testRun.Name, testRunId)
		return testRunId, nil
	} else {
		return createTestRun(testRun)
	}
}

func createTestRun(testRun *TestRun) (int, error) {
	id := 0
	url := fmt.Sprintf("%sadd_run/%s", url, projectID)
	body, err := json.Marshal(testRun)
	if err != nil {
		return id, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return id, err
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return id, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return id, fmt.Errorf("TestRail API error: %s", respBody)
	} else {
		fmt.Printf("Test Run '%s' successfully created", testRun.Name)
		var createdRun TestRun
		if err := json.NewDecoder(resp.Body).Decode(&createdRun); err != nil {
			return 0, err
		}
		return createdRun.ID, nil
	}
}

func findTestRunByName(name string) (int, error) {
	url := fmt.Sprintf("%sget_runs/%s", url, projectID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	req.SetBasicAuth(username, password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("TestRail API error: %s", respBody)
	}

	var result struct {
		Runs []TestRun `json:"runs"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	for _, run := range result.Runs {
		if strings.EqualFold(run.Name, name) {
			fmt.Printf("Test Run '%s' with ID '%d' already exists.", name, run.ID)
			return run.ID, nil
		}
	}

	return 0, nil
}

func addResultsToTestRun(results *Results, testRunId int) error {
	url := fmt.Sprintf("%sadd_results_for_cases/%d", url, testRunId)
	body, err := json.Marshal(results)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("TestRail API error: %s", respBody)
	}

	return nil
}
