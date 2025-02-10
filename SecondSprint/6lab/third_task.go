package main

import "time"

func QuizRunner(questions, answers []string, answerCh chan string) int {
	var res int
	for i := 0; i < len(questions); i++ {
		select {
		case ans := <-answerCh:
			if ans == answers[i] {
				res++
			} else {
				continue
			}
		case <-time.After(time.Second):
			continue
		}
	}
	return res
}
