import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import {  HomePageGoods , GoodsType,UploadAnyResult,UploadIImgResult,UploadGoods,UserMessage } from '../app/struct';
import {  account1, account2, MyStatus} from '../app/struct';
import { PersonalBase} from '../app/struct';
@Injectable({
  providedIn: 'root'
})

export class ServerService {

  //important config !!!
//本地开发配置
 private addr: string  = "http://localhost:4747"
 //服务器配置
  // private addr: string  = "https://www.blackcardriver.cn/server"
  constructor( 
    private http: HttpClient
  ){ }

 //=====================  重做  =====================================================================
//获取主页商品列表
 GetHomePageGoods(tag : string, index : number){
  var url = this.addr + "/homepage/goodsdata";
  var postdata = {goodstag:tag, goodsindex:index};
  return this.http.post<HomePageGoods[]>(url, JSON.stringify(postdata));
}
//主页商品类型和标签列表数据
GetHomePageType(){
  var url = this.addr + "/homepage/goodstypemsg";
  return this.http.get<GoodsType[]>(url);
}

//商品详情页面获取数据接口
GetGoodsDeta(id:number, type:string){
    var url = this.addr+"/goodsdeta";
    var data = {goodid:id, datatype:type}
    return this.http.post<any>(url,JSON.stringify(data));
}

//个人主页里得到各种信息的数据接口
GetMyMsg(username:string, tag:string){
    var url = this.addr + "/personal/data"; 
    var data = {tag:tag, name:username};
    return this.http.post<any>(url,data); 
}

//导航栏得到用户的数据
GetNavigUser(username:string){
  var url = this.addr + "/personal/data";
  var postdata = {name:username, tag:"naving"};
  return this.http.post<MyStatus>(url, JSON.stringify(postdata), {withCredentials: true});
}

//上传商品
  UploadGoodsData(data:UploadGoods){
    var url = this.addr + "/upload/newgoods"; 
    return this.http.post<UploadAnyResult>(url,data);
  }

  //上传图片到服务器得到一个访问这个图片的的url
  UploadImg(username:string , img:any){
    var postdata = new FormData();
    postdata.append("name", username);
    postdata.append("file",img)
    var url = this.addr + "/upload/images"; 
    return this.http.post<UploadIImgResult>(url,postdata);
  }     

// ==========================  the following function is related to cookie ==================================  

//use to make the cookie cant be undestant directly
encryption(code : string){
  var c=String.fromCharCode(code.charCodeAt(0)+code.length);
 for(var i=1;i<code.length;i++){      
   c+=String.fromCharCode(code.charCodeAt(i)+code.charCodeAt(i-1));
 }   
 return escape(c);
}
//restore the string that after encryption
decode(code : string ){
  code=unescape(code);      
 var c=String.fromCharCode(code.charCodeAt(0)-code.length);      
 for(var i=1;i<code.length;i++){      
  c+=String.fromCharCode(code.charCodeAt(i)-c.charCodeAt(i-1));      
 }      
 return c;  
}
//take username from cookie
Getusername(){
  var name = this.getCookie("driverlei")
  if (name=="")return "";
  name = this.decode(name);
  return name;
}
//get cookie by cookie name
getCookie(name:string){ 
  var strCookie=document.cookie; 
  var arrCookie=strCookie.split("; "); 
  for(var i=0;i<arrCookie.length;i++){ 
    var arr=arrCookie[i].split("="); 
    if(arr[0]==name)return arr[1]; 
  }
  return ""; 
}
//save a cookie as a time tag
setTimeTag(key :string, second:number){
  var exp = new Date();
  exp.setTime(exp.getTime() + 1000 * second );  //two minute  
  var ck = key+"=have;expires=";
  document.cookie = ck + exp.toUTCString();
}
//check if the tag is still in cookie, return false if cookie out of time
checkTimeTag(key:string){
  var ck = this.getCookie(key);
  //if (ck=="") alert("time tag out of time ");
  if (ck=="") return false;
  return true;
} 
//save an object in localstroge by json format
setLocalStorge(key:string, data :any){
  var jsdata = JSON.stringify(data);
  window.localStorage[key] = jsdata;
}
//take object after json.parse from localstroge by name
getLocalStorge(key:string){
  var jsdata = window.localStorage[key];
  return JSON.parse(jsdata);
}

// =======================================================================================================  
  
ChangeComfirmCode(na :string){
  var data = {name : na};
  var url = this.addr + "/getmsg/usershort/cgcfcode";
  return this.http.post<number>(url, data,{withCredentials: true});
}


  //upload goods message to server
  UploadGoods(goods:string){
    var postdata = {goodsdata:goods};
    var url = this.addr+"/upload/goods";
    return this.http.post<string>(
      url,postdata);
  }


  //get usermsg in chgmymsg page
  Getmymsg(id:string){
    var url = this.addr+"/getmsg/usermsg?id="+id;
    return this.http.get<PersonalBase>(url);
  }
  //upload and updata base message of user
  UploadMyBaseMsg(data:PersonalBase){
    var url = this.addr + "/updata/mymessage/basemsg"; 
    return this.http.post<number>(url,data);
  }
  //upload and updata base message of user
  UploadContactMsg(data:PersonalBase){
      var url = this.addr + "/updata/mymessage/contactmsg"; 
      return this.http.post<number>(url,data);
  }

  //get message of personal2 page
  GetOtherMsg(userid:string){
    var url = this.addr+"/getmsg/othermsg?id="+userid;
    return this.http.get<UserMessage>(url);
  }
  //login function used in naving
  Login(data:account2){
      var url = this.addr+"/signin";
      return this.http.post<number>(url,data,{withCredentials: true});
  }
  //send base message to server to conirm 
  ConfirmMsg(data:account1){
    var url = this.addr + "/register/confirmmsg";
    return this.http.post<number>(url,data);
  }
  //send confirm code to the server and receive the state
  ConfirmCode(data :account1){
    var url = this.addr + "/regeister/confirmcode";
    return this.http.post<number>(url,data);
  }


}
