drop table public.t_collect;
drop table public.t_comment;
drop table public.t_concern;
drop table public.t_message;
drop table public.t_upload;
drop table public.t_User;
drop table public.t_goods;


/*==============================================================*/
/* User: public                                                 */
/*==============================================================*/
/*==============================================================*/
/* Table: t_User                                                */
/*==============================================================*/
create table public.t_User (
   id                   VARCHAR(50)          not null,
   password             VARCHAR(50)          not null default '123456',
   email                VARCHAR(50)          default null,
   name                 VARCHAR(50)          default '',
   sex                  VARCHAR(50)          default 'boy',
   dorm                 VARCHAR(50)          default '',
   sign                 VARCHAR(400)         default '',
   major                VARCHAR(50)          default '',
   headimg              VARCHAR(200)         default '',
   phone                VARCHAR(50)          default '',
   qq                   VARCHAR(50)          default '',
   emails               VARCHAR(50)          default '',
   credits              INT4                 default 0,
   leave                INT4                 default 1,
   rank                 INT4                 default 99999,
   visit                INT4                 default 0,
   likes                INT4                 default 0,
   lasttime             TIMESTAMP            default now(),
   constraint PK_T_USER primary key (id)
);

comment on column t_User.id is
'�˺�';

comment on column t_User.password is
'����';

comment on column t_User.email is
'ע��ʹ�õ�����';

comment on column t_User.name is
'�û���';

comment on column t_User.sex is
'�Ա�';

comment on column t_User.dorm is
'����¥��';

comment on column t_User.sign is
'����ǩ��';

comment on column t_User.major is
'רҵ';

comment on column t_User.headimg is
'ͷ���ַ';

comment on column t_User.phone is
'�ֻ�����';

comment on column t_User.qq is
'qq����';

comment on column t_User.emails is
'�û��Լ����õ�����';

comment on column t_User.credits is
'����';

comment on column t_User.leave is
'�ȼ�';

comment on column t_User.rank is
'����';

comment on column t_User.visit is
'��ҳ������';

comment on column t_User.likes is
'�����޴���';

comment on column t_User.lasttime is
'�ϴε�¼��ʱ��';

/*==============================================================*/
/* Table: t_collect                                             */
/*==============================================================*/
create table public.t_collect (
   userid               VARCHAR(50)          not null,
   goodsid              VARCHAR(50)          not null,
   "time"               TIMESTAMP            default now(),
   constraint PK_T_COLLECT primary key (userid, goodsid)
);

comment on column t_collect.userid is
'�û�id';

comment on column t_collect.goodsid is
'��Ʒid';

comment on column t_collect."time" is
'�ղ�ʱ��';

/*==============================================================*/
/* Table: t_comment                                             */
/*==============================================================*/
create table public.t_comment (
   userid               VARCHAR(50)          not null,
   goodsid              VARCHAR(50)          not null,
   content              VARCHAR(400)         default '',
   "time"               TIMESTAMP            default now()
);

comment on column t_comment.userid is
'������id';

comment on column t_comment.goodsid is
'��Ʒid';

comment on column t_comment.content is
'��������';

comment on column t_comment."time" is
'����ʱ��';

/*==============================================================*/
/* Table: t_concern                                             */
/*==============================================================*/
create table public.t_concern (
   id1                  VARCHAR(50)          not null,
   id2                  VARCHAR(50)          not null,
   "time"              	TIMESTAMP          	 default now(),
   constraint PK_T_CONCERN primary key (id1, id2)
);

comment on column t_concern.id1 is
'���û�id';

comment on column t_concern.id2 is
'��ע���û�id';

comment on column t_concern."time" is
'��עʱ��';

/*==============================================================*/
/* Table: t_goods                                               */
/*==============================================================*/
create table public.t_goods (
   id                   VARCHAR(50)          not null,
   name                 VARCHAR(50)          default '',
   title                VARCHAR(50)          default '',
   type                 VARCHAR(50)          default '',
   tag                  VARCHAR(50)          default '',
   price                FLOAT8               default 0.0,
   file                 TEXT                 default '',
   headimg              VARCHAR(200)         default '',
   visit                INT4                 default 0,
   "like"               INT4                 default 0,
   state                INT4                 default 1,
   constraint PK_T_GOODS primary key (id)
);

comment on column t_goods.id is
'id';

comment on column t_goods.name is
'��Ʒ����';

comment on column t_goods.title is
'����';

comment on column t_goods.type is
'����';

comment on column t_goods.tag is
'��ǩ';

comment on column t_goods.price is
'���';

comment on column t_goods.file is
'��ϸ����';

comment on column t_goods.headimg is
'����ͼƬ��';

comment on column t_goods.visit is
'�������';

comment on column t_goods."like" is
'���޴���';

comment on column t_goods.state is
'��Ʒ״̬';

/*==============================================================*/
/* Table: t_message                                             */
/*==============================================================*/
create table public.t_message (
   senderid             VARCHAR(50)          null,
   receiverid           VARCHAR(50)          null,
   content              VARCHAR(400)         default '',
   "time"               TIMESTAMP            default now(),
   state                INT4                 default 0
);

comment on column t_message.senderid is
'������id';

comment on column t_message.receiverid is
'������id';

comment on column t_message.content is
'��Ϣ����';

comment on column t_message."time" is
'����ʱ��';

comment on column t_message.state is
'״̬';

/*==============================================================*/
/* Table: t_upload                                              */
/*==============================================================*/
create table public.t_upload (
   userid               VARCHAR(50)          not null,
   goodsid              VARCHAR(50)          not null,
   "time"               TIMESTAMP            default now(),
   constraint PK_T_UPLOAD primary key (userid, goodsid)
);

comment on column t_upload.userid is
'�ϴ���id';

comment on column t_upload.goodsid is
'��Ʒid';

comment on column t_upload."time" is
'�ϴ�ʱ��';

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
