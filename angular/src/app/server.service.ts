import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { GoodsType, UploadGoods } from '../app/struct';
import { RequestProto, ReplyProto } from '../app/struct';

@Injectable({
  providedIn: 'root'
})

export class ServerService {

  //global variable 🍈
  userid = "";    //this usrid only can be true usre id
  username = "";
  token = "";
  homepage_goods_perpage = 10;
  imgMaxSize = 300 * 1024;
  private addr: string  = "https://blackcardriver.cn/taobaoserver";
  private rmaddr:string = "https://blackcardriver.cn/taobaoserver";
  // private addr: string = "/localserver";
  constructor(
    private http: HttpClient,
  ) { }
  //====================================== public phsical function =================================
  //check whether the user is login, show the warm message if not 🍈
  IsNotLogin() {
    if (this.userid == "") {
      alert("你好，要先登录呦!");
      return true;
    }
    return false;
  }
  //get last section of persent url
  LastSection() {
    let rawStr = window.location.pathname;
    let lastSlash = rawStr.lastIndexOf("/");
    let result = rawStr.substring(lastSlash + 1);
    return result;
  }
  //scoll to top
  totop() {
    scroll(0, 150);
  }
  //get html element by id 🍄
  getEle(id: string) {
    return (<HTMLInputElement>document.getElementById(id));
  }
  //get a timestamp string 🍄
  formatDate() {
    var date = new Date();
    var myyear: any = date.getFullYear();
    var mymonth: any = date.getMonth() + 1;
    var myweekday: any = date.getDate();
    if (mymonth < 10) {
      mymonth = "0" + mymonth;
    }
    if (myweekday < 10) {
      myweekday = "0" + myweekday;
    }
    return (myyear + "-" + mymonth + "-" + myweekday);
  }
  //change a imgurl to out-of-focus url 🍞
  changeImgUrl(imgUrl:string){
    let id = imgUrl.lastIndexOf("/")
    let newUrl = imgUrl.slice(0,id+1) + "_" + imgUrl.slice(id+1);
    return newUrl;
  }
  //restore the image url that after compress
  restoreImg(imgUrl:string){
    let newUrl = imgUrl.replace("/_","/");
    return newUrl;
  }
  //judge if the service is mobie phone or laptop
  IsPhone(){
    let width = document.body.clientWidth;
    return width < 700;
  }
  gohome(){
    document.location.href="/homepage";
  }
  //======================================= large  interface =============================================================
  //get all kind of data in goodspage 🍌
  GetGoodsDeta(request: RequestProto) {
    var url = this.addr + "/goodsdeta";
    return this.http.post<ReplyProto>(url, JSON.stringify(request));
  }
  //request to update some simple record such as collect number 🍍🍔
  SmallUpdate(request: RequestProto) {
    request.token = this.token;
    var url = this.addr + "/smallupdate";
    return this.http.post<ReplyProto>(url, JSON.stringify(request));
  }
  //request to update some complex message such as profile 🍍🍔
  UpdateMessage(request: RequestProto) {
    request.token = this.token;
    var url = this.addr + "/update";
    return this.http.post<ReplyProto>(url, JSON.stringify(request));
  }
  //upload a images to server and receive a url to get it images 🍍🍆🍙
  UploadImg(username: string, img: any) {
    //check the name and type before send to server
    var postdata = new FormData();
    postdata.append("userid", this.userid);
    postdata.append("token", this.token);
    postdata.append("file", img)
    var url = this.rmaddr + "/upload/images"; //use remote host temply
    //post a multipart/form-data, can not use json.stringfiy
    return this.http.post<ReplyProto>(url, postdata);
  }
  //get information in personal page 🍍🌽
  GetMyMsg(request: RequestProto) {
    var url = this.addr + "/personal/data";
    return this.http.post<ReplyProto>(url, JSON.stringify(request));
  }
  //a little different from GetMyMsg 🍋🍔
  GetCredentMsg(request: RequestProto) {
    request.token = this.token;
    var url = this.addr + "/personal/data";
    return this.http.post<ReplyProto>(url, JSON.stringify(request), { withCredentials: true });
  }
  //delete something  🍑🍔
  //request send to DeleteController
  DeleteMyData(request: RequestProto) {
    request.token = this.token;
    var url = this.addr + "/deleteapi";
    return this.http.post<ReplyProto>(url, JSON.stringify(request));
  }
  //get homepage goods list data 🍋🍇🌽
  GetHomePageGoods(type: string, tag: string, page: number) {
    let postdata: RequestProto = {
      api: "gethomepagegoods",
      offset: this.homepage_goods_perpage * (page - 1),
      limit: this.homepage_goods_perpage,
      data: { goodstype: type, goodstag: tag },
      cachetime:60,
    };
    let key = "ghpg_"+type+"_"+tag+"_"+postdata.offset+"_"+postdata.limit;
    postdata.cachekey = key;
    var url = this.addr + "/homepage/goodsdata";
    return this.http.post<ReplyProto>(url, JSON.stringify(postdata));
  }
  //upload a good data 🍋🍔
  //request send to UploadGoodsController
  UploadGoodsData(data: UploadGoods) {
    var url = this.addr + "/upload/newgoods";
    let postdata: RequestProto = {
      api: "uploadgoodsdata",
      token: this.token,
      data: JSON.stringify(data),
    };
    return this.http.post<ReplyProto>(url, JSON.stringify(postdata));
  }
  //login or register interface  🍓
  Entrance(data: RequestProto) {
    var url = this.addr + "/entrance";
    return this.http.post<ReplyProto>(url, JSON.stringify(data), { withCredentials: true });
  }
  //get the list of goods type and tag,
  //note that it is a get method!
  GetHomePageType() {
    var url = this.addr + "/homepage/goodstypemsg";
    return this.http.get<GoodsType[]>(url);
  }
  //request for postfrom api 🍗
  postFormApi(form:FormData){
    var url = this.addr + "/postform";
    return this.http.post<ReplyProto>(url, form);
  }
  // ==========================  the following function is related to cookie ==================================  

