

DROP TABLE IF EXISTS 'users';
create table 'users'
(
                        'uid' bigint not null  comment '用户id',
                        'username' varchar(255) not null comment '用户名称',
                        'password' varchar(255) not null comment '用户密码',
                        'avatar_url' varchar(255) not null comment '用户头像' default 'https://th.bing.com/th/id/OIP.VlXsxUWAoGSSgksl1PTANwHaHa?rs=1&pid=ImgDetMain',
                        'created_at' varchar(255) not null comment '创建时间',
                        'updated_at' varchar(255) not null comment '登录时间',
                        'deleted_at' varchar(255) not null comment '注销时间',
                        'code_url' varchar(255) not null comment '存储位置',
                        'secret' varchar(255) not null comment '密钥',
                        primary key (uid)
) engine=InnoDB default charset=utf8mb4;

DROP TABLE IF EXISTS 'videos';
create table 'videos'(
                         'vid' bigint not null  comment '视频ID',
                         'uid' bigint not null comment '作者ID',
                         'video_url' varchar(255) not null comment '视频url',
                         'cover_url' varchar(255) not null comment '封面url',
                         'title' varchar(255) not null comment '标题',
                         'description' varchar(255) not null comment '描述',
                         'visit_count' bigint not null comment '观看量',
                         'like_count' bigint not null comment '点赞数',
                         'comment_count' bigint not null comment '评论数',
                         'created_at' varchar(255) not null comment '创建时间',
                         'updated_at' varchar(255) not null comment '修改时间',
                         'deleted_at' varchar(255) not null comment '删除时间',
                         primary key (vid),
                         foreign key (`uid`)
                            references `users`(`uid`)
                            on delete cascade

)engine=InnoDB default charset = utf8mb4;

DROP TABLE IF EXISTS 'relations';
create table 'relations'(
                            'fid' bigint not null ,
                            'from_uid' bigint not null comment '来源',
                            'to_uid' bigint not null comment '目标',
                            'created_at' varchar(255) not null comment '创建时间',
                            'updated_at' varchar(255) not null comment '修改时间',
                            'deleted_at' varchar(255) not null comment '删除时间',
                            primary key (`fid`),
                            foreign key (from_uid) references `users`(`uid`)
                            on delete cascade

                        )engine=InnoDB default charset = utf8mb4;

DROP TABLE IF EXISTS 'messages';
create table 'messages'(
                           'created_at' varchar(255) not null comment '创建时间',
                           'deleted_at' varchar(255) not null comment '删除时间',
                           'from_uid' varchar(255) not null comment '来源',
                           'to_uid' varchar(255) not null comment '目标',
                           'type' bigint not null comment '类型',
                           'content' varchar(255) not null comment '内容',
                           'read_tag' BOOLEAN DEFAULT false comment '已读标签'
                            ,
                            foreign key (`from_uid`) references users(`uid`)
                            on delete cascade
)engine=InnoDB default charset = utf8mb4;

DROP TABLE IF EXISTS 'likes';
create table 'likes'(
                        'lid' bigint not null ,
                        'vid' bigint default null ,
                        'uid' bigint not null,
                        'cid' bigint default null ,
                        'created_at' varchar(255) not null comment '创建时间',
                        'updated_at' varchar(255) not null comment '修改时间',
                        'deleted_at' varchar(255) not null comment '删除时间',
                        primary key (lid),
                        foreign key (vid)
                            references `videos`(`vid`)
                            on delete CASCADE ,
                        foreign key (cid)
                            references `comments`(cid)
                            on delete cascade
)engine=InnoDB default charset = utf8mb4;

DROP TABLE IF EXISTS  'comments';
create table 'comments'(

                           'cid' bigint not null  comment '评论id',
                           'vid' bigint not null comment '视频id',
                           'uid' bigint not null comment '用户id',
                           'parent_id' bigint not null comment '父评论id',
                         --  'like_count' bigint not null comment '点赞数',
                         -- 'child_count' bigint not null comment '子评论数',
                           'content' bigint not null comment '内容',
                           'created_at' varchar(255) not null comment '创建时间',
                           'updated_at' varchar(255) not null comment '修改时间',
                           'deleted_at' varchar(255) not null comment '删除时间',
                           primary key (cid),
                           foreign key (`vid`)
                                references `videos`(`vid`)
                                on delete cascade
)engine=InnoDB default charset = utf8mb4;