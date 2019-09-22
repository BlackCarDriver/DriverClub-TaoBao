import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import {  GoodsType,UploadGoods } from '../app/struct';
import {  RequertResult,RequestProto,ReplyProto} from '../app/struct';

@Injectable({
  providedIn: 'root'
})

export class ServerService {

  //global variable 🍈
  userid = "";    //this usrid only can be true usre id
  username = "";
  homepage_goods_perpage = 10;
  // private addr: string  = "https://blackcardriver.cn/taobaoserver"
  private addr: string  = "/localserver"

  constructor( 
    private http: HttpClient,
  ){ }
 
 //====================================== public phsical function =================================
//check whether the user is login, show the warm message if not 🍈
 IsNotLogin(){
  if(this.userid == ""){
    alert("你好，要先登录呦!");
    return true;
  }
  return false;
 }

 //get last section of persent url
 LastSection(){
  let rawStr = window.location.pathname;
  let lastSlash = rawStr.lastIndexOf("/");
  let result = rawStr.substring(lastSlash+1);
  return result;
 }

 //scoll to top
 totop(){
    scroll(0,150);
 }
 //======================================= large  interface =============================================================

 //get all kind of data in goodspage 🍌🔥
GetGoodsDeta(request : RequestProto){
  var url = this.addr+"/goodsdeta";
  return this.http.post<ReplyProto>(url,JSON.stringify(request));
}

//request to update some simple record such as collect number 🍍🔥
SmallUpdate(request : RequestProto){
  var url = this.addr + "/smallupdate"; 
  return this.http.post<ReplyProto>(url, JSON.stringify(request));
}

//request to update some complex message such as profile 🍍🔥
UpdateMessage(request : RequestProto){
  var url = this.addr + "/update"; 
  return this.http.post<ReplyProto>(url, JSON.stringify(request)); 
}

//upload a images to server and receive a url to get it images 🍍🔥
UploadImg(username:string , img:any){
  var postdata = new FormData();
  postdata.append("name", username);
  postdata.append("file",img)
  var url = this.addr + "/upload/images"; 
  //post a multipart/form-data, can not use json.stringfiy
  return this.http.post<ReplyProto>(url, postdata);
} 

//get information in personal page 🍍🔥
GetMyMsg(request : RequestProto){
  var url = this.addr + "/personal/data"; 
  return this.http.post<ReplyProto>(url, JSON.stringify(request)); 
}

//a little different from GetMyMsg 🍋🔥
GetCredentMsg(request : RequestProto){
  var url = this.addr + "/personal/data";
  return this.http.post<ReplyProto>(url, JSON.stringify(request), {withCredentials: true});
}

//delete something  🍑
DeleteMyData(request:RequestProto){
  var url = this.addr + "/deleteapi";
  return this.http.post<ReplyProto>(url, JSON.stringify(request));
}

//get homepage goods list data 🍋🔥🍇
GetHomePageGoods(type:string, tag : string, page : number){
  let postdata : RequestProto = {
    api:"gethomepagegoods",
    offset: this.homepage_goods_perpage*(page-1),
    limit:this.homepage_goods_perpage,
    data:{goodstype: type, goodstag:tag},
  };
  var url = this.addr + "/homepage/goodsdata";
  return this.http.post<ReplyProto>(url, JSON.stringify(postdata));
}

//upload a good data 🍋
UploadGoodsData(data:UploadGoods){
  var url = this.addr + "/upload/newgoods"; 
  let postdata : RequestProto = {
    api:"uploadgoodsdata",
    data:JSON.stringify(data),
  };
  return this.http.post<ReplyProto>(url,JSON.stringify(postdata));
}
 
//get the list of goods type and tag 
GetHomePageType(){
  var url = this.addr + "/homepage/goodstypemsg";
  return this.http.get<GoodsType[]>(url);
}

//login or register interface  🍓
Entrance(data:RequestProto){
  var url = this.addr + "/entrance";
  return this.http.post<ReplyProto>(url, JSON.stringify(data), {withCredentials: true});
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
