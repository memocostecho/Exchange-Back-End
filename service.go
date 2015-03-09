package secretobject

//import "net/http"
import "github.com/GoogleCloudPlatform/go-endpoints/endpoints"
import "appengine/datastore"
import apns "github.com/anachronistic/apns"
import "fmt"
//import "time"
//import "fmt"
//import "appengine/mail"


// GreetingService can sign the guesbook, list all greetings and delete
// a greeting from the guestbook.
type RegisterDevicesService struct {
}

// List responds with a list of all greetings ordered by Date field.
// Most recent greets come first.
func (gs *RegisterDevicesService) GetRegisteredDevices(c endpoints.Context) (*RegisteredDevices, error) {

  
  q := datastore.NewQuery("RegisteredDevice").Order("-Date")
  devices := make([]*RegisteredDevice, 0, 1)
  keys, err := q.GetAll(c, &devices)
  if err != nil {
    return nil,err
  }

  devices = make([]*RegisteredDevice, 0, len(keys) )


  keys, err = q.GetAll(c, &devices)

   for i, k := range keys {
    devices[i].Key = k
  }


  return &RegisteredDevices{devices},nil
  
}



func (gs *RegisterDevicesService) GetUserExchanges(c endpoints.Context, req * UserExchangesReq) (*UserExchanges, error) {

  
  q := datastore.NewQuery("Exchange_User").Filter("User =",req.UserName)
  exchanges := make([]*Exchange_User, 0, 1)
  keys, err := q.GetAll(c, &exchanges)
  if err != nil {
    return nil,err
  }

  exchanges = make([]*Exchange_User, 0, len(keys) )


  keys, err = q.GetAll(c, &exchanges)

   for i, k := range keys {
    exchanges[i].Key = k
  }


  return &UserExchanges{exchanges},nil
  
}



func (gs *RegisterDevicesService) GetExchangeUser(c endpoints.Context, req * ExchangesUserReq) (*UserExchanges, error) {

  
  q := datastore.NewQuery("Exchange_User").Filter("Exchange.ExchangeId =",req.ExchangeId)
  exchanges := make([]*Exchange_User, 0, 1)
  keys, err := q.GetAll(c, &exchanges)
  if err != nil {
    return nil,err
  }

  exchanges = make([]*Exchange_User, 0, len(keys) )


  keys, err = q.GetAll(c, &exchanges)

   for i, k := range keys {
    exchanges[i].Key = k
  }


  return &UserExchanges{exchanges},nil
  
}



func (gs *RegisterDevicesService) GenerateRaffle(c endpoints.Context, req * ExchangesUserReq) (*UserExchanges, error) {

  



q := datastore.NewQuery("Exchange_User").Filter("Exchange.ExchangeId =",req.ExchangeId)
  exchanges := make([]*Exchange_User, 0, 1)
  keys, err := q.GetAll(c, &exchanges)
  if err != nil {
    return nil,err
  }

  exchanges = make([]*Exchange_User, 0, len(keys) )


  
  keys, err = q.GetAll(c, &exchanges)

   for _, k := range exchanges {
    
  payload := apns.NewPayload()
  /*if k==0 {

    payload.Alert = "Intercambio generado te toca darle regalo a luis carino"
    //payload.Alert = `Intercambio ` +i.exchange.name+ `generado, te toca darle a Luis Carino`
  }    
  else {
    payload.Alert = "Intercambio generado te toca darle regalo a memo"
    //payload.Alert = `Intercambio `+i.exchange.name+` generado, te toca darle a Guillermo Costecho`
  }
  */
    

  payload.Alert = "Intercambio generado te toca darle regalo a memo"
    
  payload.Badge = 42
  payload.Sound = "bingbong.aiff"

  pn := apns.NewPushNotification()
  pn.DeviceToken = k.Token//"5b466398b8b9fa4b1acd9c2bcecbcb7453d5d14e1b1f7fd080ece39d595b0707"
  pn.AddPayload(payload)

  client := apns.NewClient("gateway.sandbox.push.apple.com:2195", "cert.pem", "key.pem",c)
  resp := client.Send(pn)

  alert, _ := pn.PayloadString()
  fmt.Println("  Alert:", alert)
  fmt.Println("Success:", resp.Success)
  fmt.Println("  Error:", resp.Error)


  return &UserExchanges{exchanges},resp.Error

  }



  return nil,nil

  
  

  




  
}


func (gs *RegisterDevicesService) RegisterDevice(c endpoints.Context, req *RegisteredDevice) error {


  key := datastore.NewIncompleteKey(c, "RegisteredDevice", nil)
  _, err := datastore.Put(c, key, req)

  return err
 
}



