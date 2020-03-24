package services

import (
	"context"
	"fmt"
	"log"
	"time"
)

func SampleHandler(ctx context.Context) error {
	time.Sleep(time.Second * 1)
	msg := ctx.Value("message").(int)
	if msg == 3 {
		return fmt.Errorf("message is 3!")
	}

	if msg == 5 {
		panic("message is 5!")
	}
	log.Printf("done processing, message is %d", msg)
	return nil
}
