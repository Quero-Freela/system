package utils

import (
	"encoding/json"
	"errors"
	"github.com/Quero-Freela/system/server/exceptions"
	"os"
	"reflect"
	"strconv"
)

type SecretKey string

const (
	DatabaseDsn  SecretKey = "DATABASE_DSN"
	GoogleAPI    SecretKey = "GOOGLE_API"
	MicrosoftAPI SecretKey = "MICROSOFT_API"
	GithubAPI    SecretKey = "GITHUB_API"
	FacebookAPI  SecretKey = "FACEBOOK_API"
	LinkedinAPI  SecretKey = "LINKEDIN_API"
	JwtSecret    SecretKey = "JWT_SECRET"
)

func GetSecret[T interface{}](secretKey SecretKey) (*T, error) {
	value := os.Getenv(string(secretKey))

	if value != "" {
		return parseSecretValue[T](value)
	}

	return nil, exceptions.NewNotFoundError("secret not found")
}

func parseSecretValue[T interface{}](value string) (*T, error) {
	var secret T

	if isPrimitive[T]() {
		return parseStringToPrimitive[T](value)
	}

	err := json.Unmarshal([]byte(value), &secret)

	if err != nil {
		return nil, err
	}

	return &secret, nil
}

func isPrimitive[T interface{}]() bool {
	var secret T
	tp := reflect.TypeOf(secret)

	switch tp.Kind() {
	case reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Bool:
		return true

	default:
		return false
	}
}

func parseStringToPrimitive[T interface{}](value string) (*T, error) {
	var err error
	var secint interface{}
	var secret T

	tp := reflect.TypeOf(secret)

	switch tp.Kind() {
	case reflect.String:
		secint, err = value, nil
	case reflect.Int:
		secint, err = strconv.Atoi(value)
	case reflect.Int8:
		secint, err = strconv.ParseInt(value, 10, 8)
	case reflect.Int16:
		secint, err = strconv.ParseInt(value, 10, 16)
	case reflect.Int32:
		secint, err = strconv.ParseInt(value, 10, 32)
	case reflect.Int64:
		secint, err = strconv.ParseInt(value, 10, 64)
	case reflect.Uint:
		secint, err = strconv.ParseUint(value, 10, 0)
	case reflect.Uint8:
		secint, err = strconv.ParseUint(value, 10, 8)
	case reflect.Uint16:
		secint, err = strconv.ParseUint(value, 10, 16)
	case reflect.Uint32:
		secint, err = strconv.ParseUint(value, 10, 32)
	case reflect.Uint64:
		secint, err = strconv.ParseUint(value, 10, 64)
	case reflect.Float32:
		secint, err = strconv.ParseFloat(value, 32)
	case reflect.Float64:
		secint, err = strconv.ParseFloat(value, 64)
	case reflect.Bool:
		secint, err = strconv.ParseBool(value)

	default:
		return nil, errors.New("type not supported")
	}

	secret = secint.(T)
	return &secret, err
}
