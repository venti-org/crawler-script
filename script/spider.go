package script

import (
	"sync"

	"github.com/venti-org/crawler-core/downloader"
	"github.com/venti-org/crawler-core/pipeline"
	"github.com/robertkrimen/otto"
	"github.com/sirupsen/logrus"
)

type Item = pipeline.Item
type Request = downloader.Request
type HttpRequest = downloader.HttpRequest
type Response = downloader.Response
type HttpResponse = downloader.HttpResponse

const (
	GET_REQUESTS = "get_requests"
	PARSE        = "parse"
	PROCESS_ITEM = "process_item"
)

type ScriptSpider struct {
	ScriptBase
	mutex        sync.Mutex
	seederStatus int
	requests     []Request
	requestIndex int
}

func NewScriptSpiderFromBase(sb *ScriptBase, err error) (*ScriptSpider, error) {
	if err != nil {
		return nil, err
	} else {
		return &ScriptSpider{
			ScriptBase: *sb,
		}, nil
	}
}

func NewScriptSpider(vm *otto.Otto) (*ScriptSpider, error) {
	return NewScriptSpiderFromBase(NewScriptBase(vm))
}

func NewScriptSpiderWithScript(vm *otto.Otto,
	script string) (*ScriptSpider, error) {
	return NewScriptSpiderFromBase(NewScriptBaseWithScript(vm, script))
}

func (s *ScriptSpider) NextRequest() Request {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.seederStatus == -1 {
		return nil
	}
	if s.seederStatus == 0 {
		s.seederStatus = 1
		var err error
		s.requests, err = s.getRequests()
		if err != nil {
			logrus.Errorln(err)
		}
	}
	if s.requestIndex >= len(s.requests) {
		return nil
	} else {
		request := s.requests[s.requestIndex]
		s.requestIndex += 1
		return request
	}
}

func (s *ScriptSpider) getRequests() ([]Request, error) {
	if result, err := s.Call(GET_REQUESTS); err != nil {
		return nil, err
	} else if items, err := result.Export(); err != nil {
		return nil, err
	} else {
		items = TransformSlice(items)
		switch items := items.(type) {
		case []interface{}:
			var requests []Request
			for _, item := range items {
				if item, err := TransformRequest(item); err != nil {
					return nil, err
				} else {
					requests = append(requests, item)
				}
			}
			return requests, nil
		default:
			if item, err := TransformRequest(items); err != nil {
				return nil, err
			} else {
				return []Request{item}, nil
			}
		}
	}
}

func (s *ScriptSpider) Parse(response Response) []interface{} {
	items, err := s.parse(response)
	if err != nil {
		logrus.Errorln(err)
	}
	return items
}

func (s *ScriptSpider) parse(response Response) ([]interface{}, error) {
	if result, err := s.Call(PARSE, response); err != nil {
		return nil, err
	} else if items, err := result.Export(); err != nil {
		return nil, err
	} else {
		items = TransformSlice(items)
		switch items := items.(type) {
		case []interface{}:
			var return_items []interface{}
			for _, item := range items {
				if item, err := TransformItem(item); err != nil {
					return nil, err
				} else {
					return_items = append(return_items, item)
				}
			}
			return return_items, nil
		default:
			if item, err := TransformItem(items); err != nil {
				return nil, err
			} else {
				return []interface{}{item}, nil
			}
		}
	}
}

func (s *ScriptSpider) ProcessItem(item Item) {
	if _, err := s.Call(PROCESS_ITEM, item); err != nil {
		logrus.Errorln(err)
	}
}

func (s *ScriptSpider) Close() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.seederStatus = -1
}
