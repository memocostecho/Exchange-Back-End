package exchange

//import "net/http"
import (
"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
"appengine/datastore"
"math/rand"
"appengine/mail"
apns "github.com/anachronistic/apns"
"github.com/alexjlockwood/gcm"
"appengine/urlfetch"
)
//import "fmt"
//import apns "github.com/anachronistic/apns"

//import "time"
//import "fmt"
//import "appengine/mail"


// GreetingService can sign the guesbook, list all greetings and delete
// a greeting from the guestbook.
type ExchangeService struct {
}



func (gs *ExchangeService) RegisterUser(c endpoints.Context, req *User) error {

// update the user with the mail NON REPEATED
  key := datastore.NewIncompleteKey(c, "User", nil)
  _, err := datastore.Put(c, key, req)

  return err
 
}







func (gs *ExchangeService) CreateExchange(c endpoints.Context, req *CreateExchangeReq) (*CreateExchangeResponse,error) {


  key := datastore.NewIncompleteKey(c, "Exchange", nil)
  


  // insercion del exchange en el usuario
  q := datastore.NewQuery("User").
        Filter("Mail =", req.Usermail).Limit(1)
  var users []User
  
  usersKeys, err := q.GetAll(c, &users)

  users[0].Exchanges= AppendString(users[0].Exchanges,req.Exchange.ExchangeId)

  _, errUser := datastore.Put(c, usersKeys[0], &users[0])
  if errUser!=nil{

       return &CreateExchangeResponse{"false"},errUser

   }



   // insercion del usuario en el Exchange  
  req.Exchange.Users= AppendString(req.Exchange.Users,req.Usermail)

  _, errExchange := datastore.Put(c, key, req.Exchange)
  if errExchange!=nil{

       return &CreateExchangeResponse{"false"},errExchange

   }



  return &CreateExchangeResponse{"true"},err
 
}


func (gs *ExchangeService) JoinExchange(c endpoints.Context, req *JoinExchangeReq) (*JoinExchangeResp,error) {


  //ASK FOR BELONGING
  // query the user to add the exchange
  q := datastore.NewQuery("User").
        Filter("Mail =", req.Usermail).Limit(1)
  var users []User
  
  usersKeys, err := q.GetAll(c, &users)
  if err!=nil{

       return &JoinExchangeResp{"false"},err

   }


  if !ExisteIntercambio(users[0].Exchanges,req.ExchangeId){

    users[0].Exchanges= AppendString(users[0].Exchanges,req.ExchangeId)  
    _, errUser := datastore.Put(c, usersKeys[0], &users[0])
     if errUser!=nil{

       return  &JoinExchangeResp{"false"},errUser

   }
        //return  &JoinExchangeResp{"true"},nil
  }
  
  // query the exchange to add the user
  queryExchange := datastore.NewQuery("Exchange").
        Filter("ExchangeId =", req.ExchangeId).Limit(1)
  var exchanges []Exchange
  
  exchangesKeys, errorExchanges := queryExchange.GetAll(c, &exchanges)
  if errorExchanges!=nil{

       return &JoinExchangeResp{"false"},errorExchanges

   }


  

    exchanges[0].Users= AppendString(exchanges[0].Users,req.Usermail)  
    _, errPutExchange := datastore.Put(c, exchangesKeys[0], &exchanges[0])
     if errPutExchange!=nil{

       return  &JoinExchangeResp{"false"},errPutExchange

   }


        queryAdmin := datastore.NewQuery("User").
        Filter("Mail =", exchanges[0].Users[0]).Limit(1)
        var userAdmin []User
  
        _, errorAdmin := queryAdmin.GetAll(c, &userAdmin)
        if errorAdmin!=nil{

          return &JoinExchangeResp{"false"},errorAdmin

        }



        var regIdToken string = userAdmin[0].Device

        var texto string

        if req.Language=="es"{
          texto = `Hola `+userAdmin[0].Name+`, se ha unido `+users[0].Mail+ `al intercambio `+exchanges[0].Name+`.`  
        }else if req.Language=="en" {
          texto = `Hello `+userAdmin[0].Name+`,`+users[0].Mail+ `joined to `+exchanges[0].Name+`.`  
        }
        

        if userAdmin[0].OS == "Android"{
            // notifications
          data := map[string]interface{}{"Exchange": texto }
          regIDs := []string{regIdToken}
          mensajeNotificacion := gcm.NewMessage(data, regIDs...)


        
          client := urlfetch.Client(c)
          sender := &gcm.Sender{ApiKey: "AIzaSyBxeyZ_O-eWzxRm8wrw8acr4iDr0C7qJK8", Http: client}

          _, errorNotification := sender.Send(mensajeNotificacion, 2)
          if errorNotification != nil {
          
          return &JoinExchangeResp{"false"},errorNotification
          }


        }else{



          payload := apns.NewPayload()
          payload.Alert = texto
    
          payload.Badge = 42
          payload.Sound = "bingbong.aiff"

          pn := apns.NewPushNotification()
          pn.DeviceToken =  regIdToken
          pn.AddPayload(payload)

          client := apns.NewClient("gateway.sandbox.push.apple.com:2195", "cert.pem", "key.pem",c)
          client.Send(pn)

          pn.PayloadString()




        }
        

    


  

  return  &JoinExchangeResp{"true"},nil
  


  
 
}


