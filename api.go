package exchange

import "github.com/GoogleCloudPlatform/go-endpoints/endpoints"



func init() {
exchangeServices := &ExchangeService{}
  api, err := endpoints.RegisterService(exchangeServices,
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

  
  register("RegisterUser", "user.put", "GET", "adduser", "Add a user to the system.")
  register("CreateExchange", "exchange.create", "GET", "createexchange", "Create an exchange.")
  register("GetExchanges", "exchanges.get", "GET", "getuserexchanges", "Get the user exchanges")
  register("JoinExchange", "exchange.join", "GET", "joinExchange", "Join a user to an exchange.")
  register("GetUsers", "users.get", "GET", "getusersofexchange", "Get the users of an exchange.")
  register("UpdateExchange", "exchange.update", "GET", "updateExchange", "Update an exchange.")
  register("TriggerRaffle", "raffle.trigger", "GET", "triggerraffle", "Trigger the raffle and send notifications and emails.")
  


  endpoints.HandleHTTP()
}