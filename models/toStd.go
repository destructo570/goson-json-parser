package models

func ToStd(v JsonValue) any {
	switch val := v.(type) {
	case JsonObject:
		m := make(map[string]any, len(val.Fields))
		for k, child := range val.Fields {
			m[k] = ToStd(child)
		}
		return m

	case JsonArray:
		out := make([]any, 0, len(val.Elements))
		for _, child := range val.Elements {
			out = append(out, ToStd(child))
		}
		return out

	case JsonString:
		return val.Value

	case JsonNumber:
		return val.Value

	case JsonBool:
		return val.Value

	case JsonNull:
		return nil

	default:
		// should not happen if all types covered
		return nil
	}
}
