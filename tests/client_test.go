package test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/go-phd/ssf"
	"github.com/stretchr/testify/require"
)

type User struct {
	Id           int64 //主键
	Username     string
	Password     string
	Group        int
	ExtendedPerm string
	StaticPerm   string
}

// TestGet is a sample to run an endpoint test
func TestUser(t *testing.T) {
	var err error

	// Post
	name := "cytest"
	ui := User{Username: name, Password: "pw1", Group: 1}
	resp, err := ssf.RestClient.Post("DBManager", "v1/users", ui)
	t.Log("post", ui, "success, rsp =", string(resp), err)
	require.Nil(t, err, "err must be nil")

	// Post fail
	ui = User{Username: name, Password: "pw1", Group: 1}
	resp, err = ssf.RestClient.Post("DBManager", "v1/users", ui)
	require.Nil(t, err, "err must be nil")
	require.Equal(t, strings.Contains(string(resp), "is already exist"), true, "Username is already exist")
	t.Log("post again, judge", ui, "success, yes, username is already exist")

	// Get
	uo := User{}
	resource := fmt.Sprintf("v1/users/%s", name)
	resp, err = ssf.RestClient.Get("DBManager", resource)
	require.Nil(t, err, "err must be nil")
	err = json.Unmarshal(resp, &uo)
	require.Nil(t, err, "err must be nil")
	t.Log("get", resource, "success,", uo)

	// GetAll
	us := []User{}
	resp, err = ssf.RestClient.Get("DBManager", "v1/users")
	require.Nil(t, err, "err must be nil")

	err = json.Unmarshal(resp, &us)
	require.Nil(t, err, "err must be nil")
	t.Log("get all success,", us)

	for idx, u := range us {
		t.Log(idx, u)
	}

	// Put
	newname := "cytest_new"
	resource = fmt.Sprintf("v1/users/%s", name)
	ui = User{Username: newname, Password: "pw2", Group: 2}
	resp, err = ssf.RestClient.Put("DBManager", resource, ui)
	require.Nil(t, err, "err must be nil")
	t.Log("put", resource, "success, ui = ", ui, "resp", string(resp), err)

	// Get fail
	uo = User{}
	resource = fmt.Sprintf("v1/users/%s", name)
	resp, err = ssf.RestClient.Get("DBManager", resource)
	require.Nil(t, err, "err must be nil")
	require.Equal(t, strings.Contains(string(resp), "no row found"), true, "QuerySeter no row found")
	t.Log("get old resource, judge", resource, "success, yes, no row found")

	// Get
	uo = User{}
	resource = fmt.Sprintf("v1/users/%s", newname)
	resp, err = ssf.RestClient.Get("DBManager", resource)
	require.Nil(t, err, "err must be nil")
	err = json.Unmarshal(resp, &uo)
	require.Nil(t, err, "err must be nil")
	t.Log("get", resource, "success,", uo)

	// GetAll
	us = []User{}
	resp, err = ssf.RestClient.Get("DBManager", "v1/users")
	require.Nil(t, err, "err must be nil")

	err = json.Unmarshal(resp, &us)
	require.Nil(t, err, "err must be nil")
	t.Log("get all success,", us)

	for idx, u := range us {
		t.Log(idx, u)
	}

	// Put fail
	resource = fmt.Sprintf("v1/users/%s", name)
	ui = User{Username: newname, Password: "pw2", Group: 1}
	resp, err = ssf.RestClient.Put("DBManager", resource, ui)
	require.Nil(t, err, "err must be nil")
	require.Equal(t, strings.Contains(string(resp), "User not exist"), true, "User not exist")
	t.Log("put again, judge", resource, "success, yes, user not exist")

	// Delete
	resource = fmt.Sprintf("v1/users/%s", newname)
	ui = User{Username: newname, Password: "pw2"}
	resp, err = ssf.RestClient.Delete("DBManager", resource)
	require.Nil(t, err, "err must be nil")
	t.Log("delete", ui, "success, rsp =", string(resp), err)

	// Delete fail
	resource = fmt.Sprintf("v1/users/%s", newname)
	ui = User{Username: newname, Password: "pw2"}
	resp, err = ssf.RestClient.Delete("DBManager", resource)
	require.Nil(t, err, "err must be nil")
	require.Equal(t, strings.Contains(string(resp), "User not exist"), true, "User not exist")
	t.Log("delete again, judge", resource, "success, yes, user not exist")

}
