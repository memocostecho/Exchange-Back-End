package secretobject

import "github.com/GoogleCloudPlatform/go-endpoints/endpoints"



func init() {
registerDevicesService := &RegisterDevicesService{}
  api, err := endpoints.RegisterService(registerDevicesService,
    "exchanges", "v1", "Exchange API", true)
  if err != nil {
    panic(err.Error())
  }

  register := func(orig, name, method, path, desc string) {
      m := api.MethodByName(orig)
      if m == nil {
      panic(m)
      }
      i := m.Info()
      i.Name, i.HTTPMethod, i.Path, i.Desc = name, method, path, desc
  }

  register("GetRegisteredDevices", "devices.get", "GET", "get", "List all the registered devices.")
  register("RegisterDevice", "devices.put", "GET", "register", "Insert a register on the datastore.")
  register("JoinExchange", "exchange.join", "GET", "joinExchange", "Join a user to an exchange.")
  register("RegisterUser", "user.put", "GET", "addUser", "Add a user to the system.")
  register("CreateExchange", "exchange.put", "GET", "addExchange", "Add an exchange.")
  register("GetUserExchanges", "exchange.list", "GET", "getuserexchanges", "List all the user exchanges")
  register("GetExchangeUser", "user.listexchanges", "GET", "getexchangesuser", "List all the users of an exchanges")
  register("GenerateRaffle", "raffle.trigger", "GET", "triggerrapple", "Trigger a raffle depends on the paramters")






/*
  servicioLogin := &ServicioLogin{}
  api, err = endpoints.RegisterService(servicioLogin,
    "login", "v1", "Quantity API", true)
  if err != nil {
    panic(err.Error())
  }

  


  register("RegisterLogin", "login.insertar", "GET", "registerUser", "Ingresa un usuario nuevo.")
  register("ChallengeLogin",  "login.entrar", "GET", "loginUser", "Inicia sesion en el sistema.")



  servicioRecibo:= &ServicioRecibo{}

  api, err = endpoints.RegisterService(servicioRecibo,
    "recibos", "v1", "Quantity API", true)
  if err != nil {
    panic(err.Error())
  }


  register("EnviaRecibo", "recibo.enviar", "GET", "enviarRecibo", "Envia un recibo al correo establecido")
*/

  endpoints.HandleHTTP()
}