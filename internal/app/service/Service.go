package service

import (
	"kbrprime-be/internal/app/commons"
	"kbrprime-be/internal/app/middleware/authMiddleware"
	"kbrprime-be/internal/app/repository"
	"kbrprime-be/internal/app/service/authService"
	"kbrprime-be/internal/app/service/categoriesService"
	"kbrprime-be/internal/app/service/healtyService"
	"kbrprime-be/internal/app/service/likeService"
	"kbrprime-be/internal/app/service/listenService"
	"kbrprime-be/internal/app/service/playListService"
	"kbrprime-be/internal/app/service/showService"
	"kbrprime-be/internal/app/service/userService"
)

// Option anything any service object needed
type Option struct {
	commons.Options
	*repository.Repositories
}

type Services struct {
	HealtyService     healtyService.IHealtyService
	AuthService       authService.IAuthService
	AuthMiddleware    authMiddleware.IAuthMiddleware
	ShowService       showService.IShowService
	CategoriesService categoriesService.ICategoriesService
	UserService       userService.IUserService
	ListenService     listenService.IListenService
	LikeService       likeService.ILikeService
	PlayListService   playListService.IPlayListService
}
