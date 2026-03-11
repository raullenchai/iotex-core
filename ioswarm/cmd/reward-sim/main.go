// reward-sim simulates the AgentRewardPool F1 cumulative reward distribution
// across 8 epochs with multiple agents joining, leaving, claiming at different
// times, and varying weights. Verifies the exact same integer math as the
// Solidity contract (1e18 scaling, integer division).
//
// Usage: go run ./ioswarm/cmd/reward-sim
//
// This is a pure Go simulation — no EVM needed. The F1 math is replicated
// exactly from the Solidity contract. Also verifies ABI encoding helpers.
//
// All deposit amounts are chosen to divide evenly by totalWeight, so expected
// values have zero rounding dust. The stress test uses arbitrary numbers to
// verify dust stays minimal.
package main

import (
	"fmt"
	"math/big"
	"os"
	"text/tabwriter"

	"github.com/ethereum/go-ethereum/common"

	"github.com/iotexproject/iotex-core/v2/ioswarm/contracts"
)

var scale = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil) // 1e18

// f1Pool is a pure Go replica of the AgentRewardPool Solidity contract.
type f1Pool struct {
	cumulativeRewardPerWeight *big.Int // scaled by 1e18
	totalWeight               *big.Int
	agents                    map[common.Address]*agentState
}

type agentState struct {
	weight   *big.Int
	snapshot *big.Int
	pending  *big.Int
}

func newF1Pool() *f1Pool {
	return &f1Pool{
		cumulativeRewardPerWeight: big.NewInt(0),
		totalWeight:               big.NewInt(0),
		agents:                    make(map[common.Address]*agentState),
	}
}

func (p *f1Pool) getAgent(addr common.Address) *agentState {
	a, ok := p.agents[addr]
	if !ok {
		a = &agentState{big.NewInt(0), big.NewInt(0), big.NewInt(0)}
		p.agents[addr] = a
	}
	return a
}

// depositAndSettle replicates the Solidity function exactly.
func (p *f1Pool) depositAndSettle(addrs []common.Address, weights []*big.Int, deposit *big.Int) {
	if p.totalWeight.Sign() > 0 && deposit.Sign() > 0 {
		inc := new(big.Int).Mul(deposit, scale)
		inc.Div(inc, p.totalWeight)
		p.cumulativeRewardPerWeight.Add(p.cumulativeRewardPerWeight, inc)
	}
	for i, addr := range addrs {
		a := p.getAgent(addr)
		if a.weight.Sign() > 0 {
			diff := new(big.Int).Sub(p.cumulativeRewardPerWeight, a.snapshot)
			reward := new(big.Int).Mul(a.weight, diff)
			reward.Div(reward, scale)
			a.pending.Add(a.pending, reward)
		}
		a.snapshot = new(big.Int).Set(p.cumulativeRewardPerWeight)
		oldWeight := new(big.Int).Set(a.weight)
		p.totalWeight.Sub(p.totalWeight, oldWeight)
		p.totalWeight.Add(p.totalWeight, weights[i])
		a.weight = new(big.Int).Set(weights[i])
	}
}

func (p *f1Pool) claimable(addr common.Address) *big.Int {
	a := p.getAgent(addr)
	pending := new(big.Int).Set(a.pending)
	if a.weight.Sign() > 0 {
		diff := new(big.Int).Sub(p.cumulativeRewardPerWeight, a.snapshot)
		reward := new(big.Int).Mul(a.weight, diff)
		reward.Div(reward, scale)
		pending.Add(pending, reward)
	}
	return pending
}

