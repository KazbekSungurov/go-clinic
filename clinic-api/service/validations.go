package service

import (
	"errors"
	"regexp"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

var (
	latin       = regexp.MustCompile(`\p{Latin}`)
	cyrillic    = regexp.MustCompile(`[\p{Cyrillic}]`)
	phone       = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	noSQL       = regexp.MustCompile(`\b(ALTER|CREATE|DELETE|DROP|EXEC(UTE){0,1}|INSERT( +INTO){0,1}|MERGE|SELECT|UPDATE|UNION( +ALL){0,1})\b`)
	onlyLetters = regexp.MustCompile("[^a-zA-Zа-яА-Я]+")

	ErrContainsSQL        = errors.New("no SQL commands allowed to input")
	ErrInvalidBirthDate   = errors.New("invalid date of birth. Age must be from 18 to 100. Date format RFC3339")
	ErrInvalidPhoneNumber = errors.New("invalid phone number format")
	ErrInvalidAlphabet    = errors.New("only latin or cyrillic symblos allowed")
	ErrInvalidSymbol      = errors.New("invalid symbol used. Only space and '-' symbols allowed")
)

func IsLetterHyphenSpaces(value interface{}) error {
	s := value.(string)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "-", "", -1)

	err := is.UTFLetter.Validate(s)
	if err != nil {
		return ErrInvalidSymbol
	}
	if cyrillic.MatchString(s) && !latin.MatchString(s) {
		return nil
	} else if latin.MatchString(s) && !cyrillic.MatchString(s) {
		return nil
	}
	return ErrInvalidAlphabet
}

func IsPhone(value interface{}) error {
	s := value.(string)

	if phone.MatchString(s) {
		return nil
	}
	return ErrInvalidPhoneNumber
}

// IsValidBirthDate ...
func IsValidBirthDate(value interface{}) error {
	t := time.Now()
	d := value.(*time.Time)
	err := validation.Validate(d.Format(time.RFC3339), validation.Date(time.RFC3339).Max(t.AddDate(-18, 0, 0)).Min(t.AddDate(-100, 0, 0)))
	if err != nil {
		return ErrInvalidBirthDate
	}
	return nil
}

// IsSQL ...
func IsSQL(value interface{}) error {
	s := value.(string)

	if noSQL.MatchString(strings.ToUpper(s)) {
		return ErrContainsSQL
	}

	str := onlyLetters.ReplaceAllString(s, "")

	if noSQL.MatchString(strings.ToUpper(str)) {
		return ErrContainsSQL
	}
	return nil
}
