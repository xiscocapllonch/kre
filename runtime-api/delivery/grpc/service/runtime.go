package service

import (
	"context"
	"fmt"
	"gitlab.com/konstellation/konstellation-ce/kre/runtime-api/adapter/config"
	"gitlab.com/konstellation/konstellation-ce/kre/runtime-api/domain/usecase"
	"gitlab.com/konstellation/konstellation-ce/kre/runtime-api/domain/usecase/logging"
	"gitlab.com/konstellation/konstellation-ce/kre/runtime-api/runtimepb"
)

// RuntimeService basic server
type RuntimeService struct {
	cfg        *config.Config
	logger     logging.Logger
	interactor *usecase.VersionInteractor
}

func NewRuntimeService(
	cfg *config.Config,
	logger logging.Logger,
	interactor *usecase.VersionInteractor,

) *RuntimeService {
	return &RuntimeService{
		cfg:        cfg,
		logger:     logger,
		interactor: interactor,
	}
}

func (s *RuntimeService) DeployVersion(ctx context.Context, req *runtimepb.DeployVersionRequest) (*runtimepb.DeployVersionResponse, error) {
	s.logger.Info("DeployVersionRequest received")
	runtimeName := req.GetVersion().GetName()

	message := fmt.Sprintf("Runtime %s created", runtimeName)
	success := true

	_, err := s.interactor.DeployVersion(runtimeName)
	if err != nil {
		success = false
		message = err.Error()
		s.logger.Error(message)
	}

	// Send response
	res := &runtimepb.DeployVersionResponse{
		Success: success,
		Message: message,
	}

	return res, nil
}

func (s *RuntimeService) ActivateVersion(ctx context.Context, req *runtimepb.ActivateVersionRequest) (*runtimepb.ActivateVersionResponse, error) {
	versionName := req.GetVersion().GetName()
	namespace := s.cfg.Kubernetes.Namespace
	s.logger.Info(fmt.Sprintf("Activating version \"%s\" in namespace \"%s\"...\n", versionName, namespace))

	_, err := s.interactor.ActivateVersion(versionName)
	if err != nil {
		s.logger.Error(err.Error())
		return &runtimepb.ActivateVersionResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &runtimepb.ActivateVersionResponse{
		Success: true,
		Message: "Version activated correctly.",
	}, nil
}