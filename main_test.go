package main

import (
	"fmt"
	"testing"

	"github.com/qadahss/supertest"
)

func TestSomeFunc(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		defer goleak.VerifyNone(t)
		someFunc()

	})

	t.Run("test2", func(t *testing.T) {
		good, _ := goleak.Goroutines()
		defer func() {
			leaked, _ := goleak.Goroutines()
			fmt.Println("gtest")
			if len(good) != len(leaked) {
				t.Fail()
			}
		}()

		go leak(make(chan string))
		someFunc()

	})

	t.Run("test3", func(t *testing.T) {
		good, _ := goleak.Goroutines()
		defer func() {
			leaked, _ := goleak.Goroutines()
			fmt.Println("gtest")
			if len(good) != len(leaked) {
				t.Fail()
			}
		}()
		someFunc()

	})

	t.Run("test4", func(t *testing.T) {
		good, _ := goleak.Goroutines()
		defer func() {
			leaked, _ := goleak.Goroutines()
			fmt.Println("gtest")
			if len(good) != len(leaked) {
				t.Fail()
			}
		}()
		go leak(make(chan string))
		someFunc()

	})

	t.Run("test5", func(t *testing.T) {
		good, _ := goleak.Goroutines()
		defer func() {
			leaked, _ := goleak.Goroutines()
			fmt.Println("gtest")
			if len(good) != len(leaked) {
				t.Fail()
			}
		}()
		someFunc()

	})
}

func leak(test chan string) {
	<-test
}
