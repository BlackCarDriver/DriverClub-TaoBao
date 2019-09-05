import { Component, OnInit } from '@angular/core';
import { ServerService } from '../server.service';
import { GoodsType, GoodSubType, UploadGoods } from '../struct';
declare var $: any;


@Component({
  selector: 'app-uploadgoods',
  templateUrl: './uploadgoods.component.html',
  styleUrls: ['./uploadgoods.component.css']
})


export class UploadgoodsComponent implements OnInit {

  headImgName = "未选择文件...";
  warnmsg = "";
  typearray = GoodsType[10];
  typelist = GoodSubType[100];
  //以下是打包上传到服务端的数据
  headImgUrl = "http://localhost:8090/source/images?tag=headimg&&name=testcover.jpg"
  title = "黑车司机二手交易平台大法好！";
  username = "blackcardriver";
  date = "2019-04-07";
  price = 0.0;
  typename = "";
  tagname = "";
  usenewtag = false;
  newtagname = "";
  godostext = "";

  constructor(private server: ServerService) { }

  ngOnInit() {
    //初始化富文本编辑器
    $('#summernote').summernote({
      placeholder: '<p><span style="font-size: 36px;">这里编辑你的商品展示页面</span></p>',
      tabsize: 2,
      minHeight: 300,
      toolbar: [
        ['fontname', ['fontname']], //字体系列                                 
        ['style', ['bold', 'italic', 'underline']], // 字体粗体、字体斜体、字体下划线、字体格式清除       
        ['fontsize', ['fontsize']], //字体大小                                
        ['color', ['color']], //字体颜色             
        ['style', ['style']],//样式
        ['para', ['paragraph']], //无序列表、有序列表、段落对齐方
        ['table', ['table']], //插入表格    
        ['hr', ['hr']],//插入水平线                             
        ['picture', ['picture']], //插入图片               
      ],
    });
    $(document).ready(function () {
      //上传头像框改变后，获取文件名，判断文件大小，上传文件，获得imgurl
      $("#upload").change(function (evt) {
        //如果文件为空 
        if ($(this).val() == '') {
          return;
        }
        //判断文件大小
        var files = evt.currentTarget.files;
        var filesize = files[0].size;
        //  console.log(filesize);
        if (filesize > 102400) {
          alert("请上传100kb 以下的图片");
          return;
        }
        //判断文件类型，并获取文件名到页面
        var filename = $(this).val().replace(/.*(\/|\\)/, "");
        var pos = filename.lastIndexOf(".");
        var filetype = filename.substring(pos, filename.length)  //此处文件后缀名也可用数组方式获得str.split(".") 
        if (filetype.toLowerCase() != ".jpg" && filetype.toLowerCase() != ".png") {
          alert("请上传 png 或 jpg 格式的图片");
          return;
        } else {
          $("#filename").html(filename);
          //上传图片到服务端并获imgurl
          $("#uploadbtn").trigger("click");
        }
      });
      //解决下拉菜单按钮不能下拉
      $(".dropdown-toggle").on('click', function () {
        $('.dropdown-toggle').dropdown();
      });

    });//ready() is over
    this.date = this.formatDate();
    //获得分类数据
    this.GetType();
  }//oninit() is over

  //upload select picture to server and get a url. 🍋🔥
  uploadcover() {
    var files = $("#upload").prop('files');
    this.server.UploadImg("uploadname", files[0]).subscribe(
      result => {
        if (result.statuscode == 0) {
          this.headImgUrl = result.data;
        } else {
          alert("上传失败" + result.msg);
        }
      }
    )
  };

  //upload a goods to server  🍋
  Upload() {
    //注意这里跟常规用法不同
    if ($("#check").prop("checked") == false) {
      alert("请先了解上传规则");
      return;
    }
    if (this.checkData() == true) {
      var data = new UploadGoods();
      data.userid = "blackcardriver";
      data.title = this.title;
      data.date = this.date;
      data.price = this.price;
      data.type = this.typename;
      data.usenewtag = this.usenewtag;
      data.imgurl = this.headImgUrl;
      data.text = this.godostext;
      if (this.usenewtag) {
        data.tag = $("#newtypeinput").val();
      } else {
        data.tag = this.tagname
      }
      this.server.UploadGoodsData(data).subscribe(
        result => {
          if (result.statuscode == 0) {
            this.headImgUrl = result.data;
            alert("上传成功！")
          } else {
            alert("上次失败："+result.msg);
          }
        },error=>{console.log("UploadGoodsData() fail:"+error);});
    } else {
      alert("商品描述有误，请继续完善");
    }
  }

  //get goods type list that need to show in select button. 🍋
  GetType() {
    this.server.GetHomePageType().subscribe(
      result => { this.typearray = result; });
  }

  //在页面中获得需要上传的值并且检查是否正确
  checkData() {
    if (this.headImgUrl == "http://imdg5.duitang.com/uploads/item/201601/17/20160117222537_3vCcm.jpeg") {
      this.warnmsg = "未选择商品封面"
      return false;
    }
    if (this.price < 0 || this.price > 10000) {
      this.warnmsg = "请检查出售价格是否有误";
      return false;
    }
    if (this.title.length == 0) {
      this.warnmsg = "商品标题不能为空";
      return false;
    }
    if (this.title.length > 24) {
      this.warnmsg = "商品标题太长了"
      return false;
    }
    if (this.typename == "") {
      this.warnmsg = "请选择商品分类"
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
        this.warnmsg = "请选择商品标签";
        return false;
      }
    }
    this.godostext = $('#summernote').summernote('code');
    if (this.godostext.length < 100) {
      this.warnmsg = "你的商品描叙太短，请增加一些描叙";
      return false;
    }
    if (this.godostext.length > 300 * 1024) {
      this.warnmsg = "你的商品描述超过300kb，请删减一些内容";
      return false;
    }
    this.warnmsg = "";
    return true;
  }

  //点击选择封面后激活input标签选择文件
  selectImg() {
    $("#upload").trigger("click");
  }

  //选择分类后记录这个值并更新到按钮显示
  selecttype(type: string, index: number) {
    $("#btn-type").html(type + " <span class='caret'>");
    this.typename = type;
    this.typelist = this.typearray[index].list;
    this.usenewtag = false;
  }

  //选择子分类后将子分类显示到按钮
  GetSubType(type: string) {
    $("#subtype").html(type + " <span class='caret'>")
    if (type == '新标签') this.usenewtag = true;
    else {
      this.tagname = type;
    }
  }



  //得到当日的格式化后的日期
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


}
