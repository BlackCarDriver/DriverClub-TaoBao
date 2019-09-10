
--t_collect
CREATE TABLE t_collect (
    userid character varying(50) NOT NULL,
    goodsid character varying(50) NOT NULL,
    "time" timestamp without time zone DEFAULT now(),
    constraint pk_t_collect primary key (userid, goodsid)
);
COMMENT ON COLUMN t_collect.userid IS '用户id';
COMMENT ON COLUMN t_collect.goodsid IS '商品id';
COMMENT ON COLUMN t_collect."time" IS '收藏时间';


--############################################################


--t_comment
CREATE TABLE t_comment (
    userid character varying(50) NOT NULL,
    goodsid character varying(50) NOT NULL,
    content character varying(400) DEFAULT '',
    "time" timestamp without time zone DEFAULT now()
);
COMMENT ON COLUMN t_comment.userid IS '评论者id';
COMMENT ON COLUMN t_comment.goodsid IS '商品id';
COMMENT ON COLUMN t_comment.content IS '评论内容';
COMMENT ON COLUMN t_comment."time" IS '评论时间';


--############################################################


--t_concern
CREATE TABLE t_concern (
    id1 character varying(50) NOT NULL,
    id2 character varying(50) NOT NULL,
    "time" timestamp without time zone DEFAULT now(),
    constraint pk_t_concern primary key (id1, id2)
);
COMMENT ON COLUMN t_concern.id1 IS '主用户id';
COMMENT ON COLUMN t_concern.id2 IS '关注的用户id';
COMMENT ON COLUMN t_concern."time" IS '关注时间';


--############################################################


--t_goods
CREATE TABLE t_goods (
    id character varying(50) NOT NULL,
    name character varying(50) DEFAULT '',
    title character varying(50) DEFAULT '',
    type character varying(50) DEFAULT '',
    tag character varying(50) DEFAULT '',
    price double precision DEFAULT 0.0,
    file text DEFAULT '',
    headimg character varying(200) DEFAULT '',
    visit integer DEFAULT 0,
    "like" integer DEFAULT 0,
    state integer DEFAULT 1,
    constraint pk_t_goods primary key (id)
);
COMMENT ON COLUMN t_goods.id IS 'id';
COMMENT ON COLUMN t_goods.name IS '商品名称';
COMMENT ON COLUMN t_goods.title IS '标题';
COMMENT ON COLUMN t_goods.type IS '分类';
COMMENT ON COLUMN t_goods.tag IS '标签';
COMMENT ON COLUMN t_goods.price IS '标价';
COMMENT ON COLUMN t_goods.file IS '详细描叙';
COMMENT ON COLUMN t_goods.headimg IS '封面图片名';
COMMENT ON COLUMN t_goods.visit IS '浏览次数';
COMMENT ON COLUMN t_goods."like" IS '点赞次数';
COMMENT ON COLUMN t_goods.state IS '商品状态';


--############################################################


--t_goods_like
CREATE TABLE t_goods_like (
    userid character varying(50) NOT NULL,
    goodsid character varying(50) NOT NULL,
    CONSTRAINT t_goods_like_pkey PRIMARY KEY (userid, goodsid);
);
COMMENT ON COLUMN t_goods_like.userid IS '收藏者id';
COMMENT ON COLUMN t_goods_like.goodsid IS '被收藏的商品id';


--############################################################


--t_message
CREATE TABLE t_message (
    senderid character varying(50),
    receiverid character varying(50),
    content character varying(400) DEFAULT '',
    "time" timestamp without time zone DEFAULT now(),
    state integer DEFAULT 0
);
COMMENT ON COLUMN t_message.senderid IS '发送者id';
COMMENT ON COLUMN t_message.receiverid IS '接收者id';
COMMENT ON COLUMN t_message.content IS '消息内容';
COMMENT ON COLUMN t_message."time" IS '发送时间';
COMMENT ON COLUMN t_message.state IS '状态';


--############################################################


--t_upload
CREATE TABLE t_upload (
    userid character varying(50) NOT NULL,
    goodsid character varying(50) NOT NULL,
    "time" timestamp without time zone DEFAULT now(),
    CONSTRAINT pk_t_upload PRIMARY KEY (userid, goodsid)
);
COMMENT ON COLUMN t_upload.userid IS '上传者id';
COMMENT ON COLUMN t_upload.goodsid IS '商品id';
COMMENT ON COLUMN t_upload."time" IS '上传时间';


--############################################################


