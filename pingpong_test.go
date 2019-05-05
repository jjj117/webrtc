package webrtc

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"sync"
	"testing"

	"github.com/pion/logging"
)

func TestPingPong(t *testing.T) {
	log := logging.NewDefaultLoggerFactory().NewLogger("test")
	label := "test-channel"
	testData := []byte("this is some test data")

	checkError := func(t *testing.T, err error) {
		if err != nil {
			t.Error(err.Error())
		}
	}

	t.Run("Detach", func(t *testing.T) {
		// Use Detach data channels mode
		s := SettingEngine{}
		s.DetachDataChannels()
		api := NewAPI(WithSettingEngine(s))

		// Set up two peer connections.
		config := Configuration{}
		pca, err := api.NewPeerConnection(config)
		if err != nil {
			t.Fatal(err)
		}
		pcb, err := api.NewPeerConnection(config)
		if err != nil {
			t.Fatal(err)
		}

		defer func() { checkError(t, pca.Close()) }()
		defer func() { checkError(t, pcb.Close()) }()

		var wg sync.WaitGroup

		dcChan := make(chan *DataChannel)
		pcb.OnDataChannel(func(dc *DataChannel) {
			if dc.Label() != label {
				return
			}
			log.Debug("OnDataChannel was called")
			dc.OnOpen(func() {
				dcChan <- dc
			})
		})

		wg.Add(1)
		go func() {
			defer wg.Done()

			var dc io.ReadWriteCloser
			var msg []byte

			log.Debug("Waiting for OnDataChannel")
			attached := <-dcChan
			log.Debug("data channel opened")
			dc, err = attached.Detach()
			if err != nil {
				log.Debugf("Detach failed: %s\n", err.Error())
				t.Error(err)
			}
			defer func() { checkError(t, dc.Close()) }()

			log.Debug("Waiting for ping...")
			msg, err = ioutil.ReadAll(dc)
			log.Debugf("Received ping! \"%s\"\n", string(msg))
			if err != nil {
				t.Error(err)
			}

			log.Debug("Sending pong")
			if _, err = dc.Write(msg); err != nil {
				t.Error(err)
			}
			log.Debug("Sent pong")
		}()

		if err = signalPair(pca, pcb); err != nil {
			t.Fatal(err)
		}

		attached, err := pca.CreateDataChannel(label, nil)
		if err != nil {
			t.Fatal(err)
		}
		log.Debug("Waiting for data channel to open")
		open := make(chan struct{})
		attached.OnOpen(func() {
			open <- struct{}{}
		})
		<-open
		log.Debug("data channel opened")

		var dc io.ReadWriteCloser
		dc, err = attached.Detach()
		if err != nil {
			t.Fatal(err)
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			data := []byte(fmt.Sprintf("%s - %d", testData, 0))
			log.Debug("Sending ping...")
			if _, err := dc.Write(data); err != nil {
				t.Error(err)
			}
			log.Debug("Sent ping")

			checkError(t, dc.Close())

			log.Debug("Wating for pong")
			ret, err := ioutil.ReadAll(dc)
			if err != nil {
				log.Debug("Error here")
				t.Error(err)
				return
			}
			log.Debugf("Received pong! \"%s\"\n", string(ret))
			if !bytes.Equal(data, ret) {
				log.Debug("Received pong BAD!")
				t.Errorf("expected %q, got %q", string(data), string(ret))
			} else {
				log.Debug("Received pong GOOD!")
			}
		}()

		wg.Wait()
	})

	t.Run("No detach", func(t *testing.T) {
		// streams := 1

		// Use Detach data channels mode
		s := SettingEngine{}
		//s.DetachDataChannels()
		api := NewAPI(WithSettingEngine(s))

		// Set up two peer connections.
		config := Configuration{}
		pca, err := api.NewPeerConnection(config)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { checkError(t, pca.Close()) }()

		pcb, err := api.NewPeerConnection(config)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { checkError(t, pcb.Close()) }()

		var dca, dcb *DataChannel
		doneCh := make(chan struct{})

		pcb.OnDataChannel(func(dc *DataChannel) {
			if dc.Label() != label {
				return
			}

			log.Debugf("pcb: new datachannel: %s\n", dc.Label())

			dcb = dc
			// Register channel opening handling
			dcb.OnOpen(func() {
				log.Debug("pcb: datachannel opened")
			})

			// Register the OnMessage to handle incoming messages
			log.Debug("pcb: registering onMessage callback")
			dcb.OnMessage(func(dcMsg DataChannelMessage) {
				log.Debugf("pcb: received ping: %s\n", string(dcMsg.Data))
				if !reflect.DeepEqual(dcMsg.Data, testData) {
					t.Error("data mismatch")
				}

				if err = dcb.Send(dcMsg.Data); err != nil {
					t.Error(err)
				}
				log.Debug("pcb: sent ping")
				checkError(t, dcb.Close())
			})
		})

		dca, err = pca.CreateDataChannel(label, nil)
		if err != nil {
			t.Fatal(err)
		}

		dca.OnOpen(func() {
			log.Debug("pca: data channel opened")
			log.Debugf("pca: sending \"%s\"", string(testData))
			if err := dca.Send(testData); err != nil {
				t.Fatal(err)
			}
			log.Debug("pca: sent ping")
		})

		// Register the OnMessage to handle incoming messages
		log.Debug("pca: registering onMessage callback")
		dca.OnMessage(func(dcMsg DataChannelMessage) {
			log.Debugf("pca: received pong: %s\n", string(dcMsg.Data))
			if !reflect.DeepEqual(dcMsg.Data, testData) {
				t.Error("data mismatch")
			}
			checkError(t, dca.Close())
			close(doneCh)
		})

		if err := signalPair(pca, pcb); err != nil {
			t.Fatal(err)
		}

		<-doneCh
	})
}
