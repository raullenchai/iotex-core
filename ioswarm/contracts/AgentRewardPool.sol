// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title AgentRewardPool — F1 cumulative reward distribution for IOSwarm agents
/// @notice Coordinator deposits IOTX + updates agent weights each epoch.
///         Agents call claim() to withdraw accumulated rewards.
///         Same F1 math as Cosmos SDK staking rewards.
contract AgentRewardPool {
    address public coordinator;

    // F1 core: cumulative reward per unit of weight (scaled by 1e18)
    uint256 public cumulativeRewardPerWeight;
    uint256 public totalWeight;

    struct Agent {
        uint256 weight;    // current epoch task weight (set by coordinator)
        uint256 snapshot;  // cumulativeRewardPerWeight at last settlement
        uint256 pending;   // settled but unclaimed rewards (in wei/rau)
    }

    mapping(address => Agent) public agents;

    event Deposited(uint256 amount, uint256 agentCount);
    event Claimed(address indexed agent, uint256 amount);
    event WeightUpdated(address indexed agent, uint256 oldWeight, uint256 newWeight);

    modifier onlyCoordinator() {
        require(msg.sender == coordinator, "not coordinator");
        _;
    }

    constructor(address _coordinator) {
        require(_coordinator != address(0), "zero coordinator");
        coordinator = _coordinator;
    }

    /// @notice Coordinator calls each epoch: deposit IOTX + update agent weights.
    /// @param agentAddrs  Agent wallet addresses
    /// @param newWeights  Corresponding new weights (tasks * bonus, computed by coordinator)
    function depositAndSettle(
        address[] calldata agentAddrs,
        uint256[] calldata newWeights
    ) external payable onlyCoordinator {
        require(agentAddrs.length == newWeights.length, "length mismatch");

        // 1. Update cumulative reward with deposited IOTX
        if (totalWeight > 0 && msg.value > 0) {
            cumulativeRewardPerWeight += (msg.value * 1e18) / totalWeight;
        }

        // 2. Settle each agent, then update their weight
        for (uint256 i = 0; i < agentAddrs.length; i++) {
            Agent storage a = agents[agentAddrs[i]];

            // Settle: accumulate pending reward based on old weight
            if (a.weight > 0) {
                a.pending += (a.weight * (cumulativeRewardPerWeight - a.snapshot)) / 1e18;
            }
            a.snapshot = cumulativeRewardPerWeight;

            // Incremental totalWeight update (not recompute)
            uint256 oldWeight = a.weight;
            totalWeight = totalWeight - oldWeight + newWeights[i];
            a.weight = newWeights[i];

            emit WeightUpdated(agentAddrs[i], oldWeight, newWeights[i]);
        }

        emit Deposited(msg.value, agentAddrs.length);
    }

    /// @notice Agent withdraws all accumulated rewards.
    function claim() external {
        Agent storage a = agents[msg.sender];

        // Settle to current cumulative value
        if (a.weight > 0) {
            a.pending += (a.weight * (cumulativeRewardPerWeight - a.snapshot)) / 1e18;
            a.snapshot = cumulativeRewardPerWeight;
        }

        uint256 amount = a.pending;
        require(amount > 0, "nothing to claim");

        // Clear before transfer (reentrancy safe)
        a.pending = 0;

        (bool ok, ) = payable(msg.sender).call{value: amount}("");
        require(ok, "transfer failed");

        emit Claimed(msg.sender, amount);
    }

    /// @notice View: how much an agent can claim right now.
    function claimable(address agent) external view returns (uint256) {
        Agent storage a = agents[agent];
        uint256 pending = a.pending;
        if (a.weight > 0) {
            pending += (a.weight * (cumulativeRewardPerWeight - a.snapshot)) / 1e18;
        }
        return pending;
    }

    /// @notice Allow coordinator to be updated (e.g., node migration).
    function setCoordinator(address newCoordinator) external onlyCoordinator {
        require(newCoordinator != address(0), "zero address");
        coordinator = newCoordinator;
    }

    /// @notice Allow contract to receive IOTX directly (e.g., owner top-up).
    receive() external payable {}
}
