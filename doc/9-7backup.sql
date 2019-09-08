--
-- PostgreSQL database dump
--

-- Dumped from database version 11.1 (Debian 11.1-3.pgdg90+1)
-- Dumped by pg_dump version 11.3

-- Started on 2019-09-07 14:49:59 UTC

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 197 (class 1259 OID 17239)
-- Name: t_collect; Type: TABLE; Schema: public; Owner: blackcardriver
--

CREATE TABLE public.t_collect (
    userid character varying(50) NOT NULL,
    goodsid character varying(50) NOT NULL,
    "time" timestamp without time zone DEFAULT now()
);


ALTER TABLE public.t_collect OWNER TO blackcardriver;

--
-- TOC entry 3013 (class 0 OID 0)
-- Dependencies: 197
-- Name: COLUMN t_collect.userid; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_collect.userid IS '用户id';


--
-- TOC entry 3014 (class 0 OID 0)
-- Dependencies: 197
-- Name: COLUMN t_collect.goodsid; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_collect.goodsid IS '商品id';


--
-- TOC entry 3015 (class 0 OID 0)
-- Dependencies: 197
-- Name: COLUMN t_collect."time"; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_collect."time" IS '收藏时间';


--
-- TOC entry 198 (class 1259 OID 17245)
-- Name: t_comment; Type: TABLE; Schema: public; Owner: blackcardriver
--

CREATE TABLE public.t_comment (
    userid character varying(50) NOT NULL,
    goodsid character varying(50) NOT NULL,
    content character varying(400) DEFAULT ''::character varying,
    "time" timestamp without time zone DEFAULT now()
);


ALTER TABLE public.t_comment OWNER TO blackcardriver;

--
-- TOC entry 3016 (class 0 OID 0)
-- Dependencies: 198
-- Name: COLUMN t_comment.userid; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_comment.userid IS '评论者id';


--
-- TOC entry 3017 (class 0 OID 0)
-- Dependencies: 198
-- Name: COLUMN t_comment.goodsid; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_comment.goodsid IS '商品id';


--
-- TOC entry 3018 (class 0 OID 0)
-- Dependencies: 198
-- Name: COLUMN t_comment.content; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_comment.content IS '评论内容';


--
-- TOC entry 3019 (class 0 OID 0)
-- Dependencies: 198
-- Name: COLUMN t_comment."time"; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_comment."time" IS '评论时间';


--
-- TOC entry 199 (class 1259 OID 17253)
-- Name: t_concern; Type: TABLE; Schema: public; Owner: blackcardriver
--

CREATE TABLE public.t_concern (
    id1 character varying(50) NOT NULL,
    id2 character varying(50) NOT NULL,
    "time" timestamp without time zone DEFAULT now()
);


ALTER TABLE public.t_concern OWNER TO blackcardriver;

--
-- TOC entry 3020 (class 0 OID 0)
-- Dependencies: 199
-- Name: COLUMN t_concern.id1; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_concern.id1 IS '主用户id';


--
-- TOC entry 3021 (class 0 OID 0)
-- Dependencies: 199
-- Name: COLUMN t_concern.id2; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_concern.id2 IS '关注的用户id';


--
-- TOC entry 3022 (class 0 OID 0)
-- Dependencies: 199
-- Name: COLUMN t_concern."time"; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_concern."time" IS '关注时间';


--
-- TOC entry 200 (class 1259 OID 17259)
-- Name: t_goods; Type: TABLE; Schema: public; Owner: blackcardriver
--

CREATE TABLE public.t_goods (
    id character varying(50) NOT NULL,
    name character varying(50) DEFAULT ''::character varying,
    title character varying(50) DEFAULT ''::character varying,
    type character varying(50) DEFAULT ''::character varying,
    tag character varying(50) DEFAULT ''::character varying,
    price double precision DEFAULT 0.0,
    file text DEFAULT ''::text,
    headimg character varying(200) DEFAULT ''::character varying,
    visit integer DEFAULT 0,
    "like" integer DEFAULT 0,
    state integer DEFAULT 1
);


ALTER TABLE public.t_goods OWNER TO blackcardriver;

--
-- TOC entry 3023 (class 0 OID 0)
-- Dependencies: 200
-- Name: COLUMN t_goods.id; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_goods.id IS 'id';


--
-- TOC entry 3024 (class 0 OID 0)
-- Dependencies: 200
-- Name: COLUMN t_goods.name; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_goods.name IS '商品名称';


