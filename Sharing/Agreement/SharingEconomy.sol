//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;


contract SharingEconomy {
 
 //平台合约主持人EOA
 address private host;
 
     struct Managers {
        address people;
        string name;
        string password;
        bool exist;
        bool isLogin;
    }
 
 //分类是否存在的记录
 mapping(string => bool)  goodsChk;
 
 //管理员信息
 mapping(address => Managers) managersData;

 
 //记录合约主持人
 constructor ()  {
   host = msg.sender;
 }
 	
 //只有主持人才可执行 
 modifier onlyHost() {
   require(msg.sender == host,
   "only host can do this");
    _;
 }

 //查询分类是否已经存在
 function isStickExist(string memory species) public view returns(bool) {
   return goodsChk[species];
 }
 
 //添加一种分类
 function addSticker(string memory species) public onlyHost {
  //分类不存在，才可以添加
  require(!isStickExist(species), 
		 "stick already exist");
		 
   //设置可以使用此类分类
  goodsChk[species] = true; 
 
 }
 //返回所有分类
  string[] public stickData;
  function getStick() public view  returns(string[] memory){

    return stickData;
  }
 
  //查询合约余额
  function queryBalance() public view returns (uint) {
	return address(this).balance;
  }	

    //发起人添加管理员
  function addManager(address people,string memory name,string memory password) public onlyHost{
      require(managersData[people].exist == false,"user is exist");
      managersData[people].people=people;
      managersData[people].name=name;
      managersData[people].password=password;
      managersData[people].exist=true;
  }
  
  //管理员登录
 function signIn(address people) public {
     require(people ==msg.sender,"not real user");
     require(managersData[people].people == people,"address errror");
     require(managersData[people].exist == true,"user not exist");
     managersData[people].isLogin =true;

}

//修改管理员信息
function updateUser(address people,string memory name,string memory password) public{
  require(people ==msg.sender,"not real user");
  require(managersData[people].people == people,"address errror");
     managersData[people].name=name;
     managersData[people].password=password;
}

 //管理员注销登录
 function logOut(address people) public{
     require(managersData[people].people == people,"address errror");
     require(managersData[people].isLogin == true,"people is logout");
     managersData[people].isLogin = false;
 }

}