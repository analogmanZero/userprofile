package gateway

import (
  "context"
  "userprofile/domain/entity"
  "userprofile/infrastructure/log"
  "userprofile/infrastructure/token"

  "gorm.io/gorm"
)

type indatabase2Gateway struct {
  indatabaseGateway
}

// NewIndatabase2Gateway ...
func NewIndatabase2Gateway(UserToken *token.JWTToken, db *gorm.DB) *indatabase2Gateway {

  x := NewIndatabaseGateway(UserToken, db)

  return &indatabase2Gateway{
    indatabaseGateway: *x,
  }
}

func (r *indatabase2Gateway) FindAllUser(ctx context.Context) ([]*entity.User, error) {
  log.Info(ctx, "hellooo world")

  var objs []*entity.User
  err := r.DB.
    Find(&objs).Error

  if err != nil {
    log.Error(ctx, err.Error())
    return nil, err
  }

  return objs, nil
}