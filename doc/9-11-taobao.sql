drop view v_rank;

drop view v_navingmsg;

drop view v_mymessage;

drop view v_mygoods;

drop view v_mydata;

drop view v_mycollect;

drop view v_hpgoodslist;

drop view v_goodslist;

drop view v_goods_detail;

drop view v_concern;

drop table public.t_collect;

drop table public.t_comment;

drop table public.t_concern;

drop table public.t_goods;

drop table public.t_goods_like;

drop table public.t_message;

drop table public.t_upload;

drop table public.t_user;

drop table public.t_user_like;

/*==============================================================*/
/* User: public                                                 */
/*==============================================================*/
/*==============================================================*/
/* Table: t_collect                                             */
/*==============================================================*/
create table public.t_collect (
   userid               character varying(50) not null,
   goodsid              character varying(50) not null,
   "time"               timestamp            null,
   constraint pk_t_collect primary key (userid, goodsid)
);

comment on table public.t_collect is
'商品收藏';

comment on column t_collect.userid is
'用户id';

comment on column t_collect.goodsid is
'商品id';

comment on column t_collect."time" is
'收藏时间';

/*==============================================================*/
/* Table: t_comment                                             */
/*==============================================================*/
create table public.t_comment (
   id                   SERIAL               not null,
   userid               character varying(50) not null,
   goodsid              character varying(50) not null,
   content              character varying(400) null default '',
   "time"               timestamp            null,
   constraint PK_T_COMMENT primary key (id)
);

comment on table public.t_comment is
'商品评论数据';

comment on column t_comment.id is
'评论id';

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
create table public.t_concern (
   id1                  character varying(50) not null,
   id2                  character varying(50) not null,
   "time"               timestamp            null,
   constraint pk_t_concern primary key (id1, id2)
);

comment on table public.t_concern is
'用户关注用户数据';

comment on column t_concern.id1 is
'主用户id';

comment on column t_concern.id2 is
'关注的用户id';

comment on column t_concern."time" is
'关注时间';

/*==============================================================*/
/* Table: t_goods                                               */
/*==============================================================*/
create table public.t_goods (
   id                   character varying(50) not null,
   name                 character varying(50) null default '',
   title                character varying(50) null default '',
   type                 character varying(50) null default '',
   tag                  character varying(50) null default '',
   price                double precision     null default '0.0',
   file                 text                 null default '',
   headimg              character varying(200) null default '',
   visit                integer              null default '0',
   "like"               integer              null default '0',
   state                integer              null default '1',
   constraint pk_t_goods primary key (id)
);

comment on table public.t_goods is
'商品信息';

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
'详细描叙';

comment on column t_goods.headimg is
'封面图片名';

comment on column t_goods.visit is
'浏览次数';

comment on column t_goods."like" is
'点赞次数';

comment on column t_goods.state is
'商品状态';

/*==============================================================*/
/* Table: t_goods_like                                          */
/*==============================================================*/
create table public.t_goods_like (
   userid               character varying(50) not null,
   goodsid              character varying(50) not null,
   constraint PK_T_GOODS_LIKE primary key (userid, goodsid)
);

comment on table public.t_goods_like is
'商品被用户点赞的数据';

comment on column t_goods_like.userid is
'点赞者id';

comment on column t_goods_like.goodsid is
'被点赞商品id';

/*==============================================================*/
/* Table: t_message                                             */
/*==============================================================*/
create table public.t_message (
   id                   SERIAL               not null,
   senderid             character varying(50) null,
   receiverid           character varying(50) null,
   content              character varying(400) null default '',
   "time"               timestamp            null,
   state                integer              null default '0',
   constraint PK_T_MESSAGE primary key (id)
);

comment on table public.t_message is
'用户消息数据';

comment on column t_message.id is
'消息id';

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
create table public.t_upload (
   userid               character varying(50) not null,
   goodsid              character varying(50) not null,
   "time"               timestamp            null,
   constraint pk_t_upload primary key (userid, goodsid)
);

comment on table public.t_upload is
'用户上传商品记录数据';

comment on column t_upload.userid is
'上传者id';

comment on column t_upload.goodsid is
'商品id';

comment on column t_upload."time" is
'上传时间';

