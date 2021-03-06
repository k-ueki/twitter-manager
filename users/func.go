package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

func (u *Client) ConvertIdsToUsers(ids []int64) []User {
	var users []User
	var nums string

	url := baseURL + "users/lookup.json?user_id="

	for i, v := range ids {
		if i != 0 {
			nums += "," + strconv.FormatInt(v, 10)
		}
		if i == 0 {
			nums += strconv.FormatInt(v, 10)
		}
	}
	url += nums

	resp, _ := u.HttpClient.Get(url)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &users)

	fmt.Println("users",users)

	return users
}
