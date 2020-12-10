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


func make_avp() string {
}
func make_header() string {
}
func make_sccrq() string {

        avp_data :=""
        avp_data += make_avp(attr_type=0, attr_value=1)
        avp_data += make_avp(attr_type=2, attr_value=256)
        avp_data += make_avp(attr_type=3, avp_len=10, attr_value1=0, attr_value2=3)
        avp_data += make_avp(attr_type=4, avp_len=10, attr_value1=0, attr_value2=0)
        avp_data += make_avp(attr_type=6, attr_value=1680)
        avp_data += make_avp(attr_type=7, avp_len=6+len(HOSTNAME), avp_raw_data=binascii.a2b_qp(HOSTNAME))
        avp_data += make_avp(attr_type=8, avp_len=6+len(VENDOR), avp_raw_data=binascii.a2b_qp(VENDOR))
        avp_data += make_avp(attr_type=9, attr_value=64773)
        avp_data += make_avp(attr_type=10, attr_value=4)
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