func (gs *ExchangeService) GetExchanges(c endpoints.Context, req *GetExchangesReq) (*ExchangesResp,error) {


  

  q := datastore.NewQuery("User").
        Filter("Mail =", req.Usermail).Limit(1)
  var users []*User
  
  _, err := q.GetAll(c, &users)
  if err!=nil{

    return nil,err
  }
    

  exchanges := make([]*Exchange,0,len(users[0].Exchanges))
  ex := make([]*Exchange,0,len(users[0].Exchanges))

  //exchanges = AppendExchange(exchanges,&ex)

  for _, exchange := range users[0].Exchanges {
        

        queryExchange := datastore.NewQuery("Exchange").
        Filter("ExchangeId =", exchange).Limit(1)  
        _, errorExchanges := queryExchange.GetAll(c, &ex)

        exchanges = AppendExchange(exchanges,ex[0])

        if errorExchanges!=nil{
          return nil,errorExchanges
        }
    }

  



  return &ExchangesResp{ex},nil
 
}



func (gs *ExchangeService) UpdateExchange(c endpoints.Context, req *UpdateExchangeReq) (*UpdateExchangeResp,error) {


  queryExchange := datastore.NewQuery("Exchange").
        Filter("ExchangeId =", req.ExchangeId).Limit(1)
  var exchanges []Exchange
  
  exchangesKeys, errorExchanges := queryExchange.GetAll(c, &exchanges)
  if errorExchanges!=nil{

       return &UpdateExchangeResp{"false"},errorExchanges


   }


   exchanges[0].Name = req.ExchangeNew.Name;
   exchanges[0].Reason = req.ExchangeNew.Reason;
   exchanges[0].Date = req.ExchangeNew.Date;
   exchanges[0].Ammount = req.ExchangeNew.Ammount;


    _, errExch := datastore.Put(c, exchangesKeys[0], &exchanges[0])
     if errExch!=nil{

       return  &UpdateExchangeResp{"false"},errExch

   }


   return  &UpdateExchangeResp{"true"},nil






  
 
}



func (gs *ExchangeService) GetUsers(c endpoints.Context, req *GetUsersReq) (*UsersResp,error) {


  

  q := datastore.NewQuery("Exchange").
        Filter("ExchangeId =", req.ExchangeId).Limit(1)
  var exchanges []*Exchange
  
  _, err := q.GetAll(c, &exchanges)
  if err!=nil{

    return nil,err
  }
    

  users := make([]*User,0,len(exchanges[0].Users))
  us := make([]*User,0,len(exchanges[0].Users))

  //exchanges = AppendExchange(exchanges,&ex)

  for _, user := range exchanges[0].Users {
        

        queryUser := datastore.NewQuery("User").
        Filter("Mail =", user).Limit(1)  
        _, errorUsers := queryUser.GetAll(c, &us)

        users = AppendUser(users,us[0])

        if errorUsers!=nil{
          return nil,errorUsers
        }
    }

  



  return &UsersResp{us,exchanges[0]},nil
 
}



