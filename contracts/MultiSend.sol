pragma solidity >=0.7.0 <0.9.0;

contract MultiSend {
    // to save the owner of the contract in construction
    address private owner;

    // modifier to check if the caller is owner
    modifier isOwner() {
        require(msg.sender == owner, "Caller is not owner");
        _;
    }

    /**
     * @dev Set contract deployer as owner
     */
    constructor() {
        owner = msg.sender;
    }

    function changeOwner(address newOwner) public isOwner {
        owner = newOwner;
    }

    /**
     * @dev Return owner address
     * @return address of owner
     */
    function getOwner() external view returns (address) {
        return owner;
    }

    // sum adds the different elements of the array and return its sum
    function totalAmount(
        uint[] memory amounts
    ) private pure returns (uint totalValue) {
        // the value of message should be exact of total amounts
        uint total = 0;

        for (uint i = 0; i < amounts.length; i++) {
            total += amounts[i];
        }

        return total;
    }

    // withdraw perform the transfering of ethers
    function transferOne(
        address payable receiverAddr,
        uint receiverAmnt
    ) private {
        receiverAddr.transfer(receiverAmnt);
    }

    // withdrawls enable to multiple withdraws to different accounts
    // at one call, and decrease the network fee
    function multisend(
        address payable[] memory addresses,
        uint[] memory amounts
    ) public payable isOwner {
        // the addresses and amounts should be same in length
        require(
            addresses.length == amounts.length,
            "The length of two array should be the same"
        );

        // the value of the message in addition to sotred value should be more than total amounts
        uint totalAmounts = totalAmount(amounts);

        require(
            msg.value >= totalAmounts,
            "The value is not sufficient or exceed"
        );

        for (uint i = 0; i < addresses.length; i++) {
            // send the specified amount to the recipient
            transferOne(addresses[i], amounts[i]);
        }
    }
}
