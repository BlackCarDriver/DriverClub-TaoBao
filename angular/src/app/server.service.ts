import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import {  HomePageGoods , GoodsType,UploadAnyResult,UploadIImgResult,UploadGoods } from '../app/struct';
import {  RequertResult, MyStatus,RequestProto,ReplyProto} from '../app/struct';
@Injectable({
  providedIn: 'root'
})

export class ServerService {

  //important config !!!
  //线上运行配置
  // private addr: string  = "/taobaoserver"
  //本地开发配置
  private addr: string  = "/localserver"
 //服务器配置
  // private addr: string  = "https://www.blackcardriver.cn/server"
  constructor( 
    private http: HttpClient
  ){ }

 //======================================= large  interface =============================================================

 //get all kind of data in goodspage 🍌
GetGoodsDeta(request : RequestProto){
  var url = this.addr+"/goodsdeta";
  return this.http.post<ReplyProto>(url,JSON.stringify(request));
}

//request to update some simple record such as collect number 🍍
SmallUpdate(request : RequestProto){
  var url = this.addr + "/smallupdate"; 
  return this.http.post<ReplyProto>(url, JSON.stringify(request));
}

//request to update some complex message such as profile 🍍
UpdateMessage(request : RequestProto){
  var url = this.addr + "/update"; 
  return this.http.post<ReplyProto>(url, JSON.stringify(request)); 
}

//upload a images to server and receive a url to get it images 🍍
UploadImg(username:string , img:any){
  var postdata = new FormData();
  postdata.append("name", username);
  postdata.append("file",img)
  var url = this.addr + "/upload/images"; 
  //post a multipart/form-data, can not use json.stringfiy
  return this.http.post<ReplyProto>(url, postdata);
} 

//get information in personal page 🍍
GetMyMsg(request : RequestProto){
  var url = this.addr + "/personal/data"; 
  var data = {tag:tag, name:username};
  return this.http.post<ReplyProto>(url,data); 
}


 //=======================================  重做  =====================================================================
//获取主页商品列表
GetHomePageGoods(type:string, tag : string, index : number){
  var url = this.addr + "/homepage/goodsdata";
  var postdata = {goodstype: type,goodstag:tag, goodsindex:index};
  return this.http.post<HomePageGoods[]>(url, JSON.stringify(postdata));
}

//主页商品类型和标签列表数据
GetHomePageType(){
  var url = this.addr + "/homepage/goodstypemsg";
  return this.http.get<GoodsType[]>(url);
}


//导航栏得到用户的数据
GetNavigUser(userid:string){
  var url = this.addr + "/personal/data";
  var postdata = {name:userid, tag:"naving"};
  return this.http.post<MyStatus>(url, JSON.stringify(postdata), {withCredentials: true});
}

//上传商品
UploadGoodsData(data:UploadGoods){
    var url = this.addr + "/upload/newgoods"; 
    return this.http.post<UploadAnyResult>(url,data);
}



// ================================== the following function reference to login or register ========================================================  
  
Entrance(userid:string, tag:string, data:any){
  var url = this.addr + "/entrance";
  var postdata = {userid:userid, tag:tag, data:data};
  return this.http.post<RequertResult>(url,JSON.stringify(postdata));
}

  // ==========================  the following function is related to cookie ==================================  

//use to make the cookie can't be undestant directly
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

//check if the tag is still in cookie, return false if cookie out of time
checkTimeTag(key:string){
  var ck = this.getCookie(key);
  //if (ck=="") alert("time tag out of time ");
  if (ck=="") return false;
  return true;
} 

//take object after json.parse from localstroge by name
getLocalStorge(key:string){
  var jsdata = window.localStorage[key];
  return JSON.parse(jsdata);
}


/*

//save an object in localstroge by json format
setLocalStorge(key:string, data :any){
  var jsdata = JSON.stringify(data);
  window.localStorage[key] = jsdata;
}

//save a cookie as a time tag
setTimeTag(key :string, second:number){
  var exp = new Date();
  exp.setTime(exp.getTime() + 1000 * second );  //two minute  
  var ck = key+"=have;expires=";
  document.cookie = ck + exp.toUTCString();
}

//take username from cookie
Getusername(){
  var name = this.getCookie("driverlei")
  if (name=="")return "";
  name = this.decode(name);
  return name;
}

*/

}
