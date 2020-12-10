package main

import (
"fmt"
"syscall"
)
const (
HOSTNAME = "hostname"
VENDOR   = "pyl2tp"
LNSHOST  = "some.lns.com"
LNSPORT  = 1701
SECRET   = "asecret"
USERNAME = "user@realm"
TIMEOUT  = 60
)


func make_avp(attr_type int,avp_mbit string,avp_hbit string,avp_len int,avp_vid int,avp_raw_data string,kwargs map[string]interface{}) string {

        avp_flag = int(avp_mbit+avp_hbit+'000000',2)
        avp_data = struct.pack('!BBH', avp_flag, avp_len, avp_vid)
        avp_data += struct.pack('!H', attr_type)
        for key in sorted(kwargs.iterkeys()):
            avp_data += struct.pack('!H',kwargs[key])
        if avp_raw_data != '':
            avp_data += avp_raw_data
        return avp_data
}
func make_header() string {
}
func make_sccrq() string {

        avp_data :=""
        avp_data += make_avp(0,"1","0",8,0,"", attr_value=1)
        avp_data += make_avp(2,"1","0",8,0,"", attr_value=256)
        avp_data += make_avp(3, "1","0",10,0,"", attr_value1=0, attr_value2=3)
        avp_data += make_avp(4, "1","0",10,0,"", attr_value1=0, attr_value2=0)
        avp_data += make_avp(6, "1","0",10,0,"", attr_value=1680)
        avp_data += make_avp(7, "1","0",6+len(HOSTNAME),0,"", avp_raw_data=binascii.a2b_qp(HOSTNAME))
        avp_data += make_avp(8, "1","0",6+len(VENDOR),0,"", avp_raw_data=binascii.a2b_qp(VENDOR))
        avp_data += make_avp(9, "1","0",10,0,"", attr_value=64773)
        avp_data += make_avp(10, "1","0",10,0,"", attr_value=4)
        header_data := make_header(plen=12+87)
        l2tp_data = header_data + avp_data
        return l2tp_data





}
func main() {
      var (
       sock    syscall.Handle
       addr    syscall.SockaddrInet4
        wsadata syscall.WSAData
        err     error
    )
    if err = syscall.WSAStartup(MAKEWORD(2, 2), &wsadata); err != nil {
        fmt.Println("Startup error")
        return
    }
    defer syscall.WSACleanup()

   addr.Addr = inet_addr("127.0.0.1")
   addr.Port = 8000
  sock,_=syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM,0)
  syscall.Connect(sock,&addr)
  for {
      ns:=0

  }



}