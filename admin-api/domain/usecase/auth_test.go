package usecase_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/require"

	"gitlab.com/konstellation/kre/admin-api/domain/entity"
	"gitlab.com/konstellation/kre/admin-api/domain/usecase"
	"gitlab.com/konstellation/kre/admin-api/mocks"
)

type authSuite struct {
	ctrl           *gomock.Controller
	authInteractor *usecase.AuthInteractor
	mocks          authSuiteMocks
}

type authSuiteMocks struct {
	logger                    *mocks.MockLogger
	loginLinkTransport        *mocks.MockLoginLinkTransport
	verificationCodeGenerator *mocks.MockVerificationCodeGenerator
	verificationCodeRepo      *mocks.MockVerificationCodeRepo
	userRepo                  *mocks.MockUserRepo
	settingRepo               *mocks.MockSettingRepo
	userActivityRepo          *mocks.MockUserActivityRepo
}

func newAuthSuite(t *testing.T) *authSuite {
	ctrl := gomock.NewController(t)

	logger := mocks.NewMockLogger(ctrl)
	loginLinkTransport := mocks.NewMockLoginLinkTransport(ctrl)
	verificationCodeGenerator := mocks.NewMockVerificationCodeGenerator(ctrl)
	verificationCodeRepo := mocks.NewMockVerificationCodeRepo(ctrl)
	userRepo := mocks.NewMockUserRepo(ctrl)
	settingRepo := mocks.NewMockSettingRepo(ctrl)
	userActivityRepo := mocks.NewMockUserActivityRepo(ctrl)

	mocks.AddLoggerExpects(logger)

	userActivityInteractor := usecase.NewUserActivityInteractor(
		logger,
		userActivityRepo,
		userRepo,
	)

	authInteractor := usecase.NewAuthInteractor(
		logger,
		loginLinkTransport,
		verificationCodeGenerator,
		verificationCodeRepo,
		userRepo,
		settingRepo,
		userActivityInteractor,
	)

	return &authSuite{
		ctrl: ctrl,
		mocks: authSuiteMocks{
			logger:                    logger,
			loginLinkTransport:        loginLinkTransport,
			verificationCodeGenerator: verificationCodeGenerator,
			verificationCodeRepo:      verificationCodeRepo,
			userRepo:                  userRepo,
			settingRepo:               settingRepo,
			userActivityRepo:          userActivityRepo,
		},
		authInteractor: authInteractor,
	}
}

func TestSignInWithoutDomainAndEmailValidation(t *testing.T) {
	s := newAuthSuite(t)
	defer s.ctrl.Finish()

	verificationCode := "test_verification_code"
	verificationCodeDurationInMinutes := 1
	user := &entity.User{
		Email: "userA@testdomain.com",
		ID:    "userA",
	}
	settings := &entity.Setting{
		SessionLifetimeInDays: 0,
		AuthAllowedDomains:    []string{},
		AuthAllowedEmails:     []string{},
	}

	s.mocks.userRepo.EXPECT().GetByEmail(user.Email).Return(user, nil)
	s.mocks.settingRepo.EXPECT().Get().Return(settings, nil)
	s.mocks.verificationCodeGenerator.EXPECT().Generate().Return(verificationCode)
	s.mocks.verificationCodeRepo.EXPECT().Store(verificationCode, user.ID, gomock.Any()).Return(nil)
	s.mocks.loginLinkTransport.EXPECT().Send(user.Email, verificationCode).Return(nil)

	err := s.authInteractor.SignIn(user.Email, verificationCodeDurationInMinutes)
	require.Nil(t, err)
}

