package script

import (
	"github.com/robertkrimen/otto"
	"github.com/sirupsen/logrus"
)

type ScriptExtension struct {
	ScriptBase
}

func NewScriptExtensionFromBase(sb *ScriptBase,
	err error) (*ScriptExtension, error) {
	if err != nil {
		return nil, err
	} else {
		return &ScriptExtension{
			ScriptBase: *sb,
		}, nil
	}
}

func NewScriptExtension(vm *otto.Otto) (*ScriptExtension, error) {
	return NewScriptExtensionFromBase(NewScriptBase(vm))
}

func NewScriptExtensionWithScript(vm *otto.Otto,
	script string) (*ScriptExtension, error) {
	return NewScriptExtensionFromBase(NewScriptBaseWithScript(vm, script))
}

func (s *ScriptExtension) handleError(err error) {
	logrus.Errorln(err)
}

func (s *ScriptExtension) OnSeedRequest(request Request) {
	if _, err := s.CallIfExists("on_seed_request", request); err != nil {
		s.handleError(err)
	}
}

func (s *ScriptExtension) OnScheduled(request Request) {
	if _, err := s.CallIfExists("on_scheduled", request); err != nil {
		s.handleError(err)
	}
}

func (s *ScriptExtension) OnDownloadSuccess(response Response) {
	if _, err := s.CallIfExists("on_download_success", response); err != nil {
		s.handleError(err)
	}
}

func (s *ScriptExtension) OnDownloadFailure(request Request, err error) {
	if _, err := s.CallIfExists("on_download_failure", request, err); err != nil {
		s.handleError(err)
	}
}

func (s *ScriptExtension) OnParsedRequest(request Request) {
	if _, err := s.CallIfExists("on_parsed_request", request); err != nil {
		s.handleError(err)
	}
}

func (s *ScriptExtension) OnParsedItem(item Item) {
	if _, err := s.CallIfExists("on_parsed_item", item); err != nil {
		s.handleError(err)
	}
}

func (s *ScriptExtension) OnParsedUnknown(item interface{}) {
	if _, err := s.CallIfExists("on_parsed_unknown", item); err != nil {
		s.handleError(err)
	}
}

func (s *ScriptExtension) OnParsedResult(result []interface{}) {
	if _, err := s.CallIfExists("on_parsed_result", result); err != nil {
		s.handleError(err)
	}
}

func (s *ScriptExtension) OnBeforePipeline(item Item) {
	if _, err := s.CallIfExists("on_before_pipeline", item); err != nil {
		s.handleError(err)
	}
}

func (s *ScriptExtension) OnAfterPipeline(item Item) {
	if _, err := s.CallIfExists("on_after_pipeline", item); err != nil {
		s.handleError(err)
	}
}
