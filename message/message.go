package message

import (
  model "rule-processer-go/model"
  zmq "github.com/alecthomas/gozmq"
)

var chanRegisterMap map[string]chan[]model.Reading

func ChanRegister(pluginName string,ch chan []model.Reading) {
  ChannelRegisterMap[pluginName] = ch
}

func SubscribeFromZMQ() {
  context, _ := zmq.NewContext()
  socket, _ := context.NewSocket(zmq.REP)
  defer context.Close()
  defer socket.Close()
  socket.Bind("tcp://*:5555")

  tmpMap := make(map[string][]model.Reading)
  // Wait for messages
  for {
      msg, _ := socket.Recv(0)
      println("Received : ", string(msg))
      json.Unmarshal(msg,&tmpMap)
      for _,v := range chanRegisterMap {
        v <- tmpMap["readings"]
      }
  }
}