  //use to make the cookie can't be undestant directly
  encryption(code: string) {
    var c = String.fromCharCode(code.charCodeAt(0) + code.length);
    for (var i = 1; i < code.length; i++) {
      c += String.fromCharCode(code.charCodeAt(i) + code.charCodeAt(i - 1));
    }
    return escape(c);
  }
  //restore the string that after encryption 🍄
  decode(code: string) {
    code = unescape(code);
    var c = String.fromCharCode(code.charCodeAt(0) - code.length);
    for (var i = 1; i < code.length; i++) {
      c += String.fromCharCode(code.charCodeAt(i) - c.charCodeAt(i - 1));
    }
    return c;
  }
  //save a cookie with a simple encode 🍄
  setCookie(key: string, val: string) {
    var exp = new Date();
    exp.setTime(exp.getTime() + 1000 * 86400 * 180 );  //save the cookie for six month
    document.cookie = key + "=" + this.encryption(val)+";expires=" + exp.toUTCString()+";path=/";
  }
  //get cookie by cookie name after decode 
  getCookie(name: string) {
    var arrCookie = document.cookie.split("; ");
    for (var i = 0; i < arrCookie.length; i++) {
      var arr = arrCookie[i].split("=");
      if (arr[0] == name) {
        return this.decode(arr[1]);
      }
    }
    return "";
  }
  //clear all cookie 🍄
  clearAllCookie() {
    var keys = document.cookie.match(/[^ =;]+(?=\=)/g);
    if (keys) {
      for (var i = keys.length; i--;)
        document.cookie = keys[i] + '=0;expires=Thu, 01 Jan 1970 00:00:00 GMT;path=/';
    }
  }

  //take object after json.parse from localstroge by name
  getLocalStorge(key: string) {
    var jsdata = window.localStorage[key];
    return JSON.parse(jsdata);
  }

  //============ following function is relate to input checking ===========
  //check the format of username 🍖🍚
  checkUerName(name:string, canEmail?:boolean){
    if(name=="") return "用户名不能为空";
    if(name.includes(" ")) return "用户名不能包含空格";
    let namereg = /^[\u4e00-\u9fa5_a-zA-Z0-9]{2,15}$/;
    if( namereg.test(name) ){
      return ""
    }else if (canEmail==false){
      return "用户名格式不正确,提示：不包含空格,符号,长度为2~15";
    }
    return this.checkEmail(name);
  }
  //check the format of goods name 🍚
  checkGoodsName(name:string){
    if(/^[\u4e00-\u9fa5_a-zA-Z0-9]{2,15}$/.test(name)==false){
      return "商品名不可太长太短或包含空格"
    }
    return "";
  }
  //check the title of upload goods 🍚
  checkGoodsTitle(title:string){
    if(/^[\u4e00-\u9fa5_a-zA-Z0-9 ]{5,45}$/.test(title)==false){
      return "商品标题不可太长太短或太短哦";
    }
    return "";
  }
  //check the format of password 🍖
  checkPassword(pw:string){
    if(pw=="") return "密码不能为空";
    if(pw.includes(" ")) return "密码不能包含空格";
    let passwordreg = /^[a-zA-Z._0-9]{6,20}$/;
    if( passwordreg.test(pw) == false){
     return "密码格式不正确,提示：6~20个字母或数字或._组成"
    }
    return ""
  }
  //check the format of a email 🍖
  checkEmail(email:string){
    if(email=="") return "邮箱不能为空";
    if(email.includes(" ")) return "邮箱不能包含空格";
    let regex = /^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$/
    if(regex.test(email)==false){
      return "邮箱格式不符合规则";
    }
    return ""
  }
  //check format of comfirm code 🍖
  checkCode(code:string){
    if(code=="") return "验证码不能为空";
    let regex = /^[0-9]{6}$/;
    if(regex.test(code)==false) return "验证码格式不正确";
    return "";
  }
  //check goods comment format🍚
  checkComment(cm:string) {
    if (cm.length<=2 || cm.length>=200){
        return "消息太长或为空";
    }
    return ""
  }
  //check user private message 🍚
  checkMessage(msg:string) {
    if(/^[\w\W]{2,150}$/.test(msg)==false){
      return "消息太短或太长";
    }
    return "";
  }
  //check a image file
  checkImgFile(img:File){
    let filename = img.name.replace(/.*(\/|\\)/, "");
    let filetype = filename.substring(filename.lastIndexOf("."), filename.length).toLowerCase();
    if (filetype != ".jpg" && filetype != ".png") {
      return "请选择 png 或 jpg 格式的图片";
    }
    if(img.size>this.imgMaxSize){
      return "由于本站宽带配置实在太低，请上传低于300kb的图片 :("
    }
    return "";
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