--
-- TOC entry 3025 (class 0 OID 0)
-- Dependencies: 200
-- Name: COLUMN t_goods.title; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_goods.title IS '标题';


--
-- TOC entry 3026 (class 0 OID 0)
-- Dependencies: 200
-- Name: COLUMN t_goods.type; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_goods.type IS '分类';


--
-- TOC entry 3027 (class 0 OID 0)
-- Dependencies: 200
-- Name: COLUMN t_goods.tag; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_goods.tag IS '标签';


--
-- TOC entry 3028 (class 0 OID 0)
-- Dependencies: 200
-- Name: COLUMN t_goods.price; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_goods.price IS '标价';


--
-- TOC entry 3029 (class 0 OID 0)
-- Dependencies: 200
-- Name: COLUMN t_goods.file; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_goods.file IS '详细描叙';


--
-- TOC entry 3030 (class 0 OID 0)
-- Dependencies: 200
-- Name: COLUMN t_goods.headimg; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_goods.headimg IS '封面图片名';


--
-- TOC entry 3031 (class 0 OID 0)
-- Dependencies: 200
-- Name: COLUMN t_goods.visit; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_goods.visit IS '浏览次数';


--
-- TOC entry 3032 (class 0 OID 0)
-- Dependencies: 200
-- Name: COLUMN t_goods."like"; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_goods."like" IS '点赞次数';


--
-- TOC entry 3033 (class 0 OID 0)
-- Dependencies: 200
-- Name: COLUMN t_goods.state; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_goods.state IS '商品状态';


--
-- TOC entry 213 (class 1259 OID 24746)
-- Name: t_goods_like; Type: TABLE; Schema: public; Owner: blackcardriver
--

CREATE TABLE public.t_goods_like (
    userid character varying(50) NOT NULL,
    goodsid character varying(50) NOT NULL
);


ALTER TABLE public.t_goods_like OWNER TO blackcardriver;

--
-- TOC entry 201 (class 1259 OID 17277)
-- Name: t_message; Type: TABLE; Schema: public; Owner: blackcardriver
--

CREATE TABLE public.t_message (
    senderid character varying(50),
    receiverid character varying(50),
    content character varying(400) DEFAULT ''::character varying,
    "time" timestamp without time zone DEFAULT now(),
    state integer DEFAULT 0
);


ALTER TABLE public.t_message OWNER TO blackcardriver;

--
-- TOC entry 3034 (class 0 OID 0)
-- Dependencies: 201
-- Name: COLUMN t_message.senderid; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_message.senderid IS '发送者id';


--
-- TOC entry 3035 (class 0 OID 0)
-- Dependencies: 201
-- Name: COLUMN t_message.receiverid; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_message.receiverid IS '接收者id';


--
-- TOC entry 3036 (class 0 OID 0)
-- Dependencies: 201
-- Name: COLUMN t_message.content; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_message.content IS '消息内容';


--
-- TOC entry 3037 (class 0 OID 0)
-- Dependencies: 201
-- Name: COLUMN t_message."time"; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_message."time" IS '发送时间';


--
-- TOC entry 3038 (class 0 OID 0)
-- Dependencies: 201
-- Name: COLUMN t_message.state; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_message.state IS '状态';


--
-- TOC entry 202 (class 1259 OID 17286)
-- Name: t_upload; Type: TABLE; Schema: public; Owner: blackcardriver
--

CREATE TABLE public.t_upload (
    userid character varying(50) NOT NULL,
    goodsid character varying(50) NOT NULL,
    "time" timestamp without time zone DEFAULT now()
);


ALTER TABLE public.t_upload OWNER TO blackcardriver;

--
-- TOC entry 3039 (class 0 OID 0)
-- Dependencies: 202
-- Name: COLUMN t_upload.userid; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_upload.userid IS '上传者id';


--
-- TOC entry 3040 (class 0 OID 0)
-- Dependencies: 202
-- Name: COLUMN t_upload.goodsid; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_upload.goodsid IS '商品id';


--
-- TOC entry 3041 (class 0 OID 0)
-- Dependencies: 202
-- Name: COLUMN t_upload."time"; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_upload."time" IS '上传时间';


--
-- TOC entry 196 (class 1259 OID 17214)
-- Name: t_user; Type: TABLE; Schema: public; Owner: blackcardriver
--

