//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;


contract SharingEcomomy {
 
 //平台合约主持人EOA
 address private host;
 
 //分类是否存在的记录
 mapping(string => bool)  goodsChk;
 
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
  
 
  //查询合约余额
  function queryBalance() public view returns (uint) {
	return address(this).balance;
  }	
}

