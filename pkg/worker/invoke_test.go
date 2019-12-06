package worker

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/golang/mock/gomock"
	apiv1beta1 "github.com/socketfunc/colony/proto/api/v1beta1"
	"github.com/socketfunc/colony/proto/config"
	workerv1beta1 "github.com/socketfunc/colony/proto/worker/v1beta1"
	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServer_Invoke(t *testing.T) {
	tests := []struct {
		title  string
		setup  setupFunc
		req    *workerv1beta1.InvokeRequest
		expect *expect
	}{
		{
			title: "success",
			setup: func(mocks *mocks) {
				in := &apiv1beta1.GetConfigRequest{
					Namespace: "default",
				}
				out := &apiv1beta1.GetConfigResponse{
					Config: &config.Config{
						Version: "v1beta1",
						Metadata: &config.Metadata{
							Namespace: "default",
						},
						Spec: &config.Spec{
							Functions: []*config.Function{
								{
									Name:       "add",
									Repository: "testdata/sample.wasm",
								},
							},
						},
					},
				}
				mocks.apiClient.EXPECT().GetConfig(gomock.Any(), in).Return(out, nil)
				mocks.storage.EXPECT().Download(gomock.Any(), gomock.Any()).
					DoAndReturn(func(_ context.Context, path string) ([]byte, error) {
						return ioutil.ReadFile(path)
					})
			},
			req: &workerv1beta1.InvokeRequest{
				Namespace: "default",
				Name:      "add",
				Args:      encode([]interface{}{int32(10), int32(10)}),
			},
			expect: &expect{
				code: codes.OK,
				resp: &workerv1beta1.InvokeResponse{
					Result: encode(int32(20)),
				},
			},
		},
	}
	for _, testcase := range tests {
		testcase := testcase
		t.Run(testcase.title, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			svc := newWorkerServer(ctrl, testcase.setup)
			resp, err := svc.Invoke(context.Background(), testcase.req)
			if testcase.expect.isError() {
				assert.Error(t, err)
			}
			assert.Equal(t, testcase.expect.code, status.Code(err))
			assert.Equal(t, testcase.expect.resp, resp)
		})
	}
}

func encode(data interface{}) []byte {
	res, err := msgpack.Marshal(data)
	if err != nil {
		panic(err)
	}
	return res
}
