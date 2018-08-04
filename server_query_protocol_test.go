package servers

import (
	"fmt"
	"testing"
)

const (
	unknownHostServer  string = "this.host.shouldnt.exist.either:27016"
	unresponsiveServer string = "google.com:27016"
)

func TestGetServerInfo(t *testing.T) {
	serverInfoChannel, errorChannel := GetServerInfo("81.19.212.190:27016", "500ms")

	select {
	case serverInfo := <-serverInfoChannel:
		t.Log(serverInfo)
	case error := <-errorChannel:
		t.Fatalf("Error during server info fecthing: %s", error)
	}
}

func TestGetServerInfo_InvalidTimeout(t *testing.T) {
	serverInfoChannel, errorChannel := GetServerInfo("81.19.212.190:27016", "Hudson Mowhawke is pretty good")

	select {
	case <-serverInfoChannel:
		t.Fatal("This test is supposed to fail. It hasn't. Now go fix the timeout parsing function!")
	case error := <-errorChannel:
		t.Logf("Error: %s", error)
	}
}

func TestGetServerInfo_UnknownHostServer_CI(t *testing.T) {
	serverInfoChannel, errorChannel := GetServerInfo(unknownHostServer, "500ms")

	select {
	case <-serverInfoChannel:
		t.Fatalf("This test is supposed to fail. Apparently the host '%s' exists on this network.", unknownHostServer)
	case error := <-errorChannel:
		t.Logf("Error: %s", error)
	}
}

func TestGetServerInfo_UnresponsiveServer(t *testing.T) {
	serverInfoChannel, errorChannel := GetServerInfo(unresponsiveServer, "1s")

	select {
	case <-serverInfoChannel:
		t.Fatalf("This test expects no server info response from '%s' but apparently it has in fact responded. Well, that's awkward.", unresponsiveServer)
	case error := <-errorChannel:
		t.Logf("Error: %s", error)
	}
}

func ExampleGetServerInfo() {
	serverInfoChannel, errorChannel := GetServerInfo("81.19.212.190:27016", "500ms")

	select {
	case serverInfo := <-serverInfoChannel:
		fmt.Printf("Received server info: %s", serverInfo)
	case error := <-errorChannel:
		fmt.Errorf("Error during server info fetching: %s", error)
	}
}

func TestGetPlayerInfo(t *testing.T) {
	playerInfoChannel, errorChannel := GetPlayerInfo("89.163.177.130:27022", "500ms")

	select {
	case playerInfo := <-playerInfoChannel:
		t.Log(playerInfo)
	case error := <-errorChannel:
		t.Fatalf("Error during player info fecthing: %s", error)
	}
}

func ExampleGetPlayerInfo() {
	playerInfoChannel, errorChannel := GetPlayerInfo("89.163.177.130:27022", "500ms")

	select {
	case playerInfo := <-playerInfoChannel:
		fmt.Println(playerInfo)
	case error := <-errorChannel:
		fmt.Errorf("Error during player info fetching: %s", error)
	}
}
