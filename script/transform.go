package script

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/venti-org/crawler-core/downloader"
)

func TransformRequest(item interface{}) (Request, error) {
	switch item := item.(type) {
	case string:
		return downloader.NewHttpRequestWithError(
			http.NewRequest("GET", item, nil))
	case map[string]interface{}:
		url, ok := item["url"].(string)
		if !ok {
			return nil, fmt.Errorf("Js Parse return map not contain url")
		}
		method, _ := item["method"].(string)
		body, _ := item["body"].(string)
		headers, _ := item["headers"].(map[string]interface{})
		meta, _ := item["meta"].(map[string]interface{})
		if len(method) == 0 {
			method = "GET"
		}
		var reader io.Reader
		if len(body) != 0 {
			reader = strings.NewReader(body)
		}
		if request, err := http.NewRequest(method, url, reader); err != nil {
			return nil, err
		} else {
			for name, value := range headers {
				if value, ok := value.(string); ok {
					request.Header.Set(name, value)
				}
			}
			http_request, err := downloader.NewHttpRequest(request)
			if http_request != nil {
				for k, v := range meta {
					http_request.SetMetaValue(k, v)
				}
			}
			return http_request, err
		}
	case Request:
		return item, nil
	default:
		return nil, fmt.Errorf("TransformRequest not support %T", item)
	}
}

func TransformItem(item interface{}) (interface{}, error) {
	switch item := item.(type) {
	case Item:
		return item, nil
	case Request:
		return item, nil
	default:
		return nil, fmt.Errorf("TransformItem return not support %T", item)
	}
}

func TransformSlice(item interface{}) interface{} {
	v := reflect.ValueOf(item)
	if v.Kind() == reflect.Slice {
		var items []interface{}
		for i := 0; i < v.Len(); i += 1 {
			items = append(items, v.Index(i).Interface())
		}
		return items
	}
	return item
}
