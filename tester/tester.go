package tester

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"
)

type TestCase[I any, O any] struct {
	Input    I
	Expected O
}

type Result struct {
	Name     string
	Time     time.Duration
	Total    int
	Passed   int
	Messages []string
}

type Test[I any, O any] struct {
	Name     string
	Cases    []TestCase[I, O]
	TestFunc func(I) O
}

// Implement RunnableTest
func (t *Test[I, O]) Run(ch chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	passed := 0
	messages := []string{}
	var totalTime time.Duration

	for i, c := range t.Cases {
		start := time.Now()
		var result O
		var elapsed time.Duration
		status := "Failed ❌"

		func() {
			defer func() {
				if r := recover(); r != nil {
					elapsed = time.Since(start)
					messages = append(messages, fmt.Sprintf(
						"Test %d: | Expected: %v | Got: <panic> | Panic ⚠️ (%v) | Took: %s",
						i, c.Expected, r, elapsed,
					))

				}
			}()

			result = t.TestFunc(c.Input)
			elapsed = time.Since(start)

			// Use DeepEqual so it works for slices, maps, structs
			if reflect.DeepEqual(result, c.Expected) {
				status = "Passed ✅"
				passed++
			}

			messages = append(messages, fmt.Sprintf(
				"Test %d: | Expected: %v | Got: %v | %s | Took: %s",
				i, c.Expected, result, status, elapsed,
			))

		}()

		totalTime += elapsed
	}

	ch <- Result{
		Name:     t.Name,
		Time:     totalTime,
		Total:    len(t.Cases),
		Passed:   passed,
		Messages: messages,
	}
}

type RunnableTest interface {
	Run(ch chan<- Result, wg *sync.WaitGroup)
}

type Tester struct {
	tests []RunnableTest
}

func (t *Tester) Add(test RunnableTest) {
	t.tests = append(t.tests, test)
}

func (t *Tester) Run() {
	var wg sync.WaitGroup
	ch := make(chan Result, len(t.tests))

	for _, test := range t.tests {
		wg.Add(1)
		go test.Run(ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var allResults []Result
	for r := range ch {
		printResult(r)
		allResults = append(allResults, r)
	}

	printOverallSummary(allResults)
}

func printResult(r Result) {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf(" Test Suite: %s\n", r.Name)
	fmt.Println(strings.Repeat("=", 60))

	for _, msg := range r.Messages {
		fmt.Println(msg)
	}

	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("Summary | Total: %d | Passed: %d | Failed: %d | Time: %s\n",
		r.Total, r.Passed, r.Total-r.Passed, r.Time,
	)

	fmt.Println(strings.Repeat("=", 60))
	fmt.Println()
}

func printOverallSummary(results []Result) {
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println(" OVERALL TEST SUMMARY")
	fmt.Println(strings.Repeat("=", 80))

	totalTests := 0
	totalPassed := 0
	totalFailed := 0
	var totalTime time.Duration

	for _, r := range results {
		totalTests += r.Total
		totalPassed += r.Passed
		totalFailed += (r.Total - r.Passed)
		totalTime += r.Time

		status := "✅ PASS"
		if r.Passed != r.Total {
			status = "❌ FAIL"
		}

		fmt.Printf("| %-40s | %-8s | %2d/%2d | %-10s |\n",
			r.Name, status, r.Passed, r.Total, r.Time)
	}

	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("| %-40s | %-8s | %2d/%2d | %-10s |\n",
		"TOTAL", "", totalPassed, totalTests, totalTime)
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println()
}
