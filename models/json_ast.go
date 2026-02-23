package models

type JsonValue interface {
}

type JsonObject struct {
	Fields map[string]JsonValue
}

type JsonArray struct {
	Elements []JsonValue
}

type JsonString struct {
	Value string
}

type JsonNumber struct {
	Value float64
}

type JsonBool struct {
	Value bool
}

type JsonNull struct{}
