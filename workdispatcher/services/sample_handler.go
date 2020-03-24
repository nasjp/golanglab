package services

import (
	"context"
	"log"
	"math/rand"
	"time"
)

func SampleHandler(ctx context.Context) error {
	log.Printf("start processing %d", ctx.Value("id").(int64))
	// 便宜的に時間を外部APIとする
	t := time.NewTimer(time.Duration(rand.Intn(3)) * time.Second)
	defer t.Stop()

	<-t.C

	log.Printf("done processing %d", ctx.Value("id").(int64))
	return nil
}
