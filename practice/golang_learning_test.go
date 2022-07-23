package practice

import (
	"context"
	"github.com/stretchr/testify/assert"
	"strings"
	"sync"
	"testing"
	"time"
)

// golang 학습 테스트
func TestGolang(t *testing.T) {
	t.Run("string test", func(t *testing.T) {
		str := "Ann,Jenny,Tom,Zico"
		actual := strings.Split(str, ",") // TODO str을 , 단위로 잘라주세요.
		expected := []string{"Ann", "Jenny", "Tom", "Zico"}
		//TODO assert 문을 활용해 actual과 expected를 비교해주세요.
		assert.EqualValues(t, actual, expected)
	})

	t.Run("goroutine에서 slice에 값 추가해보기", func(t *testing.T) {
		var numbers []int
		var wg sync.WaitGroup
		wg.Add(100)
		for i := 0; i < 100; i++ {
			go func(numbers []int, i int) {
				// TODO numbers에 i 값을 추가해보세요.
				defer wg.Done()
				numbers = append(numbers, i)

			}(numbers, i)
		}
		wg.Wait()

		var expected []int // actual : [0 1 2 ... 99]
		// TODO expected를 만들어주세요.
		var wg2 sync.WaitGroup
		wg2.Add(100)
		for i := 0; i < 100; i++ {
			go func(expected []int, i int) {
				defer wg2.Done()
				expected = append(expected, i)
			}(expected, i)
		}

		wg.Wait()
		wg2.Wait()
		assert.ElementsMatch(t, expected, numbers)
	})

	t.Run("fan out, fan in", func(t *testing.T) {
		/*
			TODO 주어진 코드를 수정해서 테스트가 통과하도록 해주세요!

			- inputCh에 1, 2, 3 값을 넣는다.
			- inputCh로 부터 값을 받아와, value * 10 을 한 후 outputCh에 값을 넣어준다.
			- outputCh에서 읽어온 값을 비교한다.
		*/

		inputCh := generate()
		outputCh := make(chan int)

		go func() {
			defer close(outputCh)
			for {
				select {
				case value, ok := <-inputCh:
					if !ok {
						return
					}
					outputCh <- value * 10

				}
			}
		}()

		var actual []int
		for value := range outputCh {
			actual = append(actual, value)
		}
		expected := []int{10, 20, 30}
		assert.Equal(t, expected, actual)
	})

	t.Run("context timeout", func(t *testing.T) {
		startTime := time.Now()
		add := time.Second * 3
		ctx, cancel := context.WithTimeout(context.Background(), add) // TODO 3초후에 종료하는 timeout context로 만들어주세요.
		defer cancel()
		var endTime time.Time

		select {
		case <-ctx.Done():
			endTime = time.Now()
			break
		}

		assert.True(t, endTime.After(startTime.Add(add)))
	})

	t.Run("context deadline", func(t *testing.T) {
		startTime := time.Now()
		add := time.Second * 3
		ctx, cancel := context.WithDeadline(context.Background(), startTime.Add(add)) // TODO 3초후에 종료하는 timeout context로 만들어주세요.
		defer cancel()

		var endTime time.Time
		select {
		case <-ctx.Done():
			endTime = time.Now()
			break
		}

		assert.True(t, endTime.After(startTime.Add(add)))
	})

	t.Run("context value", func(t *testing.T) {
		// context에 key, value를 추가해보세요.
		contextTest := context.Background()
		contextTest = context.WithValue(contextTest, "testKey", "testValue")
		// 추가된 key, value를 호출하여 assert로 값을 검증해보세요.
		assert.Equal(t, "testValue", contextTest.Value("testKey"))
		// 추가되지 않은 key에 대한 value를 assert로 검증해보세요.
		assert.Nil(t, contextTest.Value("can not find key"))

	})
}

func generate() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= 3; i++ {
			ch <- i
		}
	}()
	return ch
}
