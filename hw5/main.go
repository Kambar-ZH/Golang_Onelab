package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Execute(tasks []func(ctx context.Context) error, E int) error {
	errorCount := 0
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	errCh := make(chan error)
	for i := 0; i < len(tasks); i++ {
		wg.Add(1)
		id := i
		go func() {
			defer wg.Done()
			if err := tasks[id](ctx); err != nil {
				mu.Lock()
				errorCount++
				mu.Unlock()
			}
			mu.Lock()
			fmt.Println(errorCount)
			if errorCount >= E {
				cancel()
				errCh <- fmt.Errorf("error count limit exceeded")
			}
			defer mu.Unlock()
		}()
	}
	waitCh := make(chan bool)
	go func() {
		wg.Wait()
		waitCh <- true
	}()
	for {
		select {
		case err := <-errCh:
			return err
		case <-waitCh:
			return nil
		}
	}
}

func main() {
	tasks := []func(ctx context.Context) error{
		func(ctx context.Context) error {
			time.Sleep(2 * time.Second)
			fmt.Println("Task 1 done job")
			return nil
		},
		func(ctx context.Context) error {
			time.Sleep(3 * time.Second)
			fmt.Println("Task 2 done job")
			return fmt.Errorf("some problem")
		},
		func(ctx context.Context) error {
			for i := 0; i < 10; i++ {
				select {
				case <-ctx.Done():
					fmt.Println("Context canceled task 3")
					return nil
				default:
					time.Sleep(500 * time.Millisecond)
				}
			}
			fmt.Println("Task 3 done job")
			return nil
		},
		func(ctx context.Context) error {
			for i := 0; i < 5; i++ {
				select {
				case <-ctx.Done():
					fmt.Println("Context canceled task 4")
					return nil
				default:
					time.Sleep(500 * time.Millisecond)
				}
			}
			fmt.Println("Task 4 done job")
			return nil
		},
	}
	if err := Execute(tasks, 1); err != nil {
		fmt.Println(err)
	}
	time.Sleep(10 * time.Second)
}