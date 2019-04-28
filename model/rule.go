package model

type Input struct {
  DeviceName string
  DevicePropertiesName []string
}

type Action struct {
  DeviceName string
  //command name defined in device proile
  commandName string
  //only get or set
  method string
  //only used when command method is set
  value interface{}
}

type Output struct {
  Type string
  Actions []Action
}

type Reading struct {
  Origin int64
  Device string
  Name string
  Value string
}
