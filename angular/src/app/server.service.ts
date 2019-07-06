import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import {  HomePageGoods , GoodsType,GoodsState,UploadGoods,PersonalExpend } from '../app/struct';
import {  account1, account2, UserShort} from '../app/struct';
import { PersonalBase} from '../app/struct';
import { post } from 'selenium-webdriver/http';
import { JsonpCallbackContext } from '@angular/common/http/src/jsonp';
// import { ConsoleReporter } from 'jasmine';
@Injectable({
  providedIn: 'root'
})

export class ServerService {

  //important config !!!
//本地开发配置
 private addr: string  = "http://localhost:8090"
 //服务器配置
  // private addr: string  = "https://www.blackcardriver.cn/server"
  constructor( 
    private http: HttpClient
  ){ }
 
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

GetUserShort(name:string){
  var postdata = {name : name};
  var url = this.addr + "/getmsg/usershort";
  return this.http.post<UserShort>(url,postdata,{withCredentials: true});
}

  GetHomePageGoods(tag : string, index : number){
      var url = this.addr + "/data";
      return this.http.get<HomePageGoods[]>(url);
  }

  GetHomePageType(){
    var url = this.addr + "/data2";
    return this.http.get<GoodsType[]>(url);
  }

  //get homepagegoods-struct data
  GetGoodsDetail_1(id:number){
      var url = this.addr+"/getmsg/goods/souce?tag=base&goodsid="+id;
      return this.http.get<HomePageGoods>(url);
  }

  GetGoodsDetail_2(id:number){
    var url = this.addr+"/getmsg/goods/souce?tag=state&goodsid="+id;
      return this.http.get<GoodsState>(url);
  }

  GetGoodsDetail_3(id:number){
    var url =  this.addr+"/getmsg/goods/souce?tag=text&goodsid="+id;
    return this.http.get<string>(url);
  }
  //upload goods message to server
  UploadGoods(goods:string){
    var postdata = {goodsdata:goods};
    var url = this.addr+"/upload/goods";
    return this.http.post<string>(
      url,postdata);
  }
  //upload head-img of goods to server and receive an imgurl
  UploadCover(username:string , img:any){
    var postdata = new FormData();
    postdata.append("name", username);
    postdata.append("file",img)
    var url = this.addr + "/upload/cover"; 
    return this.http.post<string>(
      url,postdata
    );
  }     
 //upload head-img of user to server and receive an imgurl
 UploadHeadImg(username:string , img:any){
    var postdata = new FormData();
    postdata.append("name", username);
    postdata.append("file",img)
    var url = this.addr + "/upload/headimg"; 
    return this.http.post<string>(
      url,postdata
    );
  }

  //upload goods describe to server host used by upload goods page
  UploadGoodsData(data:UploadGoods){
    var url = this.addr + "/upload/upload/goodsdata"; 
    return this.http.post<number>(url,data);
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
  //get message of personal in pereesonal page
  GetMyMsg(username:string, key:string, tag:string){
      var url = this.addr + "/getmsg/personal/mymessage"; 
      var data = {tag:tag, name:username, key:key};
      return this.http.post<any>(url,data); 
  }
  //get message of personal2 page
  GetOtherMsg(userid:string){
    var url = this.addr+"/getmsg/othermsg?id="+userid;
    return this.http.get<PersonalExpend>(url);
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
