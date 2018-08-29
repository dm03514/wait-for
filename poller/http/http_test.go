package http

import "testing"

func TestNewHTTP_InvalidURL(t *testing.T) {
	_, err := NewHTTP("", "http://192.168.0.%31:8080/", "")
	if err == nil {
		t.Errorf("expected error")
	}
}

func Test_IsValidResponseCode(t *testing.T) {
	t.Skip()
}

func TestHTTP_CheckReady(t *testing.T) {
	t.Skip()
}
