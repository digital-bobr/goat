package client

import (
	"testing"
)

func TestGET_GetValue(t *testing.T) {
	var m Method = GET{}
	expected := "GET"
	if m.GetValue() != expected {
		t.Errorf("GET.GetValue() = %v; want %v", m.GetValue(), expected)
	}
}

func TestPOST_GetValue(t *testing.T) {
	var m Method = POST{}
	expected := "POST"
	if m.GetValue() != expected {
		t.Errorf("POST.GetValue() = %v; want %v", m.GetValue(), expected)
	}
}

func TestPUT_GetValue(t *testing.T) {
	var m Method = PUT{}
	expected := "PUT"
	if m.GetValue() != expected {
		t.Errorf("PUT.GetValue() = %v; want %v", m.GetValue(), expected)
	}
}

func TestDELETE_GetValue(t *testing.T) {
	var m Method = DELETE{}
	expected := "DELETE"
	if m.GetValue() != expected {
		t.Errorf("DELETE.GetValue() = %v; want %v", m.GetValue(), expected)
	}
}

func TestPATCH_GetValue(t *testing.T) {
	var m Method = PATCH{}
	expected := "PATCH"
	if m.GetValue() != expected {
		t.Errorf("PATCH.GetValue() = %v; want %v", m.GetValue(), expected)
	}
}

func TestHEAD_GetValue(t *testing.T) {
	var m Method = HEAD{}
	expected := "HEAD"
	if m.GetValue() != expected {
		t.Errorf("HEAD.GetValue() = %v; want %v", m.GetValue(), expected)
	}
}

func TestOPTIONS_GetValue(t *testing.T) {
	var m Method = OPTIONS{}
	expected := "OPTIONS"
	if m.GetValue() != expected {
		t.Errorf("OPTIONS.GetValue() = %v; want %v", m.GetValue(), expected)
	}
}

func TestTRACE_GetValue(t *testing.T) {
	var m Method = TRACE{}
	expected := "TRACE"
	if m.GetValue() != expected {
		t.Errorf("TRACE.GetValue() = %v; want %v", m.GetValue(), expected)
	}
}

func TestCONNECT_GetValue(t *testing.T) {
	var m Method = CONNECT{}
	expected := "CONNECT"
	if m.GetValue() != expected {
		t.Errorf("CONNECT.GetValue() = %v; want %v", m.GetValue(), expected)
	}
}
