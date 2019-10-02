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
  warnmsg = "";
  editor: any;
  //the following value will be send to server
  userid = "";
  goodsname = "名称未设置";
  title = "标题未设置";
  headImgUrl = "https://gss0.bdstatic.com/6LZ1dD3d1sgCo2Kml5_Y_D3/sys/portrait/item/_2f6de585abe7baa7e5a4a7e78b82e9a38e5a"
  date = "";
  price:number = 0.0;
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
      window.history.back();
      return;
    }
    this.initImgUpload();
    this.initEditer();
    this.GetType();
    this.date = this.server.formatDate();
    this.userid = this.server.userid;
    if(this.server.username!="")  this.username = this.server.username;
  }

  //=================== request server =======================
  //upload select picture to server and get a url. 🍋🔥🍄
  uploadcover() {
    var files = $("#upload").prop('files');
    this.server.UploadImg("uploadname", files[0]).subscribe(result => {
      if (result.statuscode == 0) {
        this.headImgUrl = result.data;
        return;
      }
      this.app.showMsgBox(-1, "封面上传失败：" + result.msg);
    }, err => { this.app.cFail(err) });
  };
  //upload a goods to server  🍋🍉🍄
  Upload() {
    if (this.checkData() != true) {
      this.app.showMsgBox(1, "商品描述有误:"+this.warnmsg);
      return;
    }
    if ($("#check").prop("checked") == false) {
      this.app.showMsgBox(1, "请先了解上传规则");
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
    //note taht Request protocol is write in UploadGoodsData
    this.server.UploadGoodsData(data).subscribe( result => {
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
    this.editor.txt.html('<p>请在这里编辑你的商品页面，建议在电脑版上进行操作并尽量使用图片链接代替上传图片。</p>')
  }
  //if images select was changed, then upload to server and get a visit url 🍄
  initImgUpload() {
    $("#upload").change(function (evt) {
      if ($(this).val() == '') return;
      //check file size, max size is 100kb
      var files = evt.currentTarget.files;
      var filesize = files[0].size;
      if (filesize > 102400) {
        alert( "服务器配置太低，请上传低于100kb的图片，谢谢！");
        return;
      }
      //check the file type 
      var filename = $(this).val().replace(/.*(\/|\\)/, "");
      var filetype = filename.substring(filename.lastIndexOf("."), filename.length).toLowerCase();
      if (filetype != ".jpg" && filetype != ".png") {
        alert("请上传 png 或 jpg 格式的图片, 谢谢！");
        return;
      } else {
        $("#filename").html(filename);
        //begain to upload images
        $("#uploadbtn").trigger("click");
      }
    });
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
  //check the upload goods data before send to server 🍄🍆
  checkData() {
    if (this.headImgUrl == "") {
      this.warnmsg = "未选择商品封面";
      return false;
    }
    if (this.goodsname == "" || this.goodsname=="名称未设置" || this.goodsname.length > 20) {
      this.warnmsg = "商品名为空或太长";
      return false
    }
    if (this.title == "" || this.title=="标题未设置" || this.title.length>49) {
      this.warnmsg = "商品标题太短或太长";
      return false;
    }
    this.price = Math.floor(this.price * 10) / 10;
    if (this.price <= 0 || this.price > 10000 || this.price==null) {
      this.warnmsg = "请检查转让价格是否填写有误";
      return false;
    }
    if (this.typename == "") {
      this.warnmsg = "请选择分类"
      return false;
    }
    if (this.usenewtag == true) {
      this.newtagname = $("#newtypeinput").val();
      if (this.newtagname.length == 0 || this.newtagname.length > 6) {
        this.warnmsg = "请检查新标签名是否有误"
        return false;
      }
    } else {
      if (this.tagname.length == 0) {
        this.warnmsg = "请选择或新增一个标签";
        return false;
      }
    }
    this.godostext = this.editor.txt.html();
    if (this.godostext.length < 100) {
      this.warnmsg = "页面的描述太简单了,请增加一些描述";
      return false;
    }
    if (this.godostext.length > 300 * 1024) {
      this.warnmsg = "对不起，描述页面超过300kb了，请删减一些内容";
      return false;
    }
    this.warnmsg = "";
    return true;
  }


}
