pragma solidity  0.7.4;

contract DataStorage {
    struct transaction {
        string fromAddr;
        string toAddr;
        uint256 amount;
    }

    mapping(string=>string) didChainAddr;
    mapping(string=>string) didClaimHash;
    mapping(string=>string) didDocumentHash;
    mapping(string=>string) didPublicKey;
    mapping(string=>transaction) transactions;

    event LogSetDidChainAddr(string did, string _didChainAddr);
    event LogSetDidClaimHash(string did, string _didClaimHash);
    event LogSetDidDocumentHash(string did, string _didDocumentHash);
    event LogSetDidPublicKey(string did, string _didPublicKey);
    event LogSetTransaction(string txId, string fromAddr, string toAddr, uint256 amount);


    constructor() public{

    }

    function setDidChainAddr(string memory did, string memory _didChainAddr) payable public {
        require(bytes(did).length  != 0);
        require(bytes(_didChainAddr).length  != 0);
        didChainAddr[did] = _didChainAddr;
        emit LogSetDidChainAddr(did, _didChainAddr);
    }

    function getDidChainAddr(string memory did) public view returns (string memory) {
        require(bytes(did).length  != 0);
        return didChainAddr[did];
    }

    function setDidClaimHash(string memory did, string memory _didClaimHash) payable public {
        require(bytes(did).length  != 0);
        require(bytes(_didClaimHash).length  != 0);
        didClaimHash[did] = _didClaimHash;
        emit LogSetDidClaimHash(did, _didClaimHash);
    }

    function getDidClaimHash(string memory did) public view returns (string memory) {
        require(bytes(did).length  != 0);
        return didClaimHash[did];
    }

    function setDidDocumentHash(string memory did, string memory _didDocumentHash) payable public {
        require(bytes(did).length  != 0);
        require(bytes(_didDocumentHash).length != 0);
        didDocumentHash[did] = _didDocumentHash;
        emit LogSetDidDocumentHash(did, _didDocumentHash);
    }

    function getDidDocumentHash(string memory did) public view returns (string memory) {
        require(bytes(did).length  != 0);
        return didDocumentHash[did];
    }

    function setDidPublicKey(string memory did, string memory _didPublicKey) payable public {
        require(bytes(did).length  != 0);
        require(bytes(_didPublicKey).length  != 0);
        didPublicKey[did] = _didPublicKey;
        emit LogSetDidPublicKey(did, _didPublicKey);
    }

    function getDidPublicKey(string memory did) public view returns (string memory) {
        require(bytes(did).length  != 0);
        return didPublicKey[did];
    }

   function setTransaction(string memory txId, string memory _fromAddr, string memory _toAddr, uint256 _amount) payable public {
        require(bytes(txId).length  != 0);
        require(bytes(_fromAddr).length  != 0);
        require(bytes(_toAddr).length  != 0);
        require(_amount != 0);
        transactions[txId] = transaction({
                fromAddr: _fromAddr,
                toAddr: _toAddr,
                amount: _amount
            });
        emit LogSetTransaction(txId, _fromAddr, _toAddr, _amount);
    }

    function getTransaction(string memory txId) public view returns (string memory, string memory, uint256) {
        require(bytes(txId).length  != 0);
        string memory _fromAddr = transactions[txId].fromAddr;
        string memory _toAddr = transactions[txId].toAddr;
        uint256 _amount = transactions[txId].amount;

        return (_fromAddr, _toAddr, _amount);
    }
}