CREATE TABLE public.t_user (
    id character varying(50) NOT NULL,
    password character varying(50) DEFAULT '123456'::character varying NOT NULL,
    email character varying(50) DEFAULT NULL::character varying,
    name character varying(50) DEFAULT ''::character varying,
    sex character varying(50) DEFAULT 'boy'::character varying,
    dorm character varying(50) DEFAULT ''::character varying,
    sign character varying(400) DEFAULT ''::character varying,
    major character varying(50) DEFAULT ''::character varying,
    headimg character varying(200) DEFAULT ''::character varying,
    phone character varying(50) DEFAULT ''::character varying,
    qq character varying(50) DEFAULT ''::character varying,
    emails character varying(50) DEFAULT ''::character varying,
    credits integer DEFAULT 0,
    leave integer DEFAULT 1,
    rank integer DEFAULT 99999,
    visit integer DEFAULT 0,
    lasttime timestamp without time zone DEFAULT now(),
    grade integer,
    colleage character varying(50),
    likes integer DEFAULT 0
);


ALTER TABLE public.t_user OWNER TO blackcardriver;

--
-- TOC entry 3042 (class 0 OID 0)
-- Dependencies: 196
-- Name: COLUMN t_user.id; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_user.id IS '账号';


--
-- TOC entry 3043 (class 0 OID 0)
-- Dependencies: 196
-- Name: COLUMN t_user.password; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_user.password IS '密码';


--
-- TOC entry 3044 (class 0 OID 0)
-- Dependencies: 196
-- Name: COLUMN t_user.email; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_user.email IS '注册使用的邮箱';


--
-- TOC entry 3045 (class 0 OID 0)
-- Dependencies: 196
-- Name: COLUMN t_user.name; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_user.name IS '用户名';


--
-- TOC entry 3046 (class 0 OID 0)
-- Dependencies: 196
-- Name: COLUMN t_user.sex; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_user.sex IS '性别';


--
-- TOC entry 3047 (class 0 OID 0)
-- Dependencies: 196
-- Name: COLUMN t_user.dorm; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_user.dorm IS '宿舍楼栋';


--
-- TOC entry 3048 (class 0 OID 0)
-- Dependencies: 196
-- Name: COLUMN t_user.sign; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_user.sign IS '个性签名';


--
-- TOC entry 3049 (class 0 OID 0)
-- Dependencies: 196
-- Name: COLUMN t_user.major; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_user.major IS '专业';


--
-- TOC entry 3050 (class 0 OID 0)
-- Dependencies: 196
-- Name: COLUMN t_user.headimg; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_user.headimg IS '头像地址';


--
-- TOC entry 3051 (class 0 OID 0)
-- Dependencies: 196
-- Name: COLUMN t_user.phone; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_user.phone IS '手机号码';


--
-- TOC entry 3052 (class 0 OID 0)
-- Dependencies: 196
-- Name: COLUMN t_user.qq; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_user.qq IS 'qq号码';


--
-- TOC entry 3053 (class 0 OID 0)
-- Dependencies: 196
-- Name: COLUMN t_user.emails; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_user.emails IS '用户自己设置的邮箱';


--
-- TOC entry 3054 (class 0 OID 0)
-- Dependencies: 196
-- Name: COLUMN t_user.credits; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_user.credits IS '积分';


--
-- TOC entry 3055 (class 0 OID 0)
-- Dependencies: 196
-- Name: COLUMN t_user.leave; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_user.leave IS '等级';


--
-- TOC entry 3056 (class 0 OID 0)
-- Dependencies: 196
-- Name: COLUMN t_user.rank; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_user.rank IS '排名';


--
-- TOC entry 3057 (class 0 OID 0)
-- Dependencies: 196
-- Name: COLUMN t_user.visit; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_user.visit IS '主页访问量';


--
-- TOC entry 3058 (class 0 OID 0)
-- Dependencies: 196
-- Name: COLUMN t_user.lasttime; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_user.lasttime IS '上次登录的时间';


--
-- TOC entry 3059 (class 0 OID 0)
-- Dependencies: 196
-- Name: COLUMN t_user.grade; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_user.grade IS '所在年级';


--
-- TOC entry 3060 (class 0 OID 0)
-- Dependencies: 196
-- Name: COLUMN t_user.colleage; Type: COMMENT; Schema: public; Owner: blackcardriver
--