func TestSignInWithValidDomain(t *testing.T) {
	s := newAuthSuite(t)
	defer s.ctrl.Finish()

	verificationCode := "test_verification_code"
	verificationCodeDurationInMinutes := 1
	domain := "testdomain.com"
	user := &entity.User{
		Email: "userA@" + domain,
		ID:    "userA",
	}
	settings := &entity.Setting{
		SessionLifetimeInDays: 0,
		AuthAllowedDomains:    []string{domain},
	}

	s.mocks.userRepo.EXPECT().GetByEmail(user.Email).Return(user, nil)
	s.mocks.settingRepo.EXPECT().Get().Return(settings, nil)
	s.mocks.verificationCodeGenerator.EXPECT().Generate().Return(verificationCode)
	s.mocks.verificationCodeRepo.EXPECT().Store(verificationCode, user.ID, gomock.Any()).Return(nil)
	s.mocks.loginLinkTransport.EXPECT().Send(user.Email, verificationCode).Return(nil)

	err := s.authInteractor.SignIn(user.Email, verificationCodeDurationInMinutes)
	require.Nil(t, err)
}

func TestSignInWithInvalidDomain(t *testing.T) {
	s := newAuthSuite(t)
	defer s.ctrl.Finish()

	verificationCodeDurationInMinutes := 1
	user := &entity.User{
		Email: "userA@testdomain.com",
		ID:    "userA",
	}
	settings := &entity.Setting{
		SessionLifetimeInDays: 0,
		AuthAllowedDomains:    []string{"anotherdomain.com"},
	}

	s.mocks.userRepo.EXPECT().GetByEmail(user.Email).Return(user, nil)
	s.mocks.settingRepo.EXPECT().Get().Return(settings, nil)

	err := s.authInteractor.SignIn(user.Email, verificationCodeDurationInMinutes)
	require.Equal(t, usecase.ErrUserNotAllowed, err)
}

func TestSignUpWithValidDomain(t *testing.T) {
	s := newAuthSuite(t)
	defer s.ctrl.Finish()

	verificationCode := "test_verification_code"
	verificationCodeDurationInMinutes := 1
	domain := "testdomain.com"
	user := &entity.User{
		Email: "userA@" + domain,
		ID:    "userA",
	}
	settings := &entity.Setting{
		SessionLifetimeInDays: 0,
		AuthAllowedDomains:    []string{domain},
	}

	s.mocks.userRepo.EXPECT().GetByEmail(user.Email).Return(nil, usecase.ErrUserNotFound)
	s.mocks.settingRepo.EXPECT().Get().Return(settings, nil)
	s.mocks.verificationCodeGenerator.EXPECT().Generate().Return(verificationCode)
	s.mocks.verificationCodeRepo.EXPECT().Store(verificationCode, user.ID, gomock.Any()).Return(nil)
	s.mocks.loginLinkTransport.EXPECT().Send(user.Email, verificationCode).Return(nil)
	s.mocks.userRepo.EXPECT().Create(gomock.Any(), user.Email, gomock.Any()).Return(user, nil)

	err := s.authInteractor.SignIn(user.Email, verificationCodeDurationInMinutes)
	require.Nil(t, err)
}

func TestSignInWithoutDomainValidationAndInvalidEmail(t *testing.T) {
	s := newAuthSuite(t)
	defer s.ctrl.Finish()

	verificationCodeDurationInMinutes := 1
	user := &entity.User{
		Email: "userA@testdomain.com",
		ID:    "userA",
	}
	settings := &entity.Setting{
		SessionLifetimeInDays: 0,
		AuthAllowedDomains:    []string{},
		AuthAllowedEmails:     []string{"userB@testdomain.com"},
	}

	s.mocks.userRepo.EXPECT().GetByEmail(user.Email).Return(user, nil)
	s.mocks.settingRepo.EXPECT().Get().Return(settings, nil)

	err := s.authInteractor.SignIn(user.Email, verificationCodeDurationInMinutes)
	require.Equal(t, usecase.ErrUserNotAllowed, err)
}

