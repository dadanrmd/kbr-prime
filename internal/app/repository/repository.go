package repository

import (
	"kbrprime-be/internal/app/commons"
	"kbrprime-be/internal/app/repository/categoriesRepository"
	"kbrprime-be/internal/app/repository/healtyRepository"
	"kbrprime-be/internal/app/repository/likeRepository"
	"kbrprime-be/internal/app/repository/listenRepository"
	"kbrprime-be/internal/app/repository/showRepository"
	"kbrprime-be/internal/app/repository/userRepository"
)

// Option anything any repo object needed
type Option struct {
	commons.Options
}

type Repositories struct {
	HealtyRepository     healtyRepository.IHealtyRepository
	ShowRepository       showRepository.IShowRepository
	CategoriesRepository categoriesRepository.ICategoriesRepository
	UserRepository       userRepository.IUserRepository
	ListenRepository     listenRepository.IListenRepository
	LikeRepository       likeRepository.ILikeRepository
}