COMMENT ON COLUMN public.t_user.colleage IS '所在学院';


--
-- TOC entry 212 (class 1259 OID 24731)
-- Name: t_user_like; Type: TABLE; Schema: public; Owner: blackcardriver
--

CREATE TABLE public.t_user_like (
    userid1 character varying(50) NOT NULL,
    userid2 character varying(50) NOT NULL
);


ALTER TABLE public.t_user_like OWNER TO blackcardriver;

--
-- TOC entry 214 (class 1259 OID 24770)
-- Name: v_concern; Type: VIEW; Schema: public; Owner: blackcardriver
--

CREATE VIEW public.v_concern AS
 SELECT u2.id AS id2,
    u2.name AS name2,
    u2.headimg AS headimg2,
    u1.id,
    u1.name,
    u1.headimg
   FROM public.t_user u1,
    public.t_user u2,
    public.t_concern c
  WHERE (((c.id1)::text = (u1.id)::text) AND ((c.id2)::text = (u2.id)::text));


ALTER TABLE public.v_concern OWNER TO blackcardriver;

--
-- TOC entry 204 (class 1259 OID 17395)
-- Name: v_goods_detail; Type: VIEW; Schema: public; Owner: blackcardriver
--

CREATE VIEW public.v_goods_detail AS
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
   FROM public.t_goods g,
    public.t_upload p,
    public.t_user u
  WHERE (((p.goodsid)::text = (g.id)::text) AND ((p.userid)::text = (u.id)::text) AND (g.state = 1));


ALTER TABLE public.v_goods_detail OWNER TO blackcardriver;

--
-- TOC entry 203 (class 1259 OID 17391)
-- Name: v_goodslist; Type: VIEW; Schema: public; Owner: blackcardriver
--

CREATE VIEW public.v_goodslist AS
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
   FROM public.t_upload p,
    public.t_goods g,
    public.t_user u
  WHERE (((p.userid)::text = (u.id)::text) AND ((p.goodsid)::text = (g.id)::text));


ALTER TABLE public.v_goodslist OWNER TO blackcardriver;

--
-- TOC entry 211 (class 1259 OID 17456)
-- Name: v_hpgoodslist; Type: VIEW; Schema: public; Owner: blackcardriver
--

CREATE VIEW public.v_hpgoodslist AS
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
   FROM public.v_goodslist
  WHERE (v_goodslist.state = 1);


ALTER TABLE public.v_hpgoodslist OWNER TO blackcardriver;

--
-- TOC entry 206 (class 1259 OID 17415)
-- Name: v_mycollect; Type: VIEW; Schema: public; Owner: blackcardriver
--

CREATE VIEW public.v_mycollect AS
 SELECT u.id AS uid,
    g.id,
    g.name,
    g.title,
    g.price,
    g.headimg
   FROM public.t_collect c,
    public.t_user u,
    public.t_goods g
  WHERE (((c.userid)::text = (u.id)::text) AND ((c.goodsid)::text = (g.id)::text));


ALTER TABLE public.v_mycollect OWNER TO blackcardriver;

--
-- TOC entry 208 (class 1259 OID 17435)
-- Name: v_mydata; Type: VIEW; Schema: public; Owner: blackcardriver
--

CREATE VIEW public.v_mydata AS
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
   FROM public.t_user;


ALTER TABLE public.v_mydata OWNER TO blackcardriver;

--
-- TOC entry 207 (class 1259 OID 17431)
-- Name: v_mygoods; Type: VIEW; Schema: public; Owner: blackcardriver
--

CREATE VIEW public.v_mygoods AS
 SELECT u.id AS uid,
    g.id,
    g.name,
    g.title,
    p."time",
    g.price,
    g.headimg
   FROM public.t_user u,
    public.t_goods g,
    public.t_upload p
  WHERE (((p.userid)::text = (u.id)::text) AND ((p.goodsid)::text = (g.id)::text) AND (g.state = 1));


ALTER TABLE public.v_mygoods OWNER TO blackcardriver;

--
-- TOC entry 205 (class 1259 OID 17411)
-- Name: v_mymessage; Type: VIEW; Schema: public; Owner: blackcardriver
--

CREATE VIEW public.v_mymessage AS
 SELECT u2.id,
    m."time",
    m.content,
    u1.name,
    u1.headimg
   FROM public.t_message m,
    public.t_user u1,
    public.t_user u2
  WHERE (((m.senderid)::text = (u1.id)::text) AND ((m.receiverid)::text = (u2.id)::text) AND (m.state = 0));


