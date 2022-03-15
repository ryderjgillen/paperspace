package grpc

import (
	"context"
	"encoding/binary"
	"fmt"
	"net"
	"time"

	pb "portService/grpc"
	"portService/server/netlink"
	"portService/server/proc"

	"github.com/kckeiks/netlink/sockdiag"
	"golang.org/x/sys/unix"
)

type ServiceData interface {
	Get() ([]*pb.PortInfoResponse, error)
}

type serviceData struct {
	ticker *time.Ticker
	data   []*pb.PortInfoResponse
	err    error
}

func NewServiceData(ctx context.Context, interval time.Duration, getData func() ([]*pb.PortInfoResponse, error)) ServiceData {

	serviceData := &serviceData{}

	//populate some data right away
	serviceData.data, serviceData.err = getData()

	//start the ticker
	serviceData.ticker = time.NewTicker(interval)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-serviceData.ticker.C:
				serviceData.data, serviceData.err = getData()
			}
		}
	}()

	return serviceData
}

func (d serviceData) Get() ([]*pb.PortInfoResponse, error) {
	return d.data, d.err
}

type inetDiagData struct{}

type protocol string

var (
	TCP protocol = "TCP"
	UDP protocol = "UDP"
)

type tcpState uint8

const (
	TCP_ESTABLISHED tcpState = 1
	TCP_LISTEN               = 10
)

//GetPortInfo returns information about the open ports on the current system
func GetPortInfo() ([]*pb.PortInfoResponse, error) {
	tcp4, err := getData(unix.AF_INET, unix.IPPROTO_TCP)
	if err != nil {
		return nil, err
	}

	tcp6, err := getData(unix.AF_INET6, unix.IPPROTO_TCP)
	if err != nil {
		return nil, err
	}

	udp4, err := getData(unix.AF_INET, unix.IPPROTO_UDP)
	if err != nil {
		return nil, err
	}

	udp6, err := getData(unix.AF_INET6, unix.IPPROTO_UDP)
	if err != nil {
		return nil, err
	}

	udpLite4, err := getData(unix.AF_INET, unix.IPPROTO_UDPLITE)
	if err != nil {
		return nil, err
	}

	udpLite6, err := getData(unix.AF_INET6, unix.IPPROTO_UDPLITE)
	if err != nil {
		return nil, err
	}

	data := make([]*pb.PortInfoResponse, 0, len(tcp4)+len(tcp6)+len(udp4)+len(udp6)+len(udpLite4)+len(udpLite6))
	data = append(data, tcp4...)
	data = append(data, tcp6...)
	data = append(data, udp4...)
	data = append(data, udp6...)
	data = append(data, udpLite4...)
	data = append(data, udpLite6...)

	return data, nil
}

func getData(family uint8, protocol uint8) ([]*pb.PortInfoResponse, error) {
	inetReq := sockdiag.InetDiagReqV2{
		Family:   family,
		Protocol: protocol,
		States:   1<<TCP_ESTABLISHED | 1<<TCP_LISTEN, //only states supported
	}
	header := unix.NlMsghdr{
		Len:   sockdiag.NlInetDiagReqV2MsgLen,
		Type:  sockdiag.SOCK_DIAG_BY_FAMILY,
		Flags: (unix.NLM_F_REQUEST | unix.NLM_F_DUMP),
		Pid:   0,
	}
	nlmsg, err := sockdiag.NewInetNetlinkMsg(header, inetReq)
	if err != nil {
		return nil, err
	}

	result, err := netlink.SendInetMessage(nlmsg)
	if err != nil {
		return nil, err
	}

	results := make([]*pb.PortInfoResponse, len(result))

	for idx, msg := range result {

		inode, err := proc.FindProcessByInode(msg.Inode)
		if err != nil {
			return nil, err
		}

		command, err := proc.GetProcessCommand(*inode)
		if err != nil {
			return nil, err
		}

		results[idx] = &pb.PortInfoResponse{
			Protocol: mapProtocol(protocol),
			Command:  command,
			Source: &pb.IpPort{
				IpAddress: toIpString(msg.ID.Src, family == unix.AF_INET6),
				Port:      uint32(binary.BigEndian.Uint16([]byte{msg.ID.SPort[0], msg.ID.SPort[1]})),
			},
			Destination: &pb.IpPort{
				IpAddress: toIpString(msg.ID.Dst, family == unix.AF_INET6),
				Port:      uint32(binary.BigEndian.Uint16([]byte{msg.ID.DPort[0], msg.ID.DPort[1]})),
			},
		}
	}

	return results, nil
}

func mapProtocol(protocol uint8) pb.Protocol {
	switch protocol {
	case unix.IPPROTO_TCP:
		return pb.Protocol_TCP
	case unix.IPPROTO_UDP:
		return pb.Protocol_UDP
	case unix.IPPROTO_UDPLITE:
		return pb.Protocol_UDP
	default:
		panic(fmt.Sprintf("unknown protocol: %d", protocol))
	}
}

func toIpString(bytes [16]byte, isIpv6 bool) string {
	if isIpv6 {
		return ((net.IP)(bytes[:])).String()
	} else {
		return net.IPv4(bytes[0], bytes[1], bytes[2], bytes[3]).To4().String()
	}
}
