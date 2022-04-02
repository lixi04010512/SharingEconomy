// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
import "./UniswapExchange.sol";


interface IUniswapFactory {
    /**
     * @notice Launch an exchange 启动一个交换  
     * @param _token the token of the exchange to be alunched 要启动的交易所的令牌  
     * @return exchange the address of the launched exchange 启动exchange的地址  
     */
    function launchExchange(address _token)
        external
        returns (address exchange);

    /**
     * @notice Returns the total number of exchanges launched 返回已启动的交易所总数  
     * @return exchangeCount the total number of exchanges launched 启动的交易所总数  
     */
    function getExchangeCount() external view returns (uint exchangeCount);

    /**
     * @notice Returns the exchange associated with a token 返回与令牌关联的交换  
     * @param _token The address of the token token的地址  
     * @return exchange The address of the exchange exchange地址  
     */
    function tokenToExchangeLookup(address _token)
        external
        view
        returns (address payable exchange);

    /**
     * @notice Returns the token associated with an exchange 返回与exchange关联的令牌  
     * @param _exchange The address of the exchange  exchange地址  
     * @return token The address of the token token的地址  
     */
    function exchangeToTokenLookup(address _exchange)
        external
        view
        returns (address token);

    event ExchangeLaunch(address indexed exchange, address indexed token);
} 


/**
 * @title UniswapFactory
 * @notice A factory for UniswapExchanges  UniswapExchanges的工厂  
 */
contract UniswapFactory is IUniswapFactory {
    // event ExchangeLaunch(address indexed exchange, address indexed token);

    // index of tokens with registered exchanges
    address[] public tokenList;
    mapping(address => address payable) tokenToExchange;
    mapping(address => address) exchangeToToken;

    function launchExchange(address _token)
        public
        virtual
        override
        returns (address exchange)
    {
        require(
            tokenToExchange[_token] == address(0),
            "Already an exchange for that."
        ); //There can only be one exchange per token
        require(
            _token != address(0) && _token != address(this),
            "Not a valid token address."
        );
        address  newExchange = address(new UniswapExchange(_token));
        tokenList.push(_token);
        tokenToExchange[_token] = payable(newExchange);
        exchangeToToken[newExchange] = _token;
        emit ExchangeLaunch(newExchange, _token);
        return newExchange;
    }

    function getExchangeCount()
        public
        virtual
        view
        override
        returns (uint exchangeCount)
    {
        return tokenList.length;
    }

    function tokenToExchangeLookup(address _token)
        public
        virtual
        view
        override
        returns (address payable exchange)
    {
        return tokenToExchange[_token];
    }

    function exchangeToTokenLookup(address _exchange)
        public
        virtual
        view
        override
        returns (address token)
    {
        return exchangeToToken[_exchange];
    }
}