ALTER TABLE public.v_mymessage OWNER TO blackcardriver;

--
-- TOC entry 210 (class 1259 OID 17452)
-- Name: v_navingmsg; Type: VIEW; Schema: public; Owner: blackcardriver
--

CREATE VIEW public.v_navingmsg AS
 SELECT t_user.id,
    t_user.name,
    t_user.headimg,
    t_user.credits,
    t_user.leave,
    t_user.lasttime,
    ( SELECT count(*) AS count
           FROM public.t_upload
          WHERE ((t_upload.userid)::text = (t_user.id)::text)) AS goodsnum,
    ( SELECT count(*) AS count
           FROM public.t_message
          WHERE ((t_message.receiverid)::text = (t_user.id)::text)) AS messagenum
   FROM public.t_user;


ALTER TABLE public.v_navingmsg OWNER TO blackcardriver;

--
-- TOC entry 209 (class 1259 OID 17448)
-- Name: v_rank; Type: VIEW; Schema: public; Owner: blackcardriver
--

CREATE VIEW public.v_rank AS
 SELECT row_number() OVER (ORDER BY t_user.credits DESC) AS rank,
    t_user.id AS userid,
    t_user.name,
    t_user.credits
   FROM public.t_user
  ORDER BY t_user.credits DESC
 LIMIT 10;


ALTER TABLE public.v_rank OWNER TO blackcardriver;

--
-- TOC entry 2852 (class 2606 OID 17244)
-- Name: t_collect pk_t_collect; Type: CONSTRAINT; Schema: public; Owner: blackcardriver
--

ALTER TABLE ONLY public.t_collect
    ADD CONSTRAINT pk_t_collect PRIMARY KEY (userid, goodsid);


--
-- TOC entry 2854 (class 2606 OID 17258)
-- Name: t_concern pk_t_concern; Type: CONSTRAINT; Schema: public; Owner: blackcardriver
--

ALTER TABLE ONLY public.t_concern
    ADD CONSTRAINT pk_t_concern PRIMARY KEY (id1, id2);


--
-- TOC entry 2856 (class 2606 OID 17276)
-- Name: t_goods pk_t_goods; Type: CONSTRAINT; Schema: public; Owner: blackcardriver
--

ALTER TABLE ONLY public.t_goods
    ADD CONSTRAINT pk_t_goods PRIMARY KEY (id);


--
-- TOC entry 2858 (class 2606 OID 17291)
-- Name: t_upload pk_t_upload; Type: CONSTRAINT; Schema: public; Owner: blackcardriver
--

ALTER TABLE ONLY public.t_upload
    ADD CONSTRAINT pk_t_upload PRIMARY KEY (userid, goodsid);


--
-- TOC entry 2850 (class 2606 OID 17238)
-- Name: t_user pk_t_user; Type: CONSTRAINT; Schema: public; Owner: blackcardriver
--

ALTER TABLE ONLY public.t_user
    ADD CONSTRAINT pk_t_user PRIMARY KEY (id);


--
-- TOC entry 2862 (class 2606 OID 24750)
-- Name: t_goods_like t_goods_like_pkey; Type: CONSTRAINT; Schema: public; Owner: blackcardriver
--

ALTER TABLE ONLY public.t_goods_like
    ADD CONSTRAINT t_goods_like_pkey PRIMARY KEY (userid, goodsid);


--
-- TOC entry 2860 (class 2606 OID 24735)
-- Name: t_user_like t_user_like_pkey; Type: CONSTRAINT; Schema: public; Owner: blackcardriver
--

ALTER TABLE ONLY public.t_user_like
    ADD CONSTRAINT t_user_like_pkey PRIMARY KEY (userid1, userid2);


--
-- TOC entry 2863 (class 2606 OID 17292)
-- Name: t_collect collect_fk; Type: FK CONSTRAINT; Schema: public; Owner: blackcardriver
--

