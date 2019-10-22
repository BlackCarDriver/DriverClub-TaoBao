import { Component, OnInit } from '@angular/core';
import { UserMessage, GoodsShort, MyMessage, Rank, User, RequestProto } from '../../struct';
import { ServerService } from '../../server.service';
import { AppComponent } from '../../app.component';
declare let $: any;

@Component({
  selector: 'app-personal',
  templateUrl: './personal.component.html',
  styleUrls: ['./personal.component.css']
})

export class PersonalComponent implements OnInit {
  msg = new UserMessage();
  emailChance = 0;
  emailBtn = "开启邮箱通知";
  mygoodslist = GoodsShort[100];
  mycollectlist = GoodsShort[100];
  mymessagelist = MyMessage[100];
  hero = Rank[20];
  icare = User[100];
  carei = User[100];
  show_no_goods = false;
  show_no_message = false;
  show_no_collect = false;
  msg_maxrow = 8;
  msg_sumpage = 0;
  msg_nowat = 1;
  msg_offset = 0;
  msg_array = new Array;
  mg_maxrow = 8;
  mg_sumpage = 0;
  mg_nowat = 1;
  mg_offset = 0;
  mg_array = new Array;
  mc_maxrow = 8;
  mc_sumpage = 0;
  mc_nowat = 1;
  mc_offset = 0;
  mc_array = new Array;
  constructor(
    private server: ServerService,
    private app: AppComponent,
  ) { }

  ngOnInit() {
    if (this.server.IsNotLogin()) {
      return;
    }
    this.getmymsg();
  }
  //show larger images 🍛
  showimg(url:string){
    this.app.ShowImg(url);
  }
  //get detail information 🍍🍈🌽🍚
  getmymsg(laster?: boolean) {
    let postdata: RequestProto = {
      api: "mymsg",
      targetid: this.server.userid,
      cachetime: 180,
      cachekey: "mymsg_" + this.server.userid,
    };
    if (laster) {
      postdata.cachetime = 0;
    }
    this.server.GetMyMsg(postdata).subscribe(result => {
      if (result.statuscode != 0) {
        this.app.showMsgBox(-1, "请求个人信息失败,请刷新试试:" + result.msg)
        return;
      }
      let temp:UserMessage = result.data;
      temp.headimg = this.server.changeImgUrl(temp.headimg);
      this.msg = temp;
      //email notification setting 🍣
      this.emailChance = result.rows;
      if (this.emailChance!=0) {
        this.emailBtn = "关闭邮箱通知";
      }
    }, err => {
      this.app.cFail(err);
      return;
    });
    this.getmymgoods();
    this.getmycollect();
    this.getmymessage();
    this.getrank();
    this.getcare();
  }