--t_user
CREATE TABLE t_user (
    id character varying(50) NOT NULL,
    password character varying(50) DEFAULT '123456' NOT NULL,
    email character varying(50) DEFAULT NULL,
    name character varying(50) DEFAULT '',
    sex character varying(50) DEFAULT 'boy',
    dorm character varying(50) DEFAULT '',
    sign character varying(400) DEFAULT '',
    major character varying(50) DEFAULT '',
    headimg character varying(200) DEFAULT '',
    phone character varying(50) DEFAULT '',
    qq character varying(50) DEFAULT '',
    emails character varying(50) DEFAULT '',
    credits integer DEFAULT 0,
    leave integer DEFAULT 1,
    rank integer DEFAULT 99999,
    visit integer DEFAULT 0,
    lasttime timestamp without time zone DEFAULT now(),
    grade integer,
    colleage character varying(50),
    likes integer DEFAULT 0,
    CONSTRAINT pk_t_user PRIMARY KEY (id)
);
COMMENT ON COLUMN t_user.id IS '账号';
COMMENT ON COLUMN t_user.password IS '密码';
COMMENT ON COLUMN t_user.email IS '注册使用的邮箱';
COMMENT ON COLUMN t_user.name IS '用户名';
COMMENT ON COLUMN t_user.sex IS '性别';
COMMENT ON COLUMN t_user.dorm IS '宿舍楼栋';
COMMENT ON COLUMN t_user.sign IS '个性签名';
COMMENT ON COLUMN t_user.major IS '专业';
COMMENT ON COLUMN t_user.headimg IS '头像地址';
COMMENT ON COLUMN t_user.phone IS '手机号码';
COMMENT ON COLUMN t_user.qq IS 'qq号码';
COMMENT ON COLUMN t_user.emails IS '用户自己设置的邮箱';
COMMENT ON COLUMN t_user.credits IS '积分';
COMMENT ON COLUMN t_user.leave IS '等级';
COMMENT ON COLUMN t_user.rank IS '排名';
COMMENT ON COLUMN t_user.visit IS '主页访问量';
COMMENT ON COLUMN t_user.lasttime IS '上次登录的时间';
COMMENT ON COLUMN t_user.grade IS '所在年级';
COMMENT ON COLUMN t_user.colleage IS '所在学院';


--############################################################


--t_user_like
CREATE TABLE t_user_like (
    userid1 character varying(50) NOT NULL,
    userid2 character varying(50) NOT NULL,
    CONSTRAINT t_user_like_pkey PRIMARY KEY (userid1, userid2)
);
COMMENT ON COLUMN t_user_like.userid1 IS '收藏者id';
COMMENT ON COLUMN t_user_like.userid2 IS '被收藏者id';


-- ####################################################### view ######################################

CREATE VIEW v_concern AS
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


CREATE VIEW v_goods_detail AS
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

CREATE VIEW v_goodslist AS
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


CREATE VIEW v_hpgoodslist AS
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



CREATE VIEW v_mycollect AS
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


CREATE VIEW v_mydata AS
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


CREATE VIEW v_mygoods AS
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


CREATE VIEW v_mymessage AS
 SELECT u2.id,
    m."time",
    m.content,
    u1.name,
    u1.headimg
   FROM t_message m,
    t_user u1,
    t_user u2
  WHERE (((m.senderid) = (u1.id)) AND ((m.receiverid) = (u2.id)) AND (m.state = 0));


CREATE VIEW v_navingmsg AS
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


CREATE VIEW v_rank AS
 SELECT row_number() OVER (ORDER BY t_user.credits DESC) AS rank,
    t_user.id AS userid,
    t_user.name,
    t_user.credits
   FROM t_user
  ORDER BY t_user.credits DESC
 LIMIT 10;


--######################################## Reference ##############################

--alter table t_collect add constraint collect_fk foreign key (userid) references t_User (id) on delete restrict on update restrict;


ALTER TABLE  t_goods_like ADD CONSTRAINT fk_goodsid FOREIGN KEY (goodsid) REFERENCES t_goods(id);

ALTER TABLE  t_user_like ADD CONSTRAINT fk_user1 FOREIGN KEY (userid1) REFERENCES t_user(id);

ALTER TABLE  t_goods_like ADD CONSTRAINT fk_userid FOREIGN KEY (userid) REFERENCES t_user(id);

ALTER TABLE  t_user_like ADD CONSTRAINT fk_usre2 FOREIGN KEY (userid2) REFERENCES t_user(id);

ALTER TABLE  t_collect ADD CONSTRAINT collect_fk FOREIGN KEY (userid) REFERENCES t_user(id) ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE  t_collect ADD CONSTRAINT collect_fk2 FOREIGN KEY (goodsid) REFERENCES t_goods(id) ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE  t_comment ADD CONSTRAINT comment_fk FOREIGN KEY (userid) REFERENCES t_user(id) ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE  t_comment ADD CONSTRAINT comment_fk2 FOREIGN KEY (goodsid) REFERENCES t_goods(id) ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE  t_concern  ADD CONSTRAINT concern_fk FOREIGN KEY (id1) REFERENCES t_user(id) ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE  t_concern ADD CONSTRAINT concern_fk2 FOREIGN KEY (id2) REFERENCES t_user(id) ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE  t_message ADD CONSTRAINT message_fk FOREIGN KEY (senderid) REFERENCES t_user(id) ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE  t_message ADD CONSTRAINT message_fk2 FOREIGN KEY (receiverid) REFERENCES t_user(id) ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE  t_upload ADD CONSTRAINT upload_fk FOREIGN KEY (userid) REFERENCES t_user(id) ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE  t_upload ADD CONSTRAINT upload_fk2 FOREIGN KEY (goodsid) REFERENCES t_goods(id) ON UPDATE RESTRICT ON DELETE RESTRICT;