ALTER TABLE ONLY public.t_collect
    ADD CONSTRAINT collect_fk FOREIGN KEY (userid) REFERENCES public.t_user(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- TOC entry 2864 (class 2606 OID 17297)
-- Name: t_collect collect_fk2; Type: FK CONSTRAINT; Schema: public; Owner: blackcardriver
--

ALTER TABLE ONLY public.t_collect
    ADD CONSTRAINT collect_fk2 FOREIGN KEY (goodsid) REFERENCES public.t_goods(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- TOC entry 2865 (class 2606 OID 17302)
-- Name: t_comment comment_fk; Type: FK CONSTRAINT; Schema: public; Owner: blackcardriver
--

ALTER TABLE ONLY public.t_comment
    ADD CONSTRAINT comment_fk FOREIGN KEY (userid) REFERENCES public.t_user(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- TOC entry 2866 (class 2606 OID 17307)
-- Name: t_comment comment_fk2; Type: FK CONSTRAINT; Schema: public; Owner: blackcardriver
--

ALTER TABLE ONLY public.t_comment
    ADD CONSTRAINT comment_fk2 FOREIGN KEY (goodsid) REFERENCES public.t_goods(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- TOC entry 2867 (class 2606 OID 17312)
-- Name: t_concern concern_fk; Type: FK CONSTRAINT; Schema: public; Owner: blackcardriver
--

ALTER TABLE ONLY public.t_concern
    ADD CONSTRAINT concern_fk FOREIGN KEY (id1) REFERENCES public.t_user(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- TOC entry 2868 (class 2606 OID 17317)
-- Name: t_concern concern_fk2; Type: FK CONSTRAINT; Schema: public; Owner: blackcardriver
--

ALTER TABLE ONLY public.t_concern
    ADD CONSTRAINT concern_fk2 FOREIGN KEY (id2) REFERENCES public.t_user(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- TOC entry 2876 (class 2606 OID 24756)
-- Name: t_goods_like fk_goodsid; Type: FK CONSTRAINT; Schema: public; Owner: blackcardriver
--

ALTER TABLE ONLY public.t_goods_like
    ADD CONSTRAINT fk_goodsid FOREIGN KEY (goodsid) REFERENCES public.t_goods(id);


--
-- TOC entry 2873 (class 2606 OID 24736)
-- Name: t_user_like fk_user1; Type: FK CONSTRAINT; Schema: public; Owner: blackcardriver
--

ALTER TABLE ONLY public.t_user_like
    ADD CONSTRAINT fk_user1 FOREIGN KEY (userid1) REFERENCES public.t_user(id);


--
-- TOC entry 2875 (class 2606 OID 24751)
-- Name: t_goods_like fk_userid; Type: FK CONSTRAINT; Schema: public; Owner: blackcardriver
--

ALTER TABLE ONLY public.t_goods_like
    ADD CONSTRAINT fk_userid FOREIGN KEY (userid) REFERENCES public.t_user(id);


--
-- TOC entry 2874 (class 2606 OID 24741)
-- Name: t_user_like fk_usre2; Type: FK CONSTRAINT; Schema: public; Owner: blackcardriver
--

ALTER TABLE ONLY public.t_user_like
    ADD CONSTRAINT fk_usre2 FOREIGN KEY (userid2) REFERENCES public.t_user(id);


--
-- TOC entry 2869 (class 2606 OID 17322)
-- Name: t_message message_fk; Type: FK CONSTRAINT; Schema: public; Owner: blackcardriver
--

ALTER TABLE ONLY public.t_message
    ADD CONSTRAINT message_fk FOREIGN KEY (senderid) REFERENCES public.t_user(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- TOC entry 2870 (class 2606 OID 17327)
-- Name: t_message message_fk2; Type: FK CONSTRAINT; Schema: public; Owner: blackcardriver
--

ALTER TABLE ONLY public.t_message
    ADD CONSTRAINT message_fk2 FOREIGN KEY (receiverid) REFERENCES public.t_user(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- TOC entry 2871 (class 2606 OID 17332)
-- Name: t_upload upload_fk; Type: FK CONSTRAINT; Schema: public; Owner: blackcardriver
--

ALTER TABLE ONLY public.t_upload
    ADD CONSTRAINT upload_fk FOREIGN KEY (userid) REFERENCES public.t_user(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- TOC entry 2872 (class 2606 OID 17337)
-- Name: t_upload upload_fk2; Type: FK CONSTRAINT; Schema: public; Owner: blackcardriver
--

ALTER TABLE ONLY public.t_upload
    ADD CONSTRAINT upload_fk2 FOREIGN KEY (goodsid) REFERENCES public.t_goods(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


-- Completed on 2019-09-07 14:50:05 UTC

--
-- PostgreSQL database dump complete
--

