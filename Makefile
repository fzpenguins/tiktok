DIR = $(shell pwd)
IDL_PATH = $(DIR)/idl

new-user:
	kitex -module tiktok ./idl/user.thrift

new-video:
	kitex -module tiktok ./idl/video.thrift

new-interaction:
	kitex -module tiktok ./idl/interaction.thrift

new-follow:
	kitex -module tiktok ./idl/follow.thrift

#update-api:
#	#hz new --module tiktok/cmd/api --out_dir ./cmd/api --idl ./idl/api.thrift
#	hz new --module tiktok/cmd/api --service "api"  --idl ./idl/api.thrift

user-serivce:
	kitex -service user ./idl/user.thrift