/*==============================================================*/
/* Table: t_user                                                */
/*==============================================================*/
create table public.t_user (
   id                   character varying(50) not null,
   password             character varying(50) not null default '123456',
   email                character varying(50) null default NULL,
   name                 character varying(50) null default '',
   sex                  character varying(50) null default 'boy',
   dorm                 character varying(50) null default '',
   sign                 character varying(400) null default '',
   major                character varying(50) null default '',
   headimg              character varying(200) null default '',
   phone                character varying(50) null default '',
   qq                   character varying(50) null default '',
   emails               character varying(50) null default '',
   credits              integer              null default '0',
   leave                integer              null default '1',
   rank                 integer              null default '99999',
   visit                integer              null default '0',
   lasttime             timestamp            null,
   grade                integer              null,
   colleage             character varying(50) null,
   likes                integer              null default '0',
   constraint pk_t_user primary key (id)
);

comment on table public.t_user is
'用户信息';

comment on column t_user.id is
'账号';

comment on column t_user.password is
'密码';

comment on column t_user.email is
'注册使用的邮箱';

comment on column t_user.name is
'用户名';

comment on column t_user.sex is
'性别';

comment on column t_user.dorm is
'宿舍楼栋';

comment on column t_user.sign is
'个性签名';

comment on column t_user.major is
'专业';

comment on column t_user.headimg is
'头像地址';

comment on column t_user.phone is
'手机号码';

comment on column t_user.qq is
'qq号码';

comment on column t_user.emails is
'用户自己设置的邮箱';

comment on column t_user.credits is
'积分';

comment on column t_user.leave is
'等级';

comment on column t_user.rank is
'排名';

comment on column t_user.visit is
'主页访问量';

comment on column t_user.lasttime is
'上次登录的时间';

comment on column t_user.grade is
'所在年级';

comment on column t_user.colleage is
'所在学院';

/*==============================================================*/
/* Table: t_user_like                                           */
/*==============================================================*/
create table public.t_user_like (
   userid1              character varying(50) not null,
   userid2              character varying(50) not null,
   constraint t_user_like_pkey primary key (userid1, userid2)
);

comment on table public.t_user_like is
'用户点赞用户数据';

comment on column t_user_like.userid1 is
'收藏者id';

comment on column t_user_like.userid2 is
'被收藏者id';

/*==============================================================*/
/* View: v_concern                                              */
/*==============================================================*/
create or replace view v_concern as
SELECT u2.id AS id2,
    u2.name AS name2,
    u2.headimg AS headimg2,
    u1.id,
    u1.name,
    u1.headimg
   FROM t_user u1,
    t_user u2,
    t_concern c
  WHERE (((c.id1) = (u1.id)) AND ((c.id2) = (u2.id)));

/*==============================================================*/
/* View: v_goods_detail                                         */
/*==============================================================*/
create or replace view v_goods_detail as
SELECT g.id AS goodsid,
    u.name AS username,
    u.id AS userid,
    g.name,
    p."time",
    g.headimg,
    g.price,
    g.title,
    g.type,
    g.tag,
    g.visit,
    g."like",
    g.file AS detail
   FROM t_goods g,
    t_upload p,
    t_user u
  WHERE (((p.goodsid) = (g.id)) AND ((p.userid) = (u.id)) AND (g.state = 1));

/*==============================================================*/
/* View: v_goodslist                                            */
/*==============================================================*/
create or replace view v_goodslist as
SELECT u.id AS uid,
    u.name AS uname,
    g.id AS gid,
    g.name AS gname,
    g.title,
    g.headimg,
    g.price,
    g.state,
    g.type,
    g.tag,
    p."time"
   FROM t_upload p,
    t_goods g,
    t_user u
  WHERE (((p.userid) = (u.id)) AND ((p.goodsid) = (g.id)));

/*==============================================================*/
/* View: v_hpgoodslist                                          */
/*==============================================================*/
create or replace view v_hpgoodslist as
SELECT v_goodslist.uid AS userid,
    v_goodslist.uname AS username,
    v_goodslist.gid AS id,
    v_goodslist.gname AS name,
    v_goodslist.title,
    v_goodslist.price,
    v_goodslist.type,
    v_goodslist.tag,
    v_goodslist."time",
    v_goodslist.headimg
   FROM v_goodslist
  WHERE (v_goodslist.state = 1);

/*==============================================================*/
/* View: v_mycollect                                            */
/*==============================================================*/
create or replace view v_mycollect as
SELECT u.id AS uid,
    g.id,
    g.name,
    g.title,
    g.price,
    g.headimg
   FROM t_collect c,
    t_user u,
    t_goods g
  WHERE (((c.userid) = (u.id)) AND ((c.goodsid) = (g.id)));

