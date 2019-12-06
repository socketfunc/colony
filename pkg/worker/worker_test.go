package worker

import (
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	storageMock "github.com/socketfunc/colony/pkg/storage/mock"
	apiv1betaMock "github.com/socketfunc/colony/proto/api/v1beta1/mock"
	workerv1beta1 "github.com/socketfunc/colony/proto/worker/v1beta1"
	"google.golang.org/grpc/codes"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

type mocks struct {
	apiClient *apiv1betaMock.MockApiServiceClient
	storage   *storageMock.MockStorage
}

type setupFunc func(*mocks)

type expect struct {
	resp interface{}
	code codes.Code
}

func (e *expect) isError() bool {
	return e.code != codes.OK
}

func newWorkerServer(ctrl *gomock.Controller, setup setupFunc) workerv1beta1.WorkerServer {
	mocks := &mocks{
		apiClient: apiv1betaMock.NewMockApiServiceClient(ctrl),
		storage:   storageMock.NewMockStorage(ctrl),
	}
	setup(mocks)
	params := &ServerParams{
		ApiClient: mocks.apiClient,
		Storage:   mocks.storage,
	}
	return NewServer(params)
}
