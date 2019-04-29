package model

type Input struct {
  // normal or urgent
  Severity string
  DeviceName string
  DevicePropertiesName []string
}

type Action struct {
  DeviceName string
  //command name defined in device proile
  CommandName string
  //only get or set
  Method string
  //only used when command method is set
  Value interface{}
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
