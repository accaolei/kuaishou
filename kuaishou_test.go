package kuaishou

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

const aToken = "ChFvYXV0aC5hY2Nlc3NUb2tlbhJQjDwaXm-VTrswJI0M2irR3GGYfaokX0iGfD8dbmxYVhy8ahyDjV3p3DpdWaqlAlTo8cN3QdXl_X2Z6yBO0D9EArbZ2wK7_JQUsfNASexN1rEaEm0bhDHcLU_umD5Zlw9oVL2loiIgrZayS8tAQK5EzwjxXQqISY-n39LCEk1Xj5ZQnEaakNYoBTAB"
const aRefToken = "ChJvYXV0aC5yZWZyZXNoVG9rZW4SoAG52cvHBYbVHvRmlPywwF5tZz0whDDk2bZ5Md_OLaVGcllM_avx0CRMJH8tX23N0rHu9UbwbSpO-kCwL9i9nuLOe8MN5YOE2wrP0aN_LHZmmZZgPE6aqYNEsAPbDgjYEm6nDdaZtuh8A1kcMNgRiRmTJsqhwmZuOKGlXpE_mE4dloHtXzujO3jSdevLloxPU4JeeeL9SX1beUiiVpRHbWdBGhIdTowfCsMwK4XsuVJP_9EZXs8iILNJAMBu8WgQ1ERCkjDp7ZQ38NkFn7OhBSPxDpkqDd34KAUwAQ"

func TestAuthURL(t *testing.T) {
	k := New("ks678354694969581791", "Wx6rJcAI3yIjgZBTfiM95Q", "https://api.seamon.life/v1/callbac/kuaishou")
	// authURL := k.AuthURL("", "12")
	// fmt.Printf("auth url: %s\n", authURL)
	// code := "a2db1151a93aa4404335c4d3e9a1773d07c71a3ac65b965a724f41328b85696ec94865c5"
	// token, err := k.Code2AccessToken(context.Background(), code)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(token)
	// accessToken := "ChFvYXV0aC5hY2Nlc3NUb2tlbhJQuwf4NZjj3iUe1OxN2z8NzJZLvRKMuhInELS_ZBbPUx08rBhQyZGe-w4dlGZtQWTAOvzEEK24YA-gQ50GGlwO0m2BxtjCkLL-0PJXlLQAjm4aEi12HLC4YkSGiKG1_lVAiiG5ayIgmvBcSt7hnnB5tI5oPhkcWVR3eIgzOSsl0K8ql-ilF-4oBTAB"
	// token, err = k.RefreshAccessToken(context.Background(), token.RefreshToken)
	// fmt.Println(token)
	userInfo, err := k.GetUserInfo(context.Background(), aToken)
	if err != nil {
		fmt.Println(err)
	}
	uByte, _ := json.Marshal(userInfo)

	fmt.Println(string(uByte))
	count, err := k.GetUserVideoCount(context.Background(), aToken)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("all_count:%d,public_count:%d", count.AllCount, count.PublicCount)
	vList, err := k.GetUserVideoInfo(context.Background(), aToken, "", 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vList)
}