/*==============================================================*/
/* View: v_mydata                                               */
/*==============================================================*/
create or replace view v_mydata as
SELECT t_user.headimg,
    t_user.id,
    t_user.name,
    t_user.sex,
    t_user.sign,
    t_user.grade,
    t_user.colleage,
    t_user.major,
    t_user.emails,
    t_user.qq,
    t_user.phone,
    t_user.lasttime,
    t_user.dorm,
    t_user.leave,
    t_user.credits,
    t_user.visit
   FROM t_user;

/*==============================================================*/
/* View: v_mygoods                                              */
/*==============================================================*/
create or replace view v_mygoods as
SELECT u.id AS uid,
    g.id,
    g.name,
    g.title,
    p."time",
    g.price,
    g.headimg
   FROM t_user u,
    t_goods g,
    t_upload p
  WHERE (((p.userid) = (u.id)) AND ((p.goodsid) = (g.id)) AND (g.state = 1));

/*==============================================================*/
/* View: v_mymessage                                            */
/*==============================================================*/
create or replace view v_mymessage as
CREATE OR REPLACE VIEW public.v_mymessage AS
 SELECT u2.id as uid,
 	m.id as mid,
    m."time",
    m.content,
    u1.name,
    u1.headimg
   FROM t_message m,
    t_user u1,
    t_user u2
  WHERE m.senderid::text = u1.id::text AND m.receiverid::text = u2.id::text AND m.state = 0;

/*==============================================================*/
/* View: v_navingmsg                                            */
/*==============================================================*/
create or replace view v_navingmsg as
SELECT t_user.id,
    t_user.name,
    t_user.headimg,
    t_user.credits,
    t_user.leave,
    t_user.lasttime,
    ( SELECT count(*) AS count
           FROM t_upload
          WHERE ((t_upload.userid) = (t_user.id))) AS goodsnum,
    ( SELECT count(*) AS count
           FROM t_message
          WHERE ((t_message.receiverid) = (t_user.id))) AS messagenum
   FROM t_user;

/*==============================================================*/
/* View: v_rank                                                 */
/*==============================================================*/
create or replace view v_rank as
SELECT row_number() OVER (ORDER BY t_user.credits DESC) AS rank,
    t_user.id AS userid,
    t_user.name,
    t_user.credits
   FROM t_user
  ORDER BY t_user.credits DESC
 LIMIT 10;

alter table t_collect
   add constraint collect_fk foreign key (userid)
      references t_user (id)
      on delete cascade on update restrict;

alter table t_collect
   add constraint collect_fk2 foreign key (goodsid)
      references t_goods (id)
      on delete cascade on update restrict;

alter table t_comment
   add constraint comment_fk foreign key (userid)
      references t_user (id)
      on delete cascade on update restrict;

alter table t_comment
   add constraint comment_fk2 foreign key (goodsid)
      references t_goods (id)
      on delete cascade on update restrict;

alter table t_concern
   add constraint concern_fk foreign key (id1)
      references t_user (id)
      on delete cascade on update restrict;

alter table t_concern
   add constraint concern_fk2 foreign key (id2)
      references t_user (id)
      on delete cascade on update restrict;

alter table t_goods_like
   add constraint FK_T_GOODS__REFERENCE_T_GOODS foreign key (goodsid)
      references t_goods (id)
      on delete restrict on update restrict;

alter table t_goods_like
   add constraint FK_T_GOODS__T_GOODS_L_T_USER foreign key (userid)
      references t_user (id)
      on delete cascade on update restrict;

alter table t_message
   add constraint message_fk foreign key (senderid)
      references t_user (id)
      on delete cascade on update restrict;

alter table t_message
   add constraint message_fk2 foreign key (receiverid)
      references t_user (id)
      on delete cascade on update restrict;

alter table t_upload
   add constraint upload_fk foreign key (userid)
      references t_user (id)
      on delete cascade on update restrict;

alter table t_upload
   add constraint upload_fk2 foreign key (goodsid)
      references t_goods (id)
      on delete cascade on update restrict;

alter table t_user_like
   add constraint fk_user1 foreign key (userid1)
      references t_user (id)
      on delete cascade on update restrict;

alter table t_user_like
   add constraint fk_usre2 foreign key (userid2)
      references t_user (id)
      on delete cascade on update restrict;
