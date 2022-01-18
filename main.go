package main

import (
	"bd_tp_V2/register"
	userDelivery "bd_tp_V2/user/delivery/http"
	userRepo "bd_tp_V2/user/repository"
	userUsecase "bd_tp_V2/user/usecase"

	/*
	forumDelivery "bd_tp_V2/forum/delivery/http"
	forumRepo "bd_tp_V2/forum/repository"
	forumUsecase "bd_tp_V2/forum/usecase"

	threadDelivery "bd_tp_V2/thread/delivery/http"
	threadRepo "bd_tp_V2/thread/repository"
	threadUsecase "bd_tp_V2/thread/usecase"

	postDelivery "bd_tp_V2/post/delivery/http"
	postRepo "bd_tp_V2/post/repository"
	postUsecase "bd_tp_V2/post/usecase"

	serviceDelivery "bd_tp_V2/service/delivery/http"
	serviceRepo "bd_tp_V2/service/repository"
	serviceUsecase "bd_tp_V2/service/usecase"
	*/

	"bd_tp_V2/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	db, err := utils.InitPostgresDB()
	if err != nil {
		fmt.Println(err)
		return 
	}
	err = utils.Prepare(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	userR := userRepo.NewRepository(db)
	/*
	forumR := forumRepo.NewForumRepository(db)
	threadR := threadRepo.NewThreadRepository(db)
	postR := postRepo.NewPostRepository(db)
	serviceR := serviceRepo.NewServiceRepository(db)
	*/

	userU := userUsecase.NewUserUsecase(userR)
	/*
	forumU := forumUsecase.NewForumUsecase(forumR, userR, threadR)
	threadU := threadUsecase.NewThreadUsecase(threadR, postR, userR, forumR)
	postU := postUsecase.NewPostUsecase(postR, userR, forumR, threadR)
	serviceU := serviceUsecase.NewServiceUseCase(serviceR)
	*/

	userD := userDelivery.NewUserDelivery(userU)
	/*
	forumD := forumDelivery.NewForumDelivery(forumU)
	threadD := threadDelivery.NewThreadDelivery(threadU)
	postD := postDelivery.NewPostDelivery(postU)
	serviceD := serviceDelivery.NewServiceDelivery(serviceU)
	*/

	r := mux.NewRouter()
	rApi := r.PathPrefix("/api").Subrouter()
	userRouter := rApi.PathPrefix("/user").Subrouter()
	register.UserEndpoints(userRouter, userD)
	/*
	forumRouter := rApi.PathPrefix("/forum").Subrouter()
	register.ForumEndpoints(forumRouter, forumD)
	threadRouter := rApi.PathPrefix("/thread").Subrouter()
	register.ThreadEndpoints(threadRouter, threadD)
	postRouter := rApi.PathPrefix("/post").Subrouter()
	register.PostEndpoints(postRouter, postD)
	serviceRouter := rApi.PathPrefix("/service").Subrouter()
	register.ServiceEndpoints(serviceRouter, serviceD)
	*/

	err = http.ListenAndServe(":5000", r)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
}
