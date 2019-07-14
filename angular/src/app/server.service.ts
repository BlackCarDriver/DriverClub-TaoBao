import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import {  HomePageGoods , GoodsType,UploadAnyResult,UploadIImgResult,UploadGoods,UserMessage } from '../app/struct';
import {  RequertResult, MyStatus, UpdateResult} from '../app/struct';
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

 //=======================================  重做  =====================================================================
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
    return this.http.post<UploadIImgResult>(url,JSON.stringify(postdata));
}     

//更新信息接口
UpdateMessage(userid:string, tag:string, data:any){
  var postdata = {userid:userid, tag:tag, data:data};
  var url = this.addr + "/update"; 
  return this.http.post<UpdateResult>(url, JSON.stringify(postdata)); 
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