  //get the list of user i care and which acre me🍍🍈🍞🌽🍚
  getcare() {
    let postdata: RequestProto = {
      api: "mycare",
      targetid: this.server.userid,
      cachetime: 600,
      cachekey: "mycare_" + this.server.userid,
    };
    this.server.GetMyMsg(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        let temp: User[] = result.data[0];
        temp.forEach(row => { row.headimg = this.server.changeImgUrl(row.headimg); });
        let temp2 = result.data[1];
        temp2.forEach(row => { row.headimg = this.server.changeImgUrl(row.headimg); });
        this.icare = temp;
        this.carei = temp2;
      } else {
        this.app.showMsgBox(-1, "请求关注信息失败,请刷新试试" + result.msg)
      }
    }, err => { this.app.cFail(err); });
  }

  //get my goods information 🍍 🍉🍈 🍇🍏 🍞🌽🍚
  getmymgoods() {
    let postdata: RequestProto = {
      api: "mygoods",
      targetid: this.server.userid,
      offset: (this.mg_nowat - 1) * this.mg_maxrow,
      limit: this.mg_maxrow,
      cachetime: 180,
    };
    postdata.cachekey = "mygoods_" + postdata.userid + "_" + postdata.offset + "_" + postdata.limit;
    this.server.GetMyMsg(postdata).subscribe(result => {
      if (result.statuscode != 0) {
        this.app.showMsgBox(-1, "无法获取商品列表，请稍后再试:" + result.msg);
        return;
      }
      if (result.rows == 0) {
        this.show_no_goods = true;
        return;
      }
      //change image url to make it load faster
      let temp: GoodsShort[] = result.data;
      temp.forEach(row => { row.headimg = this.server.changeImgUrl(row.headimg); });
      this.mygoodslist = temp;
      //change the page box
      this.mg_sumpage = Math.ceil(result.sum / this.mg_maxrow);
      if (this.mg_sumpage > 1) {
        this.mg_array = new Array;
        for (let i = 1; i <= this.mg_sumpage && i <= 9; i++) {
          this.mg_array.push(i);
        }
      }
    }, err => { this.app.cFail(err); });
  }

  //get my collect goods information 🍍 🍉 🍈 🍇🍏 🍞🌽🍖🍚
  getmycollect() {
    let postdata: RequestProto = {
      api: "mycollect",
      targetid: this.server.userid,
      offset: (this.mc_nowat - 1) * this.mc_maxrow,
      limit: this.mc_maxrow,
      cachetime: 180,
    };
    postdata.cachekey = "mycollect_" + postdata.targetid + "_" + postdata.offset + "_" + postdata.limit;
    this.server.GetMyMsg(postdata).subscribe(result => {
      if (result.statuscode != 0) {
        this.app.showMsgBox(-1, "获取收藏数据失败,请稍后再试：" + result.msg);
        return;
      }
      if (result.rows == 0) {
        this.show_no_collect = true;
        return;
      }
      //change image url to make it load faster
      let temp: GoodsShort[] = result.data;
      temp.forEach(row => { row.headimg = this.server.changeImgUrl(row.headimg); });
      this.mycollectlist = temp;
      //change the page box display
      if (result.sum > 1) {
        this.mc_array = new Array;
        this.mc_sumpage = Math.ceil(result.sum / this.mc_maxrow);
        for (let i = 1; i <= this.mc_sumpage && i <= 9; i++) {
          this.mc_array.push(i);
        }
      }
    }, err => { this.app.cFail(err) });
  }

  // get my mail message  🍍 🍉🍈 🍇 🍏 🍑🌽🍚
  getmymessage() {
    let postdata: RequestProto = {
      api: "message",
      targetid: this.server.userid,
      offset: (this.msg_nowat - 1) * this.msg_maxrow,
      limit: this.msg_maxrow,
      cachetime: 300,
    };
    postdata.cachekey = "mymsgs_" + postdata.targetid + "_" + postdata.offset + "_" + postdata.limit;
    this.server.GetMyMsg(postdata).subscribe(result => {
      if (result.statuscode != 0) {
        this.app.showMsgBox(-1, "请求消息数据失败,请刷新试试: " + result.msg);
        return;
      }
      this.mymessagelist = result.data;
      if (result.rows == 0){
        this.show_no_message = true;
        return;
      }
      if (result.sum > 1) {
        this.msg_array = new Array;
        this.msg_sumpage = Math.ceil(result.sum / this.msg_maxrow);
        for (let i = 1; i <= this.msg_sumpage && i <= 9; i++) {
          this.msg_array.push(i);
        }
      }
    }, err => { this.app.cFail(err)});
  }

  //get users rank message  🍍🍈🍚
  getrank() {
    if(this.server.IsPhone()){
      return;
    }
    let postdata: RequestProto = {
      api: "rank",
      targetid: this.server.userid,
    };
    this.server.GetMyMsg(postdata).subscribe(result => {
      if (result.statuscode != 0) {
        this.app.showMsgBox(-1, "获取请求排名数据失败,请刷新试试", result.msg);
        return;
      } 
      this.hero = result.data;
    }, err => { this.app.cFail(err); });
  }

  //delete my upload goods 🍑
  deleteMyGoods(gid: string) {
    if (!confirm("确定删除这个商品吗？")) return;
    let postdata: RequestProto = {
      api: "deletemygoods",
      targetid: gid,
      userid: this.server.userid,
    };
    this.server.DeleteMyData(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        this.app.showMsgBox(0, "删除成功!");
        this.getmymgoods();
      } else {
        this.app.showMsgBox(-1, "请求删除商品失败,请刷新试试", result.msg);
      }
    }, error => { console.log(error) });
  }

  //cancel collect a goods 🍑
  cancelCollect(gid: string) {
    if (!confirm("确定取消收藏吗？")) return;
    let postdata: RequestProto = {
      api: "uncollectgoods",
      targetid: gid,
      userid: this.server.userid,
    };
    this.server.DeleteMyData(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        this.app.showMsgBox(0, "取消收藏成功");
        this.getmycollect();
      } else {
        this.app.showMsgBox(-1, "请求取消收藏失败,请刷新试试", result.msg);
      }
    }, error => { console.log(error) });
  }

  //cancel collect a goods 🍑
  deleteMessage(mid: string) {
    if (!confirm("确定要删除这条消息吗？")) return;
    let postdata: RequestProto = {
      api: "deletemymessage",
      targetid: mid,
      userid: this.server.userid,
    };
    this.server.DeleteMyData(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        this.app.showMsgBox(0, "删除成功");
        this.getmymessage();
      } else {
        this.app.showMsgBox(-1, "请求删除消息失败，请稍后重试" + result.msg);
      }
    }, error => { console.log(error) });
  }

  //set a message as already read  🍞
  setIsRead(index: number, id: string) {
    if (this.mymessagelist[index].state == 1) {
      return;
    }
    this.mymessagelist[index].state = 1;
    let postdata: RequestProto = {
      api: "msgisread",
      targetid: id,
      userid: this.server.userid,
    };
    this.server.SmallUpdate(postdata).subscribe(result => {
      if (result.statuscode != 0) {
        this.app.showMsgBox(1, result.msg);
      }
    }, error => { console.log(error) });

  }

  //update the setting of notification email receive 🍣
  updateEmailReveive(){
    if (this.server.IsNotLogin()){
      return;
    }
    let postdata: RequestProto = {
      userid: this.server.userid,
      cachekey: "setemail"+this.server.userid,
      cachetime:60,
    };
    if(this.emailChance==0){
      postdata.api = "SetReceiveEmail";
      if (!confirm("开启本功能后, 若收到其他用户的私信将会以邮箱的形式通知您,\
      为了避免消息太多打扰到用户，发送三次后需回到本页重新开启本功能，是否继续？")){
        return;
      }
    }else{
      if (!confirm("确认关闭邮箱通知功能？")){
        return;
      }
      postdata.api = "cancelReceiveEmail";
    }
    this.server.UpdateMessage(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        if(this.emailChance==0){
          this.app.showMsgBox(0, "开启成功");
          this.emailChance=3;
        }else{
          this.app.showMsgBox(0, "已关闭");
          this.emailBtn = "开启邮箱通知";
          this.emailChance=0;
        }
      } else {
        this.app.showMsgBox(-1, "设置失败，请稍后重试:" + result.msg);
      }
    });
  }
  //#################### reference to pagebox #######################

  //functions reference to my_message area pageboxx 🍏
  setMsgPagebox(topage: number) {
    if (topage < 0 || topage > this.msg_sumpage) return;
    this.msg_nowat = topage;
    this.getmymessage();
  }
  MsgPrepage() {
    if (this.msg_nowat == 0) return;
    this.msg_nowat--;
    this.getmymessage();
    this.adjustMsgPage();
  }
  MsgNextpage() {
    if (this.msg_nowat + 1 > this.msg_sumpage) return;
    this.msg_nowat++;
    this.getmymessage();
    this.adjustMsgPage();
  }
  adjustMsgPage() {
    if (this.msg_sumpage <= 9) return;
    if (this.msg_nowat > 5) {
      this.msg_offset = this.msg_nowat - 5;
    }
  }
  //functions reference to my_goods area pagebox 🍏
  setMgPagebox(topage: number) {
    if (topage < 0 || topage > this.mg_sumpage) return;
    this.mg_nowat = topage;
    this.getmymgoods();
  }
  MgPrepage() {
    if (this.mg_nowat == 0) return;
    this.mg_nowat--;
    this.getmymgoods();
    this.adjustMgPage();
  }
  MgNextpage() {
    if (this.mg_nowat + 1 > this.mg_sumpage) return;
    this.mg_nowat++;
    this.getmymgoods();
    this.adjustMgPage();
  }
  adjustMgPage() {
    if (this.mg_sumpage <= 9) return;
    if (this.mg_nowat > 5) {
      this.mg_offset = this.mg_nowat - 5;
    }
  }
  //functions reference to my_collect area pagebox 🍏
  setMcPagebox(topage: number) {
    if (topage < 0 || topage > this.mc_sumpage) return;
    this.mc_nowat = topage;
    this.getmycollect();
  }
  McgPrepage() {
    if (this.mc_nowat == 0) return;
    this.mc_nowat--;
    this.getmycollect();
    this.adjustMcgPage();
  }
  McgNextpage() {
    if (this.mc_nowat + 1 > this.mc_sumpage) return;
    this.mc_nowat++;
    this.getmycollect();
    this.adjustMcgPage();
  }
  adjustMcgPage() {
    if (this.mc_sumpage <= 9) return;
    if (this.mc_nowat > 5) {
      this.mc_offset = this.mc_nowat - 5;
    }
  }
}
