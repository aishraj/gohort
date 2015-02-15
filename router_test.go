package khukuri

import "testing"

func TestRouter(t *testing.T) {
	if t == nil {
		t.Error("t is nil")
	}
	//TODO: Need to figure out a better way to test this
	//httptest seems fine but I'm not yet fully convinced about its use with gorilla/mux.
	//RegisterAndStart()
}
