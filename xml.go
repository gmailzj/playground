package main

import (
	"fmt"

	"encoding/xml"
)

type TaxCtrlPrepareResult2 struct {
	TerminalNO string `json:"terminal_no" xml:"nsrsbh"`
}

//税控接口回复结构体
type TaxCtrlResponse struct {
	XMLName xml.Name       `xml:"business"`
	Id      string         `xml:"id,attr"`
	Comment string         `xml:"comment,attr"`
	Body    TaxCtrlRetBody `xml:"body"`
}

type TaxCtrlRetBody struct {
	Code    string      `xml:"returncode"`
	Message string      `xml:"returnmsg"`
	Data    interface{} `xml:"returndata"` //此处填充接口对应的结构体
}


func main() {
	res :=`
<?xml version="1.0" encoding="UTF-8"?>
<business id="92001" comment="申领准备信息查询">
  <body>
    <returncode>返回代码</returncode>
    <returnmsg>返回信息</returnmsg>
    <returndata>
      <nsrsbh>纳税人识别号</nsrsbh>
      <xnsbh>虚拟设备号</xnsbh>
      <sjrxx count="2">
        <group xh="1">
          <sjrmc>收件人名称</sjrmc>
          <sjdz>收件地址</sjdz>
          <yddh>移动电话</yddh>
          <gddh>固定电话</gddh>
          <yb>邮编</yb>
          <mrdz>默认地址</mrdz>
          <bzxx>备注信息</bzxx>
          <by1>备用1</by1>
          <by2>备用2</by2>
        </group>
        <group xh="2">
          <sjrmc>收件人名称</sjrmc>
          <sjdz>收件地址</sjdz>
          <yddh>移动电话</yddh>
          <gddh>固定电话</gddh>
          <yb>邮编</yb>
          <mrdz>默认地址</mrdz>
          <bzxx>备注信息</bzxx>
          <by1>备用1</by1>
          <by2>备用2</by2>
        </group>
      </sjrxx>
      <jbrxx count="2">
        <jbr xh="1">
          <jbrxm>经办人姓名</jbrxm>
          <zjlx>证件类型</zjlx>
          <zjhm>证件号码</zjhm>
        </jbr>
        <jbr xh="2">
          <jbrxm>经办人姓名</jbrxm>
          <zjlx>证件类型</zjlx>
          <zjhm>证件号码</zjhm>
        </jbr>
      </jbrxx>
      <pzhdxx count="2">
        <group xh="1">
          <fpzl>发票种类</fpzl>
          <fpzldm>发票种类代码</fpzldm>
          <fpzlmc>发票种类名称</fpzlmc>
          <fplxdm>发票类型代码</fplxdm>
          <mcfpxe>每次购票限额</mcfpxe>
          <myfpxe>每月购票限额</myfpxe>
          <zgcpl>最高持票量</zgcpl>
        </group>
        <group xh="2">
          <fpzl>发票种类</fpzl>
          <fpzldm>发票种类代码</fpzldm>
          <fpzlmc>发票种类名称</fpzlmc>
          <fplxdm>发票类型代码</fplxdm>
          <mcfpxe>每次购票限额</mcfpxe>
          <myfpxe>每月购票限额</myfpxe>
          <zgcpl>最高持票量</zgcpl>
        </group>
      </pzhdxx>
    </returndata>
  </body>
</business>


`
result := TaxCtrlPrepareResult2{}

	var response TaxCtrlResponse
	response.Body.Data = &result
	
	
err := xml.Unmarshal([]byte(res), &response)
    if err != nil {
        fmt.Printf("error:%s", err)
        return
    }
fmt.Printf("%s", response.Body.Data)
}


