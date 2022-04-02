// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;
import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/IERC20.sol";  
contract a {
    
    address ercaddress;
    
   function setErc(address ercaddress_ ) public {
        ercaddress =ercaddress_; 
   } 
  
   function transfer2other(address to_ ,uint ercnum_) public {
        IERC20(ercaddress).transfer(to_,ercnum_);
    //   return string(abi.encodePacked(a, b,a, b));
   }
   
   function getERCBalance() public view returns (uint){
       return IERC20(ercaddress).balanceOf(address(this));
   }
   
   function getUserBalance(address UserAddr) public view returns (uint){
       require ( msg.sender == UserAddr , "please sure this address is belong to you");
       return IERC20(ercaddress).balanceOf(UserAddr);
   }

   function ERCTranfer(address _from , address _to , uint256 _ercnum) public {
       IERC20(ercaddress).transferFrom(_from,_to,_ercnum);
   }
}