func (gs *RegisterDevicesService) JoinExchange(c endpoints.Context, req *Exchange_User) error {


  key := datastore.NewIncompleteKey(c, "Exchange_User", nil)
  _, err := datastore.Put(c, key, req)

  return err
 
}


func (gs *RegisterDevicesService) RegisterUser(c endpoints.Context, req *User) error {


  key := datastore.NewIncompleteKey(c, "User", nil)
  _, err := datastore.Put(c, key, req)

  return err
 
}




func (gs *RegisterDevicesService) CreateExchange(c endpoints.Context, req *Exchange) error {


  key := datastore.NewIncompleteKey(c, "Exchange", nil)
  _, err := datastore.Put(c, key, req)

  return err
 
}



/*

type ServicioLogin struct {
}



func (gs *ServicioLogin) ChallengeLogin(c endpoints.Context, req *Login) (*RespuestaLogin,error) {

  
  q := datastore.NewQuery("Login").Filter("Password=", req.Password).Filter("Usuario=", req.Usuario)
  login:= make([]*Login, 0, 2) //debe ser nada mas uno... que coincida
  keys, err := q.GetAll(c, &login)
  if err != nil {
    return &RespuestaLogin{Respuesta:false},err
  }

  bandera := false

 // fmt.Println(keys)

  for _, k := range keys {
    bandera = true
    fmt.Println(k,bandera)
  }

  
  
  return &RespuestaLogin{Respuesta:bandera},nil
}



func (gs *ServicioLogin) RegisterLogin(c endpoints.Context, req *Login) error {

 


  q := datastore.NewQuery("Login").Filter("Password =", req.Password).Filter("Usuario =", req.Password)
  logines:= make([]*Login, 0, 1) //debe ser nada mas uno... que coincida
  keys, err := q.GetAll(c, &logines)
  if err != nil {
    return err
  }

  bandera := false

  for _, k := range keys {
    bandera = true
    fmt.Println(k,bandera)
  }


  if bandera==false {

    key := datastore.NewIncompleteKey(c, "Login", nil)
  _, err := datastore.Put(c, key, req)

    if err != nil {
    return err
  }

  return nil

  }
 // else {

  //  resp = "Usuario duplicado!"
//}
  

  

  return nil

  

  

  

  
}




type ServicioRecibo struct {
}


func (gs *ServicioRecibo) EnviaRecibo(c endpoints.Context, req *CorreoReq) error {
  


  mensaje := &ReciboCorreo{

  Nombre : req.Nombre,//"Guillermo Romero ", 
  Correo  : req.Correo,//"memo@hola",
  Mensaje : req.Mensaje,//"hola",
  Asunto :req.Asunto,//"saludo",
  Date: time.Now(),   
    
  }

  key := datastore.NewIncompleteKey(c, "ReciboCorreo", nil)
  _, err := datastore.Put(c, key, mensaje)


if err != nil {
    return err
  }

var textoMensaje = 
`Estimado Cliente de Quantity:
 El cliente `+req.Nombre+` ha enviado un recibo electronico:
 asunto:  `+req.Asunto+` 
 mensaje: `+req.Mensaje+`
 contacto:`+req.Correo+`
 PD: Este mensaje ya ha sido guardado en el datastore.`



const textoMensaje2 = 

`Gracias por comprar a traves de Quantity! 

Tu compra ha sido realizada exitosamente.

Rancheritos :   $20


Recibe un cordial saludo,
Quantity
`


  

        addr := "memogrr.dohko@gmail.com"//r.FormValue("email")
        //url := createConfirmationURL(r)
        msg := &mail.Message{
                Sender:  "memogrr.dohko@gmail.com",
                To:      []string{addr,"luiscrmz@gmail.com","alonsoauriazul@gmail.com","ozkarramdev@gmail.com"},
                Subject: "Quantity owner",
                Body:    fmt.Sprintf(textoMensaje),
        }
        if err2 := mail.Send(c, msg); err2 != nil {
                c.Errorf("Couldn't send email: %v", err2)
        }



        
        //url := createConfirmationURL(r)
        msg2 := &mail.Message{
                Sender:  "memogrr.dohko@gmail.com",
                To:      []string{req.Correo,"luiscrmz@gmail.com","alonsoauriazul@gmail.com","ozkarramdev@gmail.com"},
                Subject: "Quantity receipt",
                Body:    fmt.Sprintf(textoMensaje2),
        }
        if err3 := mail.Send(c, msg2); err3 != nil {
                c.Errorf("Couldn't send email: %v", err3)
        }


  return nil
}




*/



