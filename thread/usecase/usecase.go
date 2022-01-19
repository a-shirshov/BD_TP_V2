package usecase

import (
	forumRepo "bd_tp_V2/forum/repository"
	"bd_tp_V2/models"
	//postRepo "bd_tp_V2/post/repository"
	threadRepo "bd_tp_V2/thread/repository"
	userRepo "bd_tp_V2/user/repository"
	//"errors"
	//"strconv"
)

type Usecase struct {
	threadRepo *threadRepo.Repository
	//postRepo   *postRepo.Repository
	userRepo   *userRepo.Repository
	forumRepo  *forumRepo.Repository
}

func NewThreadUsecase(tR *threadRepo.Repository, /*pR *postRepo.Repository,*/ uR *userRepo.Repository, fR *forumRepo.Repository) *Usecase {
	return &Usecase{
		threadRepo: tR,
		//postRepo:   pR,
		userRepo:   uR,
		forumRepo:  fR,
	}
}

func (tU *Usecase) ThreadDetails(slug_or_id string) (*models.Thread, error) {
	thread,err := tU.threadRepo.ThreadDetails(slug_or_id) 
	if err != nil {
		return nil, err
	}
	return thread, nil
}

func (tU *Usecase) CreatePostsNew(posts *models.Posts, slug_or_id string) (*models.Posts, error) {
	thread, err := tU.threadRepo.ThreadDetails(slug_or_id)
	if err != nil {
		return nil, err
	}

	if len(posts.Posts) == 0 {
		return &models.Posts{}, nil
	}

	newPosts, err := tU.threadRepo.CreatePostsNew(thread.ID,thread.Forum, posts)
	if err != nil {
		return nil, err
	}
	return newPosts, nil
}

func (tU *Usecase) ThreadVote(vote *models.Vote, slug_or_id string) (*models.Thread, error) {
	user ,err := tU.userRepo.GetUserByNickname(vote.Nickname)
	if err != nil {
		return nil, err
	}

	threadInfo, err := tU.threadRepo.ThreadDetails(slug_or_id)
	if err != nil {
		return nil, err
	}
	vote.Nickname = user.Nickname
	err = tU.threadRepo.ThreadVote(vote, threadInfo.ID)
	if err != nil {
		return nil, err
	}

	thread, err := tU.threadRepo.ThreadDetails(slug_or_id)
	if err != nil {
		return nil, err
	}
	return thread, err
}

func (tU *Usecase) ThreadGetPosts(slug_or_id, limit, since, sort, desc string) (*models.Posts, error) {
	thread, err := tU.ThreadDetails(slug_or_id)
	if err != nil {
		return nil,err
	}

	if limit == "" {
		limit = "100"
	}
	posts, err := tU.threadRepo.ThreadGetPosts(thread.ID, limit, since, sort, desc)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (tU *Usecase) UpdateThreadDetails(slug_or_id string, thread *models.Thread) (*models.Thread, error) {
	oldThread, err := tU.threadRepo.ThreadDetails(slug_or_id)
	if err != nil {
		return nil, err
	}
	thread.ID = oldThread.ID
	if thread.Title == "" {
		thread.Title = oldThread.Title
	}

	if thread.Message == "" {
		thread.Message = oldThread.Message
	}

	newThread, err := tU.threadRepo.UpdateThreadDetails(thread)
	if err != nil {
		return nil, err
	}
	return newThread, nil
}