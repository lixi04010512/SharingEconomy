//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./User.sol";

contract Community is User  {
  //平台合约主持人EOA
 address private host_community;
 
 //公益是否存在的记录
 mapping(uint => bool) communitiesChk;
 
 //记录合约主持人
 constructor ()  {
   host_community = msg.sender;
 }
 	
 //只有主持人才可执行 
 modifier OnlyHost_community() {
   require(msg.sender == host_community,
   "only host can do this");
    _;
 }
 
 //公益活动的编号
//  uint number;
 
 //公益信息
 struct Communitys{
     string communityName;//公益名称
     string communityIntroduce;//公益介绍
     address beneficiaryAddr;//受益人
     uint communityAmounts;//公益总积分
     uint communitySchedule;//公益进度
 }
 
 //存储所有公益信息
 mapping(uint => Communitys) CommunitysData;
 
 //捐赠者信息
 struct donators{
   address donatorAddr; //捐赠者
     uint recordCnt;   //捐赠记录总量  
     mapping(uint => donateRecords) donateData;//存储所有捐赠记录信息
     bool exist;
 }
 
  //捐赠记录
 struct donateRecords{
     string donatorName;//捐赠者
     uint donationAmount;//捐赠金额
     uint time;//捐赠时间
     string donateCommunityName;//捐赠公益名称
     bool exist;
 }
 
 address[] public donatorsAddr;
  function donatorData() public view  returns(address[] memory){

    return donatorsAddr;
  }
 
 //存储所有捐赠者信息
 mapping(address => donators) donatorsData;
 
  //查询公益活动是否已经存在
 function isCommunityExist(uint number) public view returns(bool) {
   return communitiesChk[number];
 }
 
 //添加一个公益活动
 function addCommunity(uint number,address beneficiaryAddr,string memory communityName,string memory communityIntroduce,uint communityAmounts) public OnlyHost_community {
  //公益活动不存在，才可以添加
  require(!isCommunityExist(number), 
		 "stick already exist");
  require(integralData[beneficiaryAddr].exist == true,"user not exist");
		 
   //设置可以使用此类分类
  communitiesChk[number] = true;
  
  number++;
  CommunitysData[number].communityName=communityName;
  CommunitysData[number].beneficiaryAddr=beneficiaryAddr;
  CommunitysData[number].communityIntroduce=communityIntroduce;
  CommunitysData[number].communityAmounts=communityAmounts;
  CommunitysData[number].communitySchedule=0;
 }
 
  //捐赠积分  
  function donate(uint number,address donator,uint amount,string memory donatorName) public payable returns(uint){
      require(isCommunityExist(number), 
		 "stick not exist");
      
      //转账
    //   erc20().transferFrom(donator,CommunitysData[number].beneficiaryAddr,amount);
      
      donatorsData[donator].recordCnt+=1;
    uint inx=donatorsData[donator].recordCnt;
      
      //添加捐赠记录
      donateRecords memory record = donateRecords({
	 donatorName: donatorName,	//捐赠者名称
	 donationAmount:amount,
	 time: block.timestamp,
	 donateCommunityName:CommunitysData[number].communityName,
	 exist: true		//确认信息存在
   });
   
   donatorsData[donator].donateData[inx] = record;
   donatorsAddr.push(donator);
   
   //捐赠进度
   CommunitysData[number].communitySchedule+=amount;
      
      return inx;
  }
  
  //下架公益活动
  function delCommunity(uint number) public OnlyHost_community{
  //活动必须存在 
  require(isCommunityExist(number), 
		 "stick not exist"); 
  
  communitiesChk[number] =false;
  
 }
 
 //积分进度
 function CommunitySchedule(uint number) public view returns(uint){
     //活动必须存在 
  require(isCommunityExist(number), 
		 "stick not exist");
	return CommunitysData[number].communitySchedule;
	
 }
  
}