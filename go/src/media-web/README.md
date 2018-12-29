获取用户所有视频 /user/:username/videos		GET 200 400 500
获取一个视频 /user/:username/videos/:vid-id  GET 200 400 500
删除一个视频 /user/:username/videos/:vid-id  DELETE 204 400 401 403 500

显示评论 /videos/:vid-id/comments GET 200 400 500
添加评论 POST 201 400 500
删除 /videos/:vid-id/comment/:comment-id  DELETE 204 400 4010 403 500

handle - validation(request,user) [data model , error handle]- business logic - reponse

处理程序 验证 逻辑 响应结果

user 用户表：id login_name pwd 
video_info 视频表：id (varchar) author_id(u int) name(text) display_ctime(text) create_time (datetime)

comments 评论： id （varchar） video_id author_id content time
sessions 登录： session_id login_name pwd
