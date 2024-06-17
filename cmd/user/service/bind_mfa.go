package service

import (
	"github.com/pquerna/otp/totp"
	"log"
	"tiktok/kitex_gen/user"
	"tiktok/pkg/errno"
)

func (s *UserService) BindMFA(req *user.BindMFAReq) error {
	valid := totp.Validate(req.GetSecret(), req.GetCode())

	if !valid {
		log.Println("Invalid passcode")
		return errno.UserMFAInvalid
	}

	return nil
}