import { Component, OnInit } from '@angular/core';
import { ServerService } from '../../server.service';
import { GoodsType, GoodSubType, UploadGoods } from '../../struct';
import { AppComponent } from '../../app.component';

import * as wangEditor from '../../../assets/wangEditor.min.js';
// import * as wangEditor from 'https://unpkg.com/wangeditor/release/wangEditor.min.js';

declare var $: any;

@Component({
  selector: 'app-uploadgoods',
  templateUrl: './uploadgoods.component.html',
  styleUrls: ['./uploadgoods.component.css']
})

export class UploadgoodsComponent implements OnInit {
  typearray = GoodsType[10];
  typelist = GoodSubType[100];
  headImgName = "未选择文件...";
  username = "username";
  editor: any;
  //the following value will be send to server
  userid = "";
  goodsname = "名称未设置";
  title = "标题未设置";
  headImgUrl = "https://img-blog.csdnimg.cn/20191003114954113.jpg"
  date = "";
  price: number = 0.0;
  typename = "";
  tagname = "";
  usenewtag = false;
  newtagname = "";
  godostext = "";

  constructor(
    private server: ServerService,
    private app: AppComponent,
  ) { }

  ngOnInit() {
    if (this.server.IsNotLogin()) {
      return;
    }
    if (this.server.IsPhone()){
      this.app.showMsgBox(1,"你好,本页面上传的图片将被限制大小,建议在电脑版上进行操作并通过截图的方式来降低图片封面的大小哦！");
    }
    this.initImgUpload();
    this.initEditer();
    this.GetType();
    this.date = this.server.formatDate();
    this.userid = this.server.userid;
    if (this.server.username != "") this.username = this.server.username;
  }

  //=================== request server =======================
  //upload select picture to server and get a url. 🍋🔥🍄🍚
  //it function is called by a hidden button which will be clicked after image checking 
  uploadcover() {
    if (this.server.IsNotLogin()) {
      return;
    }
    var files = $("#upload").prop('files');
    this.server.UploadImg("uploadname", files[0]).subscribe(result => {
      if (result.statuscode == 0) {
        this.headImgUrl = result.data;
        return;
      }
      this.app.showMsgBox(-1, "封面上传失败：" + result.msg);
    }, err => { this.app.cFail(err) });
  };

  //upload a goods to server  🍋🍉🍄🍚
  Upload() {
    if (this.server.IsNotLogin()) {
      return;
    }
    if ($("#check").prop("checked") == false) {
      this.app.showMsgBox(1, "请先了解上传规则");
      return;
    }
    let warn = this.checkData();
    if (warn!="") {
      this.app.showMsgBox(1, "商品描述有误:" + warn);
      return;
    }
    let data: UploadGoods = {
      userid: this.userid,
      imgurl: this.headImgUrl,
      name: this.goodsname,
      title: this.title,
      price: this.price,
      date: this.date,
      text: this.godostext,
      type: this.typename,
      usenewtag: this.usenewtag,
      tag: (this.usenewtag ? $("#newtypeinput").val() : this.tagname),
    };
    this.app.showMsgBox(1,"开始发送，请稍等！");
    //note taht Request protocol is write in UploadGoodsData
    this.server.UploadGoodsData(data).subscribe(result => {
      if (result.statuscode != 0) {
        this.app.showMsgBox(-1, "对不起,上传失败,请稍后再试试：" + result.msg);
        return;
      }
      alert("上传成功");
      this.app.showMsgBox(0, "上传成功");
      window.history.back();
    }, err => { this.app.cFail(err) });
  }

  //get goods type list that need to show in select button. 🍋🍄
  GetType() {
    this.server.GetHomePageType().subscribe(
      result => { this.typearray = result; });
  }

  //=================== init component =================
  //deiter setting up : https://www.kancloud.cn/wangfupeng/wangeditor3/332599🍄
  initEditer() {
    this.editor = new wangEditor('#div3');
    this.editor.customConfig.uploadImgShowBase64 = true; //allowed to save image in base64-encoding
    this.editor.customConfig.menus = [
      'head',
      'fontSize',
      'bold',
      'foreColor',
      'backColor',
      'image',
      'emoticon',
      'link',
      'justify',
      'table',
      'code',
    ]
    this.editor.customConfig.zIndex = 1;
    this.editor.create();
    this.editor.txt.html('<b style="color:#ff0000a6;">请在这里编辑你的商品页面，建议在电脑版上进行操作并尽量使用图片链接代替上传图片。</b>')
  }
  //if images select was changed, then upload to server and get a visit url 🍄🍚🍙
  initImgUpload() {
    if (this.server.IsNotLogin()) {
      return;
    }
    $("#upload").change(function(){
      let goodsImg:File = $("#upload").prop('files')[0];
      let imgName = goodsImg.name;
      if(imgName=="") return;
      let err = this.server.checkImgFile(goodsImg);
      if (err!="") {
          alert(err);
          return;
      }
      $("#filename").html(imgName);
      $("#uploadbtn").trigger("click");
    }.bind(this));
  }
  //trigger to open the images select dialogue 
  selectImg() {
    $("#upload").trigger("click");
  }
  //set up the data display in type select box
  selecttype(type: string, index: number) {
    $("#btn-type").html(type + " <span class='caret'>");
    this.typename = type;
    this.typelist = this.typearray[index].list;
    this.usenewtag = false;
  }
  //set up the data display in tag select box
  selectTag(type: string) {
    $("#subtype").html(type + " <span class='caret'>")
    if (type == '新标签') this.usenewtag = true;
    else {
      this.tagname = type;
    }
  }

  //=================== input checking =================
  //check the upload goods data before send to server 🍄🍆🍚
  checkData() {
    if (this.headImgUrl == "") {
      return  "未选择商品封面";
    }
    if (this.goodsname == "" || this.goodsname == "名称未设置") {
      return "商品名为空或太长";
    }
    let err = this.server.checkGoodsName(this.goodsname);
    if (err!=""){
      return err;
    }
    if (this.title == "" || this.title == "标题未设置" || this.server.checkGoodsTitle(this.title)!="") {
     return "商品标题不可太短或太长哦";
    }
    this.price = Math.floor(this.price * 10) / 10;
    if (this.price <= 0 || this.price > 10000 || this.price == null) {
      return"请检查转让价格是否填写有误";
    }
    if (this.typename == "") {
      return"请选择分类"
    }
    if (this.usenewtag == true) {
      this.newtagname = $("#newtypeinput").val();
      if (this.newtagname.length == 0 || this.newtagname.length > 6) {
       return"标签名不可太长或太短哦"
      }
    } else {
      if (this.tagname.length == 0) {
       return"请为商品选择或新增一个标签";
      }
    }
    this.godostext = this.editor.txt.html();
    if (this.godostext.length < 100) {
     return"页面的描述太简单了,请增加一些描述";
    }
    if (this.godostext.length > 500 * 1024) {
     return"对不起，描述页面超过 500kb 了，请删减一些内容";
    }
   return"";
  }


}
