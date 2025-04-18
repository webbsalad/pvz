package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/webbsalad/pvz/internal/app"
)

var (
	baseURL   string
	modToken  string
	empToken  string
	appCtx    context.Context
	cancelApp context.CancelFunc
)

func doRequest(t *testing.T, method, url string, body interface{}, out interface{}, token string) {
	t.Helper()
	b, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("marshal body: %v", err)
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(b))
	if err != nil {
		t.Fatalf("new request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("do request %s %s: %v", method, url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status %d from %s %s", resp.StatusCode, method, url)
	}

	if out != nil {
		if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
			t.Fatalf("decode response from %s %s: %v", method, url, err)
		}
	}
}

func TestMain(m *testing.M) {
	_, thisFile, _, _ := runtime.Caller(0)
	root := filepath.Dir(filepath.Dir(filepath.Dir(thisFile)))

	envTest := filepath.Join(root, ".env.test")
	if err := godotenv.Load(envTest); err != nil { // idk why he cant see it without absolute path (inserted root is literaly ../)
		log.Fatalf("failed to load %s: %v", envTest, err)
	}

	dsn := os.Getenv("testDSN")
	secret := os.Getenv("testJWT_SECRET")
	if dsn == "" || secret == "" {
		log.Fatalf("testDSN or testJWT_SECRET is not set in .env.test")
	}

	os.Setenv("DSN", dsn)
	os.Setenv("JWT_SECRET", secret)

	appCtx, cancelApp = context.WithTimeout(context.Background(), time.Minute)
	a := app.NewApp()
	errCh := make(chan error, 1)
	go func() {
		errCh <- a.Start(appCtx)
	}()
	select {
	case err := <-errCh:
		if err != nil {
			log.Fatalf("app start error: %v", err)
		}
	case <-time.After(500 * time.Millisecond):
	}
	baseURL = "http://localhost:8080"

	modToken = login("moderator")
	empToken = login("employee")

	code := m.Run()

	_ = a.Stop(appCtx)
	cancelApp()

	os.Exit(code)
}

func login(role string) string {
	var resp struct {
		Token string `json:"token"`
	}
	doRequest(&testing.T{}, http.MethodPost, baseURL+"/dummyLogin", map[string]string{"role": role}, &resp, "")
	if resp.Token == "" {
		log.Fatalf("empty %s token", role)
	}
	return resp.Token
}
