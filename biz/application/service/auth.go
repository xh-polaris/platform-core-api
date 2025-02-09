package service

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
	"github.com/xh-polaris/platform-core-api/biz/adaptor"
	"github.com/xh-polaris/platform-core-api/biz/application/dto/platform/core_api"
	"github.com/xh-polaris/platform-core-api/biz/infra/config"
	"github.com/xh-polaris/platform-core-api/biz/infra/consts"
	"github.com/xh-polaris/platform-core-api/biz/infra/rpc/platform_sts"
	"github.com/xh-polaris/platform-core-api/biz/infra/util"
	"github.com/xh-polaris/platform-core-api/biz/infra/util/log"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/sts"
	"time"
)

type IAuthService interface {
	SignIn(ctx context.Context, req *core_api.SignInReq) (*core_api.SignInResp, error)
	SetPassword(ctx context.Context, req *core_api.SetPasswordReq) (*core_api.SetPasswordResp, error)
	SendVerifyCode(ctx context.Context, req *core_api.SendVerifyCodeReq) (*core_api.SendVerifyCodeResp, error)
}

type AuthService struct {
	Config *config.Config
	Sts    platform_sts.IPlatformSts
}

var AuthServiceSet = wire.NewSet(
	wire.Struct(new(AuthService), "*"),
	wire.Bind(new(IAuthService), new(*AuthService)),
)

func (s *AuthService) SignIn(ctx context.Context, req *core_api.SignInReq) (*core_api.SignInResp, error) {
	resp := new(core_api.SignInResp)
	rpcResp, err := s.Sts.SignIn(ctx, &sts.SignInReq{
		AuthType:   req.GetAuthType(),
		AuthId:     req.GetAuthId(),
		Password:   req.Password,
		VerifyCode: req.VerifyCode,
	})
	if err != nil {
		return nil, err
	}

	auth := s.Config.Auth
	resp.AccessToken, resp.AccessExpire, err = generateJwtToken(req, rpcResp, auth.SecretKey, auth.AccessExpire)
	if err != nil {
		log.CtxError(ctx, "[generateJwtToken] fail, err=%v, config=%s, resp=%s", err, util.JSONF(s.Config.Auth), util.JSONF(rpcResp))
		return nil, err
	}
	resp.UserId = rpcResp.GetUserId()
	if rpcResp.Options != nil {
		resp.Option = rpcResp.Options
	}
	return resp, nil
}

func generateJwtToken(req *core_api.SignInReq, resp *sts.SignInResp, secret string, expire int64) (string, int64, error) {
	key, err := jwt.ParseECPrivateKeyFromPEM([]byte(secret))
	if err != nil {
		return "", 0, err
	}
	iat := time.Now().Unix()
	exp := iat + expire
	claims := make(jwt.MapClaims)
	claims["exp"] = exp
	claims["iat"] = iat
	claims["userId"] = resp.GetUserId()
	claims["appId"] = req.GetAppId()
	claims["deviceId"] = req.GetDeviceId()
	claims["wechatUserMeta"] = &basic.WechatUserMeta{
		AppId:   resp.GetAppId(),
		OpenId:  resp.GetOpenId(),
		UnionId: resp.GetUnionId(),
	}
	token := jwt.New(jwt.SigningMethodES256)
	token.Claims = claims
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", 0, err
	}
	return tokenString, exp, nil
}

func (s *AuthService) SetPassword(ctx context.Context, req *core_api.SetPasswordReq) (*core_api.SetPasswordResp, error) {
	user := adaptor.ExtractUserMeta(ctx)
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.SetPasswordResp)
	_, err := s.Sts.SetPassword(ctx, &sts.SetPasswordReq{
		UserId:   user.UserId,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *AuthService) SendVerifyCode(ctx context.Context, req *core_api.SendVerifyCodeReq) (*core_api.SendVerifyCodeResp, error) {
	_, err := s.Sts.SendVerifyCode(ctx, &sts.SendVerifyCodeReq{
		AuthType: req.AuthType,
		AuthId:   req.AuthId,
	})
	if err != nil {
		return &core_api.SendVerifyCodeResp{
			Code: 1005,
			Msg:  "验证码发送失败，请重试",
		}, nil
	}

	return &core_api.SendVerifyCodeResp{
		Code: 0,
		Msg:  "验证码发送成功",
	}, nil
}
