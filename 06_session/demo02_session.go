package main

import (
	"net/http"
	"html/template"
	session "06_session/session"
	_ "06_session/memory"
	"fmt"
	md5 "crypto/md5"
	"time"
	"io"
)

func main() {
	http.HandleFunc("/login",login)
	http.HandleFunc("/count",count)
	http.ListenAndServe(":8080",nil)
}

//然后在init函数中初始化
var globalSessions *session.Manager
func init() {
	globalSessions ,_= session.NewManager("memory","gosessionid",3600)
	fmt.Println("globalSessions,",globalSessions)
	//go globalSessions.GC()
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login....")
	sess := globalSessions.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
	} else {
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
}



func count(w http.ResponseWriter, r *http.Request) {



	fmt.Println("count.......")
	sess := globalSessions.SessionStart(w, r)
	fmt.Println("sess:",sess)


	//方案1
	//h := md5.New()
	//salt:="hanru%^7&8888"
	//io.WriteString(h,salt+time.Now().String())
	//token:=fmt.Sprintf("%x",h.Sum(nil))
	//if r.Form["token"][0]!=token{
	//	//提示登录
	//}
	//sess.Set("token",token)
	//方案2
	createtime := sess.Get("createtime")
	if createtime == nil {
		sess.Set("createtime", time.Now().Unix())
	} else if (createtime.(int64) + 60) < (time.Now().Unix()) {
		globalSessions.SessionDestroy(w, r)
		sess = globalSessions.SessionStart(w, r)
	}


	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	fmt.Println("...")
	t, _ := template.ParseFiles("count.gtpl")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get("countnum"))


}