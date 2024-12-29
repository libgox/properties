package properties

import (
	"fmt"
	"strconv"
	"strings"
)

type Properties struct {
	properties map[string]string
}

func LoadFromBytes(bytes []byte) (*Properties, error) {
	properties := &Properties{
		properties: make(map[string]string),
	}

	split := strings.Split(string(bytes), "\n")
	for _, line := range split {
		if strings.HasPrefix(line, "#") {
			continue
		}
		str := strings.TrimSpace(line)
		if str == "" {
			continue
		}
		parts := strings.Split(str, "=")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid property line: %s", line)
		}
		properties.properties[parts[0]] = parts[1]
	}
	return properties, nil
}

func (p *Properties) Get(key string) (string, error) {
	val, ok := p.properties[key]
	if !ok {
		return "", &NoSuchPropertyError{Key: key}
	}
	return val, nil
}

func (p *Properties) GetOr(key, defaultValue string) string {
	val, ok := p.properties[key]
	if !ok {
		return defaultValue
	}
	return val
}

func (p *Properties) GetInt(key string) (int, error) {
	aux, ok := p.properties[key]
	if !ok {
		return 0, &NoSuchPropertyError{Key: key}
	}
	res, err := strconv.Atoi(aux)
	if err != nil {
		return 0, fmt.Errorf("invalid integer value for %s: %s", key, aux)
	}
	return res, nil
}

func (p *Properties) GetIntOr(key string, defaultValue int) (int, error) {
	aux, ok := p.properties[key]
	if !ok {
		return defaultValue, nil
	}
	res, err := strconv.Atoi(aux)
	if err != nil {
		return defaultValue, fmt.Errorf("invalid integer value for %s: %s", key, aux)
	}
	return res, nil
}

func (p *Properties) MustGetInt(key string) int {
	val, err := p.GetInt(key)
	if err != nil {
		panic(err)
	}
	return val
}

func (p *Properties) MustGetIntOr(key string, defaultValue int) int {
	val, err := p.GetIntOr(key, defaultValue)
	if err != nil {
		panic(err)
	}
	return val
}
