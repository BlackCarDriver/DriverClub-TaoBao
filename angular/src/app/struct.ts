

//#######################    homepage    ##########################################

//商品显示的基本信息，见首页封面
export class HomePageGoods{
    headimg:string;
    userid:string;
    time:string;
    title:string;
    price:number;
    id:string;
    name:string;
}

//显示在首页的商品主分类
export class GoodsType{
    type :string;
    list:GoodSubType[];
}
//显示在首页的商品子分类，分类名和商品数量
export class GoodSubType{
    tag:string;
    number: number;
}



//#########################  goodsdetail  ######################################################

//商品详情页需要的数据
export class GoodsDetail{
    headimg:string;
    userid:string;     //到时需要变成用户名
    time:string;
    title:string;
    price:number;
    id:string;
    name:string;
    visit:number; 
    like:number;
    collect:number;     //precial
    talk:number;        //precial
}

//#########################  personal  ######################################################

//个人主页里面需要展示的详细信息
export class UserMessage{
    //基本信息和联系信息
    headimg:string; 
    name:string; 
    id:string;
    sex:string;
    sign:string; 
    grade:string;
    colleage:string;
    major:string;
    emails:string;
    qq:string ;
    phone:string; 
    //首部数据
    leave:number;  
    credits:number;   //积分
    rank:number ;    
    becare:number ;    //关注我的人数       (special)
    like:number ;      //被点赞次数         (special)
    //其他数据
    lasttime:number;  //上次登录的时间间隔（小时）
    visit:number;    //主页访问次数
    goodsnum:number;  //拥有商品的数量      (special)
    scuess:number;    //成功交易的商品数量
    care:number ;    //我关注的人数         (special)
}

//个人主页里我的商品，以及我的收藏 数据结构
export class GoodsShort{
    id:string;
    headimg:string; 
    name:string
    title:string;
    price:number;
}

//个人主页里面的我的消息
export class MyMessage{
    senderid:string;
    sendername:string;      //(precial)
    content:string;
    time:string;
}

//我关注的和关注我的用户列表
export class User{
    id2:string;
    name:string;    
    headimg:string; 
}

//#########################  修改个人信息页面 ######################################################

//个人信息设置页上传的信息数据
export class PersonalBase{
    username:string;
    userid:string ;
    usersex:string ;
    sign:string ;
    grade:string;
    colleage:string ;
    email:string ;
    qq:string ;
    phone:string;
}




//用户排名数据元素
export class Rank{
    rank:number;
    name:string;
    userid:string;
}




//首页-个人信息下拉框
export class UserShort{
    imgurl :string ;
    grade : number ;
    score : number ;
    message:number ;
    goods:number ;
    lastime:string ;
}


//长传商品页面上传的信息
export class UploadGoods{
    username: string;
    title:string;
    date :string ;
    price:number;
    imgurl:string;
    type:string;
    tag:string;
    usenewtag:boolean = false;
    newtagname:string ;
    text:string;
}


//注册账号需要用到的用户信息
export class account1{
    name:string;
    password:string;
    email:string;
    code:string;
}
//登录账号信息
export class account2{
    name:string;
    password:string;
}