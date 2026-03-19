package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func main() {
	//todo := context.TODO()
	//ctx := context.Background()
	//
	//ctx, cancel := context.WithCancel(ctx)
	//defer cancel()
	//
	//
	//ctx, cancel := context.WithCancelCause(ctx)
	//defer cancel(fmt.Errorf("job not needed"))
	//
	//ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	//defer cancel()
	//
	//ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))
	//defer cancel()
}

//По соглашению Context всегда передается первым параметром в функции, обычно именуясь ctx.
//
//database/sql.(*DB).QueryContext(ctx context.Context, query string, args ...interface{}) (*Rows, error)
//database/sql.(*DB).ExecContext(ctx context.Context, query string, args ...interface{}) (Result, error)
//net/http.NewRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*Request, error)
//golang.org/x/sync/errgroup.WithContext(ctx context.Context) (*Group, context.Context)
//...

type User struct {
	ID   int
	Name string
}

func fetchUserData(userID string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	url := fmt.Sprintf("https://api.example.com/users/%s", userID)
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, fmt.Errorf("API timeout after 2s")
		}
		return nil, err
	}
	defer resp.Body.Close()

	var user User
	json.NewDecoder(resp.Body).Decode(&user)
	return &user, nil
}
