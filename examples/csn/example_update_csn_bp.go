package csnexamples

import (
	"fmt"

	"github.com/wenyining/bce-sdk-go/services/csn"
	"github.com/wenyining/bce-sdk-go/util"
)

func UpdateCsnBp() {
	client, err := csn.NewClient("Your AK", "Your SK", "csn.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new csn client, err: %v.\n", err)
		return
	}
	request := &csn.UpdateCsnBpRequest{
		Name: "csn_bp_test_update", // 带宽包名称
	}
	if err = client.UpdateCsnBp("Your csnBpId", request, util.NewUUID()); err != nil {
		fmt.Printf("Failed to update csn bp, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully updated csn bp.")
}
