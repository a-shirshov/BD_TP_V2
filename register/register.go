package register

import (
	"github.com/gorilla/mux"
	forumD "bd_tp_V2/forum/delivery/http"
	userD "bd_tp_V2/user/delivery/http"	
	threadD "bd_tp_V2/thread/delivery/http"
	postD "bd_tp_V2/post/delivery/http"
	serviceD "bd_tp_V2/service/delivery/http"

)

func UserEndpoints(r *mux.Router, userD *userD.UserDelivery) {
	r.HandleFunc("/{nickname}/create",userD.CreateUserV2).Methods("POST")
	r.HandleFunc("/{nickname}/profile",userD.ProfileInfoV2).Methods("GET")
	r.HandleFunc("/{nickname}/profile",userD.UpdateProfileV2).Methods("POST")
}

func ForumEndpoints(r *mux.Router, forumD *forumD.ForumDelivery) {
	r.HandleFunc("/create",forumD.CreateForumV2).Methods("POST")
	r.HandleFunc("/{slug}/details",forumD.ForumDetailsV2).Methods("GET")
	r.HandleFunc("/{slug}/create", forumD.ForumThreadCreateV2).Methods("POST")
	r.HandleFunc("/{slug}/threads", forumD.GetThreadsByForum).Methods("GET")
	r.HandleFunc("/{slug}/users",forumD.GetForumUsers).Methods("GET")
}

func ThreadEndpoints(r *mux.Router, threadD *threadD.ThreadDelivery) {
	r.HandleFunc("/{slug_or_id}/create",threadD.CreatePostsNew).Methods("POST") 
	r.HandleFunc("/{slug_or_id}/details",threadD.ThreadDetails).Methods("GET")
	r.HandleFunc("/{slug_or_id}/details",threadD.ThreadDetailsUpdate).Methods("POST")
	r.HandleFunc("/{slug_or_id}/vote",threadD.ThreadVote).Methods("POST")
	r.HandleFunc("/{slug_or_id}/posts",threadD.ThreadGetPosts).Methods("GET")
}

func PostEndpoints(r *mux.Router, postD *postD.PostDelivery) {
	r.HandleFunc("/{id}/details",postD.PostDetails).Methods("GET")
    r.HandleFunc("/{id}/details",postD.UpdatePost).Methods("POST")
}

func ServiceEndpoints(r *mux.Router, serviceD *serviceD.ServiceDelivery) {
	r.HandleFunc("/clear",serviceD.Clear).Methods("POST")
	r.HandleFunc("/status",serviceD.GetStatus).Methods("GET")
}
