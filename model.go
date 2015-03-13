package exchange

import "time"
import "appengine/datastore"







type UserExchangesReq struct {

  UserName string  `json:"username"`

}

type ExchangesUserReq struct {

  ExchangeId string  `json:"exchangeId"`
  Username string `json:"username"`

}

type TriggerRippleReq struct {

  ExchangeId string  `json:"exchangeId"`

}


type TriggerRippleRespDummy struct {

  Arreglo []int  `json:"arreglo"`


}


type User struct{

  Key     *datastore.Key `json:"id" datastore:"-"`
  Mail  string         `json:"mail"`
  Name  string         `json:"name"`
  Date time.Time  `json:"date"`
  OS string `json:"os"`
  Device string `json:"deviceRegistered"`
  Exchanges []string  `json:"exchanges"`  //`datastore:"-"`


}

type Exchange struct{

  Key     *datastore.Key `json:"id" datastore:"-"`
  Name  string         `json:"name"`
  Reason string         `json:"reason"`
  ExchangeId string         `json:"exchangeId"`
  Ammount string        `json:"ammount"`
  Users []string `json:"users"` //`datastore:"-"`


}


type ExchangesResp struct{

  UserExchanges []*Exchange `json:"userExchanges"`

}



type UsersResp struct{

  ExchangeUsers []*User `json:"exchangeUsers"`

}





type CreateExchangeReq struct{

  Exchange *Exchange  `json:"exchange"`
  Usermail  string  `json:"useremail"`

}


type GetExchangesReq struct{

  
  Usermail  string  `json:"useremail"`

}


type GetUsersReq struct{

  
  ExchangeId  string  `json:"exchangeId"`

}

type JoinExchangeReq struct{

  
  Usermail  string  `json:"useremail"`
  ExchangeId  string  `json:"exchangeId"`

}


type JoinExchangeResp struct{

  
  ExitoJoin  string  `json:"exitoBool"`
  

}






