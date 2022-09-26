package check

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	levelD = iota
	LevelC
	LevelB
	LevelA
	LevelS
)

func Check(minLength, maxLength, minLevel int, pwd string) error {
	if len(pwd) < minLength {
		return fmt.Errorf("BAD PASSWORD: The password is shorter than %d characters", minLength)
	}
	if len(pwd) > maxLength {
		return fmt.Errorf("BAD PASSWORD: The password is logner than %d characters", maxLength)
	}

	var level int = levelD
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`, `[~!@#$%^&*?_-]+`}
	for _, pattern := range patternList {
		match, _ := regexp.MatchString(pattern, pwd)
		if match {
			level++
		}
	}

	if level < minLevel {
		return errors.New("The password does not satisfy the current policy requirements. ")
	}
	return nil
}
