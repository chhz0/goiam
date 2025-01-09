package v1

import (
	"context"
	"regexp"
	"sync"

	"github.com/chhz0/goiam/internal/apisvr/dal"
	errcode "github.com/chhz0/goiam/internal/pkg/errorscore/code"
	"github.com/chhz0/goiam/internal/pkg/logger"
	"github.com/chhz0/goiam/internal/pkg/model"
	"github.com/chhz0/goiam/pkg/errors"
	"github.com/chhz0/goiam/pkg/log"
	"github.com/chhz0/goiam/pkg/meta"
)

//go:generate mockgen -self_package=github.com/chhz0/goiam/internal/apisvr/service/v1 -destination=../mock/mock_v1.go -package=mock github.com/chhz0/goiam/internal/apisvr/service/v1 UserSrv,SecretSrv,PolicySrv
type UserSrv interface {
	Create(ctx context.Context, user *model.User, opts meta.CreateOptions) error
	Update(ctx context.Context, user *model.User, opts meta.UpdateOptions) error
	Delete(ctx context.Context, username string, opts meta.DeleteOptions) error
	DeleteCollection(ctx context.Context, username []string, opts meta.DeleteOptions) error
	Get(ctx context.Context, username string, opts meta.GetOptions) (*model.User, error)
	List(ctx context.Context, opts meta.ListOptions) (*model.UserList, error)
	ListWithBadPerformance(ctx context.Context, opts meta.ListOptions) (*model.UserList, error)
	ChangePassword(ctx context.Context, user *model.User) error
}

type userService struct {
	dal dal.Factory
}

// ChangePassword implements UserSrv.
func (u *userService) ChangePassword(ctx context.Context, user *model.User) error {
	if err := u.dal.Users().Update(ctx, user, meta.UpdateOptions{}); err != nil {
		return errors.WithCode(errcode.ErrDatabase, err)
	}

	return nil
}

// Create implements UserSrv.
func (u *userService) Create(ctx context.Context, user *model.User, opts meta.CreateOptions) error {
	if err := u.dal.Users().Create(ctx, user, opts); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'idx_name'", err.Error()); match {
			return errors.WithCode(errcode.ErrUserAlreadyExist, err)
		}

		return errors.WithCode(errcode.ErrDatabase, err)
	}

	return nil
}

// Delete implements UserSrv.
func (u *userService) Delete(ctx context.Context, username string, opts meta.DeleteOptions) error {
	if err := u.dal.Users().Delete(ctx, username, opts); err != nil {
		return err
	}

	return nil
}

// DeleteCollection implements UserSrv.
func (u *userService) DeleteCollection(ctx context.Context, username []string, opts meta.DeleteOptions) error {
	if err := u.dal.Users().DeleteCollection(ctx, username, opts); err != nil {
		return errors.WithCode(errcode.ErrDatabase, err)
	}

	return nil
}

// Get implements UserSrv.
func (u *userService) Get(ctx context.Context, username string, opts meta.GetOptions) (*model.User, error) {
	user, err := u.dal.Users().Get(ctx, username, opts)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// List implements UserSrv.
func (u *userService) List(ctx context.Context, opts meta.ListOptions) (*model.UserList, error) {
	users, err := u.dal.Users().List(ctx, opts)
	if err != nil {
		log.L(ctx, logger.UseKeys...).Errorf("list users server failed, %s", err.Error())

		return nil, errors.WithCode(errcode.ErrDatabase, err)
	}

	wg := sync.WaitGroup{}
	errChan := make(chan error, 1)
	finished := make(chan struct{})

	var m sync.Map
	for _, user := range users.Items {
		wg.Add(1)

		go func(user *model.User) {
			defer wg.Done()

			// some cost time process
			policies, err := u.dal.Policies().List(ctx, user.Name, meta.ListOptions{})
			if err != nil {
				errChan <- errors.WithCode(errcode.ErrDatabase, err)

				return
			}

			m.Store(user.ID, &model.User{
				ObjectMeta: &meta.ObjectMeta{
					ID:         user.ID,
					InstanceID: user.InstanceID,
					Name:       user.Name,
					ExtenAttrs: user.ExtenAttrs,
					CreatedAt:  user.CreatedAt,
					UpdatedAt:  user.UpdatedAt,
				},
				Nickname:    user.Nickname,
				Email:       user.Email,
				Phone:       user.Phone,
				IsAdmin:     user.IsAdmin,
				TotalPolicy: policies.TotalCount,
				LoginedAt:   user.LoginedAt,
			})
		}(user)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, err
	}

	infos := make([]*model.User, 0, len(users.Items))
	for _, user := range users.Items {
		if info, ok := m.Load(user.ID); ok {
			infos = append(infos, info.(*model.User))
		}
	}

	log.L(ctx, logger.UseKeys...).Debugf("list users server success, %d users", len(infos))

	return &model.UserList{
		ListMeta: users.ListMeta,
		Items:    infos,
	}, nil
}

// ListWithBadPerformance implements UserSrv.
// 这个函数是对 List 函数的一个低性能的实现，仅用于测试性能
func (u *userService) ListWithBadPerformance(ctx context.Context, opts meta.ListOptions) (*model.UserList, error) {
	users, err := u.dal.Users().List(ctx, opts)
	if err != nil {
		return nil, errors.WithCode(errcode.ErrDatabase, err)
	}

	infos := make([]*model.User, 0)
	for _, user := range users.Items {
		policy, err := u.dal.Policies().List(ctx, user.Name, meta.ListOptions{})
		if err != nil {
			return nil, errors.WithCode(errcode.ErrDatabase, err)
		}

		infos = append(infos, &model.User{
			ObjectMeta: &meta.ObjectMeta{
				ID:         user.ID,
				InstanceID: user.InstanceID,
				Name:       user.Name,
				ExtenAttrs: user.ExtenAttrs,
				CreatedAt:  user.CreatedAt,
				UpdatedAt:  user.UpdatedAt,
			},
			Nickname:    user.Nickname,
			Email:       user.Email,
			Phone:       user.Phone,
			IsAdmin:     user.IsAdmin,
			TotalPolicy: policy.TotalCount,
		})
	}

	return &model.UserList{
		Items:    infos,
		ListMeta: users.ListMeta,
	}, nil
}

// Update implements UserSrv.
func (u *userService) Update(ctx context.Context, user *model.User, opts meta.UpdateOptions) error {
	if err := u.dal.Users().Update(ctx, user, opts); err != nil {
		return errors.WithCode(errcode.ErrDatabase, err)
	}

	return nil
}

func NewUsers(f dal.Factory) *userService {
	return &userService{
		dal: f,
	}
}
