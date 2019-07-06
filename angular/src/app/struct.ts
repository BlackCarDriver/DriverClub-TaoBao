//商品显示的基本信息，见首页封面
export class HomePageGoods{
    imgurl:string="https://i03piccdn.sogoucdn.com/5b84e0a94f71278f";
    owner:string="找不到数据...";
    date:string ="2019-1-1";
    title:string ="找不到数据...";
    price:number = 0;
    goodsid:number=0;
    gname:string="找不到数据...";
}
//显示在首页的商品子分类，分类名和商品数量
export class GoodSubType{
    name:string="";
    num: number=0;
}
//显示在首页的商品主分类
export class GoodsType{
    typename :string = "";
    list:GoodSubType[];
}
//首页-个人信息下拉框
export class UserShort{
    imgurl :string = "";
    grade : number = 0;
    score : number = 0;
    message:number = 0;
    goods:number = 0;
    lastime:string ="";
}

//商品状况，包括观看数，收藏数等
export class GoodsState{
    see:number =0;
    like:number =0;
    collect:number =0;
    talk:number=0;
}
//长传商品页面上传的信息
export class UploadGoods{
    username: string="";
    title:string="";
    date :string = "";
    price:number=0;
    imgurl:string="";
    type:string="";
    tag:string="";
    usenewtag:boolean = false;
    newtagname:string = "";
    text:string="";
}
//个人信息设置页上传的信息数据
export class PersonalBase{
    username:string= "";
    userid:string = "";
    usersex:string = "";
    sign:string = "";
    grade:string="";
    colleage:string ="";
    email:string = "";
    qq:string ="";
    phone:string="";
}
//个人主页里面需要展示的详细信息
export class PersonalExpend{
    //基本信息和联系信息
    imgurl:string="";
    username:string= "";
    userid:string = "";
    usersex:string = "";
    sign:string = "";
    grade:string="";
    colleage:string ="";
    email:string = "";
    qq:string ="";
    phone:string="";
    //首部数据
    leave:number = 0;  
    sorce:number = 0;   //积分
    rank:number =0;    
    becare:number =0;    //关注我的人数
    like:number =0;      //被点赞次数
    //其他数据
    lasttime:number=0;  //上次登录的时间间隔（小时）
    goodsnum:number=0;  //拥有商品的数量
    scuess:number=0;    //成功交易的商品数量
    visnum:number=0;    //主页访问次数
    care:number = 0;    //我关注的人数
}
//个人主页里我的商品和我的收藏，商品简要显示数据
export class Goods2{
    imgurl:string="https://img.alicdn.com/bao/uploaded/i1/TB1WlfKOFXXXXbJapXXXXXXXXXX_!!0-item_pic.jpg_160x160xz.jpg";
    goodsid:string="000";
    describ:string="鸿星尔克男鞋板鞋休";
    price:number=0;
}
//个人主页里面的我的消息
export class MyMessage{
    messageid:string;
    senderimgurl:string;
    sendername:string;
    title:string;
    time:string;
    describ:string;
}
//用户排名数据元素
export class Rank{
    rank:number;
    name:string;
    userid:string;
}
//我收藏的和关注我的用户数据元素
export class User{
    userid:string;
    username:string;
    imgurl:string;
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