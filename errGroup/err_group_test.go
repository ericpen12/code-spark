package errGroup

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(1)

	for i := 0; i < 10; i++ {
		v := i
		g.Go(func() error {

			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return nil
			default:
			}
			if v == 4 {
				//return errors.New("v == 4")
			}
			fmt.Println(v)
			time.Sleep(time.Second)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println("Error happened: ", err)
	}
}