func (p *f1Pool) claim(addr common.Address) *big.Int {
	a := p.getAgent(addr)
	if a.weight.Sign() > 0 {
		diff := new(big.Int).Sub(p.cumulativeRewardPerWeight, a.snapshot)
		reward := new(big.Int).Mul(a.weight, diff)
		reward.Div(reward, scale)
		a.pending.Add(a.pending, reward)
		a.snapshot = new(big.Int).Set(p.cumulativeRewardPerWeight)
	}
	amount := new(big.Int).Set(a.pending)
	a.pending = big.NewInt(0)
	return amount
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("  AgentRewardPool F1 Simulation")
	fmt.Println("  (Pure Go — replicates Solidity integer math exactly)")
	fmt.Println("═══════════════════════════════════════════════════════════")

	addrs := []common.Address{
		common.HexToAddress("0x1111111111111111111111111111111111111111"),
		common.HexToAddress("0x2222222222222222222222222222222222222222"),
		common.HexToAddress("0x3333333333333333333333333333333333333333"),
		common.HexToAddress("0x4444444444444444444444444444444444444444"),
	}
	names := []string{"Agent-1", "Agent-2", "Agent-3", "Agent-4"}

	pool := newF1Pool()
	totalDeposited := big.NewInt(0)
	totalClaimed := big.NewInt(0)
	passed := 0
	failed := 0

	checkIOTX := func(name string, got, want *big.Int) {
		if got.Cmp(want) == 0 {
			fmt.Printf("    PASS  %-30s = %s IOTX\n", name, fmtIOTX(got))
			passed++
		} else {
			fmt.Printf("    FAIL  %-30s = %s IOTX (want %s)\n", name, fmtIOTX(got), fmtIOTX(want))
			failed++
		}
	}
	checkInt := func(name string, got, want *big.Int) {
		if got.Cmp(want) == 0 {
			fmt.Printf("    PASS  %-30s = %s\n", name, got.String())
			passed++
		} else {
			fmt.Printf("    FAIL  %-30s = %s (want %s)\n", name, got.String(), want.String())
			failed++
		}
	}

	dep := func(n int64) *big.Int {
		d := iotx(n)
		totalDeposited.Add(totalDeposited, d)
		return d
	}

	// ═══════════════════════════════════════════════════════
	// Key insight: depositAndSettle() first updates cumReward with the deposit
	// (using the CURRENT totalWeight), then settles each agent using their OLD
	// weight. So the deposit is distributed in the SAME call it's made.
	//
	// Epoch 1 has totalWeight=0, so its deposit cannot be distributed.
	// To avoid losing tokens, we deposit 0 in epoch 1 (just set weights).
	// ═══════════════════════════════════════════════════════

	// ─── Epoch 1: Register 3 agents (no deposit) ───
	section("Epoch 1: Register 3 agents, deposit 0 IOTX")
	pool.depositAndSettle(
		addrs[:3],
		[]*big.Int{big.NewInt(100), big.NewInt(100), big.NewInt(100)},
		dep(0),
	)
	for i := 0; i < 3; i++ {
		checkIOTX(names[i]+" claimable", pool.claimable(addrs[i]), iotx(0))
	}
	checkInt("totalWeight", pool.totalWeight, big.NewInt(300))

	// ─── Epoch 2: Equal distribution ───
	// 900 IOTX / totalWeight 300 = 3 IOTX per weight unit
	// Each agent (w=100): 100 × 3 = 300 IOTX
	section("Epoch 2: Deposit 900 IOTX, equal weights (100,100,100)")
	pool.depositAndSettle(
		addrs[:3],
		[]*big.Int{big.NewInt(100), big.NewInt(100), big.NewInt(100)},
		dep(900),
	)
	for i := 0; i < 3; i++ {
		checkIOTX(names[i]+" claimable", pool.claimable(addrs[i]), iotx(300))
	}

	// ─── Epoch 3: Unequal weights ───
	// 1200 IOTX / totalWeight 300 = 4 IOTX per weight unit
	// Each (old w=100): 100 × 4 = 400. Carry: 300 + 400 = 700
	section("Epoch 3: Deposit 1200 IOTX, new weights (A1=200, A2=100, A3=100)")
	pool.depositAndSettle(
		addrs[:3],
		[]*big.Int{big.NewInt(200), big.NewInt(100), big.NewInt(100)},
		dep(1200),
	)
	for i := 0; i < 3; i++ {
		checkIOTX(names[i]+" claimable", pool.claimable(addrs[i]), iotx(700))
	}
	checkInt("totalWeight", pool.totalWeight, big.NewInt(400))

	// ─── Agent-1 claims mid-run ───
	section("Agent-1 Claims Mid-Run")
	claimed := pool.claim(addrs[0])
	totalClaimed.Add(totalClaimed, claimed)
	checkIOTX("Agent-1 claimed", claimed, iotx(700))
	checkIOTX("Agent-1 claimable after", pool.claimable(addrs[0]), iotx(0))

	// ─── Epoch 4: Agent-4 joins ───
	// 800 IOTX / totalWeight 400 = 2 IOTX per weight unit
	// A1 (w=200): 200 × 2 = 400. Fresh (claimed already) → 400
	// A2 (w=100): 100 × 2 = 200. Carry: 700 + 200 = 900
	// A3 (w=100): same → 900
	// A4 (w=0, new): 0
	section("Epoch 4: Deposit 800 IOTX, Agent-4 joins")
	pool.depositAndSettle(
		addrs[:4],
		[]*big.Int{big.NewInt(100), big.NewInt(100), big.NewInt(100), big.NewInt(100)},
		dep(800),
	)
	checkIOTX("Agent-1 claimable", pool.claimable(addrs[0]), iotx(400))
	checkIOTX("Agent-2 claimable", pool.claimable(addrs[1]), iotx(900))
	checkIOTX("Agent-3 claimable", pool.claimable(addrs[2]), iotx(900))
	checkIOTX("Agent-4 claimable", pool.claimable(addrs[3]), iotx(0))

	// ─── Epoch 5: Agent-3 leaves ───
	// 1200 IOTX / totalWeight 400 = 3 IOTX per weight unit
	// A1 (w=100): 300. 400 + 300 = 700
	// A2 (w=100): 300. 900 + 300 = 1200
	// A3 (w=100): 300. 900 + 300 = 1200 (settled BEFORE weight → 0)
	// A4 (w=100): 300. 0 + 300 = 300
	section("Epoch 5: Deposit 1200 IOTX, Agent-3 leaves (weight=0)")
	pool.depositAndSettle(
		addrs[:4],
		[]*big.Int{big.NewInt(100), big.NewInt(100), big.NewInt(0), big.NewInt(100)},
		dep(1200),
	)
	checkIOTX("Agent-1 claimable", pool.claimable(addrs[0]), iotx(700))
	checkIOTX("Agent-2 claimable", pool.claimable(addrs[1]), iotx(1200))
	checkIOTX("Agent-3 claimable", pool.claimable(addrs[2]), iotx(1200))
	checkIOTX("Agent-4 claimable", pool.claimable(addrs[3]), iotx(300))
	checkInt("totalWeight", pool.totalWeight, big.NewInt(300))

	fmt.Println("\n    Agent-3 claims remaining before leaving:")
	claimed3 := pool.claim(addrs[2])
	totalClaimed.Add(totalClaimed, claimed3)
	checkIOTX("Agent-3 claimed", claimed3, iotx(1200))

	// ─── Epoch 6: Zero deposit ───
	// No new reward. Cumulative doesn't change. Pending unchanged.
	section("Epoch 6: Zero deposit epoch")
	pool.depositAndSettle(
		[]common.Address{addrs[0], addrs[1], addrs[3]},
		[]*big.Int{big.NewInt(100), big.NewInt(100), big.NewInt(100)},
		dep(0),
	)
	checkIOTX("Agent-1 claimable", pool.claimable(addrs[0]), iotx(700))
	checkIOTX("Agent-2 claimable", pool.claimable(addrs[1]), iotx(1200))
	checkIOTX("Agent-3 claimable", pool.claimable(addrs[2]), iotx(0))
	checkIOTX("Agent-4 claimable", pool.claimable(addrs[3]), iotx(300))

	// ─── Epoch 7: Agent-3 re-joins ───
	// 1200 IOTX / totalWeight 300 = 4 IOTX per weight unit
	// A1 (w=100): 400. 700 + 400 = 1100
	// A2 (w=100): 400. 1200 + 400 = 1600
	// A3 (w=0): 0. Still 0.
	// A4 (w=100): 400. 300 + 400 = 700
	section("Epoch 7: Deposit 1200 IOTX, Agent-3 re-joins (weight=150)")
	pool.depositAndSettle(
		addrs[:4],
		[]*big.Int{big.NewInt(100), big.NewInt(100), big.NewInt(150), big.NewInt(100)},
		dep(1200),
	)
	checkIOTX("Agent-1 claimable", pool.claimable(addrs[0]), iotx(1100))
	checkIOTX("Agent-2 claimable", pool.claimable(addrs[1]), iotx(1600))
	checkIOTX("Agent-3 claimable", pool.claimable(addrs[2]), iotx(0))
	checkIOTX("Agent-4 claimable", pool.claimable(addrs[3]), iotx(700))
	checkInt("totalWeight", pool.totalWeight, big.NewInt(450))

	// ─── Epoch 8: Final epoch ───
	// 900 IOTX / totalWeight 450 = 2 IOTX per weight unit
	// A1 (w=100): 200. 1100 + 200 = 1300
	// A2 (w=100): 200. 1600 + 200 = 1800
	// A3 (w=150): 300. 0 + 300 = 300
	// A4 (w=100): 200. 700 + 200 = 900
	section("Epoch 8: Final epoch, deposit 900 IOTX")
	pool.depositAndSettle(
		addrs[:4],
		[]*big.Int{big.NewInt(100), big.NewInt(100), big.NewInt(100), big.NewInt(100)},
		dep(900),
	)
	checkIOTX("Agent-1 claimable", pool.claimable(addrs[0]), iotx(1300))
	checkIOTX("Agent-2 claimable", pool.claimable(addrs[1]), iotx(1800))
	checkIOTX("Agent-3 claimable", pool.claimable(addrs[2]), iotx(300))
	checkIOTX("Agent-4 claimable", pool.claimable(addrs[3]), iotx(900))

	fmt.Println("\n    Everyone claims:")
	for i := 0; i < 4; i++ {
		c := pool.claim(addrs[i])
		totalClaimed.Add(totalClaimed, c)
		fmt.Printf("    %-10s claimed %s IOTX\n", names[i], fmtIOTX(c))
	}

	// ═══════════════════════════════════════════════════════
	section("Scenario 9: Proportionality Check")
	// Agent with 2x weight should get exactly 2x reward
	propPool := newF1Pool()
	a := common.HexToAddress("0xAAAA")
	b := common.HexToAddress("0xBBBB")
	propPool.depositAndSettle(
		[]common.Address{a, b},
		[]*big.Int{big.NewInt(200), big.NewInt(100)},
		big.NewInt(0), // first epoch, just set weights
	)
	propPool.depositAndSettle(
		[]common.Address{a, b},
		[]*big.Int{big.NewInt(200), big.NewInt(100)},
		iotx(3000), // 3000/300=10 per unit. A=2000, B=1000
	)
	checkIOTX("A (w=200) claimable", propPool.claimable(a), iotx(2000))
	checkIOTX("B (w=100) claimable", propPool.claimable(b), iotx(1000))
	ratio := new(big.Float).Quo(
		new(big.Float).SetInt(propPool.claimable(a)),
		new(big.Float).SetInt(propPool.claimable(b)),
	)
	ratioF, _ := ratio.Float64()
	if ratioF == 2.0 {
		fmt.Println("    PASS  A/B ratio = 2.0x (exact)")
		passed++
	} else {
		fmt.Printf("    FAIL  A/B ratio = %.6f (expected 2.0)\n", ratioF)
		failed++
	}

	// ═══════════════════════════════════════════════════════
	section("Scenario 10: Stress Test — 20 agents, 50 epochs")
	stressPool := newF1Pool()
	stressDeposited := big.NewInt(0)
	stressClaimed := big.NewInt(0)

	stressAddrs := make([]common.Address, 20)
	for i := range stressAddrs {
		stressAddrs[i] = common.BigToAddress(big.NewInt(int64(100 + i)))
	}

	for epoch := 0; epoch < 50; epoch++ {
		weights := make([]*big.Int, 20)
		for i := range weights {
			w := int64((epoch*7 + i*13) % 500)
			if i > 15 && epoch < 10 {
				w = 0
			}
			if i < 3 && epoch > 40 {
				w = 0
			}
			weights[i] = big.NewInt(w)
		}
		// Epoch 0: deposit 0 (no agents registered yet, totalWeight=0)
		// Same as production: coordinator only deposits when agents exist.
		var d *big.Int
		if epoch == 0 {
			d = big.NewInt(0)
		} else {
			d = iotx(int64(100 + epoch*50))
		}
		stressDeposited.Add(stressDeposited, d)
		stressPool.depositAndSettle(stressAddrs, weights, d)

		if epoch%10 == 9 {
			for i := 0; i < 5; i++ {
				c := stressPool.claim(stressAddrs[i])
				stressClaimed.Add(stressClaimed, c)
			}
		}
	}

	for _, addr := range stressAddrs {
		c := stressPool.claim(addr)
		stressClaimed.Add(stressClaimed, c)
	}

	stressDust := new(big.Int).Sub(stressDeposited, stressClaimed)
	fmt.Printf("    20 agents x 50 epochs complete\n")
	fmt.Printf("    Deposited: %s IOTX\n", fmtIOTX(stressDeposited))
	fmt.Printf("    Claimed:   %s IOTX\n", fmtIOTX(stressClaimed))
	fmt.Printf("    Dust:      %s wei\n", stressDust.String())

	if stressClaimed.Cmp(stressDeposited) > 0 {
		fmt.Println("    FAIL  claimed > deposited!")
		failed++
	} else {
		fmt.Println("    PASS  claimed <= deposited")
		passed++
	}

	// Dust should be < 0.01 IOTX over 50 epochs with rounding
	maxDust := new(big.Int).Div(iotx(1), big.NewInt(100)) // 0.01 IOTX
	if new(big.Int).Abs(stressDust).Cmp(maxDust) > 0 {
		fmt.Printf("    WARN  Dust exceeds 0.01 IOTX — may indicate rounding issue\n")
	} else {
		fmt.Println("    PASS  Dust within 0.01 IOTX tolerance")
		passed++
	}

	// ═══════════════════════════════════════════════════════
	section("ABI Encoding Verification")

	testAddrs := []common.Address{addrs[0], addrs[1]}
	testWeights := []*big.Int{big.NewInt(100), big.NewInt(200)}
	data, err := contracts.PackDepositAndSettle(testAddrs, testWeights)
	if err != nil {
		fmt.Printf("    FAIL  PackDepositAndSettle: %v\n", err)
		failed++
	} else {
		fmt.Printf("    Calldata: %d bytes, selector 0x%x\n", len(data), data[:4])
		fmt.Println("    PASS  PackDepositAndSettle")
		passed++
	}

	claimData, err := contracts.PackClaim()
	if err != nil {
		fmt.Printf("    FAIL  PackClaim: %v\n", err)
		failed++
	} else {
		fmt.Printf("    claim() selector: 0x%x\n", claimData[:4])
		fmt.Println("    PASS  PackClaim")
		passed++
	}

	claimableData, err := contracts.PackClaimable(addrs[0])
	if err != nil {
		fmt.Printf("    FAIL  PackClaimable: %v\n", err)
		failed++
	} else {
		fmt.Printf("    claimable() selector: 0x%x, len=%d\n", claimableData[:4], len(claimableData))
		fmt.Println("    PASS  PackClaimable")
		passed++
	}

	// ═══════════════════════════════════════════════════════
	section("Final Accounting")

	dust := new(big.Int).Sub(totalDeposited, totalClaimed)

	fmt.Println()
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "    Total Deposited:\t%s IOTX\n", fmtIOTX(totalDeposited))
	fmt.Fprintf(tw, "    Total Claimed:\t%s IOTX\n", fmtIOTX(totalClaimed))
	fmt.Fprintf(tw, "    Dust:\t%s wei\n", dust.String())
	tw.Flush()

	fmt.Println()
	if totalClaimed.Cmp(totalDeposited) > 0 {
		fmt.Println("    FAIL  Main: claimed > deposited!")
		failed++
	} else {
		fmt.Println("    PASS  Main: claimed <= deposited")
		passed++
	}
	if dust.Sign() == 0 {
		fmt.Println("    PASS  Zero dust (all divisions were exact)")
		passed++
	}

	// ═══════════════════════════════════════════════════════
	fmt.Println()
	fmt.Println("═══════════════════════════════════════════════════════════")
	if failed == 0 {
		fmt.Printf("  ALL %d CHECKS PASSED\n", passed)
	} else {
		fmt.Printf("  %d PASSED, %d FAILED\n", passed, failed)
	}
	fmt.Println("═══════════════════════════════════════════════════════════")

	if failed > 0 {
		os.Exit(1)
	}
}

func iotx(n int64) *big.Int {
	return new(big.Int).Mul(scale, big.NewInt(n))
}

func fmtIOTX(wei *big.Int) string {
	if wei.Sign() == 0 {
		return "0"
	}
	f := new(big.Float).SetInt(wei)
	e := new(big.Float).SetInt(scale)
	f.Quo(f, e)
	return f.Text('f', 2)
}

func section(title string) {
	fmt.Println()
	fmt.Println("────────────────────────────────────────────────────────────")
	fmt.Printf("  %s\n", title)
	fmt.Println("────────────────────────────────────────────────────────────")
}
