// SPDX-License-Identifier: GPL-3.0
 
pragma solidity ^0.8.0;
 
 
contract demoContract{
    uint256 public val1 ;
    bool public bool1;
    address public addr1;
    string public str1;
    uint256[] public arr;
    mapping(address => uint256) balances;
    mapping(address => mapping(address => uint256)) balances2;
 
 
    function changeValues(uint256 newValue,bool newBool,address newAddress,string newString,uint256 newArrayElement,uint256 newBalances2,address addr1Balances2,address addr2Balances2,uint256 newBalances,address addrBalances) external{
        val1 = newValue;
        bool1 = newBool;
        addr1= newAddress;
        str1=newString;
        arr.push(newArrayElement);
        arr.push(15);
        balances2[addr1Balances2][addr2Balances2]=newBalances2;
        balances[addrBalances] = newBalances;
    }
    function changeArrayValueAtIndex(uint256 index,uint256 newValue) external{
        arr[index]= newValue;
    }
 
}