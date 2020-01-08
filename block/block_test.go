package block_test

import (
	"testing"
	"time"
  "fmt"

	"github.com/stretchr/testify/assert"

	// "github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum/rlp"
	. "github.com/floydeconomy/blockchain/block"
  // "github.com/floydeconomy/blockchain/types"
	"github.com/vechain/thor/thor"
)

func TestBlock(t *testing.T) {

	// tx1 := new(tx.Builder).Clause(tx.NewClause(&thor.Address{})).Clause(tx.NewClause(&thor.Address{})).Build()
	// tx2 := new(tx.Builder).Clause(tx.NewClause(nil)).Build()

	// privKey := string("dce1443bd2ef0c2631adc1c67e5c93f13dc23a41c18b536effbbdcbcdb96fb65")
	now := uint64(time.Now().UnixNano())

	var (
		emptyRoot   thor.Bytes32 = thor.BytesToBytes32([]byte("0"))
		beneficiary thor.Address = thor.BytesToAddress([]byte("abc"))
	)

	block := new(Builder).
		Timestamp(now).
		ParentID(emptyRoot).
		Beneficiary(beneficiary).
		Build()

  h := block.Header()
  // b := block.Body()
  // txs := block.Body().Transactions()

  assert.Equal(t, now, h.Timestamp())
  assert.Equal(t, emptyRoot, h.ParentID())
	assert.Equal(t, beneficiary, h.Beneficiary())

  fmt.Println(h.ID())



	// h := block.Header()
  //
	// txs := block.Transactions()
	// body := block.Body()
	// txsRootHash := txs.RootHash()
  //
	// fmt.Println(h.ID())
  //
	// assert.Equal(t, body.Txs, txs)
	// assert.Equal(t, Compose(h, txs), block)
	// assert.Equal(t, gasLimit, h.GasLimit())
	// assert.Equal(t, gasUsed, h.GasUsed())
	// assert.Equal(t, totalScore, h.TotalScore())
	// assert.Equal(t, emptyRoot, h.StateRoot())
	// assert.Equal(t, emptyRoot, h.ReceiptsRoot())
	// assert.Equal(t, now, h.Timestamp())
	// assert.Equal(t, emptyRoot, h.ParentID())
	// assert.Equal(t, beneficiary, h.Beneficiary())
	// assert.Equal(t, txsRootHash, h.TxsRoot())
  //
	// key, _ := crypto.HexToECDSA(privKey)
	// sig, _ := crypto.Sign(block.Header().SigningHash().Bytes(), key)
  //
	// block = block.WithSignature(sig)
  //
	// data, _ := rlp.EncodeToBytes(block)
	// fmt.Println(Raw(data).DecodeHeader())
	// fmt.Println(Raw(data).DecodeBody())
  //
	// fmt.Println(block.Size())
  //
	// b := Block{}
	// rlp.DecodeBytes(data, &b)
	// fmt.Println(b.Header().ID())
	// fmt.Println(&b)
  //
	// block = new(Builder).
	// 	GasUsed(gasUsed).
	// 	GasLimit(gasLimit).
	// 	TotalScore(totalScore).
	// 	StateRoot(emptyRoot).
	// 	ReceiptsRoot(emptyRoot).
	// 	Timestamp(now).
	// 	ParentID(emptyRoot).
	// 	Beneficiary(beneficiary).
	// 	TransactionFeatures(1).
	// 	Build()
  //
	// assert.Equal(t, tx.Features(1), block.Header().TxsFeatures())
	// data, _ = rlp.EncodeToBytes(block)
	// var bx Block
	// assert.Nil(t, rlp.DecodeBytes(data, &bx))
	// assert.Equal(t, block.Header().ID(), bx.Header().ID())
	// assert.Equal(t, block.Header().TxsFeatures(), bx.Header().TxsFeatures())
}
