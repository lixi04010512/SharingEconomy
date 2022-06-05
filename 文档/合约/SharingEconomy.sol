//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;


contract SharingEconomy {
 
 //平台合约主持人EOA
 address public host;

 //分类编号
 uint stick_id;

 //分类信息
 struct Sticks{
     string name;//分类名称
     string img;//分类图标
     bool exist;//是否存在
     
 }
 
 //存储所有分类信息
 mapping(uint => Sticks) stickData;

 
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
 
 //添加一种分类
 function addSticker(string memory species,string memory img) public onlyHost {
   stick_id++;
   stickData[stick_id].name=species;
   stickData[stick_id].img=img;
   stickData[stick_id].exist=true;
 
 }
 
 //返回分类
  function getStick(uint id) public view  returns(string memory name,string memory img){
   if(stickData[id].exist==true){
    return (stickData[id].name,stickData[id].img);
    
   }
 }
  
  //修改分类
 function updateStick(uint id,string memory newName) public {
     require(stickData[id].exist==true,"not exist stick");
     stickData[id].name=newName;
 }
 
 //删除分类
function delStick(uint id) public {
     require(stickData[id].exist==true,"not exist stick");
     stickData[id].exist=false;
}
  
 //管理员登录
 function signIn(address people) public view returns(string memory){
     require(people==host,"not real user");
     return "yes";

}
 

}