func TestSignUpWithValidEmailAddress(t *testing.T) {
	s := newAuthSuite(t)
	defer s.ctrl.Finish()

	verificationCode := "test_verification_code"
	verificationCodeDurationInMinutes := 1
	userId := "userA"
	domain := "testdomain.com"
	email := fmt.Sprintf("%s@%s", userId, domain)

	user := &entity.User{
		Email: email,
		ID:    userId,
	}
	settings := &entity.Setting{
		SessionLifetimeInDays: 0,
		AuthAllowedDomains:    []string{"anotherdomain.com"},
		AuthAllowedEmails:     []string{email},
	}

	s.mocks.userRepo.EXPECT().GetByEmail(user.Email).Return(nil, usecase.ErrUserNotFound)
	s.mocks.settingRepo.EXPECT().Get().Return(settings, nil)
	s.mocks.verificationCodeGenerator.EXPECT().Generate().Return(verificationCode)
	s.mocks.verificationCodeRepo.EXPECT().Store(verificationCode, user.ID, gomock.Any()).Return(nil)
	s.mocks.loginLinkTransport.EXPECT().Send(user.Email, verificationCode).Return(nil)
	s.mocks.userRepo.EXPECT().Create(gomock.Any(), user.Email, gomock.Any()).Return(user, nil)

	err := s.authInteractor.SignIn(user.Email, verificationCodeDurationInMinutes)
	require.Nil(t, err)
}

func TestSignInErrGettingUser(t *testing.T) {
	s := newAuthSuite(t)
	defer s.ctrl.Finish()

	email := "userA@testdomain.com"
	unexpectedErr := errors.New("unexpected error")

	s.mocks.userRepo.EXPECT().GetByEmail(email).Return(nil, unexpectedErr)

	err := s.authInteractor.SignIn(email, 1)
	require.Equal(t, unexpectedErr, err)
}

func TestSignInErrGettingSettings(t *testing.T) {
	s := newAuthSuite(t)
	defer s.ctrl.Finish()

	unexpectedErr := errors.New("unexpected error")
	verificationCodeDurationInMinutes := 1
	user := &entity.User{
		Email: "userA@testdomain.com",
		ID:    "userA",
	}

	s.mocks.userRepo.EXPECT().GetByEmail(user.Email).Return(user, nil)
	s.mocks.settingRepo.EXPECT().Get().Return(&entity.Setting{}, unexpectedErr)

	err := s.authInteractor.SignIn(user.Email, verificationCodeDurationInMinutes)
	require.Equal(t, unexpectedErr, err)
}

func TestSignInErrInvalidEmail(t *testing.T) {
	s := newAuthSuite(t)
	defer s.ctrl.Finish()

	email := "invalidemailaddress"

	err := s.authInteractor.SignIn(email, 1)
	require.Equal(t, usecase.ErrUserEmailInvalid, err)
}

func TestSignInErrStoringValidationCode(t *testing.T) {
	s := newAuthSuite(t)
	defer s.ctrl.Finish()

	unexpectedErr := errors.New("unexpected error")
	verificationCode := "test_verification_code"
	verificationCodeDurationInMinutes := 1
	user := &entity.User{
		Email: "userA@testdomain.com",
		ID:    "userA",
	}
	settings := &entity.Setting{
		SessionLifetimeInDays: 0,
		AuthAllowedDomains:    []string{},
	}

	s.mocks.userRepo.EXPECT().GetByEmail(user.Email).Return(user, nil)
	s.mocks.settingRepo.EXPECT().Get().Return(settings, nil)
	s.mocks.verificationCodeGenerator.EXPECT().Generate().Return(verificationCode)
	s.mocks.verificationCodeRepo.EXPECT().Store(verificationCode, user.ID, gomock.Any()).Return(unexpectedErr)

	err := s.authInteractor.SignIn(user.Email, verificationCodeDurationInMinutes)
	require.Equal(t, unexpectedErr, err)
}

func TestSignUpErrCreatingUser(t *testing.T) {
	s := newAuthSuite(t)
	defer s.ctrl.Finish()

	unexpectedErr := errors.New("unexpected error")
	verificationCodeDurationInMinutes := 1
	domain := "testdomain.com"
	user := &entity.User{
		Email: "userA@" + domain,
		ID:    "userA",
	}
	settings := &entity.Setting{
		SessionLifetimeInDays: 0,
		AuthAllowedDomains:    []string{domain},
	}

	s.mocks.userRepo.EXPECT().GetByEmail(user.Email).Return(nil, usecase.ErrUserNotFound)
	s.mocks.settingRepo.EXPECT().Get().Return(settings, nil)
	s.mocks.userRepo.EXPECT().Create(gomock.Any(), user.Email, gomock.Any()).Return(nil, unexpectedErr)

	err := s.authInteractor.SignIn(user.Email, verificationCodeDurationInMinutes)
	require.Equal(t, unexpectedErr, err)
}

