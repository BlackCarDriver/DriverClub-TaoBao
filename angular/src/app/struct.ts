

//#######################    homepage    ##########################################

//商品显示的基本信息，见首页封面
export class HomePageGoods{
    headimg:string;
    userid:string;
    username:string;
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
export class GoodsDetail {
    headimg:string;
    userid:string;    
    username:string;
    time:string;
    title:string;
    price:number;
    id:string;
    name:string;
    visit:number; 
    like:number;
    collect:number;     //precial
    talk:number;        //precial
    detail:string;
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
    dorm:string
    //首部数据
    leave:number;  
    credits:number;   //积分
    rank:number ;    
    becare:number ;    //关注我的人数       (special)
    likes:number ;      //被点赞次数         (special)
    //其他数据
    lasttime:string;  
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
    title:string;
    time:string;
    name:string;      //(precial)
    headimg:string;
    content:string;
}

//我关注的和关注我的用户列表
export class User{
    id:string;
    name:string;    
    headimg:string; 
}

//用户排名数据元素
export class Rank{
    rank:number;
    name:string;
    userid:string;
}

//#########################  uploadgoods ######################################################


//长传商品页面上传的信息
export class UploadGoods{
    userid: string;
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

//上传图片返回的结果
export class UploadIImgResult{
    status:number;
    describe:string;
    imgurl:string;
}

//上传数据返回结果
export class UploadAnyResult {
    status:number;
    describe:string;
}

//#########################  chgmymsg ######################################################


//个人信息设置页上传的信息数据
export class PersonalSetting{
    headimg:string;
    name:string;
    id:string ;
    sex:string ;
    sign:string ;
    grade:string;
    colleage:string ;
    dorm:string;
    emails:string ;
    qq:string ;
    phone:string;
}
//上传数据返回结果
export class UpdateResult {
    status:number;
    describe:string;
}

//#########################  naving ######################################################



//首页-个人信息下拉框
export class MyStatus{
    headimg :string ;
    leave : number ;
    credits : number ;
    messagenum:number ;    //precial
    goodsnum:number ;      //pricial
    lasttime:string ;    
}

//注册账号发送数据
export class RegisterData {
    name:string;
    password:string;
    email:string;
    code:string;
}

//登录账号发送数据
export class LoginData{
    name:string;
    password:string;
}

export class RequertResult {
    status:number;
    describe:string;
}