package async

import "context"

// Future interface has the method signature for await
type Future interface {
	Await() (interface{}, error)
}

type future struct {
	await func(ctx context.Context) (interface{}, error)
}

func (f future) Await() (interface{}, error) {
	return f.await(context.Background())
}

// Exec executes the async function
func Exec(f func() (interface{}, error)) Future {
	var result interface{}
	var err error
	c := make(chan struct{})
	go func() {
		defer close(c)
		result, err = f()
	}()
	return future{
		await: func(ctx context.Context) (interface{}, error) {
			select {
			case <-ctx.Done():
				return nil,ctx.Err()
			case <-c:
				return result, err
			}
		},
	}
}
