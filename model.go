package secretobject

import "time"
import "appengine/datastore"


type RegisteredDevice struct {
  Key     *datastore.Key `json:"id" datastore:"-"`
  Rid  string         `json:"registrationID"`
  Date time.Time  `json:"date"`   

}

// GreetingsList is a response type of GreetingService.List method
type RegisteredDevices struct {
  RegisteredDevices []*RegisteredDevice `json:"devices"`
}


type UserExchanges struct {
  UserExchanges []*Exchange_User `json:"exchanges"`
}

type UserExchangesReq struct {

  UserName string  `json:"username"`

}

type ExchangesUserReq struct {

  ExchangeId string  `json:"exchangeId"`
  Username string `json:"username"`

}


type User struct{

  Key     *datastore.Key `json:"id" datastore:"-"`
  Mail  string         `json:"mail"`
  Name  string         `json:"name"`
  Date time.Time  `json:"date"`   
  Device string `json:"deviceRegistered"`


}

type Exchange struct{

  Key     *datastore.Key `json:"id" datastore:"-"`
  Name  string         `json:"name"`
  Reason string         `json:"reason"`
  ExchangeId string         `json:"exchangeId"`
  Ammount string        `json:"ammount"`


}


type Exchange_User struct{

 Key     *datastore.Key `json:"id" datastore:"-"`
 User string  `json:"user"`
 Exchange Exchange  `json:"exchange"`
 Administrator string        `json:"administrator"`
 Token string   `json:"token"`



}







