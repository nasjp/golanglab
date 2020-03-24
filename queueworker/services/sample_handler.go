package services

import (
	"context"
	"fmt"
	"log"
	"time"
)

func SampleHandler(ctx context.Context) error {
	msg := ctx.Value("message").(int)
	if msg == 3 {
		return fmt.Errorf("message is 3!")
	}

	if msg == 5 {
		panic("message is 5!")
	}

	time.Sleep(time.Second * 1)
	log.Printf("Done message: %d", msg)
	return nil
}
