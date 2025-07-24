package gnb

import (
	"net"

	"github.com/free5gc/aper"
)

type XnUe struct {
	ulTeid aper.OctetString
	dlTeid aper.OctetString

	dataPlaneConn net.Conn
}

func NewXnUe(dlTeid aper.OctetString, dataPlaneConn net.Conn) *XnUe {
	return &XnUe{
		ulTeid: aper.OctetString{},
		dlTeid: dlTeid,

		dataPlaneConn: dataPlaneConn,
	}
}

func (x *XnUe) GetUlTeid() aper.OctetString {
	return x.ulTeid
}

func (x *XnUe) GetDlTeid() aper.OctetString {
	return x.dlTeid
}

func (x *XnUe) GetDataPlaneConn() net.Conn {
	return x.dataPlaneConn
}

func (x *XnUe) SetUlTeid(ulTeid aper.OctetString) {
	x.ulTeid = ulTeid
}

func (x *XnUe) SetDataPlaneConn(dataPlaneConn net.Conn) {
	x.dataPlaneConn = dataPlaneConn
}