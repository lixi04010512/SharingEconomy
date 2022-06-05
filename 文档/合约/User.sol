//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./SharingEconomy.sol";

contract User is SharingEconomy{
 
  //所有分类物品笔数
 mapping(string => uint)  goodsInx;

 //所有分物品订单笔数
 mapping(uint => uint)  ordersInx;

 //所有订单的编号
 uint orderId; 
 
 //所有物品的编号
 uint number; 
 
 //用户信息
 struct Integral {
     string name; //用户姓名
     string password; //用户密码
     address people; //用户地址
     string email;//用户邮箱
     string headImg;//头像
     string sign;//个性签名
     uint goodsNum;//物品数量
     bool exist; //是否存在
     bool isLogin;//是否登录
 }
 //租借者信息
 struct Borrower {
     address borrower; //租借者地址
 }
 
 //物品信息
 struct Goods {
   address owner;	//出借人EOA
   Borrower borrowers; //借用人信息
   uint id; //物品id
   string name; //物品名称
   string species; //物品所在的种类
   uint rent; //租金
   uint ethPledge;	//押金
     bool isAgree; //是否审核中
   uint deal;//交易记录
   mapping(uint => dealRecords) dealData;
    uint backs;//归还记录
    mapping(uint => backRecords) backData;
   bool available;  //是否已上架
   bool isBorrow;   //是否已借出
 }

 //订单信息
 struct order {
   address orderOwner;	//出借人
   address orderBorrower; //租借者地址
   uint oId; //交易id
   uint id; //物品id
   uint deal;//借出记录
   uint back;//租用记录
 }
mapping (uint => order) public orderData;

 //物品图片
 struct goodsImg{
   string goodSign;
   string[] goodsImgs;
 }

 //存储物品图片
 mapping (uint => goodsImg) goodsImgData;

  struct Needs {
   uint id;
   address people;
   string species;
   string content;
   string issueTimes;
   
 }

 //交易记录
 struct dealRecords{
   uint blockNum;// 区块号
   uint dealDate; //交易时间
   bytes32 signHash; //签名哈希
   string since; //租借开始时间
   uint borrowDays;//租用天数
   
 }
 //归还记录
 struct backRecords{
     string over; //租借截止时间
     bool isBackAgree; //是否同意归还押金
 }
  
 //存储所有物品信息
 mapping(uint => Goods) public goodsData;
 
 //存储所有用户信息
 mapping(address => Integral) public integralData;

 //每次转账金额
	uint eachAmount =1 ether;
 
 //用户注册
 function register(string memory name,address people,string memory email,string memory password) payable public {
     require(integralData[people].exist == false,"user is exist");
     
     integralData[people].name=name;
     integralData[people].people=people;
     integralData[people].email=email;
     integralData[people].password=password;
     integralData[people].exist=true;
     payable(people).transfer(eachAmount);
 }

 //获取余额
 function getBalance(address people) public view returns (uint){
         return address(people).balance;
}
 
 //用户登录
 function login(address people) public {
     require(people ==msg.sender,"not real user");
     require(integralData[people].people == people,"address errror");
     require(integralData[people].exist == true,"user not exist");
     integralData[people].isLogin =true;
    //  loginUser.push(people);

 }
//修改用户信息
function updateUser(address people,string memory name,string memory email,string memory sign) public{
  require(people ==msg.sender,"not real user");
  require(integralData[people].people == people,"address errror");
     integralData[people].name=name;
     integralData[people].sign=sign;
     integralData[people].email=email;
}

//添加头像
function addUserImg(address people,string memory headImg) public{
  require(people ==msg.sender,"not real user");
  require(integralData[people].people == people,"address errror");
  integralData[people].headImg=headImg;
}

 //用户注销登录
 function logout(address people) public{
     require(integralData[people].people == people,"address errror");
     require(integralData[people].isLogin == true,"people is logout");
     integralData[people].isLogin = false;
 }
 
 //物品上架
 function addGoods(address owner,string memory name,string memory species,uint rent, uint ethPledge,string[] memory goodsImgs,string memory goodSign) public returns(uint){
  require(integralData[owner].isLogin == true,"people is logout");
		 
   //物品序号加1   
   goodsInx[species] +=1;
   uint inx = goodsInx[species];
   
   number++;
   
  goodsData[number].owner=owner;
  goodsData[number].id=number;
  goodsData[number].name=name;
  goodsData[number].species=species;
  goodsData[number].rent=rent;
  goodsData[number].ethPledge=ethPledge;
  goodsData[number].available=true;
  goodsImgData[number].goodSign=goodSign;
  goodsImgData[number].goodsImgs=goodsImgs;
  integralData[owner].goodsNum+=1;
  goodsId.push(number);
   //返回数据索引
   return inx;
 }
 
 //修改物品
 function updGoods(uint id,string memory name,string memory species,uint rent, uint ethPledge) public {
    //必须是借出人才可以修改
    require(goodsData[id].owner ==msg.sender,"owner isn't msg");
    
  goodsData[id].name=name;
  goodsData[id].species=species;
  goodsData[id].rent=rent;
  goodsData[id].ethPledge=ethPledge;
 }

 //充值
 function topUp(uint myAmount,address people ) public payable{
  require(integralData[people].people == people,"address errror");
  payable(people).transfer(myAmount*1000000000000000000);
 }

 //返回用户信息
 function getUser(address user) public view returns (
     string memory name,address people,string memory email,string memory sign,uint goodsNum){
         return (integralData[user].name,integralData[user].people,integralData[user].email,integralData[user].sign,integralData[user].goodsNum);
 }

 //返回交易记录
 function getDealRec(uint id,uint deal)public view returns(uint,uint,bytes32,uint,string memory){
   return(goodsData[id].dealData[deal].blockNum,goodsData[id].dealData[deal].dealDate,goodsData[id].dealData[deal].signHash,goodsData[id].dealData[deal].borrowDays,goodsData[id].dealData[deal].since);
 }
 //返回归还记录
 function getBackRec(uint id,uint back)public view returns(bool,string memory){
   return (goodsData[id].backData[back].isBackAgree,goodsData[id].backData[back].over);
 }

 //返回用户头像
 function getUserImg(address user)public view returns(string memory img){
   return integralData[user].headImg;
 }
 //按照物品ID返回物品信息
 function getGoods(uint id) public view returns (
     address owner,string memory name,string memory species,uint rent,uint ethPledge,string[] memory goodImg,string memory goodSign){
         return (goodsData[id].owner,goodsData[id].name,goodsData[id].species,goodsData[id].rent,goodsData[id].ethPledge,goodsImgData[id].goodsImgs,goodsImgData[id].goodSign);
 }
 //按照订单ID返回订单信息
 function getOrders(uint oid) public view returns (
     address,address,uint,uint,uint){
         return (orderData[oid].orderOwner,orderData[oid].orderBorrower,orderData[oid].oId,orderData[oid].id,orderData[oid].back);
 }

uint[] public goodsId;
 //返回所有物品id
 function getGoodsId() public view returns(uint[] memory){
     
     return goodsId; 
 }

 uint[] public ordersId;
 //返回所有物品id
 function getOrdersId() public view returns(uint[] memory){
     
     return ordersId; 
 }
  
 //物品下架
 function outGoods(address owner,uint id) public {
  //物品必须存在 
   require(goodsData[id].available==true, 
		 "goods not exist");
		 
  //只有拥有者能够下架
  require(owner == goodsData[id].owner);
  
  //操作者必须是拥有者
  require(goodsData[id].owner == msg.sender);
  
  goodsData[id].available=false;
  delete goodsData[id];
  delete goodsImgData[id];
  integralData[owner].goodsNum-=1;
 }

 //交易开始时间
    uint public auctionStart;
	
	//交易期限
    uint public auctionLimit;
 
 
//  借用物品
 function borrowGoods(uint id,uint borrowDays) public payable{ 
   //物品必须存在 
   require(goodsData[id].available==true, 
		 "goods not exist");
   
   //物品必须是可用状态
   require(goodsData[id].available,
           "goods not available");
		   
   //物品必须没被借出	
   require(!goodsData[id].isBorrow,
           "goods already lend");
   //物品必须不在审核中
   require(goodsData[id].isAgree==false,
           "goods already lend");
  
   auctionStart = block.timestamp;
		
	 //设置交易期限	
   auctionLimit = block.timestamp + 12 hours;
		   
   goodsData[id].deal += 1;
   goodsData[id].isAgree=true;
   goodsData[id].dealData[goodsData[id].deal].borrowDays=borrowDays;
   goodsData[id].borrowers.borrower = msg.sender;
    payable(goodsData[id].owner).transfer((borrowDays*goodsData[id].rent+goodsData[id].ethPledge)*1000000000000000000);
 }

 //同意借用
 function agreeBorrow(uint id,uint deal,string memory since) public returns(uint){

   //区块时间需早于交易期限
    require(block.timestamp <= auctionLimit, "auction ended");
   require(goodsData[id].owner ==msg.sender,"owner isn't msg");
   //设置借用人EOA
   
   goodsData[id].dealData[deal].since= since;
   goodsData[id].isAgree=false;
   goodsData[id].isBorrow = true;
   //订单序号加1   
   ordersInx[id] +=1;
   uint inx = ordersInx[id];

   orderId++;
   orderData[orderId].orderOwner=goodsData[id].owner;
   orderData[orderId].orderBorrower=goodsData[id].borrowers.borrower;
   orderData[orderId].id=id;
   orderData[orderId].oId=orderId;
   orderData[orderId].deal=deal;
   orderData[orderId].back=goodsData[id].backs+1;
   goodsData[id].dealData[deal].blockNum = block.number;
   goodsData[id].dealData[deal].dealDate = block.timestamp;

   ordersId.push(orderId);

   return inx;
   
 }

 //不同意借用
 function disagreeBorrow(uint id,uint deal) public payable{
   require(goodsData[id].owner ==msg.sender,"owner isn't msg");
   goodsData[id].isAgree=false;
   payable(goodsData[id].borrowers.borrower).transfer((goodsData[id].dealData[deal].borrowDays*goodsData[id].rent+goodsData[id].ethPledge)*1000000000000000000);
   goodsData[id].borrowers.borrower = address(0);
 }

  //交易哈希
 function dealHash(uint id,uint deal,uint oid) public{
   require(goodsData[id].isAgree=true,"disagree");
   goodsData[id].dealData[deal].signHash = keccak256(abi.encodePacked(id, deal, oid));
 }

 //归还
 function doGoodsReturn(uint id) public { 
   //物品必须存在 
   require(goodsData[id].available==true, 
		 "goods not exist");
	
   //必须是出借人才可以改变状态	
   require(goodsData[id].borrowers.borrower == msg.sender,
           "not goods owner");
		   
   //物品必须已被借出	
   require(goodsData[id].isBorrow,
           "goods not lend");

   goodsData[id].backs += 1;
 }

 //同意归还
 function agreeBack(uint id,uint backs,string memory over) public payable {
   require(goodsData[id].owner ==msg.sender,"owner isn't msg");
   uint pledge = goodsData[id].ethPledge;
   goodsData[id].backData[backs].isBackAgree=true;
   payable( goodsData[id].borrowers.borrower).transfer(pledge*1000000000000000000);
   
   goodsData[id].borrowers.borrower = address(0);
   goodsData[id].backData[backs].over=over;
   goodsData[id].isBorrow = false;
   
 }
 //不同意归还
 function disagreeBack(uint id,uint backs,string memory over) public{
   require(goodsData[id].owner ==msg.sender,"owner isn't msg");

   goodsData[id].backData[backs].isBackAgree=false;
   
     //设置借用人EOA
   goodsData[id].borrowers.borrower = address(0);
   
   //设置归还时间
   goodsData[id].backData[backs].over=over;
   
   //设置为未借出
   goodsData[id].isBorrow = false;
 }

  

}
 