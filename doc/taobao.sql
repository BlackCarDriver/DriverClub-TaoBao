drop table t_message;

drop table t_upload;

drop table t_collect;

drop table t_comment;

drop table t_concern;

drop table t_goods;

drop table t_user;

/*==============================================================*/
/* Table: t_User                                                */
/*==============================================================*/
create table t_User (
   id                   VARCHAR(50)          not null,
   password             VARCHAR(50)          not null default '123456',
   email                VARCHAR(50)          null,
   name                 VARCHAR(50)          null,
   sex                  VARCHAR(50)          null default 'boy',
   dorm                 VARCHAR(50)          null,
   sign                 VARCHAR(400)         null,
   major                VARCHAR(50)          null,
   headimg              VARCHAR(100)         null,
   phone                VARCHAR(50)          null,
   qq                   VARCHAR(50)          null,
   emails               VARCHAR(50)          null,
   credits              INT4                 not null default 0,
   leave                INT4                 not null default 1,
   rank                 INT4                 not null,
   visit                INT4                 not null default 0,
   likes                INT4                 not null default 0,
   lasttime             TIMESTAMP            not null default now(),
   constraint PK_T_USER primary key (id)
);

comment on column t_User.id is
'账号';

comment on column t_User.password is
'密码';

comment on column t_User.email is
'注册使用的邮箱';

comment on column t_User.name is
'用户名';

comment on column t_User.sex is
'性别';

comment on column t_User.dorm is
'宿舍楼栋';

comment on column t_User.sign is
'个性签名';

comment on column t_User.major is
'专业';

comment on column t_User.headimg is
'头像地址';

comment on column t_User.phone is
'手机号码';

comment on column t_User.qq is
'qq号码';

comment on column t_User.emails is
'用户自己设置的邮箱';

comment on column t_User.credits is
'积分';

comment on column t_User.leave is
'等级';

comment on column t_User.rank is
'排名';

comment on column t_User.visit is
'主页访问量';

comment on column t_User.likes is
'被点赞次数';

comment on column t_User.lasttime is
'上次登录的时间';

/*==============================================================*/
/* Table: t_collect                                             */
/*==============================================================*/
create table t_collect (
   userid               VARCHAR(50)          not null,
   goodsid               VARCHAR(50)          not null,
   id                   VARCHAR(50)          null,
   goo_id               VARCHAR(50)          null,
   time                 TIMESTAMP            not null default now(),
   constraint PK_T_COLLECT primary key (userid, goodsid)
);

comment on column t_collect.userid is
'用户id';

comment on column t_collect.goodsid is
'商品id';

comment on column t_collect.id is
'账号';

comment on column t_collect.goo_id is
'id';

comment on column t_collect."time" is
'收藏时间';

/*==============================================================*/
/* Table: t_comment                                             */
/*==============================================================*/
create table t_comment (
   userid               VARCHAR(50)          not null,
   goodsid              VARCHAR(50)          not null,
   content              VARCHAR(400)         null,
   time               	TIMESTAMP            not null default now()
);

comment on column t_comment.userid is
'评论者id';

comment on column t_comment.goodsid is
'商品id';

comment on column t_comment.content is
'评论内容';

comment on column t_comment."time" is
'评论时间';

/*==============================================================*/
/* Table: t_concern                                             */
/*==============================================================*/
create table t_concern (
   id1                  VARCHAR(50)          not null,
   id2                  VARCHAR(50)          not null,
   time               	VARCHAR(50)          null,
   constraint PK_T_CONCERN primary key (id1, id2)
);

comment on table t_concern is
'关注';

comment on column t_concern.id1 is
'主用户id';

comment on column t_concern.id2 is
'关注的用户id';

comment on column t_concern."time" is
'关注时间';

/*==============================================================*/
/* Table: t_goodss                                               */
/*==============================================================*/
create table t_goods (
   id                   VARCHAR(50)          not null,
   name                 VARCHAR(50)          null,
   title                VARCHAR(50)          null,
   type                 VARCHAR(50)          null,
   tag                  VARCHAR(50)          null,
   price                FLOAT8               null,
   file                 VARCHAR(100)         null,
   headimg              VARCHAR(100)         null,
   visit                INT4                 null default 0,
   "like"              	INT4                 default 0,
   state                INT4                 default 1,
   constraint PK_T_GOODS primary key (id)
);

comment on column t_goods.id is
'id';

comment on column t_goods.name is
'商品名称';

comment on column t_goods.title is
'标题';

comment on column t_goods.type is
'分类';

comment on column t_goods.tag is
'标签';

comment on column t_goods.price is
'标价';

comment on column t_goods.file is
'文件名';

comment on column t_goods.headimg is
'封面图片名';

comment on column t_goods.visit is
'浏览次数';

comment on column t_goods."like" is
'点赞次数';

comment on column t_goods.state is
'商品状态';

/*==============================================================*/
/* Table: t_message                                             */
/*==============================================================*/
create table t_message (
   senderid             VARCHAR(50)          null,
   receiverid           VARCHAR(50)          null,
   content              VARCHAR(400)         null,
   time               	TIMESTAMP            not null default now(),
   state                INT4                 null
);

comment on column t_message.senderid is
'发送者id';

comment on column t_message.receiverid is
'接收者id';

comment on column t_message.content is
'消息内容';

comment on column t_message."time" is
'发送时间';

comment on column t_message.state is
'状态';

/*==============================================================*/
/* Table: t_upload                                              */
/*==============================================================*/
create table t_upload (
   userid               VARCHAR(50)          not null,
   goodsid              VARCHAR(50)          not null,
   time               	TIMESTAMP            not null default now(),
   constraint PK_T_UPLOAD primary key (userid, goodsid)
);

comment on column t_upload.userid is
'上传者id';

comment on column t_upload.goodsid is
'商品id';

comment on column t_upload."time" is
'上传时间';



alter table t_collect
   add constraint collect_fk foreign key (userid)
      references t_User (id)
      on delete restrict on update restrict;

alter table t_collect
   add constraint collect_fk2 foreign key (goodsid)
      references t_goods (id)
      on delete restrict on update restrict;

alter table t_comment
   add constraint comment_fk foreign key (userid)
      references t_User (id)
      on delete restrict on update restrict;

alter table t_comment
   add constraint comment_fk2 foreign key (goodsid)
      references t_goods (id)
      on delete restrict on update restrict;

alter table t_concern
   add constraint concern_fk foreign key (id1)
      references t_User (id)
      on delete restrict on update restrict;

alter table t_concern
   add constraint concern_fk2 foreign key (id2)
      references t_User (id)
      on delete restrict on update restrict;

alter table t_message
   add constraint message_fk foreign key (senderid)
      references t_User (id)
      on delete restrict on update restrict;

alter table t_message
   add constraint message_fk2 foreign key (receiverid)
      references t_User (id)
      on delete restrict on update restrict;

alter table t_upload
   add constraint upload_fk foreign key (userid)
      references t_User (id)
      on delete restrict on update restrict;

alter table t_upload
   add constraint upload_fk2 foreign key (goodsid)
      references t_goods (id)
      on delete restrict on update restrict;
