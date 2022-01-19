package usecase

import (
	forumRepo "bd_tp_V2/forum/repository"
	"bd_tp_V2/models"
	threadRepo "bd_tp_V2/thread/repository"
	userRepo "bd_tp_V2/user/repository"
	"github.com/satori/go.uuid"
)

type Usecase struct {
	forumRepo  *forumRepo.Repository
	userRepo   *userRepo.Repository
	threadRepo *threadRepo.Repository 
}

func NewForumUsecase(fR *forumRepo.Repository, uR *userRepo.Repository, tR *threadRepo.Repository) *Usecase {
	return &Usecase{
		forumRepo:  fR,
		userRepo:   uR,
		threadRepo: tR,
	}
}

func (fU *Usecase) CreateForum(forum *models.Forum) (*models.Forum, error) {
	user, err := fU.userRepo.GetUserByNickname(forum.User)
	if err != nil {
		return nil, err
	}
	//
	forum.User = user.Nickname
	if forum.Slug != "" {
		oldForum, err := fU.forumRepo.ForumDetails(forum.Slug)
		if err == nil {
			return oldForum, models.ErrorForumExists
		} else if err == models.ErrorForumNotFound {
			newForum, err := fU.forumRepo.CreateForum(forum)
			if err != nil {
				return nil, err
			}
			return newForum, nil
		} else {
			return nil, err
		}
	} else {
		forum.Slug = uuid.NewV4().String()
	}
	newForum, err := fU.forumRepo.CreateForum(forum)
	if err != nil {
		return nil, err
	}
	return newForum, err
}

func (fU *Usecase) ForumDetails (slug string) (*models.Forum, error) {
	forum,err := fU.forumRepo.ForumDetails(slug)
	if err != nil {
		return nil,err
	}
	return forum, err
}

func (fU *Usecase) ForumThreadCreate (slug string, thread *models.Thread) (*models.Thread, error) {
	forum, err := fU.forumRepo.ForumDetails(slug)
	if err != nil {
		return nil, err
	}
	user, err := fU.userRepo.GetUserByNickname(thread.Author)
	if err != nil {
		return nil, err
	}
	thread.Forum = forum.Slug
	//
	thread.Author = user.Nickname
	if thread.Slug != "" {
		oldThread, err := fU.threadRepo.ThreadDetails(thread.Slug)
		if err == nil {
			return oldThread, models.ErrorThreadExists
		} else if err == models.ErrorThreadNotFound {
			newThread, err := fU.forumRepo.ForumThreadCreate(thread)
			if err != nil {
				return nil, err
			}
			return newThread, err
		}
	}
	newThread, err := fU.forumRepo.ForumThreadCreate(thread)
	if err != nil {
		return nil, err
	}
	return newThread, err
}

func (fU *Usecase) GetThreadsByForum(slug, limit, since, desc string) (*models.Threads, error) {
	//Прверка на наличие форума
	_, err := fU.forumRepo.ForumDetails(slug)
	if err != nil {

		return nil, err
	}
	if limit == "" {
		limit = "100"
	}
	//Здесь форум есть, но мб нет постов - не ошибка
	threads, err := fU.forumRepo.GetThreadsByForum(slug, limit, since, desc)
	if err != nil {

		return nil, err
	}
	return threads, nil
}

func (fU *Usecase) GetForumUsers(slug, limit, since, desc string) (*models.Users, error) {
	_, err := fU.ForumDetails(slug)
	if err != nil {
		return nil, err
	}
	if limit == "" {
		limit = "100"
	}
	users, err := fU.forumRepo.GetForumUsers(slug, limit, since, desc)
	if err != nil {
		return nil, err
	}
	return users, nil
}