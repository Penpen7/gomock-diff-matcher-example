package cmd_test

import (
	"testing"
	"time"

	"github.com/Penpen7/gomock-diff-matcher-example/cmd"
	"github.com/Penpen7/gomock-diff-matcher-example/matcher"
	"github.com/Penpen7/gomock-diff-matcher-example/mock_cmd"
	"go.uber.org/mock/gomock"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_cmd.NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().FindByID(1).Return(
		&cmd.User{
			ID:        "1",
			Name:      "Alice",
			CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	)
	mockUserRepository.EXPECT().Create(
		matcher.DiffEq(
			cmd.User{
				ID:        "1",
				Name:      "Penpen7",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		),
	).Return(nil)

	if err := cmd.Run(mockUserRepository, 1); err != nil {
		t.Errorf("Run() error = %v, want nil", err)
	}
}