func (gs *ExchangeService) TriggerRaffle(c endpoints.Context, req *TriggerRippleReq) (*TriggerRippleRespDummy,error) {
  
//Falta hacer el update del arreglo assignments.

  this := ExchangeService{}

  users,err := this.GetUsers(c, &GetUsersReq{req.ExchangeId})



  arregloEnteros := make([]int, len(users.ExchangeUsers),len(users.ExchangeUsers))  

  for i := range arregloEnteros {
        arregloEnteros[i] = i
  }

  perm := rand.Perm(len(arregloEnteros))
  for i, v := range perm {
    arregloEnteros[i] = v
  }


  for i, v := range arregloEnteros {







      var receiver int = 0

      receiver = i+1

      if i==len(arregloEnteros)-1{
        receiver = 0
      }
        
      var  texto  string 
      var textoMensaje string

        if req.Language=="es"{
      
          texto =  `En el intercambio `+users.ExchangeDetail.Name+` te toca regalarle a `+users.ExchangeUsers[arregloEnteros[receiver]].Name
          textoMensaje = 
          ` <h1>ExchangeApp</h1> 
          <b>
          Hola `+ users.ExchangeUsers[v].Name +`</b><br>
          En el intercambio `+users.ExchangeDetail.Name+` te toca regalarle a `+users.ExchangeUsers[arregloEnteros[receiver]].Name+`.`

        }else if req.Language=="en" {
          texto = `In the exchange `+users.ExchangeDetail.Name+` you give to `+users.ExchangeUsers[arregloEnteros[receiver]].Name
          textoMensaje = 
         ` <h1>ExchangeApp</h1> 
          <b>
          Hello `+ users.ExchangeUsers[v].Name +`</b><br>
          In the exchange `+users.ExchangeDetail.Name+` you give to `+users.ExchangeUsers[arregloEnteros[receiver]].Name+`.`

        }
      

      

        //create an exchange account 
        addr := "memogrr.dohko@gmail.com"
        msg := &mail.Message{
                Sender:  "memogrr.dohko@gmail.com",
                To:      []string{users.ExchangeUsers[v].Mail,addr},
                Subject: "Exchange Intercambio",                
                HTMLBody : textoMensaje ,
        }
        if err2 := mail.Send(c, msg); err2 != nil {
                c.Errorf("Couldn't send email: %v", err2)
        }


        
        var regIdToken string = users.ExchangeUsers[v].Device

        if users.ExchangeUsers[v].OS == "Android"{
            // notifications
          data := map[string]interface{}{"Exchange": texto }
          regIDs := []string{regIdToken}
          mensajeNotificacion := gcm.NewMessage(data, regIDs...)


        
          client := urlfetch.Client(c)
          sender := &gcm.Sender{ApiKey: "AIzaSyBxeyZ_O-eWzxRm8wrw8acr4iDr0C7qJK8", Http: client}

          _, errorNotification := sender.Send(mensajeNotificacion, 2)
          if errorNotification != nil {
          
          return nil,errorNotification
          }


        }else{



          payload := apns.NewPayload()
          payload.Alert = texto
    
          payload.Badge = 42
          payload.Sound = "bingbong.aiff"

          pn := apns.NewPushNotification()
          pn.DeviceToken =  regIdToken
          pn.AddPayload(payload)

          client := apns.NewClient("gateway.sandbox.push.apple.com:2195", "cert.pem", "key.pem",c)
          client.Send(pn)

          pn.PayloadString()




        }

        queryExchange := datastore.NewQuery("Exchange").
        Filter("ExchangeId =", req.ExchangeId).Limit(1)
        var exchanges []Exchange
  
       exchangesKeys, errorExchanges := queryExchange.GetAll(c, &exchanges)
       if errorExchanges!=nil{

        return &TriggerRippleRespDummy{arregloEnteros},errorExchanges

       }


  

      exchanges[0].Launched= "true"  
      _, errPutExchange := datastore.Put(c, exchangesKeys[0], &exchanges[0])
      if errPutExchange!=nil{

          return  &TriggerRippleRespDummy{arregloEnteros},errPutExchange

      }

        

    



      
  }



  return &TriggerRippleRespDummy{arregloEnteros},err
 
}






func ExtendString(slice []string, element string) []string {
    n := len(slice)
    if n == cap(slice) {
        // Slice is full; must grow.
        // We double its size and add 1, so if the size is zero we still grow.
        newSlice := make([]string, len(slice), 2*len(slice)+1)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0 : n+1]
    slice[n] = element
    return slice
}

func AppendString(slice []string, items ...string) []string {
    for _, item := range items {
        slice = ExtendString(slice, item)
    }
    return slice
}


func ExtendExchange(slice []*Exchange, element * Exchange) []*Exchange {
    n := len(slice)
    if n == cap(slice) {
        // Slice is full; must grow.
        // We double its size and add 1, so if the size is zero we still grow.
        newSlice := make([]*Exchange, len(slice), 2*len(slice)+1)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0 : n+1]
    slice[n] = element
    return slice
}

func AppendExchange(slice []*Exchange, items ... *Exchange) []*Exchange {
    for _, item := range items {
        slice = ExtendExchange(slice, item)
    }
    return slice
}


func ExtendUser(slice []*User, element * User) []*User {
    n := len(slice)
    if n == cap(slice) {
        // Slice is full; must grow.
        // We double its size and add 1, so if the size is zero we still grow.
        newSlice := make([]*User, len(slice), 2*len(slice)+1)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0 : n+1]
    slice[n] = element
    return slice
}

func AppendUser(slice []*User, items ... *User) []*User {
    for _, item := range items {
        slice = ExtendUser(slice, item)
    }
    return slice
}


func ExisteIntercambio(arreglo []string,intercambio string) bool {

  var existe bool = false


  for _, cadena := range arreglo {
        if cadena == intercambio{
          existe = true

        }
    }

  return existe

}