func TestVerifyCode(t *testing.T) {
	s := newAuthSuite(t)
	defer s.ctrl.Finish()

	code := "test_verification_code"
	verificationCode := &entity.VerificationCode{
		Code:      code,
		UID:       "userA",
		ExpiresAt: time.Now().Add(time.Duration(1) * time.Minute),
	}

	s.mocks.verificationCodeRepo.EXPECT().Get(code).Return(verificationCode, nil)
	s.mocks.verificationCodeRepo.EXPECT().Delete(code).Return(nil)
	s.mocks.userActivityRepo.EXPECT().Create(gomock.Any()).Return(nil)

	userId, err := s.authInteractor.VerifyCode(code)
	require.Nil(t, err)
	require.Equal(t, verificationCode.UID, userId)
}

func TestVerifyNotFoundCode(t *testing.T) {
	s := newAuthSuite(t)
	defer s.ctrl.Finish()

	code := "test_verification_code"
	s.mocks.verificationCodeRepo.EXPECT().Get(code).Return(nil, usecase.ErrVerificationCodeNotFound)

	_, err := s.authInteractor.VerifyCode(code)
	require.Equal(t, usecase.ErrVerificationCodeNotFound, err)
}

func TestVerifyExpiredCode(t *testing.T) {
	s := newAuthSuite(t)
	defer s.ctrl.Finish()

	code := "test_verification_code"
	verificationCode := &entity.VerificationCode{
		Code:      code,
		UID:       "userA",
		ExpiresAt: time.Now().Add(time.Duration(-1) * time.Minute),
	}

	s.mocks.verificationCodeRepo.EXPECT().Get(code).Return(verificationCode, nil)
	s.mocks.verificationCodeRepo.EXPECT().Delete(code).Return(nil)

	_, err := s.authInteractor.VerifyCode(code)
	require.Equal(t, usecase.ErrExpiredVerificationCode, err)
}

func TestVerifyCodeErrCreatingUserActivity(t *testing.T) {
	s := newAuthSuite(t)
	defer s.ctrl.Finish()

	unexpectedErr := errors.New("unexpected error")
	code := "test_verification_code"
	verificationCode := &entity.VerificationCode{
		Code:      code,
		UID:       "userA",
		ExpiresAt: time.Now().Add(time.Duration(1) * time.Minute),
	}

	s.mocks.verificationCodeRepo.EXPECT().Get(code).Return(verificationCode, nil)
	s.mocks.verificationCodeRepo.EXPECT().Delete(code).Return(nil)
	s.mocks.userActivityRepo.EXPECT().Create(gomock.Any()).Return(unexpectedErr)

	_, err := s.authInteractor.VerifyCode(code)
	require.Equal(t, fmt.Errorf("error creating userActivity: %w", unexpectedErr), err)
}

func TestVerifyCodeErrDeletingValidationCode(t *testing.T) {
	s := newAuthSuite(t)
	defer s.ctrl.Finish()

	unexpectedErr := errors.New("unexpected error")
	code := "test_verification_code"
	verificationCode := &entity.VerificationCode{
		Code:      code,
		UID:       "userA",
		ExpiresAt: time.Now().Add(time.Duration(1) * time.Minute),
	}

	s.mocks.verificationCodeRepo.EXPECT().Get(code).Return(verificationCode, nil)
	s.mocks.verificationCodeRepo.EXPECT().Delete(code).Return(unexpectedErr)
	s.mocks.userActivityRepo.EXPECT().Create(gomock.Any()).Return(nil)

	userId, err := s.authInteractor.VerifyCode(code)
	require.Nil(t, err)
	require.Equal(t, verificationCode.UID, userId)
}
