//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./SharingEconomy.sol";

contract User is SharingEconomy{
 
  //所有分类物品笔数
 mapping(string => uint)  goodsInx;
 
 //所有物品的编号
 uint number; 
 
 //用户信息
 struct Integral {
     string name; //用户姓名
     string password; //用户密码
     address people; //用户地址
     uint integral; //用户积分
     bool exist; //是否存在
     bool isLogin;//是否登录
 }
 
 //租借者信息
 struct Renter {
     address renter; //租借者地址
     uint since; //租借开始时间
     uint over; //租借截止时间
     
 }
 
 //物品信息
 struct Goods {
   address owner;	//出借人EOA
   Renter borrower; //借用人信息
   uint id; //物品id
   string name; //物品名称
   string species; //物品所在的种类
   uint rent; //租金
   uint ethPledge;	//押金
   uint theNew; //物品几成新
   uint count; //借用次数
   bool available;  //是否已上架
   bool isBorrow;   //是否已借出
   
 }
 
  
 //存储所有物品信息
 mapping(uint => Goods) goodsData;
 
 //存储所有用户信息
 mapping(address => Integral) integralData;
 
 
 //用户注册
 function register(string memory name,address people,string memory password)  public {
     require(integralData[people].exist == false,"user is exist");
     
     integralData[people].name=name;
     integralData[people].people=people;
     integralData[people].password=password;
     integralData[people].exist=true;
     loginUser.push(people);
 }
 
 //用户登录
 function login(address people) public{
     require(people ==msg.sender,"not real user");
     require(integralData[people].people == people,"address errror");
     require(integralData[people].exist == true,"user not exist");
     integralData[people].isLogin =true;

 }

 //用户注销登录
 function logout(address people) public{
     require(integralData[people].people == people,"address errror");
     require(integralData[people].isLogin == true,"people is logout");
     integralData[people].isLogin = false;
 }

 address[] public loginUser;
 //返回用户信息
 function getUser() public view returns(address[] memory){
     
     return loginUser;
 }
 
 //物品上架
 function addGoods(address owner,string memory name,string memory species,uint rent, uint ethPledge, uint theNew) public returns(uint){
  //分类必须存在 
  require(isStickExist(species), "stick not exist");
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
  goodsData[number].theNew=theNew;
  goodsData[number].count=goodsData[number].count+1;
  goodsData[number].available=true;
   goodsId.push(number);
   //返回数据索引
   return inx;
 }
 
 //修改物品
 function updGoods(uint id,string memory name,string memory species,uint rent, uint ethPledge, uint theNew) public {
    //必须是借出人才可以修改
    require(goodsData[id].owner ==msg.sender);
    
  goodsData[id].name=name;
  goodsData[id].species=species;
  goodsData[id].rent=rent;
  goodsData[id].ethPledge=ethPledge;
  goodsData[id].theNew=theNew;
 }
 
 
 //按照物品ID返回物品信息
 function getGoods(uint id) public view returns (
     address owner,string memory name,string memory species,uint rent,uint ethPledge,uint theNew){
         return (goodsData[id].owner,goodsData[id].name,goodsData[id].species,goodsData[id].rent,goodsData[id].ethPledge,goodsData[id].theNew);
 }

uint[] public goodsId;
 //返回所有物品id
 function getGoodsId() public view returns(uint[] memory){
     
     return goodsId; 
 }

 
 //判断物品是否上架
 function isGoodExist(string memory species, uint id) public view returns(bool){
   //分类必须存在 
   require(isStickExist(species), 
		 "stick not exist");	 
        return goodsData[id].available;   
 }
 
  
 //物品下架
 function delGoods(address owner,string memory species,uint id) public {
  //分类必须存在 
  require(isStickExist(species), "stick not exist"); 
  //物品必须存在 
   require(isGoodExist(species,id), 
		 "goods not exist");
		 
  //只有拥有者能够下架
  require(owner == goodsData[id].owner);
  
  //操作者必须是拥有者
  require(goodsData[id].owner == msg.sender);
  
  goodsData[id].available=false;
  
  
 }
 
 
 //查询物品借出状态
 function isGoodsLend(string  memory species, uint id) public view returns(bool) { 
   //物品必须存在 
   require(isGoodExist(species,id), 
		 "goods not exist");
	
   //返回借出状态	
   return goodsData[id].isBorrow;
 }
 
 
 //借出物品
 function borrowGoods(address borrower,uint since,string memory species, uint id) public payable { 
   //物品必须存在 
   require(isGoodExist(species,id), 
		 "goods not exist");
   
   //物品必须是可用状态
   require(goodsData[id].available,
           "goods not available");
		   
   //物品必须没被借出	
   require(!goodsData[id].isBorrow,
           "goods already lend");
		   
   //设置借用人EOA
   goodsData[id].borrower.renter =borrower;
   goodsData[id].borrower.since= since;
   
   //设置为已借出
   goodsData[id].isBorrow = true;
   

 }
 
 //查询物品借出人
 function queryBorrower(string memory species,uint id) public view returns(address) { 
   //物品必须存在 
   require(isGoodExist(species,id), 
		 "goods not exist");
	
   //物品必须已被借出	
   require(goodsData[id].isBorrow,
           "goods not lend");
		   
   //返回借出人
   return goodsData[id].borrower.renter;
 }
 
 
 //归还
 function doGoodsReturn(address borrower,uint over,string memory species, uint id) public { 
   //物品必须存在 
   require(isGoodExist(species,id), 
		 "goods not exist");
	
   //必须是出借人才可以改变状态	
   require(goodsData[id].borrower.renter == borrower,
           "not goods owner");
		   
   //物品必须已被借出	
   require(goodsData[id].isBorrow,
           "goods not lend");
	
   //将押金返还借用人
   uint pledge = goodsData[id].ethPledge;
    payable( goodsData[id].borrower.renter).transfer(pledge);
   
   
   //设置借用人EOA
   goodsData[id].borrower.renter = address(0);
   
   //设置归还时间
   goodsData[id].borrower.over=over;
   
   //设置为未借出
   goodsData[id].isBorrow = false;
   
   //借入人获得积分
   uint integral =goodsData[id].rent *2;
   integralData[borrower].integral=integralData[borrower].integral + integral;
 }
 
 
//   //购买理赔保险
//  function BuyInsurance(address addr,string memory species,uint id) public {
//      //必须本人购买
//     require(addr == msg.sender); 
//      //物品必须存在 
//    require(isGoodExist(species,id), 
// 		 "goods not exist");
		 
		 
//     if (addr == goodsData[id].owner){
//         goodsData[id].buy_owner=true;
//     }else if (addr == goodsData[id].borrower.renter){
//         goodsData[id].buy_borrower=true;
//     }
    
//  }

//  //查看是否买了保险
//  function searchInsurance(address addr,uint id) public view returns (bool a){
     	 
//     if (addr == goodsData[id].owner){
//        return goodsData[id].buy_owner;
//     }else if (addr == goodsData[id].borrower.renter){
//        return goodsData[id].buy_borrower;
//     }
//    }
   
 }
 