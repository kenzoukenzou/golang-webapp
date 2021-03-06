package main

import (
	"testing"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	client := new(client)
	url, err := authAvatar.GetAvatarURL(client)
	if err != ErroNoAvatarURL {
		t.Error("Authavatar.GetAvatarURL should return ErrorNoAvatarURL when no value present")
	}

	testURL := "http://url-to-gravatar/"
	client.userData = map[string]interface{}{"avatar_url": testURL}
	url, err = authAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("AuthAvatar.GetAvatarURL should return no error when value present.")
	}
	if url != testURL {
		t.Error("AuthAvatar.GetAvatarURL should return correct URL.")
	}
}
